package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestApiTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Test Suite Example")
}
