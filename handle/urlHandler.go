package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/felipetlg/hire.me/model"
	"github.com/felipetlg/hire.me/service"
	"github.com/gorilla/mux"
)

type UrlHandler struct {
	Service *service.UrlService
}

func (uh *UrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var url model.Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	err := uh.Service.InsertNewAlias(&url)
	fmt.Print(url)

	if err != nil {
		// TODO diferenciar erro de "problema" de erro de inserção

		http.Error(w, "error", http.StatusInternalServerError)
		log.Print("error")
		return
	}

	// url.shortUrl foi inserida na variável url
	json.NewEncoder(w).Encode(url)
}

func (uh *UrlHandler) RedirectToLongUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var alias string = vars["alias"]

	url, err := uh.Service.RetrieveLongUrl(alias)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Print(err)
		return
	}

	defer uh.Service.UpdateVisitUrl(url)
	w.Header().Set("Cache-Control", "no-cache")
	http.Redirect(w, r, url.LongUrl, http.StatusMovedPermanently)
}

func (uh *UrlHandler) GetMostVisited(w http.ResponseWriter, r *http.Request) {
	urls, err := uh.Service.MostVisited()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Print(err)
		return
	}
	// TODO pegar slice de string e retornar json
	json.NewEncoder(w).Encode(urls)
}
