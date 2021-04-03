package getonlinelist

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
	GroupName string      `json:"groupName" validate:"required"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
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

	systemId := r.Header.Get("SystemId")
	ret := servers.GetOnlineList(&systemId, &inputData.GroupName)

	api.Render(w, retcode.SUCCESS, "success", ret)
	return
}
