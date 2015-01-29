package ws

import (
	"net/http"

	"github.com/unrolled/render"
)

var (
	Render = render.New()
)

func JSON(w http.ResponseWriter, status int, v interface{}) {
	Render.JSON(w, status, v)
}

func Text(w http.ResponseWriter, status int, v string) {
	w.Header().Set("Content-Type", "text/plain")
	Data(w, status, []byte(v))
}

func Data(w http.ResponseWriter, status int, v []byte) {
	Render.Data(w, status, v)
}

func Respond(w http.ResponseWriter, status int, v interface{}) {
	switch resp := v.(type) {
	case error:
		JSON(w, status, map[string]interface{}{"error": resp.Error()})
		return
	case []byte:
		Data(w, status, resp)
		return
	default:
		JSON(w, status, v)
	}
}
