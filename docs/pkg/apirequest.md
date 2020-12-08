# apirequest
--
    import "github.com/marco-ostaska/sl1cli-tools/pkg/apirequest.go"

Package apirequest makes http request calls on sl1api

## Usage

#### type APIData

```go
type APIData struct {
	API    string // API section as : /api/account
	ARGS   string // any parameter need to be sent to api
	Result []byte // result from call
}
```

APIData an abstraction to API

#### func (*APIData) NewRequest

```go
func (a *APIData) NewRequest(v interface{}, as ...string) error
```
NewRequest make new request to sl1 API
