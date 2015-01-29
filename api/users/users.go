package users

import (
	"net/http"

	"github.com/pkieltyka/godo-app/data"
	"github.com/pkieltyka/godo-app/lib/ws"
	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2/bson"
)

func SignUp(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	// TODO: implement some rate-limiting for signup

	u := data.NewUser()
	u.Username = username
	u.SetPassword(password)

	err := u.Validate()
	if err != nil {
		ws.Respond(w, http.StatusBadRequest, err)
		return
	}

	oid, err := data.UserCollection.Append(u)
	if err != nil {
		ws.Respond(w, http.StatusBadRequest, err)
		return
	}
	u.ID = oid.(bson.ObjectId)

	ws.Respond(w, http.StatusCreated, u)
}
