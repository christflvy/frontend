package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserData struct {
	ID   int
	Name string
	Age  int
	Flag bool
}

func PaymentPageHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")

	data := map[string]interface{}{
		"Name":        "Alvin",
		"TotalAmount": 30000,
		"ThisArray": []string{
			"Apple",
			"Orange",
		},
		"ThisObject": map[string]string{
			"Animal": "Dog",
			"Name":   "Doggy",
			"Hobby":  "Bite Bone",
		},
		"RelationFlag": false,
		"ArrayOfObjects": []UserData{
			UserData{ID: 1, Name: "Alvin", Age: 16, Flag: true},
			UserData{ID: 2, Name: "Alviani", Age: 15},
		},
		"Input": input,
	}

	err := renderEngine.ExecuteTemplate(w, "payment_page", data)
	if err != nil {
		log.Printf("Failed to Render Page because: %s", err.Error())
	}
}

func AjaxPaymentPageHandler(w http.ResponseWriter, r *http.Request) {
	stype := r.FormValue("type")

	status := 500
	errors := []string{"Ada sesuatu yang gagal"}
	if stype == "success" {
		status = 200
		errors = []string{}
	}

	data := map[string]interface{}{
		"status": status,
		"errors": errors,
		"data": map[string]interface{}{
			"Name": "Alvin",
			"Age":  16,
		},
	}

	str, _ := json.Marshal(data)
	w.Write(str)
}
