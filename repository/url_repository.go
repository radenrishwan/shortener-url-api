package repository

import (
	"errors"
	"gorm.io/gorm"
	"seior-shortener-url/model/entity"
)

type UrlRepository interface {
	Create(url entity.Url) (entity.Url, error)
	FindByAlias(url entity.Url) (entity.Url, error)
	FindById(url entity.Url) (entity.Url, error)
	Update(url entity.Url) entity.Url
	Delete(url entity.Url) entity.Url
}

type urlRepository struct {
	*gorm.DB
}

func NewUrlRepository(DB *gorm.DB) UrlRepository {
	return &urlRepository{DB: DB}
}

func (repository *urlRepository) Create(url entity.Url) (entity.Url, error) {
	result := repository.DB.Create(&url)

	if result.Error != nil {
		return url, errors.New("alias already exist")
	}

	return url, nil
}

func (repository *urlRepository) FindByAlias(url entity.Url) (entity.Url, error) {
	result := repository.DB.Where("alias = ?", url.Alias).First(&url)

	if result.RowsAffected < 1 {
		return url, errors.New("url not found")
	}

	return url, nil
}

func (repository *urlRepository) FindById(url entity.Url) (entity.Url, error) {
	result := repository.DB.Where("id = ?", url.Id).First(&url)

	if result.RowsAffected < 1 {
		return url, errors.New("url not found")
	}

	return url, nil
}

func (repository *urlRepository) Update(url entity.Url) entity.Url {
	repository.DB.Model(&url).Where("id = ?", url.Id).Updates(&url)

	return url
}

func (repository *urlRepository) Delete(url entity.Url) entity.Url {
	repository.DB.Where("id = ?", url.Id).Delete(&url)

	return url
}
