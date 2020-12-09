# apipost
--
    import "github.com/marco-ostaska/sl1cli-tools/pkg/apipost.go"

Package apipost post into sl1api

## Usage

#### type APIData

```go
type APIData struct {
	API     string          // API section as : /api/account
	ARGS    string          // any parameter need to be sent to api
	Payload *strings.Reader //payload for posting
	Result  []byte          // result from call
}
```

APIData an abstraction to API

#### func (*APIData) APIPost

```go
func (a *APIData) APIPost() error
```
APIPost make the http calls

#### func (*APIData) NewPost

```go
func (a *APIData) NewPost(v interface{}, as ...string) error
```
NewPost make new post to sl1 API
