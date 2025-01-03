package main

import (
	"os"
	"testing"

	"github.com/rishabhsvats/go-patterns/go-breeders/configuration"
	"github.com/rishabhsvats/go-patterns/go-breeders/models"
)

var testApp application

func TestMain(m *testing.M) {

	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}
	testApp = application{
		App:        configuration.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct {
}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{ID: 1, Breed: "Tomcat", Details: "Details about description of Tomcat"},
	}

	return breeds, nil
}
