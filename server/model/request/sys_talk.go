package request

type RequestNotDistraub struct {
	Not_disturb int `json:"not_disturb"`
	RequestCreate
}

type RequestCreate struct {
	Receive_id string `json:"receive_id"`
	Type       int    `json:"type"`
}

type RequestUpdateNoread struct {
	Receive string `json:"receive"`
	Type    string `json:"type"`
}
