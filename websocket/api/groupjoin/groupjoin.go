package groupjoin

import (
	"encoding/json"
	"net/http"

	"websocket/api"
	"websocket/define/retcode"
	"websocket/servers"
)

type Controller struct {
}

//发送入群消息，websocket 接收
func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData servers.GroupList
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	inputData.EventDo()
	api.Render(w, retcode.SUCCESS, "success", map[string]string{})
	return
}
