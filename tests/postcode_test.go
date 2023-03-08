package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang-ginkgo-tests/internal/config"
	"io"
	"net/http"
	"strings"
)

var _ = Describe("Postcode", func() {
	It("should return a valid lat/long", func() {
		testPostCode := "SE1 7QD"
		baseUrl := config.PostCodeIOEndPoint() + "/postcodes/%s"
		url := fmt.Sprintf(baseUrl, testPostCode)
		method := "GET"

		payload := strings.NewReader("")

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		Ω(err).ShouldNot(HaveOccurred())

		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		Ω(err).ShouldNot(HaveOccurred())

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				Ω(err).ShouldNot(HaveOccurred())
			}
		}(res.Body)
		Ω(res.StatusCode).Should(Equal(200), "Get Postcode lat long Status code failed")
	})
})
