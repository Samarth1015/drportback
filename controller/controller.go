package controller

import (
	
	"back/model"
	
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	
	"gopkg.in/gomail.v2"
)



func SendData(w http.ResponseWriter, r *http.Request) {
  
    

    
    var data model.Connect
    err:= json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

  
    fmt.Println("Received Data:", data)

   
    pass := os.Getenv("PASSWORD")
    if pass == "" {
        fmt.Println("Environment variable PASSWORD is missing")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

   
    m := gomail.NewMessage()
    m.SetHeader("From", "testdrport@gmail.com") 
    m.SetHeader("To", "samarth.acharya2005@gmail.com") 
    m.SetHeader("Subject", "New User Submission")
    m.SetBody("text/plain", fmt.Sprintf(
        "Name: %s\nEmail: %s\nPhone: %s\nMessage: %s",
        data.Name, data.Email, data.Phone, data.Message,
    ))

   
    d := gomail.NewDialer("smtp.gmail.com", 587, "testdrport@gmail.com", pass)

  
    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Error sending email:", err)
        http.Error(w, "Failed to send email", http.StatusInternalServerError)
        return
    }

  
    json.NewEncoder(w).Encode("Email sent successfully")
}
