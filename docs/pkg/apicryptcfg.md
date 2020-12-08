# apicryptcfg
--
    import "github.com/marco-ostaska/sl1cli-tools/internal/pkg/apicryptcfg"

Package apicryptcfg crypt the config file used by sl1tools.

## Usage

#### type UserInfo

```go
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
```

UserInfo got the configuration for user

#### func (*UserInfo) ReadCryptFile

```go
func (u *UserInfo) ReadCryptFile() error
```
ReadCryptFile read the crypt file to be used by sl1tools

#### func (*UserInfo) SetInfo

```go
func (u *UserInfo) SetInfo(user, passwd, url string) error
```
SetInfo set basic UserInfo to be used by sl1tools
