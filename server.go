package fifthsquare

import (
	"fmt"
	"github.com/aykutaras/gosquare"
	"log"
	"net/http"
)

const (
	Uri          = "http://localhost:4001"
	ClientId     = ""
	ClientSecret = ""
)

var api *gosquare.Api = &gosquare.Api{
	Auth:  &gosquare.Auth{ClientId: ClientId, ClientSecret: ClientSecret},
	Users: new(gosquare.Users),
}

func ConnectToFoursquare(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if code := query.Get("code"); code != "" || api.Auth.AccessToken != "" {
		code := query.Get("code")
		api.Auth.GetAccessToken(Uri, code)
		fmt.Fprint(w, "<a href='profile'>Profile</a><br />")
		fmt.Fprint(w, "<a href='checkins'>CheckIns</a><br />")
		fmt.Fprint(w, "<a href='friends'>Friends</a><br />")
	} else {
		authUri := api.Auth.Authenticate(Uri)
		fmt.Fprintf(w, "<a href='%s'>Connect to Foursquare</a>", authUri)
	}
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	api.Users.Profile.Get(api.Auth.AccessToken)
	fmt.Fprintf(w, "%s", api.Users.CheckIns)
}

func GetUserCheckIns(w http.ResponseWriter, r *http.Request) {
	api.Users.CheckIns.Get(api.Auth.AccessToken)
	fmt.Fprintf(w, "%s", api.Users.CheckIns)
}

func GetUserFriends(w http.ResponseWriter, r *http.Request) {
	api.Users.Friends.Get(api.Auth.AccessToken)
	fmt.Fprintf(w, "%s", api.Users.Friends)
}

func InitHttpService() {
	http.HandleFunc("/", ConnectToFoursquare)
	http.HandleFunc("/checkins", GetUserCheckIns)
	http.HandleFunc("/friends", GetUserFriends)
	http.HandleFunc("/profile", GetUserProfile)
	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
