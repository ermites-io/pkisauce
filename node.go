//go:build go1.17
// +build go1.17

package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/hkdf"

	"ermites.io/pkisauce/config"
	"ermites.io/pkisauce/policy"
	"ermites.io/pkisauce/scramble"
)

// it's fairly simple, but a little bit complicated.
// because we know all the serving servers and whom they allow to connect,
// we derivate the clients from these definitions and nothing more.
// by default as long as you go through these helpers, nothing else should be allowed.
// the overhead is negligible, since those connections are supposed to be long lived.
// we will add GRPC load balancing code also but later and nothing will change as.
// each services, will still follow the same identities/nodes.

//
// POLICY MAP
//

type NodePkiInfo map[string]string

type NodeInfo struct {
	Name   string // real name
	HName  string // the "human" name, really the code name
	PHName string // the policy name
	UUID   string // the UUID.
}

func NewNodeInfo(ph *scramble.Hasher, name, uuid string) *NodeInfo {
	h := sha256.Sum256([]byte(name))
	hname := hex.EncodeToString(h[:])

	phname := ""
	if ph != nil {
		phname = ph.Sum64([]byte(uuid))
	}

	return &NodeInfo{
		Name:   name,
		HName:  hname,
		PHName: phname,
		UUID:   uuid,
	}
}

type Node struct {
	// Template part
	Name     string // name of the node.
	HName    string
	Time     string
	DebugLog bool

	Server     bool
	ServerUUID string // serving UUID (the listening part)
	ServerCert *Cert

	Client     bool
	ClientUUID string // client certificate UUID (what we use to connect to another services)
	ClientCert *Cert

	PemCA string // PEM CA
	// New Policy
	// 64 bytes, [0:32] bytes for hmac [32:64] for private key encryption
	//keyReader io.Reader // this is the constructed HKDF
	//key        []byte    // MasterKey we will derivate the policy and the cert key from this.
	k1, k2, k3 []byte // all the keys we need.
	Key        string // the string version for template

	//Policy       *policy.Policy // the node policy for clients..
	Policy     policy.Hosts // server template policy
	PolicyBlob string       // compiled and stored policy.
	PkiBlob    string
	PkiInfo    NodePkiInfo

	Servers map[string]*NodeInfo // list of server I can connect to.

	Path   string // path where to write
	Flavor int    // the node flavor, basically what to generate
}

func (n *Node) PkiClientName() string {
	return n.ClientUUID
}

func (n *Node) CertClient() *Cert {
	return n.ClientCert
}

func (n *Node) CertServer() *Cert {
	return n.ServerCert
}

// this export data and marshall the pki in a map["ca"] -> caPEM
// ca, cc, ck, sc, sk : respectively client cert, client key, server cert, server key, ca

func (n *Node) Export() (err error) {
	cltcert, cltkey := n.CertClient().PEM(nil)
	srvcert, srvkey := n.CertServer().PEM(nil)

	n.PkiInfo["ca"] = n.PemCA
	n.PkiInfo["cc"] = cltcert
	n.PkiInfo["ck"] = cltkey
	n.PkiInfo["sc"] = srvcert
	n.PkiInfo["sk"] = srvkey

	ihname := sha256.Sum256([]byte(n.Name))
	ihk64 := sha256.Sum256([]byte(n.Key))

	ad := ihk64[:]
	ad = append(ad, ihname[:]...)

	n.PkiBlob, err = scramble.Marshal(n.k1, ad, n.PkiInfo)
	if err != nil {
		return
	}

	return
}

func (n *Node) PkiServerName() string {
	return n.ServerUUID
}

func (n *Node) String() (str string) {
	str += "\n---\n"
	str += fmt.Sprintf("name: %s\n", n.Name)
	str += fmt.Sprintf("flavor: %v\n", n.Flavor)
	str += fmt.Sprintf("uuid server: %s\n", n.ServerUUID)
	//	str += fmt.Sprintf("Pem server: %.100s\n", n.PemSrv)
	//	str += fmt.Sprintf("Pem key: %.100s\n", n.PemSrvKey)
	str += fmt.Sprintf("uuid client: %s\n", n.ClientUUID)
	//	str += fmt.Sprintf("Pem client: %.100s\n", n.PemClt)

	str += fmt.Sprintf("policy: %v\n", n.Policy)
	//str += fmt.Sprintf("policy map: %v\n", n.PolicyServerGRPC)
	return
}

func (n *Node) writeTemplate(outfile string, tmpl *template.Template) error {
	// if absolute, use the path and check we can write in it.
	// if relative concatenate and try.
	outputPath, err := filepath.Abs(n.Path)
	if err != nil {
		return err
	}
	outputFilePath := filepath.Join(outputPath, outfile)
	//tmplOutputFilename := fmt.Sprintf("pkisauce_server_%s.go.tmp", n.Name)
	tmplOutputFd, err := os.OpenFile(outputFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer tmplOutputFd.Close()

	err = tmpl.Execute(tmplOutputFd, n)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) WriteCertificate(outfile string) error {
	return nil
}

func (n *Node) setPolicy(h policy.Hosts) {
	n.Policy = h

	// build adh..
	adh := sha256.New()

	adhca := sha256.Sum256([]byte(n.PkiInfo["ca"]))
	adh.Write(adhca[:])
	adhsc := sha256.Sum256([]byte(n.PkiInfo["sc"]))
	adh.Write(adhsc[:])
	adhsk := sha256.Sum256([]byte(n.PkiInfo["sk"]))
	adh.Write(adhsk[:])
	adhcc := sha256.Sum256([]byte(n.PkiInfo["cc"]))
	adh.Write(adhcc[:])
	adhck := sha256.Sum256([]byte(n.PkiInfo["ck"]))
	adh.Write(adhck[:])

	ad := adh.Sum(nil)

	// marshal our stuff
	pstr, err := scramble.Marshal(n.k2, ad, n.Policy)
	if err != nil {
		panic(err)
	}
	n.PolicyBlob = pstr
}

func (n *Node) setServers(s map[string]*NodeInfo) {
	n.Servers = s
}

func (n *Node) generateDebug() (err error) {
	tmpl, err := template.ParseFS(tFS, TemplateNameDebug)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, n)
	return
}

func (n *Node) generateFilesGo() (err error) {
	//
	// main files
	//
	err = n.generateFilesGoHelpersMain()
	if err != nil {
		return
	}
	//
	// scramble templates
	//
	err = n.generateFilesGoHelpersScramble()
	if err != nil {
		return
	}

	return
}

func (n *Node) generateFilesGoHelpersMain() (err error) {
	//
	// common
	//
	cmntmpl, err := template.ParseFS(tFS, TemplateHelpersCommon)
	if err != nil {
		return err
	}
	outputFilename := "helpers_common.pksc.go"
	err = n.writeTemplate(outputFilename, cmntmpl)
	if err != nil {
		return err
	}
	//fmt.Printf("-> generated %s\n", outputFilename)

	//
	// now the helpers
	//
	hlprtmpl, err := template.ParseFS(tFS, TemplateHelpers)
	if err != nil {
		return err
	}

	// to help defer close each file..
	outputFilename = "helpers_" + n.Name + ".pksc.go"
	err = n.writeTemplate(outputFilename, hlprtmpl)
	if err != nil {
		return err
	}
	//fmt.Printf("-> generated %s\n", outputFilename)

	return
}

func (n *Node) generateFilesGoHelpersScramble() (err error) {
	fileSuffixes := []string{"error.go", "hash.go", "marshal.go"}
	for _, fileSuffix := range fileSuffixes {
		templateFile := fmt.Sprintf("templates/helpers_scramble_%s.template", fileSuffix)

		// parse template first
		cmntmpl, err := template.ParseFS(tFS, templateFile)
		if err != nil {
			return err
		}
		//outputFilename := "helpers_scramble.pksc.go"
		outputFilename := fmt.Sprintf("helpers_scramble_%s.pksc.go", fileSuffix)
		err = n.writeTemplate(outputFilename, cmntmpl)
		if err != nil {
			return err
		}
	}
	return
}

func (n *Node) generateFilesPEM() (err error) {
	//
	// CA
	//
	catmpl, err := template.ParseFS(tFS, TemplateCAPEM)
	if err != nil {
		return err
	}
	outputFilename := "ca.pksc.pem"
	err = n.writeTemplate(outputFilename, catmpl)
	if err != nil {
		return err
	}
	//fmt.Printf("-> generated %s\n", outputFilename)

	//
	// client
	//
	//if len(n.To) > 0 {
	if n.Client {
		clttmpl, err := template.ParseFS(tFS, TemplateClientPEM)
		if err != nil {
			return err
		}
		outputFilename = n.Name + "_client_cert.pksc.pem"
		err = n.writeTemplate(outputFilename, clttmpl)
		if err != nil {
			return err
		}
		//fmt.Printf("-> generated %s\n", outputFilename)
	}

	//
	// servers
	//
	//if len(n.From) > 0 {
	if n.Server {
		srvtmpl, err := template.ParseFS(tFS, TemplateServerPEM)
		if err != nil {
			return err
		}
		outputFilename = n.Name + "_server_cert.pksc.pem"
		err = n.writeTemplate(outputFilename, srvtmpl)
		if err != nil {
			return err
		}
		//fmt.Printf("-> generated %s\n", outputFilename)
	}
	return
}

func (n *Node) generateFiles() (err error) {

	switch n.Flavor {
	case config.FlavorGoTLS, config.FlavorGoGRPC:
		err = n.generateFilesGo()
	case config.FlavorPEM:
		err = n.generateFilesPEM()
	}

	return
}

func NewNodeWithUUIDs(name, path, userdata string, flavor int, pol policy.Hosts, debug bool, srvuuid, cltuuid string) (n *Node, err error) {
	srvcsr, err := NewCSR(srvuuid, userdata, "pkisauce@ermites.io")
	if err != nil {
		return
	}

	cltcsr, err := NewCSR(cltuuid, userdata, "pkisauce@ermites.io")
	if err != nil {
		return
	}

	// just a random key we will use.
	// OLD
	// k1 := [0:16] + [32:48] --> CERT + gob blob Key
	// k2 := [16:32] + [48:64] --> hmac key
	//
	// NEW:
	// hkdf salt := [0:32]
	// hkdf key := [32:64]
	//
	// then later:
	// hkdf salt := [0:16] + [32:48]
	// hkdf key := [16:32] + [48:64]
	k := make([]byte, 64)
	_, err = rand.Read(k)
	if err != nil {
		return
	}
	k64 := b64.EncodeToString(k)

	ihname := sha256.Sum256([]byte(name))
	ihk64 := sha256.Sum256([]byte(k64))

	info := ihname[:]
	info = append(info, ihk64[:]...)

	krdr := hkdf.New(sha256.New, k[32:64], k[0:32], info)

	// encode to have names in the code.
	h := sha256.Sum256([]byte(name))
	hname := hex.EncodeToString(h[:])

	// read the keys

	// PKI key
	k1 := make([]byte, 32)
	_, err = krdr.Read(k1)
	if err != nil {
		return
	}

	// policy key
	k2 := make([]byte, 32)
	_, err = krdr.Read(k2)
	if err != nil {
		return
	}

	// hmac key
	k3 := make([]byte, 32)
	_, err = krdr.Read(k3)
	if err != nil {
		return
	}

	// shall we add PEM only definition ?
	// as we will describe the trust relationships using pkisauce
	// and generate code as well as a graph to understand the links...
	server := false
	if len(pol) > 0 {
		server = true
	}

	n = &Node{
		Name:     name,
		HName:    hname,
		Time:     time.Now().String(),
		DebugLog: debug,

		k1:  k1,
		k2:  k2,
		k3:  k3,
		Key: k64, // visible

		Server:     server,
		ServerUUID: srvuuid,
		ServerCert: srvcsr,
		PkiInfo:    make(NodePkiInfo),
		// by default all element can be clients...
		// let's not discriminate, since the servers ARE doing the filtering when they can.
		Client:     true,
		ClientUUID: cltuuid,
		ClientCert: cltcsr,

		Path:   path,
		Flavor: flavor,
		Policy: pol,
	}
	return
}

func NewNode(name, path, userdata string, flavor int, pol policy.Hosts, debug bool) (*Node, error) {
	return NewNodeWithUUIDs(name, path, userdata, flavor, pol, debug, uuid.New().String(), uuid.New().String())
}
