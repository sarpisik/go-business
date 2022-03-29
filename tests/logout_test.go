package main_test

import (
	"net/http"
	"testing"

	test_helpers "github.com/sarpisik/go-business/utils/test"
)

func TestLogoutGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
	res := test_helpers.ExecuteRequest(&a, req)
	expected := http.StatusFound
	received := res.Code

	test_helpers.CheckResponseCode(t, expected, received)
	test_helpers.TestEmptyCookie(res, t)
}
