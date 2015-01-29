package sessions

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"

	"upper.io/db"

	"github.com/pkieltyka/godo-app"
	"github.com/pkieltyka/godo-app/data"
	"github.com/pkieltyka/godo-app/lib/ws"
	"github.com/zenazn/goji/web"
)

var (
	ErrAuthFailed  = errors.New("authentication failed")
	ErrUnknownUser = errors.New("unknown user")
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	var username, password string
	if r.Method == "POST" {
		r.ParseForm()
		username = r.FormValue("username")
		password = r.FormValue("password")
	} else {
		username = r.URL.Query().Get("username")
		password = r.URL.Query().Get("password")
	}

	u, err := data.FindUser(db.Cond{"username": username})
	if err != nil {
		ws.Respond(w, http.StatusUnauthorized, ErrAuthFailed)
		return
	}
	if u.HashPassword(password) != u.Password {
		ws.Respond(w, http.StatusUnauthorized, ErrAuthFailed)
		return
	}

	// Build JWT token
	token, err := godo.App.TokenAuth.Encode(map[string]interface{}{"user_id": u.ID.Hex()})
	if err != nil {
		ws.Respond(w, http.StatusUnauthorized, ErrAuthFailed)
		return
	}

	// Store token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name: "jwt", Path: "/", Value: token,
	})

	// Build response
	resp := struct {
		Jwt string `json:"token"`
	}{Jwt: token}
	ws.Respond(w, 200, resp)
}

func UserContext(c *web.C, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := c.Env["token"].(*jwt.Token)
		if !ok {
			ws.Respond(w, http.StatusUnauthorized, errors.New("bad token"))
			return
		}

		userId, ok := token.Claims["user_id"].(string)
		if !ok {
			ws.Respond(w, http.StatusUnauthorized, ErrUnknownUser)
			return
		}

		user, err := data.FindUserById(bson.ObjectIdHex(userId))
		if err != nil && err == db.ErrNoMoreRows {
			ws.Respond(w, 403, ErrUnknownUser)
			return
		}
		c.Env["user"] = user
		h.ServeHTTP(w, r)
	})
}
