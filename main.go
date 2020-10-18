package main

import (
	"log"
	"net/http"

	"github.com/felipetlg/hire.me/service"

	"github.com/felipetlg/hire.me/repository"

	"github.com/felipetlg/hire.me/handle"

	"github.com/gorilla/mux"
)

func main() {
	repository := &repository.Repository{}
	repository.SetupDatabase()

	//TODO captar sinais do sistema para notificações de kill...

	urlService := &service.UrlService{
		Repo: repository,
	}
	urlHandler := &handle.UrlHandler{
		Service: urlService,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", urlHandler.CreateShortUrl).Methods("POST")
	router.HandleFunc("/get/{alias}", urlHandler.RedirectToLongUrl).Methods("GET")
	router.HandleFunc("/topvisit", urlHandler.GetMostVisited).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
