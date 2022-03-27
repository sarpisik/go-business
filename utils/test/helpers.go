package test_helpers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/sarpisik/go-business/app"
)

func ExecuteRequest(a *app.App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func GetByTestID(t *testing.T, testID string, res *httptest.ResponseRecorder) *goquery.Selection {
	var el *goquery.Selection

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	query := strings.Replace(`[data-testid="testID"]`, "testID", testID, 1)
	if el = doc.Find(query); el.Length() == 0 {
		return nil
	}

	return el
}
