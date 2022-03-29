package test_helpers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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

func GetByTestID2(t *testing.T, testID string, doc *goquery.Document) *goquery.Selection {
	var el *goquery.Selection

	query := strings.Replace(`[data-testid="testID"]`, "testID", testID, 1)
	if el = doc.Find(query); el.Length() == 0 {
		return nil
	}

	return el
}

func GetDocFromRes(t *testing.T, res *httptest.ResponseRecorder) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return doc
}

func Signup(t *testing.T, a *app.App, data *url.Values) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res := ExecuteRequest(a, req)

	expected := http.StatusOK
	received := res.Code
	CheckResponseCode(t, expected, received)

	if el := GetByTestID(t, "signup-success-msg", res); el.Length() == 0 {
		t.Fatal("Success signup message not rendered.")
	}

	return res
}

func TestEmptyCookie(res *httptest.ResponseRecorder, t *testing.T) {
	for _, cookie := range res.Result().Cookies() {
		if cookie.Name == "session" {
			if len(cookie.Value) != 0 {
				t.Fatal(fmt.Sprintf("Cookie value should be empty. Received %s", cookie.Value))
			}
		}
	}
}

func SignupAndLogin(t *testing.T, a *app.App, data *url.Values) string {
	Signup(t, a, data)

	// Login and keep the cookie
	data.Del("name")
	data.Del("confirmPassword")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res := ExecuteRequest(a, req)

	expected := http.StatusFound
	received := res.Code
	CheckResponseCode(t, expected, received)

	var cookieValue string
	cookies := res.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "session" {
			cookieValue = cookie.Value
			break
		}
	}

	if len(cookieValue) == 0 {
		t.Fatal("Cookie not received.")
	}

	return cookieValue
}
