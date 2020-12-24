# httpcalls
--
    import "github.com/marco-ostaska/sl1cmd/pkg/httpcalls"

Package httpcalls makes http request calls on sl1api

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

#### func (*APIData) NewPost

```go
func (a *APIData) NewPost(v interface{}, as ...string) error
```
NewPost make new post to sl1 API

#### func (*APIData) NewRequest

```go
func (a *APIData) NewRequest(v interface{}, as ...string) error
```
NewRequest make new call to sl1 API
