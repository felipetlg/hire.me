package service_test

import (
	"errors"
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
	url = rm.url
	return rm.err
}
func (rm *RepositoryMock) RetrieveLongUrl(alias string) (*model.Url, error) {

	return rm.url, rm.err
}
func (rm *RepositoryMock) UpdateUrl(url *model.Url) {
	url = rm.url
}
func (rm *RepositoryMock) MostVisited(n int) ([]model.Url, error) {

	return rm.urls, rm.err
}

func TestInsertNewAlias_success(t *testing.T) {
	expectedAlias := "alias"
	expectedLongUrl := "test"
	expectedVisits := 0
	var rMock repository.Repository = &RepositoryMock{
		url: &model.Url{
			Alias:   expectedAlias,
			LongUrl: expectedLongUrl,
			Visits:  expectedVisits},
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	url := model.Url{Alias: expectedAlias,
		LongUrl: expectedLongUrl}
	err := urlService.InsertNewAlias(&url)

	if err != nil {
		t.Fatal("Expected no error retrieving longUrl")
	}
	if url.Alias != expectedAlias {
		t.Fatal("Alias was modified")
	}
	if url.LongUrl != expectedLongUrl {
		t.Fatal("LongUrl was modified")
	}
	baseShortUrl := "http://localhost:8080/s/"
	if url.ShortUrl != baseShortUrl+expectedAlias {
		t.Fatal("ShortUrl different than expected.")
	}
	if url.Visits != expectedVisits {
		t.Fatal("Visits was modified")
	}
}

func TestInsertNewAlias_failEmpty(t *testing.T) {
	var rMock repository.Repository = &RepositoryMock{
		err: errors.New("Error case"),
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	var url model.Url
	err := urlService.InsertNewAlias(&url)

	if err == nil {
		t.Fatal("Expected error from repository")
	}
}

func TestInsertNewAlias_failAliasExists(t *testing.T) {
	expectedAlias := "alias"
	var rMock repository.Repository = &RepositoryMock{
		url: &model.Url{Alias: expectedAlias},
		err: errors.New("Error case"),
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	url := model.Url{Alias: expectedAlias}
	err := urlService.InsertNewAlias(&url)

	if url.Alias != "" && err == nil {
		t.Fatal("Expected error from repository")
	}
}

func TestRetrieveLongUrl_success(t *testing.T) {
	expectedLongUrl := "test"
	var rMock repository.Repository = &RepositoryMock{
		url: &model.Url{LongUrl: expectedLongUrl},
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	url, err := urlService.RetrieveLongUrl("alias")

	if url == nil || err != nil {
		t.Fatal("Expected no error retrieving longUrl")
	}
	if url.LongUrl != expectedLongUrl {
		t.Fatal("LongUrl retrieved with error")
	}
}

func TestRetrieveLongUrl_fail(t *testing.T) {
	var rMock repository.Repository = &RepositoryMock{
		err: errors.New("Error case"),
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	url, err := urlService.RetrieveLongUrl("")

	if url != nil || err == nil {
		t.Fatal("Expected error retrieving longUrl")
	}
}

func TestUpdateVisitUrl(t *testing.T) {
	expectedAlias := "alias"
	expectedLongUrl := "long"
	expectedShortUrl := "short"
	expectedVisits := 2
	var rMock repository.Repository = &RepositoryMock{
		url: &model.Url{
			Alias:    expectedAlias,
			LongUrl:  expectedLongUrl,
			ShortUrl: expectedShortUrl,
			Visits:   expectedVisits},
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	url := model.Url{
		Alias:    expectedAlias,
		LongUrl:  expectedLongUrl,
		ShortUrl: expectedShortUrl,
		Visits:   1}
	urlService.UpdateVisitUrl(&url)

	if url.Visits != expectedVisits {
		t.Fatal("Expected different visits count")
	}

	// side effects
	if url.Alias != expectedAlias {
		t.Fatal("Alias was modified")
	}
	if url.LongUrl != expectedLongUrl {
		t.Fatal("LongUrl was modified")
	}
	if url.ShortUrl != expectedShortUrl {
		t.Fatal("ShortUrl was modified")
	}
}

func TestMostVisited_success(t *testing.T) {
	expectedUrls := []model.Url{{LongUrl: "expectedLongUrl1"},
		{LongUrl: "expectedLongUrl2"}, {LongUrl: "expectedLongUrl3"}, {LongUrl: "expectedLongUrl4"},
		{LongUrl: "expectedLongUrl5"}, {LongUrl: "expectedLongUrl6"}, {LongUrl: "expectedLongUrl7"},
		{LongUrl: "expectedLongUrl8"}, {LongUrl: "expectedLongUrl9"}, {LongUrl: "expectedLongUrl10"}}
	var rMock repository.Repository = &RepositoryMock{
		urls: expectedUrls,
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	urls, err := urlService.MostVisited()

	if err != nil {
		t.Fatal("Expected no error retrieving longUrl")
	}
	topVisitQuantity := 10
	if len(urls) != topVisitQuantity || len(urls) != len(expectedUrls) {
		t.Fatal("Quantity returned wrong")
	}
}

func TestMostVisited_fail(t *testing.T) {
	var rMock repository.Repository = &RepositoryMock{
		err: errors.New("Error case"),
	}
	urlService := &service.UrlService{
		Repo: rMock,
	}

	urls, err := urlService.MostVisited()

	if err == nil || urls != nil {
		t.Fatal("Expected error retrieving most visited list")
	}
}
