# vault
--
    import "github.com/marco-ostaska/sl1cmd/pkg/sl1/vault"

Package vault manage encryption for sl1cmd credentials

## Usage

#### type Credential

```go
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
```

Credential is an abstraction to credential vault

#### func (*Credential) ReadFile

```go
func (c *Credential) ReadFile() error
```
ReadFile reads the credential vault and unmarshal it.

#### func (*Credential) SetInfo

```go
func (c *Credential) SetInfo(user, passwd, url string) error
```
SetInfo set provided information to credential vault
