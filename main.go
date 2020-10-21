package main

import (
	"log"
	"net/http"

	"github.com/felipetlg/hire.me/service"

	"github.com/felipetlg/hire.me/repository"

	"github.com/felipetlg/hire.me/handle"

	"github.com/gorilla/mux"
)

const staticFilesLocation = "/template/static/"

func main() {
	var repo repository.Repository = &repository.UrlRepository{}
	err := repo.SetupDatabase()
	if err != nil {
		log.Print(err)
	}

	//TODO captar sinais do sistema para notificações de kill...

	urlService := &service.UrlService{
		Repo: repo,
	}
	urlHandler := &handle.UrlHandler{
		Service: urlService,
	}

	router := mux.NewRouter()
	router.PathPrefix(staticFilesLocation).Handler(http.StripPrefix(staticFilesLocation, http.FileServer(http.Dir("."+staticFilesLocation))))
	router.HandleFunc("/", urlHandler.Index).Methods("GET")
	router.HandleFunc("/", urlHandler.CreateShortUrl).Methods("POST")
	router.HandleFunc("/s/{alias}", urlHandler.RedirectToLongUrl).Methods("GET")
	router.HandleFunc("/top", urlHandler.GetMostVisited).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
