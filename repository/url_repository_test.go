package repository

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"seior-shortener-url/database"
	"seior-shortener-url/model/entity"
	"strconv"
	"testing"
	"time"
)

type urlRepositorySuite struct {
	suite.Suite
	*gorm.DB
	UrlRepository
}

func (suite *urlRepositorySuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	suite.DB = database.NewDB()

	err = suite.DB.AutoMigrate(&entity.Url{})
	if err != nil {
		panic(err)
	}

	suite.UrlRepository = NewUrlRepository(suite.DB)

	suite.DB.Exec("delete from urls")
}

func (suite *urlRepositorySuite) TestDelete() {
	date := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

	result, err := suite.UrlRepository.Create(entity.Url{
		Id:          uuid.NewString(),
		Destination: "www.google.com",
		Alias:       "google",
		Clicked:     0,
		CreatedAt:   date,
		UpdatedAt:   date,
	})

	suite.UrlRepository.Delete(result)

	_, err = suite.UrlRepository.FindByAlias(entity.Url{
		Alias: "google",
	})

	assert.NotNil(suite.T(), err)
}

func (suite *urlRepositorySuite) TestCreateDuplicate() {
	createdAt := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

	suite.UrlRepository.Create(entity.Url{
		Id:          uuid.NewString(),
		Destination: "www.google.com",
		Alias:       "google",
		Clicked:     0,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	})

	_, err := suite.UrlRepository.Create(entity.Url{
		Id:          uuid.NewString(),
		Destination: "www.google.com",
		Alias:       "google",
		Clicked:     0,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	})

	assert.NotNil(suite.T(), err)
}

func (suite *urlRepositorySuite) TestCreate() {
	createdAt := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

	suite.UrlRepository.Create(entity.Url{
		Id:          uuid.NewString(),
		Destination: "www.google.com",
		Alias:       "google",
		Clicked:     0,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	})

	_, err := suite.UrlRepository.FindByAlias(entity.Url{
		Alias: "google",
	})

	assert.Nil(suite.T(), err)
}

func (suite *urlRepositorySuite) TestUpdate() {
	date := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

	result, err := suite.UrlRepository.Create(entity.Url{
		Id:          uuid.NewString(),
		Destination: "www.google.com",
		Alias:       "google",
		Clicked:     0,
		CreatedAt:   date,
		UpdatedAt:   date,
	})

	result.Clicked = 1
	result.Alias = "google2"

	suite.UrlRepository.Update(result)

	editedResult, err := suite.UrlRepository.FindByAlias(entity.Url{
		Alias: "google2",
	})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "google2", editedResult.Alias)
	assert.Equal(suite.T(), uint64(1), editedResult.Clicked)
}

//func (suite *urlRepositorySuite) TearDownSuite() {
//	suite.DB.Exec("delete from urls")
//}

func TestUrlRepository(t *testing.T) {
	suite.Run(t, new(urlRepositorySuite))
}
