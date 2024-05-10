package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GMcD/api-semaphore/module/api"
)

var a api.App

func TestEmptyTable(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	req, _ := http.NewRequest("GET", "/productbyname/fred", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestCreateProduct(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test product" {
		t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
	}

	if m["price"] != 11.22 {
		t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
	}
}

func TestGetProductByName(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	const tesla = "Tesla"
	teslaNameUrl := fmt.Sprintf("/productbyname/%v/", tesla)

	req, _ := http.NewRequest("GET", teslaNameUrl, nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if response.Code != http.StatusOK {
		t.Errorf("Failed to retrieve: '%v'", tesla)
	}
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)
	if originalProduct["name"].(string) != tesla {
		t.Errorf("Failed to retrieve details of '%v'", tesla)
	}
}

func TestUpdateProduct(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	const tesla = "Tesla"
	teslaNameUrl := fmt.Sprintf("/productbyname/%v/", tesla)

	// const ford = "Ford"
	// ford_url := fmt.Sprintf("/productbyname/%v/", ford)

	req, _ := http.NewRequest("GET", teslaNameUrl, nil)
	response := executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Failed to retrieve: '%v'", tesla)
	}
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	teslaIdUrl := fmt.Sprintf("/product/%v/", originalProduct["id"].(string))
	var jsonStr = []byte(`{"name":"ford", "price": 11.22}`)
	req, _ = http.NewRequest("PUT", teslaIdUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalProduct["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["id"], m["id"])
	}

	if m["name"] == originalProduct["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalProduct["name"], m["name"], m["name"])
	}

	if m["price"] == originalProduct["price"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", originalProduct["price"], m["price"], m["price"])
	}
}

func TestDeleteProduct(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	tesla := "Tesla"

	teslaNameUrl := fmt.Sprintf("/productbyname/%v/", tesla)

	req, _ := http.NewRequest("GET", teslaNameUrl, nil)

	response := executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Failed to retrieve: '%v'", tesla)
	}
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	teslaIdUrl := fmt.Sprintf("/product/%v/", originalProduct["id"].(string))

	req, _ = http.NewRequest("DELETE", teslaIdUrl, nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", teslaNameUrl, nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

//
// Db Utilities
//

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
