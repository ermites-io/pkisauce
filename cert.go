//go:build go1.16
// +build go1.16

package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

type CA struct {
	cert    *x509.Certificate
	key     ed25519.PrivateKey
	dercert []byte
}

type Cert struct {
	csr     *x509.CertificateRequest
	cert    *x509.Certificate
	key     ed25519.PrivateKey
	dercert []byte
	derkey  []byte
}

// XXX let's be realistic we cannot defend against state actors and we run in the cloud haha :)
// but.. all communication needs basic strong encryption without too much overhead, ECC will help us with that.
// key exchange is much smaller than RSA and much quicker
// symmetric communication we will rely on AES-GCM (AEAD) in 128 bits
// AND ALWAYS PROTECT our HEADER with AD.
// to define our basis we rely on ECRYPT-CSA recommendations 2018 / for the next 10 years
// here :
// http://www.ecrypt.eu.org/csa/documents/D5.4-FinalAlgKeySizeProt.pdf
// https://www.keylength.com/en/3/
// Symmetric -> 128 bits (AES-GCM) (16 bytes)
// RSA -> 3072
// ECC -> 256
// Hash -> 256
//
// summarize as:
// CA RSA 3K
// Server & Clients ECC secp1P256
//
func NewSerial() (*big.Int, error) {
	randomBytesFeed := make([]byte, 1024)
	_, err := rand.Read(randomBytesFeed)
	if err != nil {
		return nil, err
	}

	// we limit serial number to a certain size.
	randomBytes := sha256.Sum256(randomBytesFeed)
	i := new(big.Int)
	i.SetBytes(randomBytes[:])

	return i, nil
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func computeSKI(pub crypto.PublicKey) ([]byte, error) {
	pubKeyBit, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		//sLog.Printf(1, "RET GetSHA1KeyId(%p) -> [Error: %s]\n", pub, err.Error())
		return nil, err
	}

	var spki subjectPublicKeyInfo
	_, err = asn1.Unmarshal(pubKeyBit, &spki)
	if err != nil {
		return nil, err
	}

	pubHash := sha256.Sum224(spki.SubjectPublicKey.Bytes)
	return pubHash[:], nil
}

func doPEM(dercert, derkey, password []byte) (pemcert, pemkey string) {
	// TODO: encrypt the PEM using password
	if len(dercert) > 0 {
		certPemBlock := &pem.Block{
			Type:    defaultPEMHeaderCert,
			Headers: nil,
			Bytes:   dercert,
		}
		pemcert = string(pem.EncodeToMemory(certPemBlock))
	}

	if len(derkey) > 0 {
		keyPemBlock := &pem.Block{
			Type:    defaultPEMHeaderKey,
			Headers: nil,
			Bytes:   derkey,
		}
		pemkey = string(pem.EncodeToMemory(keyPemBlock))
	}

	return
}

//CreateCertificate creates a new certificate based on a template. The following
//members of template are used: SerialNumber, Subject, NotBefore, NotAfter,
//KeyUsage, ExtKeyUsage, UnknownExtKeyUsage, BasicConstraintsValid, IsCA,
//MaxPathLen, SubjectKeyId, DNSNames, PermittedDNSDomainsCritical,
//PermittedDNSDomains, SignatureAlgorithm.

// we take info from the CA if we have nothing location based..
func newCSR(s pkix.Name, emails []string) (csr *x509.CertificateRequest, priv ed25519.PrivateKey, err error) {
	// generate key first
	//key, err = ecdsa.GenerateKey(DefaultECDSAPublicKeySize, rand.Reader)
	//pubkey, privkey, err := ed25519.GenerateKey(rand.Reader)
	_, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}

	// let's build the CSR and create the certificate request
	csr = new(x509.CertificateRequest)
	csr.Subject = s
	//csr.PublicKeyAlgorithm = DefaultECDSAPublicKeyAlgorithm
	csr.PublicKeyAlgorithm = x509.Ed25519
	//csr.SignatureAlgorithm = DefaultECDSASignatureAlgorithm
	csr.SignatureAlgorithm = x509.PureEd25519

	csr.DNSNames = []string{s.CommonName}
	csr.EmailAddresses = emails
	csr.Version = Defaultx509version

	// create the CertificateRequest
	der, err := x509.CreateCertificateRequest(rand.Reader, csr, priv)
	if err != nil {
		return
	}

	// replace our template by the generate CSR
	csr, err = x509.ParseCertificateRequest(der)
	if err != nil {
		return
	}

	//keyDer, err = x509.MarshalECPrivateKey(key)
	/*
		keyDer, err = x509.MarshalPKCS8PrivateKey(privkey)
		c = &Cert{
			csr: csr,
			pub: pubkey,
			key: privkey,
		}
	*/
	return
}

func NewCSR(cn, userdata string, email ...string) (cert *Cert, err error) {
	// adding userdata
	ou := defaultCertOrgUnit
	if len(userdata) > 0 {
		ou = append(ou, userdata)
	}

	csrSubject := pkix.Name{
		Country:            defaultCertCountry,
		Organization:       defaultCertOrg,
		OrganizationalUnit: ou,
		Locality:           defaultCertLocality,
		Province:           defaultCertProvince,
		StreetAddress:      defaultCertStreetAddress,
		PostalCode:         defaultCertPostalCode,
		CommonName:         cn,
	}

	csr, key, err := newCSR(csrSubject, email)
	if err != nil {
		return
	}

	keyder, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return
	}

	cert = &Cert{
		csr:    csr,
		key:    key,
		derkey: keyder,
	}

	return
}

func (c *Cert) PEM(password []byte) (pemcert, pemkey string) {
	// TODO: encrypt the PEM using password
	return doPEM(c.dercert, c.derkey, password)
}

// GenerateCA generate a CA certificate using s as subject.
func newCA(s pkix.Name, expdays int) (ca *x509.Certificate, cakey ed25519.PrivateKey, cader, cakeyder []byte, err error) {
	ca = new(x509.Certificate)

	//certRootTemplate.PublicKeyAlgorithm = DefaultECDSAPublicKeyAlgorithm
	ca.PublicKeyAlgorithm = x509.Ed25519
	//certRootTemplate.SignatureAlgorithm = DefaultECDSASignatureAlgorithm
	ca.SignatureAlgorithm = x509.PureEd25519
	ca.Subject = s

	randomSerial, err := NewSerial()
	if err != nil {
		//return nil, nil, nil, nil, err
		return
	}
	ca.SerialNumber = randomSerial
	ca.Subject.SerialNumber = randomSerial.Text(16)

	// TIME
	ca.NotBefore = time.Now()
	// 30 days.
	ca.NotAfter = ca.NotBefore.AddDate(0, 0, expdays)

	ca.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	ca.ExtKeyUsage = nil //[]x509.ExtKeyUsage{x509.ExtKeyUsageOCSPSigning}
	ca.UnknownExtKeyUsage = nil

	ca.BasicConstraintsValid = true
	ca.IsCA = true
	ca.MaxPathLen = 1

	ca.DNSNames = nil
	ca.EmailAddresses = nil
	ca.IPAddresses = nil

	// first create an RSA key
	//ecdsakey, err := ecdsa.GenerateKey(DefaultECDSAPublicKeySize, rand.Reader)
	pubkey, privkey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		//return nil, nil, nil, nil, err
		return
	}

	// according to RFC 3280 SHA-1 of BIT STRING of publickey
	//certRootTemplate.SubjectKeyId, err = computeSKI(ecdsakey.Public())
	/*
		certRootTemplate.SubjectKeyId, err = computeSKI(pubkey)
		if err != nil {
			return nil, nil, nil, nil, err
		}
	*/

	// then create the certificate
	//caPubDer, err := x509.CreateCertificate(rand.Reader, certRootTemplate, certRootTemplate, &ecdsakey.PublicKey, ecdsakey)
	cader, err = x509.CreateCertificate(rand.Reader, ca, ca, pubkey, privkey)
	if err != nil {
		//	return nil, nil, nil, nil, err
		return
	}

	// this is just a verification.
	ca, err = x509.ParseCertificate(cader)
	if err != nil {
		return
	}

	//caPrivDer, err := x509.MarshalECPrivateKey(ecdsakey)
	cakeyder, err = x509.MarshalPKCS8PrivateKey(privkey)
	if err != nil {
		//return nil, nil, nil, nil, err
		return
	}
	cakey = privkey
	return
}

func NewCA(expDays int) (ca *CA, err error) {
	// create the CA
	caSubject := pkix.Name{
		Country:            defaultCertCountry,
		Organization:       defaultCertOrg,
		OrganizationalUnit: defaultCaCertOrgUnit,
		Locality:           defaultCertLocality,
		Province:           defaultCertProvince,
		StreetAddress:      defaultCertStreetAddress,
		PostalCode:         defaultCertPostalCode,
		CommonName:         defaultCaCommonName,
	}

	cacert, cakey, capubder, _, err := newCA(caSubject, expDays)
	if err != nil {
		return nil, err
	}

	ca = &CA{
		cert:    cacert,
		key:     cakey,
		dercert: capubder,
	}
	return
}

func (ca *CA) PEM() (cacert string) {
	cacert, _ = doPEM(ca.dercert, nil, nil)
	return
}

func (ca *CA) Sign(entry *Cert, expDays, purpose int) (c *Cert, err error) {
	if entry.csr == nil {
		err = fmt.Errorf("invalid input")
		return
	}
	cert := new(x509.Certificate)

	cert.Subject = entry.csr.Subject
	cert.Version = entry.csr.Version

	cert.SignatureAlgorithm = ca.cert.SignatureAlgorithm
	cert.PublicKeyAlgorithm = entry.csr.PublicKeyAlgorithm

	cert.EmailAddresses = entry.csr.EmailAddresses
	realEmail := pkix.AttributeTypeAndValue{
		Type:  asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1},
		Value: cert.EmailAddresses[0],
	}
	cert.Subject.ExtraNames = append(cert.Subject.ExtraNames, realEmail)

	cert.DNSNames = entry.csr.DNSNames
	cert.IsCA = false
	cert.BasicConstraintsValid = true

	// TODO check that all is ok..
	/*
		cert.SubjectKeyId, err = computeSKI(entry.pub)
		if err != nil {
			return
		}
	*/

	cert.AuthorityKeyId = ca.cert.SubjectKeyId
	cert.SerialNumber, err = NewSerial()
	if err != nil {
		return
	}

	// default key usage for a CLIENT certificate
	cert.KeyUsage = x509.KeyUsageContentCommitment | x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement
	// TODO XXX rework it to allow all to NOT be mutually exclusive anymore.
	switch {
	case (purpose & PurposeClientCert) == PurposeClientCert:
		cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageTimeStamping, x509.ExtKeyUsageEmailProtection}
	case (purpose & PurposeServerCert) == PurposeServerCert:
		cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageTimeStamping, x509.ExtKeyUsageEmailProtection}
	}

	// Expiration...
	cert.NotBefore = time.Now()
	cert.NotAfter = cert.NotBefore.AddDate(0, 0, expDays)

	entry.dercert, err = x509.CreateCertificate(rand.Reader, cert, ca.cert, entry.csr.PublicKey, ca.key)
	if err != nil {
		return
	}

	// replace our template by the generate CSR
	entry.cert, err = x509.ParseCertificate(entry.dercert)
	if err != nil {
		return
	}
	c = entry
	return
}
