package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	if err := generateCertAndKey(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func generateCertAndKey() error {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNo, err := rand.Int(rand.Reader, max)
	if err != nil {
		return err
	}
	subject := pkix.Name{
		Organization:       []string{"TestOrg Co."},
		OrganizationalUnit: []string{"Cyber"},
		CommonName:         "Stamp CLI",
	}

	template := x509.Certificate{
		SerialNumber: serialNo,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	if err != nil {
		return err
	}

	// Create cert
	certOut, err := os.Create("cert.pem")
	if err != nil {
		return err
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return err
	}

	// Create key
	keyOut, err := os.Create("key.pem")
	if err != nil {
		return err
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "RSA_PRIVATE_KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}); err != nil {
		return err
	}

	return nil
}
