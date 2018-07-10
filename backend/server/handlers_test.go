package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/businesscard/backend/data"
)

var testVcard = "{\"name\":\"Сергей\",\"position\":\"заведующий\",\"phone\":\"+7-999-999-99-99\",\"email\":\"mail@domain.ru\",\"workplace\":\"VNII\",\"address\":\"street\"}"

func TestCrudHandlers(t *testing.T) {
	vl := data.NewVcardList()
	router := NewRouter(vl)
	req, err := http.NewRequest("GET", "/api/businesscard/vcards", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testVcard)
	req, err = http.NewRequest("POST", "/api/businesscard/vcards", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/businesscard/vcards/1" {
		t.Error("Expected another location")
	}

	if len(vl.GetVcards()) != 1 {
		t.Error("Expected new value")
	}
	testData = strings.NewReader(testVcard)
	req, err = http.NewRequest("PUT", "/api/businesscard/vcards/1", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/businesscard/vcards/1" {
		t.Error("Expected another location")
	}

	if len(vl.GetVcards()) != 1 {
		t.Error("Expected old value")
	}

	testData = strings.NewReader(testVcard)
	req, err = http.NewRequest("DELETE", "/api/businesscard/vcards/1", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten: %d)", resp.StatusCode)

	}

	if len(vl.GetVcards()) != 0 {
		t.Error("Expected old value")
	}

}
