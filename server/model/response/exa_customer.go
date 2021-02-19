package response

import "newchat/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
