package router

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/gorilla/mux"
)

func ViewCookies(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cookies Cookies
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Cookie data.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Cookies
	//
	// responses:
	//   '200':
	//     description: Cookie data.

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "cookies"))
}

func ViewSetCookies(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cookies/set Cookies
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Sets cookie data and redirects to /cookies.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Cookies
	//
	// parameters:
	// - in: query
	//   name: whoami
	//   required: true
	//   type: string
	//
	// responses:
	//   '302':
	//     description: Sets cookie data and redirects to /cookies.

	for key, vals := range r.URL.Query() {
		c := http.Cookie{
			Name:  key,
			Value: vals[0],
			Path:  "/",
		}
		http.SetCookie(w, &c)
	}
	http.Redirect(w, r, "/cookies", http.StatusFound)
}

func ViewSetCookie(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cookies/set/{name}/{value} Cookies
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Sets cookie data and redirects to /cookies.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Cookies
	//
	// parameters:
	// - in: path
	//   name: name
	//   required: true
	//   type: string
	//
	// - in: path
	//   name: value
	//   required: true
	//   type: string
	//
	// responses:
	//   '302':
	//     description: Sets cookie data and redirects to /cookies.

	v := mux.Vars(r)
	http.SetCookie(w, &http.Cookie{
		Name:  v["name"],
		Value: v["value"],
		Path:  "/",
	})
	http.Redirect(w, r, "/cookies", http.StatusFound)
}

func ViewDeleteCookies(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cookies/delete Cookies
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Deletes cookies specified in query args and redirects to /cookies.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Cookies
	//
	// parameters:
	// - in: query
	//   name: names
	//   required: true
	//   type: string
	//
	// responses:
	//   '302':
	//     description: Deletes cookie data and redirects to /cookies.

	names, ok := r.URL.Query()["names"]
	if !ok {
		http.Redirect(w, r, "/cookies", http.StatusFound)
	}

	keys := strings.Split(names[0], ",")

	for _, name := range keys {
		http.SetCookie(w, &http.Cookie{
			Name:    name,
			Value:   helpers.RandomStr(6),
			Expires: time.Now(),
			Path:    "/",
		})
	}
	http.Redirect(w, r, "/cookies", http.StatusFound)
}
