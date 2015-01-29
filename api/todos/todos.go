package todos

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"upper.io/db"

	"github.com/pkieltyka/godo-app/data"
	"github.com/pkieltyka/godo-app/lib/ws"
	"github.com/pressly/cji"

	"github.com/zenazn/goji/web"
)

var (
	Ctx = cji.Use(TodoContext)
)

func TodoContext(c *web.C, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oid := bson.ObjectIdHex(c.URLParams["id"])
		todo, err := data.FindTodoById(oid)
		if err != nil {
			ws.Respond(w, 404, errors.New("not found"))
			return
		}
		c.Env["todo"] = todo
		h.ServeHTTP(w, r)
	})
}

// List the Todos
func Index(c web.C, w http.ResponseWriter, r *http.Request) {
	todos, err := data.FindAllTodos()
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}
	ws.Respond(w, 200, todos)
}

// Save a new Todo
func Create(c web.C, w http.ResponseWriter, r *http.Request) {
	user := c.Env["user"].(*data.User)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}

	body := struct {
		data.Todo

		// omit fields..
		UserId interface{} `json:"user_id"`
	}{}

	err = json.Unmarshal(b, &body)
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}
	todo := &body.Todo
	todo.UserID = user.ID

	oid, err := data.TodoCollection.Append(todo)
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}
	todo.ID = oid.(bson.ObjectId)
	ws.Respond(w, http.StatusCreated, todo)
}

// Fetch a Todo by id
func Read(c web.C, w http.ResponseWriter, r *http.Request) {
	todo := c.Env["todo"].(*data.Todo)
	ws.Respond(w, 200, todo)
}

// Update a Todo
func Update(c web.C, w http.ResponseWriter, r *http.Request) {
	todo := c.Env["todo"].(*data.Todo)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}

	body := struct {
		*data.Todo

		// omit fields..
		UserId interface{} `json:"user_id"`
	}{Todo: todo}

	err = json.Unmarshal(b, &body)
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}
	todo = body.Todo

	res := data.TodoCollection.Find(db.Cond{"_id": todo.ID})
	err = res.Update(todo)
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}

	ws.Respond(w, 200, todo)
}

// Delete a Todo
func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	todo := c.Env["todo"].(*data.Todo)

	res := data.TodoCollection.Find(db.Cond{"_id": todo.ID})
	err := res.Remove()
	if err != nil {
		ws.Respond(w, 400, err)
		return
	}
	ws.Respond(w, 204, []byte{})
}
