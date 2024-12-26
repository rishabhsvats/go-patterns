package main

import (
	"os"
	"testing"

	"github.com/rishabhsvats/go-patterns/go-breeders/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
