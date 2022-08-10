package api

import (
	"testing"
)

/* type test struct {
	server *httptest.Server
} */

func TestAddImage(t *testing.T) {

}

func TestFindImage(t *testing.T) {

}

/* func TestFindImages(t *testing.T) {
	req, err := http.NewRequest("GET", "/idk", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	a := &Application{}

	//handler := http.HandlerFunc(a.FindImage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

} */

func TestDeleteImage(t *testing.T) {

}

func TestUpdateImage(t *testing.T) {

}
