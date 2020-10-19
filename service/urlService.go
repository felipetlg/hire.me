package service

import (
	"time"

	"github.com/felipetlg/hire.me/model"
	"github.com/felipetlg/hire.me/repository"
	"github.com/speps/go-hashids"
)

const (
	baseUrl          = "http://localhost:8080/"
	topVisitQuantity = 10
)

type UrlService struct {
	Repo repository.Repository
}

func (s *UrlService) InsertNewAlias(url *model.Url) error {
	// Verifica se não enviou um alias customizado para criar uma hash
	if url.Alias == "" {
		url.Alias = createHash(url.LongUrl)
	}

	url.ShortUrl = baseUrl + url.Alias

	err := s.Repo.InsertNewEntry(url)

	return err
}

func (s *UrlService) RetrieveLongUrl(alias string) (*model.Url, error) {
	return s.Repo.RetrieveLongUrl(alias)
}

func (s *UrlService) UpdateVisitUrl(url *model.Url) {
	url.Visits += 1
	s.Repo.UpdateUrl(url)
}

func (s *UrlService) MostVisited() ([]model.Url, error) {
	return s.Repo.MostVisited(topVisitQuantity)
}

func createHash(seed string) string {
	hd := hashids.NewData()
	hd.Salt = seed
	h, err1 := hashids.NewWithData(hd)
	hash, err2 := h.Encode([]int{int(time.Now().Unix())})

	if err1 != nil || err2 != nil {
		//TODO: logError
	}

	return hash
}
