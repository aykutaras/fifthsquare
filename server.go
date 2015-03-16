package fifthsquare

import (
	"fmt"
	"github.com/aykutaras/gosquare"
	"github.com/aykutaras/gosquare/oauth2"
	"github.com/aykutaras/gosquare/users"
	"log"
	"net/http"
)

const (
	Uri          = "localhost:4001"
	ClientId     = "32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC"
	ClientSecret = "SUSVS34WO3ANGHYSYOSA5HOPKWB0M0E1LN3SICZKFCF1ZZIZ"
)

var api *gosquare.Api = &gosquare.Api{
	Auth:  &oauth2.Auth{ClientId: ClientId, ClientSecret: ClientSecret},
	Users: &users.Users{},
}

func SendLinkHandler(w http.ResponseWriter, r *http.Request) {
	authUri := api.Auth.Authenticate(Uri)
	fmt.Fprintf(w, "<a href='%s'>Connect to Foursquare</a>", authUri)
}

func ListenForCodeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := fmt.Sprintf("%s", query.Get("code"))

	api.Auth.GetAccessToken(Uri, code)
	api.Users.GetCheckIns(api.Auth.AccessToken)

	fmt.Fprintf(w, "%s", api.Users.CheckIns)
}

func InitHttpService() {
	http.HandleFunc("/", SendLinkHandler)
	http.HandleFunc("/code", ListenForCodeHandler)
	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
