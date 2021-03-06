package sendtoclient

import (
	"encoding/json"
	"net/http"

	"websocket/api"
	"websocket/define/retcode"
	"websocket/servers"
)

type Controller struct {
}

type inputData struct {
	Rep string `json:"rep"`
}

//服务器图片发消息，websocket 接收
func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData servers.ReponseFromTalkEvent
	var rep servers.DResultTalkEvent

	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rep.Event = inputData.Event
	rep.Status = 200
	rep.Receivedlist = inputData.Receive_list
	rep.DataResponse.Data = inputData.Messagedata.Data
	rep.DataResponse.Data.File = inputData.Messagedata.Data.File
	rep.DataResponse.Receive_user = inputData.Messagedata.Receive_user
	rep.DataResponse.Send_user = inputData.Messagedata.Send_user
	rep.DataResponse.Source_type = inputData.Messagedata.Source_type
	servers.SendMessage(rep)
	api.Render(w, retcode.SUCCESS, "success", map[string]string{})
	return
}
