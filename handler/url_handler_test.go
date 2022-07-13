package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"seior-shortener-url/database"
	"seior-shortener-url/model/entity"
	"seior-shortener-url/model/web"
	"seior-shortener-url/repository"
	"seior-shortener-url/service"
	"strings"
	"testing"
)

type urlHandlerSuite struct {
	suite.Suite
	*gorm.DB
	UrlHandler
	repository.UrlRepository
	service.UrlService
	*fiber.App
}

func (suite *urlHandlerSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	suite.DB = database.NewDB()

	err = suite.DB.AutoMigrate(&entity.Url{})
	if err != nil {
		panic(err)
	}

	suite.UrlRepository = repository.NewUrlRepository(suite.DB)
	suite.UrlService = service.NewUrlService(suite.UrlRepository)

	suite.UrlHandler = NewUrlHandler(suite.UrlService)

	suite.App = fiber.New()
	suite.App.Use(recover.New())
	suite.App.Use(logger.New())

	suite.App.Post("/api/url", suite.UrlHandler.CreateUrl)
	suite.App.Delete("/api/url/", suite.UrlHandler.DeleteUrl)
	suite.App.Get("/api/url/", suite.UrlHandler.GetUrlInfo)
	suite.App.Put("/api/url/", suite.UrlHandler.UpdateUrl)

	suite.App.Get("/:alias", suite.UrlHandler.Redirect)

	suite.DB.Exec("delete from urls")
}

func (suite *urlHandlerSuite) TestCreateSuccess() {
	body := strings.NewReader(`{"destination":"https://www.google.com", "alias": "google"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/url", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.App.Test(req)
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	var result fiber.Map
	err = json.Unmarshal(bytes, &result)

	data := result["data"].(map[string]interface{})

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), "https://www.google.com", data["destination"])
	assert.Equal(suite.T(), "google", data["alias"])
}

func (suite *urlHandlerSuite) TestDeleteSuccess() {
	suite.DB.Exec("delete from urls")
	createUrl := suite.UrlService.Create(web.CreateUrlRequest{Destination: "https://www.google.com", Alias: "google"})

	req := httptest.NewRequest(http.MethodDelete, "/api/url?id="+createUrl.Data.Id, nil)

	resp, err := suite.App.Test(req)
	if err != nil {
		panic(err)
	}

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}

func (suite *urlHandlerSuite) TestUpdateSuccess() {
	suite.DB.Exec("delete from urls")
	createUrl := suite.UrlService.Create(web.CreateUrlRequest{Destination: "https://www.google.com", Alias: "google"})

	bodyString := fmt.Sprintf(
		`{"id": "%s", "destination": "https://www.facebook.com", "alias": "facebook"}`, createUrl.Data.Id,
	)

	body := strings.NewReader(bodyString)
	req := httptest.NewRequest(http.MethodPut, "/api/url", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.App.Test(req)
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	var result fiber.Map
	err = json.Unmarshal(bytes, &result)

	data := result["data"].(map[string]interface{})

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), "https://www.facebook.com", data["destination"])
	assert.Equal(suite.T(), "facebook", data["alias"])
}

func (suite *urlHandlerSuite) TestRedirectSuccess() {
	suite.DB.Exec("delete from urls")
	suite.UrlService.Create(web.CreateUrlRequest{Destination: "https://www.google.com", Alias: "google"})

	req := httptest.NewRequest(http.MethodGet, "/google", nil)

	resp, err := suite.App.Test(req)
	if err != nil {
		panic(err)
	}

	assert.Equal(suite.T(), http.StatusFound, resp.StatusCode)
}

func (suite *urlHandlerSuite) TearDownSuite() {
	suite.DB.Exec("delete from urls")
}

func TestUrlService(t *testing.T) {
	suite.Run(t, new(urlHandlerSuite))
}
