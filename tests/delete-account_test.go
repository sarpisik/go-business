package main_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/sarpisik/go-business/models"
	test_helpers "github.com/sarpisik/go-business/utils/test"
)

func TestDeleteAccountGetFail(t *testing.T) {
	// Fails when trying to delete a non exist user.
	req, _ := http.NewRequest(http.MethodGet, "/delete-account", nil)
	res := test_helpers.ExecuteRequest(&a, req)

	expected := http.StatusInternalServerError
	received := res.Code

	test_helpers.CheckResponseCode(t, expected, received)

	// Fails when trying to delete an already deleted user.
	email := "already@deleted.com"

	data := url.Values{}
	data.Set("name", "test user")
	data.Set("email", email)
	data.Set("password", "123456")
	data.Set("confirmPassword", "123456")

	test_helpers.SignupAndLogin(t, &a, &data)

	// Delete the user
	u := models.User{
		Email: email,
	}
	u.GetUserByEmail(a.DB)
	u.DeleteUserByID(a.DB)

	// Test api response
	req, _ = http.NewRequest(http.MethodGet, "/delete-account", nil)
	res = test_helpers.ExecuteRequest(&a, req)

	expected = http.StatusInternalServerError
	received = res.Code

	test_helpers.CheckResponseCode(t, expected, received)
}

func TestDeleteAccountGetSuccess(t *testing.T) {
	email := "user@deleted.com"

	data := url.Values{}
	data.Set("name", "test user")
	data.Set("email", email)
	data.Set("password", "123456")
	data.Set("confirmPassword", "123456")

	cookieValue := test_helpers.SignupAndLogin(t, &a, &data)

	// Test api response
	req, _ := http.NewRequest(http.MethodGet, "/delete-account", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: cookieValue})
	res := test_helpers.ExecuteRequest(&a, req)

	expected := http.StatusFound
	received := res.Code

	test_helpers.CheckResponseCode(t, expected, received)
	test_helpers.TestEmptyCookie(res, t)

	// Validate user not exist in DB.
	u := models.User{Email: email}
	err := u.GetUserByEmail(a.DB)
	if err == nil {
		t.Fatal("User should not exist in DB.")
	}
}
