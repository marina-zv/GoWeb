package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GreetingRequest struct {
	FirstName string //`json: "first_name"`
	LastName  string //`json: "last_name"`
}

func main() {

	//Server
	rt := chi.NewRouter()

	//Endpoints

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`pong`))
	})

	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {

		// Decodificar el cuerpo JSON de la solicitud en la estructura GreetingRequest
		var greetingReq GreetingRequest
		err := json.NewDecoder(r.Body).Decode(&greetingReq)
		if err != nil {
			http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
			return
		}

		// Construir el mensaje de saludo
		greeting := fmt.Sprintf("Hello %s %s", greetingReq.FirstName, greetingReq.LastName)

		// Responder con el mensaje de saludo
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(greeting))

	})
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}

}