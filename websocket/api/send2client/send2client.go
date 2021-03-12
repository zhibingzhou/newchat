package send2client

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
	ClientId   string `json:"clientId" validate:"required"`
	SendUserId string `json:"sendUserId"`
	Code       int    `json:"code"`
	Event      string `json:"event"`
	Data       string `json:"data"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData inputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := api.Validate(inputData)
	if err != nil {
		api.Render(w, retcode.FAIL, err.Error(), []string{})
		return
	}

	//发送信息
	
	messageId := servers.SendMessage2Client(inputData.ClientId, inputData.SendUserId, inputData.Code, inputData.Event, &inputData.Data)
	api.Render(w, retcode.SUCCESS, "success", map[string]string{
		"messageId": messageId,
	})
	return
}
