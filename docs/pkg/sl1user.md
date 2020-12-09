# sl1user
--
    import "github.com/marco-ostaska/sl1cli-tools/pkg/sl1user"

Package sl1user have the routines for /api/account

## Usage

#### type AccessHooks

```go
type AccessHooks struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}
```

AccessHooks subset of UserDetails

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
ID returns a specific user ID index from UserAcct

#### func (*UserAcct) PrintID

```go
func (ua *UserAcct) PrintID(args []string)
```
PrintID print userId from /api/account

#### func (*UserAcct) SearchByURI

```go
func (ua *UserAcct) SearchByURI(uri string) (string, error)
```
SearchByURI search user name by provided URI

#### func (*UserAcct) Sl1UserID

```go
func (ua *UserAcct) Sl1UserID(user string) (string, error)
```
Sl1UserID returns sl1id from user

#### type UserDetails

```go
type UserDetails struct {
	Organization         string        `json:"organization"`
	User                 string        `json:"user"`
	Email                string        `json:"email"`
	PasswdExpiration     string        `json:"passwd_expiration"`
	PasswdSetDate        string        `json:"passwd_set_date"`
	PasswdPrevCount      string        `json:"passwd_prev_count"`
	PasswdResetRequired  string        `json:"passwd_reset_required"`
	PasswdStrength       string        `json:"passwd_strength"`
	LoginState           string        `json:"login_state"`
	RestrictIP           string        `json:"restrict_ip"`
	Admin                string        `json:"admin"`
	Active               string        `json:"active"`
	CreateDate           string        `json:"create_date"`
	CreatedBy            string        `json:"created_by"`
	EditDate             string        `json:"edit_date"`
	UpdatedBy            string        `json:"updated_by"`
	UserPolicy           string        `json:"user_policy"`
	Country              string        `json:"country"`
	ContactFname         string        `json:"contact_fname"`
	ContactLname         string        `json:"contact_lname"`
	AlignedOrganizations []string      `json:"aligned_organizations"`
	AlignedTicketQueues  []interface{} `json:"aligned_ticket_queues"`
	AccessHooks          AccessHooks   `json:"access_hooks"`
}
```

UserDetails is an abstraction of /api/account/x

#### func (*UserDetails) LoadUserDetails

```go
func (ud *UserDetails) LoadUserDetails(id string) error
```
LoadUserDetails get user details from /api/account/x and load it to *UserDetails

#### func (*UserDetails) PrintUserDetails

```go
func (ud *UserDetails) PrintUserDetails()
```
PrintUserDetails prints user details from /api/account/x
