package service

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"seior-shortener-url/database"
	"seior-shortener-url/model/entity"
	"seior-shortener-url/model/web"
	"seior-shortener-url/repository"
	"testing"
)

type urlServiceSuite struct {
	suite.Suite
	*gorm.DB
	UrlService
}

func (suite *urlServiceSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	suite.DB = database.NewDB()

	err = suite.DB.AutoMigrate(&entity.Url{})
	if err != nil {
		panic(err)
	}

	urlRepository := repository.NewUrlRepository(suite.DB)

	suite.UrlService = NewUrlService(urlRepository)

	suite.DB.Exec("delete from urls")
}

func (suite *urlServiceSuite) TestCreateSuccess() {
	url := suite.UrlService.Create(web.CreateUrlRequest{
		Destination: "https://google.com",
		Alias:       "google",
	})

	alias := suite.UrlService.FindByAlias(web.FindByAliasUrlRequest{
		Alias: url.Data.Alias,
	})

	assert.Equal(suite.T(), alias, url)
}

func (suite *urlServiceSuite) TestCreateFailedDestination() {
	assert.Panics(suite.T(), func() {
		suite.UrlService.Create(web.CreateUrlRequest{
			Destination: "asdlkjasldkaasdljha",
			Alias:       "google",
		})
	})
}

func (suite *urlServiceSuite) TestCreateFailedDuplicate() {
	assert.Panics(suite.T(), func() {
		suite.UrlService.Create(web.CreateUrlRequest{
			Destination: "https://google.com",
			Alias:       "google",
		})

		suite.UrlService.Create(web.CreateUrlRequest{
			Destination: "https://google.com",
			Alias:       "google",
		})
	})
}

func (suite *urlServiceSuite) TestDeleteSuccess() {
	suite.DB.Exec("delete from urls")

	url := suite.UrlService.Create(web.CreateUrlRequest{
		Destination: "https://google.com",
		Alias:       "google",
	})

	suite.UrlService.Delete(web.DeleteUrlRequest{
		Id: url.Data.Id,
	})

	assert.Panics(suite.T(), func() {
		suite.UrlService.FindByAlias(web.FindByAliasUrlRequest{
			Alias: url.Data.Alias,
		})
	})
}

func (suite *urlServiceSuite) TestDeleteFailedNotFound() {
	suite.DB.Exec("delete from urls")

	suite.UrlService.Create(web.CreateUrlRequest{
		Destination: "https://google.com",
		Alias:       "google",
	})

	assert.Panics(suite.T(), func() {
		suite.UrlService.Delete(web.DeleteUrlRequest{
			Id: "asjdhasdh",
		})
	})
}

func (suite *urlServiceSuite) TestDeleteFailedInput() {
	suite.DB.Exec("delete from urls")

	assert.Panics(suite.T(), func() {
		suite.UrlService.Create(web.CreateUrlRequest{
			Destination: "asjdhaskjh",
			Alias:       "google",
		})
	})
}

func (suite *urlServiceSuite) TestUpdateSuccess() {
	result := suite.UrlService.Create(web.CreateUrlRequest{
		Destination: "https://google.com",
		Alias:       "google",
	})

	suite.UrlService.Update(web.UpdateUrlRequest{
		Id:          result.Data.Id,
		Destination: "https://facebook.com",
		Alias:       "facebook",
	})

	assert.NotPanics(suite.T(), func() {
		suite.UrlService.FindByAlias(web.FindByAliasUrlRequest{
			Alias: "facebook",
		})
	})
}

func (suite *urlServiceSuite) TestUpdateFailedNotFound() {
	suite.UrlService.Create(web.CreateUrlRequest{
		Destination: "https://google.com",
		Alias:       "google",
	})

	assert.Panics(suite.T(), func() {
		suite.UrlService.Update(web.UpdateUrlRequest{
			Id:          "alskdjasljdlas",
			Destination: "https://facebook.com",
			Alias:       "facebook",
		})

	})
}

func (suite *urlServiceSuite) TestUpdateFailedInput() {
	assert.Panics(suite.T(), func() {
		suite.UrlService.Create(web.CreateUrlRequest{
			Destination: "aslkdhasldjh",
			Alias:       "",
		})
	})
}

func (suite *urlServiceSuite) TearDownSuite() {
	suite.DB.Exec("delete from urls")
}

func TestUrlService(t *testing.T) {
	suite.Run(t, new(urlServiceSuite))
}
