package authentication

import (
	"golang.org/x/crypto/ssh"
)

// Authentication represents ssh auth methods.
type Authentication []ssh.AuthMethod

// Password returns password auth method.
func Password(password string) (authentication Authentication) {
	authentication = Authentication{
		ssh.Password(password),
	}

	return
}

// KeyWithPassphrase returns key with a passphrase auth method.
func KeyWithPassphrase(privateKey, passphrase string) (authentication Authentication, err error) {
	var signer ssh.Signer

	signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(privateKey), []byte(passphrase))
	if err != nil {
		return
	}

	authentication = Authentication{
		ssh.PublicKeys(signer),
	}

	return
}

// KeyWithoutPassphrase returns key without a passphrase auth method.
func KeyWithoutPassphrase(privateKey string) (authentication Authentication, err error) {
	var signer ssh.Signer

	signer, err = ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return
	}

	authentication = Authentication{
		ssh.PublicKeys(signer),
	}

	return
}
