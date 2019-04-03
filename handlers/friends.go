package handlers

import (
	"net/http"

	"go.isomorphicgo.org/go/isokit"

	"github.com/JayneJacobs/gopherclub/common"
)

func FriendsHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render("friends_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}
