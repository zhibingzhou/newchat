package response

import "newchat/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
