package database

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}

func TestNewDB(t *testing.T) {
	setUp()

	db := NewDB()

	assert.Nil(t, db.Error)
}
