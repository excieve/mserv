package http_funcs

import (
	"encoding/json"
	"github.com/TykTechnologies/mserv/api"
	"io/ioutil"
	"net/http"

	"github.com/TykTechnologies/mserv/coprocess/bindings/go"
	"golang.org/x/net/context"
	"time"
)

// Test endpoints
func (h *HttpServ) Execute(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.HandleError(err, w, r)
		return
	}

	cp := &coprocess.Object{}
	err = json.Unmarshal(body, cp)
	if err != nil {
		h.HandleError(err, w, r)
		return
	}

	d := api.Dispatcher{}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	retObj, err := d.Dispatch(ctx, cp)
	if err != nil {
		h.HandleError(err, w, r)
		return
	}

	h.HandleOK(retObj, w, r)
}
