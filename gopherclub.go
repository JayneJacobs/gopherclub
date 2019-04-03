package main

import (
	"log"
	"os"

	"github.com/JayneJacobs/gopherclub/common"
	"github.com/JayneJacobs/gopherclub/common/datastore"
	"github.com/JayneJacobs/gopherclub/endpoints"
	"github.com/JayneJacobs/gopherclub/handlers"
	"github.com/JayneJacobs/gopherclub/middleware"

	"go.isomorphicgo.org/go/isokit"

	"net/http"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

const (
	//	WEBSERVERPORT = ":8443"
	WEBSERVERPORT = ":8081"
)

var WebAppRoot = os.Getenv("GOPHERCLUB_APP_ROOT")

func main() {

	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherclub:gopherclub@/gopherclubdb")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{}
	isokit.TemplateFilesPath = WebAppRoot + "/templates"
	isokit.TemplateFileExtension = ".html"
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
	env.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "POST")

	r.Handle("/feed", middleware.GatedContentHandler(handlers.FeedHandler(&env))).Methods("GET")
	r.Handle("/friends", middleware.GatedContentHandler(handlers.FriendsHandler(&env))).Methods("GET")
	r.Handle("/myprofile", middleware.GatedContentHandler(handlers.MyProfileHandler(&env))).Methods("GET")
	r.Handle("/profile/{username}", middleware.GatedContentHandler(handlers.ProfileHandler(&env))).Methods("GET")

	// Register REST API Endpoints
	r.Handle("/restapi/get-user-profile", middleware.GatedContentHandler(endpoints.GetUserProfileEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/save-user-profile", middleware.GatedContentHandler(endpoints.SaveUserProfileEndpoint(&env))).Methods("POST")
	r.Handle("/restapi/save-user-profile-image", middleware.GatedContentHandler(endpoints.SaveUserProfileImageEndpoint(&env))).Methods("POST")
	r.Handle("/restapi/find-gophers", middleware.GatedContentHandler(endpoints.FindGophersEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/follow-gopher", middleware.GatedContentHandler(endpoints.FollowGopherEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/unfollow-gopher", middleware.GatedContentHandler(endpoints.UnfollowGopherEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/get-friends-list", middleware.GatedContentHandler(endpoints.FriendsListEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/save-post", middleware.GatedContentHandler(endpoints.SavePostEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/fetch-posts", middleware.GatedContentHandler(endpoints.FetchPostsEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/get-gopher-profile", middleware.GatedContentHandler(endpoints.GetGopherProfileEndpoint(&env))).Methods("GET", "POST")

	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot))
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(WebAppRoot+"/static"))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	//	err = http.ListenAndServeTLS(WEBSERVERPORT, WebAppRoot+"/certs/gopherclubcert.pem", WebAppRoot+"/certs/gopherclubkey.pem", nil)
	//	http.Handle("/", r)
	err = http.ListenAndServe(WEBSERVERPORT, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
