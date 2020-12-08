# sl1user
--
    import "github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"

Package sl1user have the routines for /api/account

## Usage

#### type UserAcct

```go
type UserAcct []struct {
        URI         string `json:"URI"`
        Description string `json:"description"`
}
```

UserAcct is an abstraction to /api/account

#### func (*UserAcct) GetIDs

```go
func (ua *UserAcct) GetIDs() error
```
GetIDs get user IDs from /api/account

#### func (*UserAcct) ID

```go
func (ua *UserAcct) ID(user string) (int, error)
```
ID returns a specific user ID index on UserAcct

#### func (*UserAcct) PrintID

```go
func (ua *UserAcct) PrintID(args []string)
```
PrintID print userId from /api/account