package sendtoclient

import (
	"encoding/json"
	"net/http"

	"github.com/woodylan/go-websocket/api"
	"github.com/woodylan/go-websocket/define/retcode"
	"github.com/woodylan/go-websocket/servers"
)

type Controller struct {
}

type inputData struct {
	Rep string `json:"rep"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData servers.ReponseFromWeb
	var rep servers.DResult
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rep.Event = inputData.Data.Event
	rep.Status = inputData.Code
	rep.Receivedlist = inputData.Data.Receive_list
	rep.DataResponse.Data = inputData.Data.Messagedata.Data
	servers.SendMessage(rep)
	api.Render(w, retcode.SUCCESS, "success", map[string]string{})
	return
}
