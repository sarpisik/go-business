package main_test

import (
	"net/http"
	"net/url"
	"testing"

	test_helpers "github.com/sarpisik/go-business/utils/test"
)

func TestIndexGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := test_helpers.ExecuteRequest(&a, req)

	expected := http.StatusOK
	received := res.Code

	test_helpers.CheckResponseCode(t, expected, received)

	if el := test_helpers.GetByTestID(t, "login-button", res); el.Length() == 0 {
		t.Fatal("Login button not rendered.")
	}
}

func TestIndexAuth(t *testing.T) {
	// Sign up
	data := url.Values{}
	data.Set("name", "test user")
	data.Set("email", "user@test.com")
	data.Set("password", "123456")
	data.Set("confirmPassword", "123456")

	cookieValue := test_helpers.SignupAndLogin(t, &a, &data)

	// Get index page
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: cookieValue})
	res := test_helpers.ExecuteRequest(&a, req)

	expected := http.StatusOK
	received := res.Code

	test_helpers.CheckResponseCode(t, expected, received)

	if el := test_helpers.GetByTestID(t, "login-button", res); el != nil {
		t.Fatal("Login button should not rendered.")
	}
}
