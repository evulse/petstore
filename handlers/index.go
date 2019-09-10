package handlers

import (
	"net/http"
)

type IndexHandler struct {}

func (handler *IndexHandler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type","text/html")
	w.Write([]byte(`<h3>PETS API</h3><ul>
		<li><a href="/pet/">GET /pet</a></li>
		<li><a href="/pet/1">GET /pet/1</a></li>
		<li>You can POST /pet</li>
		<li>You can PUT /pet</li>
	</ul>`))
}
