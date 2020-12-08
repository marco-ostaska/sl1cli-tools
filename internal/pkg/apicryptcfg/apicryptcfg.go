// December 2020
// v1.0.0

// Package apicryptcfg crypt the config file used by sl1tools.
package apicryptcfg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
)

// UserInfo got the configuration for user
type UserInfo struct {
	HomeDir    string
	CFGFile    string
	CFGDir     string
	Username   string
	CryptP     string
	DcryptP    string
	CryptJSON  string
	DcryptJSON string
	B64        string `json:"b64"`
	UserAPI    string `json:"user"`
	URL        string `json:"url"`
}

func getHash(s string) (bs []byte, err error) {
	hash := sha512.New()
	if _, err = hash.Write([]byte(s)); err != nil {
		return bs, err
	}

	bs = []byte(hex.EncodeToString(hash.Sum(nil))[:32])
	return bs, nil
}

func (u *UserInfo) getInfo() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	u.HomeDir = usr.HomeDir
	u.Username = usr.Username
}

func (u *UserInfo) setDir() error {
	u.CFGDir = u.HomeDir + "/.local/sl1api/"
	u.CFGFile = u.CFGDir + "sl1api.cfg"
	return os.MkdirAll(u.CFGDir, 0700)
}

func (u *UserInfo) newGCM() (gcm cipher.AEAD, err error) {
	hash, err := getHash(u.HomeDir + u.Username)
	if err != nil {
		return gcm, err
	}

	cBlock, err := aes.NewCipher(hash)
	if err != nil {
		return gcm, err
	}

	return cipher.NewGCM(cBlock)

}

func (u *UserInfo) encrypt(s string) (string, error) {
	data := []byte(s)
	gcm, err := u.newGCM()
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	enc := gcm.Seal(nonce, nonce, data, nil)

	//u.CryptP =
	return base64.StdEncoding.EncodeToString(enc), nil
}

func (u *UserInfo) decrypt(s string) (bs []byte, err error) {
	data := []byte(s)
	gcm, err := u.newGCM()
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

func (u *UserInfo) apiB64(user, pass string) error {
	sDec, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return err
	}

	bs, err := u.decrypt(string(sDec))
	if err != nil {
		return err
	}

	u.DcryptP = string(bs)
	up := user + ":" + u.DcryptP
	u.B64 = base64.StdEncoding.EncodeToString([]byte(up))
	return nil
}

// SetInfo set basic UserInfo to be used by sl1tools
func (u *UserInfo) SetInfo(user, passwd, url string) error {
	u.getInfo()
	u.URL = url
	u.UserAPI = user
	err := u.setDir()
	if err != nil {
		return err
	}

	enc, err := u.encrypt(passwd)
	if err != nil {
		return err
	}

	u.CryptP = enc

	err = u.apiB64(u.UserAPI, u.CryptP)
	if err != nil {
		return err
	}

	err = u.toJSON()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserInfo) toJSON() error {
	bs, err := json.Marshal(u)
	if err != nil {
		return err
	}

	f, err := os.Create(u.CFGFile)

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
	enc, err := u.encrypt(string(bs))
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(enc))
	if err != nil {
		return err
	}
	return nil
}

//ReadCryptFile read the crypt file to be used by sl1tools
func (u *UserInfo) ReadCryptFile() error {
	u.getInfo()
	u.CFGDir = u.HomeDir + "/.local/sl1api/"
	u.CFGFile = u.CFGDir + "sl1api.cfg"
	data, err := ioutil.ReadFile(u.CFGFile)
	if err != nil {
		return err
	}

	sDec, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}

	bs, err := u.decrypt(string(sDec))
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &u)
	if err != nil {
		return err
	}
	return nil

}
