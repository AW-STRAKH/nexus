package quic

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

func generateTLSConfig() (*tls.Config, error) {

	key, err := rsa.GenerateKey(
		rand.Reader,
		2048,
	)

	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),

		Subject: pkix.Name{
			CommonName: "nexus",
		},

		NotBefore: time.Now(),
		NotAfter: time.Now().Add(
			365 * 24 * time.Hour,
		),

		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature,

		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
	}

	certDER, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&key.PublicKey,
		key,
	)

	if err != nil {
		return nil, err
	}

	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type: "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(
				key,
			),
		},
	)

	certPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certDER,
		},
	)

	cert, err := tls.X509KeyPair(
		certPEM,
		keyPEM,
	)

	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
		NextProtos: []string{
			"nexus",
		},
	}, nil
}

func generateClientTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"nexus"},
	}
}
