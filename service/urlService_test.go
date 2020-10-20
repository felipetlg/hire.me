package service_test

import (
	"testing"

	"github.com/felipetlg/hire.me/model"
	"github.com/felipetlg/hire.me/repository"
	"github.com/felipetlg/hire.me/service"
)

type RepositoryMock struct {
	url  *model.Url
	urls []model.Url
	err  error
}

func (rm *RepositoryMock) SetupDatabase() error {

	return rm.err
}
func (rm *RepositoryMock) InsertNewEntry(url *model.Url) error {

	return nil
}
func (rm *RepositoryMock) RetrieveLongUrl(alias string) (*model.Url, error) {

	return rm.url, rm.err
}
func (rm *RepositoryMock) UpdateUrl(url *model.Url) {

}
func (rm *RepositoryMock) MostVisited(n int) ([]model.Url, error) {

	return nil, nil
}

func TestRetrieveLongUrl_success(t *testing.T) {
	var rMock repository.Repository = &RepositoryMock{
		url: &model.Url{LongUrl: "test"},
		err: nil,
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	url, err2 := urlService.RetrieveLongUrl("alias")

	if url == nil || err2 != nil {
		t.Fatal("Expected no error retrieving longUrl")
	}
	if url.LongUrl != "test" {
		t.Fatal("LongUrl retrieved with error")
	}
}
