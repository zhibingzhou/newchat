package response

import "newchat/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
