package pkg

import (
	"encoding/json"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
)

const (
	hostUrl     = "smtp.gmail.com"
	hostPort    = 587
	emailSender = "pbobby001@st.ug.edu.gh"
	password    = "yoforreal.com"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {

	emailReciever := "pbobby001@st.ug.edu.gh"

	req := &Request{}

	if r.Method == http.MethodPost {
		err := req.FromJson(r)
		if err != nil {
			errmess := ErrorMessage{
				Comment: "Request object is empty",
				Message: err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			err = json.NewEncoder(w).Encode(errmess)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}

	details := "First Name: " + req.FirstName +
		" Last Name: " + req.LastName +
		" Email: " + req.Email +
		" Phone Number: " + req.PhoneNumber +
		" Comment: " + req.Comment +
		" School: " + req.School

	mail := gomail.NewMessage()
	mail.SetHeader("From", emailSender)
	mail.SetHeader("To", emailReciever)
	mail.SetHeader("Subject", "New Hostel Connect Client Request")
	mail.SetBody("text/plain", details)

	dialer := gomail.NewDialer(hostUrl, hostPort, emailSender, password)

	err := dialer.DialAndSend(mail)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Email sent.")

	response := &Response{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
