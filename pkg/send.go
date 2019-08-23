package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
)

const (
	hostUrl = "smtp.gmail.com"
	hostPort = "587"
	emailSender = "pbobby001@st.ug.edu.gh"
	password = "yoforreal.com"
)

func SendEmail(w http.ResponseWriter, r* http.Request)  {

	emailReciever := "pbobby001@st.ug.edu.gh"

	emailAuth := smtp.PlainAuth("", emailSender, password, hostUrl)

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
		"/r/n" + "Last Name: " + req.LastName +
		"/r/n" + "Email: " + req.Email +
		"/r/n" + "Phone Number: " + req.PhoneNumber +
		"/r/n" + "Comment: " + req.Comment +
		"/r/n" + "School: " + req.School

	msg := []byte("To: " + emailReciever +
		"/r/n" + "Subject: " + "New Hostel Connect Client Request" +
		"/r/n" + details)

	err := smtp.SendMail(
		hostUrl + ":" + hostPort,
		emailAuth,
		emailSender,
		[]string{emailReciever},
		msg)

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