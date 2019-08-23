package pkg

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	FirstName string 	`json:"first_name"`
	LastName string 	`json:"last_name"`
	Email string 		`json:"email"`
	PhoneNumber string	`json:"phone_number"`
	Comment string 		`json:"comment"`
	School string 		`json:"school"`
}

type Response struct {
	Status string 		`json:"status"`
	Message string 		`json:"message"`
}

type ErrorMessage struct {
	Comment string 		`json:"comment"`
	Message string 		`json:"message"`
}

func (request *Request) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return err
	}

	return nil
}