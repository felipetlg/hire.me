package handle

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/felipetlg/hire.me/model"
	"github.com/felipetlg/hire.me/service"
	"github.com/gorilla/mux"
)

const (
	templateLocation = "template/index.html"
)

type UrlHandler struct {
	Service *service.UrlService
}

func (uh *UrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var url model.Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	err := uh.Service.InsertNewAlias(&url)

	if err != nil {
		if url.Alias != "" {
			// Envia código de erro em caso de o Alias já estar em uso
			m := `{"alias": "%s", "err_code": "001", "description": "CUSTOM ALIAS ALREADY EXISTS"}`
			send := fmt.Sprintf(m, url.Alias)

			w.Write([]byte(send))
		} else {
			http.Error(w, "error", http.StatusInternalServerError)
			log.Print(err)
		}

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

	json.NewEncoder(w).Encode(urls)
}

func (uh *UrlHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateLocation))
	tmpl.Execute(w, nil)
}
