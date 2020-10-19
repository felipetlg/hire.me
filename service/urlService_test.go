package service_test

import (
	"testing"

	"github.com/felipetlg/hire.me/model"
	"github.com/felipetlg/hire.me/repository"
	"github.com/felipetlg/hire.me/service"
)

type RepositoryMock struct{}

func (rm *RepositoryMock) SetupDatabase() error {

	return nil
}
func (rm *RepositoryMock) InsertNewEntry(url *model.Url) error {

	return nil
}
func (rm *RepositoryMock) RetrieveLongUrl(alias string) (*model.Url, error) {

	return nil, nil
}
func (rm *RepositoryMock) UpdateUrl(url *model.Url) {

}
func (rm *RepositoryMock) MostVisited(n int) ([]model.Url, error) {

	return nil, nil
}

func TestRetrieveLongUrl(t *testing.T) {
	var rMock repository.Repository = &RepositoryMock{}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	err1, err2 := urlService.RetrieveLongUrl("alias")

	if err1 != nil || err2 != nil {
		t.Errorf("Expected no error retrieving longUrl")
	}
}
