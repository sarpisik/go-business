package main_test

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/sarpisik/go-business/models"
	test_helpers "github.com/sarpisik/go-business/utils/test"
)

func TestSignupGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/signup", nil)
	res := test_helpers.ExecuteRequest(&a, req)

	expected := http.StatusOK
	received := res.Code

	test_helpers.CheckResponseCode(t, expected, received)

	if el := test_helpers.GetByTestID(t, "signup-submit-btn", res); el.Length() == 0 {
		t.Fatal("Signup submit button not rendered.")
	}
}

func TestSignupPostFail(t *testing.T) {
	email := "existing@user.com"
	u := models.User{Email: email}
	u.CreateUser(a.DB)

	defer u.DeleteUserByID(a.DB)

	type field struct {
		name string
		id   string
	}

	type test struct {
		email           string
		password        string
		confirmPassword string
		expected        int
		fields          []field
	}

	tests := []test{
		{email: "", password: "", confirmPassword: "", expected: http.StatusOK, fields: []field{
			{name: "email", id: "email-error-msg"},
			{name: "password", id: "password-error-msg"},
			{name: "confirm password", id: "confirm-password-error-msg"},
		}},
		{email: "abcdef", password: "", confirmPassword: "", expected: http.StatusOK, fields: []field{
			{name: "email", id: "email-error-msg"},
			{name: "password", id: "password-error-msg"},
			{name: "confirm password", id: "confirm-password-error-msg"},
		}},
		{email: "test@mail.com", password: "", confirmPassword: "", expected: http.StatusOK, fields: []field{
			{name: "password", id: "password-error-msg"},
			{name: "confirm password", id: "confirm-password-error-msg"},
		}},
		{email: "test@mail.com", password: "1234", confirmPassword: "", expected: http.StatusOK, fields: []field{
			{name: "password", id: "password-error-msg"},
			{name: "confirm password", id: "confirm-password-error-msg"},
		}},
		{email: "test@mail.com", password: "123456", confirmPassword: "654321", expected: http.StatusOK, fields: []field{
			{name: "confirm password", id: "confirm-password-error-msg"},
		}},
		{email: email, password: "123456", confirmPassword: "123456", expected: http.StatusOK, fields: []field{
			{name: "user not found error", id: "other-error-msg"},
		}},
	}

	for _, tc := range tests {
		data := url.Values{}
		data.Set("email", tc.email)
		data.Set("password", tc.password)
		data.Set("confirmPassword", tc.confirmPassword)

		req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		res := test_helpers.ExecuteRequest(&a, req)
		received := res.Code
		test_helpers.CheckResponseCode(t, tc.expected, received)

		doc := test_helpers.GetDocFromRes(t, res)

		for _, field := range tc.fields {
			if el := test_helpers.GetByTestID2(t, field.id, doc); el == nil {
				t.Fatal(strings.Replace("Invalid field_name message should rendered.", "field_name", field.name, 1))
			}
		}
	}
}

func TestSignupPostSuccess(t *testing.T) {
	email := "new@user.com"

	// Delete the user if exist
	u := models.User{
		Email: email,
	}
	u.GetUserByEmail(a.DB)
	u.DeleteUserByID(a.DB)

	data := url.Values{}
	data.Set("name", "test user")
	data.Set("email", email)
	data.Set("password", "123456")
	data.Set("confirmPassword", "123456")
	test_helpers.Signup(t, &a, &data)
}
