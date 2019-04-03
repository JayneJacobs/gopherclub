package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JayneJacobs/gopherclub/common/authenticate"

	"github.com/JayneJacobs/gopherclub/common"
)

func FetchPostsEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherclub-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		posts, err := env.DB.FetchPosts(uuid)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)

	})
}
