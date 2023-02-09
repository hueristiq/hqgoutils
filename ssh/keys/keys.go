package keys

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

// Read
func Read(keyPair string) (string, string, error) {
	_, err := os.Stat(keyPair)
	if err != nil {
		return "", "", err
	}

	priv, err := ioutil.ReadFile(keyPair)
	if err != nil {
		return "", "", err
	}

	pub, err := ioutil.ReadFile(keyPair + ".pub")
	if err != nil {
		return "", "", err
	}

	return string(pub), string(priv), nil
}

// Generate
func Generate() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	var private bytes.Buffer

	if err := pem.Encode(&private, privateKeyPEM); err != nil {
		return "", "", err
	}

	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	public := ssh.MarshalAuthorizedKey(pub)

	return string(public), private.String(), nil
}

// ReadOrGenerate
func ReadOrGenerate(keyPairName string) (string, string, error) {
	pub, priv, err := Read(keyPairName)
	if err != nil {
		goto GENERATE
	} else {
		return string(pub), string(priv), nil
	}

GENERATE:
	pub, priv, err = Generate()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate keys - %s", err)
	}

	if err = Write(keyPairName, pub, priv); err != nil {
		return "", "", fmt.Errorf("failed to write file - %s", err)
	}

	return pub, priv, nil
}

// Write
func Write(keyPairName, pub, priv string) (err error) {
	directory := filepath.Dir(keyPairName)

	if _, err = os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if directory != "" {
				if err = os.MkdirAll(directory, os.ModePerm); err != nil {
					return
				}
			}
		} else {
			return
		}
	}

	if err = ioutil.WriteFile(keyPairName, []byte(priv), 0600); err != nil {
		return
	}

	if err = ioutil.WriteFile(keyPairName+".pub", []byte(pub), 0644); err != nil {
		return
	}

	return
}
