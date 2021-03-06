//go:build go1.17
// +build go1.17

package main

import (
	"fmt"

	"ermites.io/pkisauce/config"
	"ermites.io/pkisauce/policy"
	"ermites.io/pkisauce/scramble"
)

type PKI struct {
	ca       *CA                  // the CA
	userdata string               // data to add in each certificate
	hasWc    bool                 // does the PKI has a wildcard defined host, if yes all hosts are potentially clients.
	M        map[string]*Node     // the map of nodes.
	S        map[string]*NodeInfo // the additionnal map info.
	C        map[string]bool
}

// this clones the current server list to be copied into a node server infos.
// to become "node relative" as each node has its own cryptographic key to protect
// certificates, policies and names..
func (p *PKI) Servers(ph *scramble.Hasher) (s map[string]*NodeInfo) {
	s = make(map[string]*NodeInfo)
	for k, v := range p.S {
		s[k] = NewNodeInfo(ph, v.Name, v.UUID)
	}
	return
}

func (p *PKI) Debug(name string) (err error) {
	node, ok := p.Node(name)
	if !ok {
		err = fmt.Errorf("!! definition error: '%s' is NOT defined", name)
		return
	}
	if p.IsClient(node.Name) {
		node.Client = true
	}

	// build the policy map
	err = p.BuildNodePolicy(node)
	if err != nil {
		return
	}

	// display that node
	err = node.generateDebug()
	return
}

func (p *PKI) String() (str string) {
	for n, node := range p.M {
		str += fmt.Sprintf("- node: %s -\n", n)
		str += fmt.Sprintf("%s\n", node)
	}

	return
}

func (p *PKI) Sign(n *Node) (err error) {
	_, err = p.ca.Sign(n.CertClient(), 180, PurposeClientCert)
	if err != nil {
		return
	}
	_, err = p.ca.Sign(n.CertServer(), 180, PurposeServerCert)
	if err != nil {
		return
	}
	return
}

func (p *PKI) SignAndAdd(n *Node) (err error) {
	err = p.Sign(n)
	if err != nil {
		return
	}
	p.Add(n)
	return
}

func (p *PKI) Add(n *Node) {
	n.PemCA = p.ca.PEM()
	p.M[n.Name] = n

	// is it a server? yes add it!
	if n.Server {
		p.S[n.Name] = NewNodeInfo(nil, n.Name, n.ServerUUID)
	}

	// add all policy hosts as client
	// add wc and trigger bool
	for h := range n.Policy {
		p.C[h] = true
	}

	// if ONE of our PKI node has a wildcard host, then all host are clients.
	// we use that later to see if we include or NOT the client certificate in the resulting blob
	if n.Policy.HasWc() {
		p.hasWc = true
	}
}

func (p *PKI) IsClient(name string) (ok bool) {
	_, cok := p.C[name]
	wcok := p.hasWc
	if cok || wcok {
		ok = true
	}
	return
}

func (p *PKI) IncludeConfig(cfg config.Config) (err error) {
	for name, opt := range cfg {
		err = opt.IsValid(name)
		if err != nil {
			err = fmt.Errorf("!! definition error: %s err: %v", name, err)
			return
		}

		_, ok := p.Node(name)
		if ok {
			err = fmt.Errorf("!! definition error: %s is already defined", name)
			return
		}

		// TODO: handle faulty flavor.
		//flavor, _ := config.FlavorKeyword[opt.Flavor]
		//fmt.Printf("HOSTS[%s]: %#v\n", name, opt.Policy.Hosts())

		//fmt.Printf("NAME: %s POLICY: %v\n", name, opt.Policy)
		node, nerr := NewNode(name, opt.Path, p.userdata, opt.Flavor, opt.Policy.Hosts(), opt.Debug)
		if nerr != nil {
			err = nerr
			return
		}

		// XXX TODO: create p.PrepareNode(node)
		// which will create the node, sign with the pki
		// and export the PEM certificate for that node directly.
		err = p.SignAndAdd(node)
		if err != nil {
			return
		}

		// make it ready for templating
		/*
			err = node.Export()
			if err != nil {
				return
			}
		*/
	} // we're done with cfg.
	return
}

func (p *PKI) Node(name string) (n *Node, ok bool) {
	n, ok = p.M[name]
	return
}

func (p *PKI) Nodes() map[string]*Node {
	return p.M
}

/// XXX not implemented.
func (pki *PKI) Validate() {
}

// TODO: simplify... no need of this switch anymore :)
func (pki *PKI) BuildNodePolicy(n *Node) (err error) {
	switch n.Flavor {
	case config.FlavorGoGRPC, config.FlavorGoTLS:
		return pki.buildNodePolicyServer(n)
	case config.FlavorJavaGRPC: // XXX not implemented
		fallthrough
	case config.FlavorJavaTLS: // XXX not implemented
		fallthrough
	case config.FlavorCppGRPC: // XXX not implemented
		fallthrough
	case config.FlavorPEM: // nothing to do.. regular CA:)
		return
	}
	return
}

// I always wanted to pretend i was "compiling" :))
func (pki *PKI) buildNodePolicyServer(n *Node) (err error) {
	// the new version will do it this way..
	// we create an encrypted policy for THAT NODE.
	// using hmac etc..
	ph := scramble.NewHasher(n.k3)
	if err != nil {
		return
	}

	//hosts := n.Policy
	np := policy.New(ph, "*")

	// XXX temporary
	/*
		fmt.Printf("POLICY[%s]: %#v\n", n.HName, n.Policy)
		fmt.Printf("POLICY[%s] Client: %v Server: %v\n", n.Name, n.Client, n.Server)
	*/

	for host, svcs := range n.Policy {
		// we don't need the wildcard pki node anymore..
		// by default we consider all is wildcard
		// so all nodes are client
		clientuuid := "*"
		if host != "*" {
			client, ok := pki.Node(host)
			if !ok {
				err = fmt.Errorf("missing node '%s'", host)
				return
			}
			//fmt.Printf("CLIENT: %s\n", client.Name)
			//client.Client = true
			clientuuid = client.PkiClientName()
		}
		for svc, rpcs := range svcs {
			for rpc := range rpcs {
				np.Entry(clientuuid, svc, rpc)
			}
		}
	}

	// that's our policy now.
	n.setPolicy(np.Hosts())
	n.setServers(pki.Servers(ph))
	return
}

func NewPKI(ca *CA, userdata string) (p *PKI, err error) {
	p = &PKI{
		ca:       ca,
		userdata: userdata,
		hasWc:    false,
		M:        make(map[string]*Node),
		S:        make(map[string]*NodeInfo),
		C:        make(map[string]bool),
	}
	return
}
