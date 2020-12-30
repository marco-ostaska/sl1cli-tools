/*
Copyright © 2020 Marco Ostaska

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnc.org/licenses/>.
*/

// Package vault manage encryption for sl1cmd credentials
package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const vaultFile string = "/.local/sl1cmd/sl1cmd.cfg"

// Credential is an abstraction to credential vault
type Credential struct {
	HomeDir    string // local user home directory
	Username   string // local user name
	Hostname   string // local hostname
	File       string // credential vault file full path
	CryptP     string // encrypted API password
	DcryptP    string // decrypted API password
	CryptJSON  string
	DcryptJSON string
	B64        string `json:"b64"`  // base64 mask to be used by API calls
	UserAPI    string `json:"user"` // API username
	URL        string `json:"url"`  // API URL
}

func getHash(s string) (bs []byte, err error) {
	hash := sha512.New()
	if _, err = hash.Write([]byte(s)); err != nil {
		return bs, err
	}

	bs = []byte(hex.EncodeToString(hash.Sum(nil))[:32])
	return bs, nil
}

// UserInfo parse some local information to vault.Credential
func (c *Credential) UserInfo() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	c.HomeDir = usr.HomeDir
	c.Username = usr.Username
	c.Hostname = hostname
	c.File = usr.HomeDir + vaultFile

	return nil
}

func (c *Credential) setDir() error {
	return os.MkdirAll(path.Dir(c.File), 0600)
}

func (c *Credential) newGCM() (gcm cipher.AEAD, err error) {
	hash, err := getHash(c.HomeDir + c.Username + c.Hostname)
	if err != nil {
		return gcm, err
	}

	cBlock, err := aes.NewCipher(hash)
	if err != nil {
		return gcm, err
	}

	return cipher.NewGCM(cBlock)

}

func (c *Credential) encrypt(s string) (string, error) {
	data := []byte(s)
	gcm, err := c.newGCM()
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	enc := gcm.Seal(nonce, nonce, data, nil)

	return base64.StdEncoding.EncodeToString(enc), nil
}

func (c *Credential) decrypt(s string) (bs []byte, err error) {
	data := []byte(s)
	gcm, err := c.newGCM()
	if err != nil {
		return bs, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	ptxt, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return bs, err
	}

	return ptxt, err

}

func (c *Credential) apiB64(user, pass string) error {
	sDec, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return err
	}

	bs, err := c.decrypt(string(sDec))
	if err != nil {
		return err
	}

	c.DcryptP = string(bs)
	up := user + ":" + c.DcryptP
	c.B64 = base64.StdEncoding.EncodeToString([]byte(up))
	return nil
}

// SetInfo set provided information to credential vault
func (c *Credential) SetInfo(user, passwd, url string) error {
	if err := c.UserInfo(); err != nil {
		return err
	}
	c.URL = url
	c.UserAPI = user
	err := c.setDir()
	if err != nil {
		return err
	}

	enc, err := c.encrypt(passwd)
	if err != nil {
		return err
	}

	c.CryptP = enc

	err = c.apiB64(c.UserAPI, c.CryptP)
	if err != nil {
		return err
	}

	err = c.toJSON()
	if err != nil {
		return err
	}

	return nil
}

func (c *Credential) toJSON() error {
	bs, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(c.File)

	if err != nil {
		return err
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	if err != nil {
		return err
	}
	enc, err := c.encrypt(string(bs))
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(enc))
	if err != nil {
		return err
	}

	if err := os.Chmod(c.File, 0600); err != nil {
		return err
	}
	return nil
}

// ReadFile reads the credential vault and unmarshal it.
func (c *Credential) ReadFile() error {
	if err := c.UserInfo(); err != nil {
		return err
	}
	data, err := ioutil.ReadFile(c.File)
	if err != nil {
		return err
	}

	sDec, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}

	bs, err := c.decrypt(string(sDec))
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &c)
	if err != nil {
		return err
	}
	return nil

}
