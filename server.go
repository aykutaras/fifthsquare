package fifthsquare

import (
	"fmt"
	"github.com/aykutaras/gosquare"
	"log"
	"net/http"
	"os"
)

const (
	uri                = "http://localhost:4001"
	clientIdEnvKey     = "FOURSQUARE_CLIENTID"
	clientSecretEnvKey = "FOURSQUARE_CLIENTSECRET"
)

var api *gosquare.Api = &gosquare.Api{
	Auth:  &gosquare.Auth{ClientId: os.Getenv(clientIdEnvKey), ClientSecret: os.Getenv(clientSecretEnvKey)},
	Users: new(gosquare.Users),
}

func ConnectToFoursquare(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if code := query.Get("code"); code != "" || api.Auth.AccessToken != "" {
		code := query.Get("code")
		api.Auth.GetAccessToken(uri, code)
		fmt.Fprint(w, "<a href='profile'>Profile</a><br />")
		fmt.Fprint(w, "<a href='checkins'>CheckIns</a><br />")
		fmt.Fprint(w, "<a href='friends'>Friends</a><br />")
	} else {
		authUri := api.Auth.Authenticate(uri)
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

func InitHttpService(serverUrl string) {
	http.HandleFunc("/", ConnectToFoursquare)
	http.HandleFunc("/checkins", GetUserCheckIns)
	http.HandleFunc("/friends", GetUserFriends)
	http.HandleFunc("/profile", GetUserProfile)
	err := http.ListenAndServe(serverUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
}
