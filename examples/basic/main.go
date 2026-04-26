package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gowrisankar-GO/go-mini-di/di"
)

type DB struct {
	Name string
}

type Service struct {
	DB      DB
	Port    Port
	Handler http.Handler
}

type Port struct {
	PortNumber string
}

func NewDB() DB {
	return DB{Name: "PostgreSQL"}
}

func NewPort() Port {
	return Port{PortNumber: "7070"}
}

func NewService(db DB, port Port, handler http.Handler) Service {
	return Service{
		DB:      db,
		Port:    port,
		Handler: handler,
	}
}

func GetRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	return mux
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Health string `json:"health"`
	}{Health: "healthy"}

	jbytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error in marshalling the health response: %v\n", err)
		return
	}

	w.Write(jbytes)
}

func (s Service) Run() {
	log.Printf("Service starting in port <%s>\n", s.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", s.Port.PortNumber), s.Handler)
}

func StartService(s Service) {
	s.Run()
}

func main() {
	c := di.New()

	c.Provide(NewPort)

	c.Provide(NewDB)

	c.Provide(GetRoutes)

	c.Provide(NewService)

	c.Invoke(StartService)
}
