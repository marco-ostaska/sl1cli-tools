package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/marco-ostaska/sl1cli-tools/internal/pkg/apirequest.go"
)

const version string = "v1.0.0 dec-2020"

type sl1UserLst []struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

type sl1UserDetails struct {
	Organization           string        `json:"organization"`
	User                   string        `json:"user"`
	Email                  string        `json:"email"`
	PasswdExpiration       string        `json:"passwd_expiration"`
	PasswdSetDate          string        `json:"passwd_set_date"`
	PasswdPrevCount        string        `json:"passwd_prev_count"`
	PasswdResetRequired    string        `json:"passwd_reset_required"`
	PasswdStrength         string        `json:"passwd_strength"`
	LoginState             string        `json:"login_state"`
	RestrictIP             string        `json:"restrict_ip"`
	Admin                  string        `json:"admin"`
	Active                 string        `json:"active"`
	CreateDate             string        `json:"create_date"`
	CreatedBy              string        `json:"created_by"`
	EditDate               string        `json:"edit_date"`
	UpdatedBy              string        `json:"updated_by"`
	Timezone               string        `json:"timezone"`
	AutodetectTimezone     string        `json:"autodetect_timezone"`
	RecentTimezone         string        `json:"recent_timezone"`
	DefaultMap             string        `json:"default_map"`
	DefaultMapType         string        `json:"default_map_type"`
	Theme                  string        `json:"theme"`
	Refresh                string        `json:"refresh"`
	Barred                 string        `json:"barred"`
	PageResults            string        `json:"page_results"`
	EventSeverity          string        `json:"event_severity"`
	UserPolicy             string        `json:"user_policy"`
	Ldap                   string        `json:"ldap"`
	ConsoleHeight          string        `json:"console_height"`
	DateFormat             string        `json:"date_format"`
	IflabelPref            string        `json:"iflabel_pref"`
	AllOrgs                string        `json:"all_orgs"`
	TicketNoteSort         string        `json:"ticket_note_sort"`
	Codehighlight          string        `json:"codehighlight"`
	EventsDefaultView      string        `json:"events_default_view"`
	IfGraphPerc            string        `json:"if_graph_perc"`
	DevHTML5Graph          string        `json:"dev_html5_graph"`
	TableRowHeight         string        `json:"table_row_height"`
	ShowSeverityBadges     string        `json:"show_severity_badges"`
	EventsBeta             string        `json:"events_beta"`
	MfUser                 string        `json:"mf_user"`
	IsEmailSelect          string        `json:"is_email_select"`
	Address                string        `json:"address"`
	City                   string        `json:"city"`
	State                  string        `json:"state"`
	Zip                    string        `json:"zip"`
	Country                string        `json:"country"`
	ContactFname           string        `json:"contact_fname"`
	ContactLname           string        `json:"contact_lname"`
	Title                  string        `json:"title"`
	Dept                   string        `json:"dept"`
	Office                 string        `json:"office"`
	BillingID              string        `json:"billing_id"`
	CrmID                  string        `json:"crm_id"`
	Phone                  string        `json:"phone"`
	Fax                    string        `json:"fax"`
	Tollfree               string        `json:"tollfree"`
	Email2                 string        `json:"email_2"`
	Email3                 string        `json:"email_3"`
	Pager                  string        `json:"pager"`
	Cell                   string        `json:"cell"`
	Im                     string        `json:"im"`
	ImType                 string        `json:"im_type"`
	Role                   string        `json:"role"`
	Critical               string        `json:"critical"`
	Notes                  string        `json:"notes"`
	VerificationQuestion   string        `json:"verification_question"`
	VerificationAnswer     string        `json:"verification_answer"`
	EventsShowMasked       string        `json:"events_show_masked"`
	EventsCollapseOrgs     string        `json:"events_collapse_orgs"`
	GraphScaleTo100        string        `json:"graph_scale_to_100"`
	NavbarAutohideDisabled string        `json:"navbar_autohide_disabled"`
	NetworkHideEmpty       string        `json:"network_hide_empty"`
	TicketCommentCloaking  string        `json:"ticket_comment_cloaking"`
	TicketViewAssignedOnly string        `json:"ticket_view_assigned_only"`
	PermissionKeys         []interface{} `json:"permission_keys"`
	AlignedOrganizations   []string      `json:"aligned_organizations"`
	AlignedTicketQueues    []interface{} `json:"aligned_ticket_queues"`
	AccessHooks            AccessHooks   `json:"access_hooks"`
}
type AccessHooks struct {
	URI         string `json:"URI"`
	Description string `json:"description"`
}

func usage() {
	usage := `Print sl1 user information for the specified USERS,
or (when USER omitted) prints a list of all users.

  -v             version
  -h             display this and exit
`
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [USERS]\n%v", filepath.Base(os.Args[0]), usage)
}

func chkArgs() bool {
	if len(os.Args) == 1 {
		return false
	}

	switch os.Args[1] {
	case "-v":
		fmt.Printf("%s: %s\n", filepath.Base(os.Args[0]), version)
		return false
	case "-h":
		usage()
		return false
	default:
		return true
	}

}

func listIDs() {
	var api apirequest.APIData
	api.API = "/api/account/"

	if err := api.APIRequest(); err != nil {
		fmt.Println(err)
	}

	var dat sl1UserLst

	if err := json.Unmarshal(api.Result, &dat); err != nil {
		fmt.Println(err)
	}

	for _, u := range dat {
		if u.Description == os.Args[1] {
			getDetails(filepath.Base(u.URI))
		}
	}
	return

}

// need to create a func to avoid copy and paste
func getDetails(id string) error {
	var api apirequest.APIData
	api.API = "/api/account/" + id

	if err := api.APIRequest(); err != nil {
		return err
	}

	var dat sl1UserDetails

	if err := json.Unmarshal(api.Result, &dat); err != nil {
		return err
	}

	fmt.Println("User                :", dat.User)
	fmt.Println("Email               :", dat.Email)
	fmt.Println("Organization        :", path.Base(dat.Organization))
	fmt.Println("PasswdExpiration    :", path.Base(dat.PasswdExpiration))

	i, _ := strconv.ParseInt(dat.PasswdSetDate, 10, 64)
	t := time.Unix(i, 0)
	fmt.Println("PasswdSetDate       :", t)

	// PasswdPrevCount     : 5
	// PasswdResetRequired : 0
	// PasswdStrength      : 0
	// LoginState          : 1
	// RestrictIP          :
	// Admin               : 0
	// Active              : 0
	// CreateDate          : 1603145206
	// CreatedBy           : /api/account/4
	// EditDate            : 1607101184
	// UpdatedBy           : /api/account/2
	// UserPolicy          : /api/account_policy/11
	// AllOrgs             : 1
	// ContactFname        : Marco Testetetete
	// ContactLname        : Aurelio Najar Ostaka
	// AlignedOrganizations: [/api/organization/0 /api/organization/2 /api/organization/3 /api/organization/4]

	return nil
}

func main() {

	if chkArgs() {
		listIDs()
	}

}
