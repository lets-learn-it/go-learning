package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"avabodha.in/concurrency/subscription/data"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login",
		url:                "/login",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout",
		url:                "/logout",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.LoginPage,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func Test_Page(t *testing.T) {
	pathToTemplates = "./templates"

	for _, e := range pageTests {
		rr := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", e.url, nil)

		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for key, value := range e.sessionData {
				testApp.Session.Put(ctx, key, value)
			}
		}

		e.handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error(e.name, " failed to get status OK, got ", rr.Code)
		}

		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Error(e.name, " expected html ", e.expectedHTML, " not found.")
			}
		}
	}

}

func TestConfig_PostLoginPage(t *testing.T) {
	pathToTemplates = "./templates"

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"adbasjahdhuisdiuh"},
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("Failed to login")
	}

	// check userID in session
	if !testApp.Session.Exists(ctx, "userID") {
		t.Error("did not find userId in session")
	}
}

func TestConfig_SubscribeToPlan(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subscribe?id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "user", data.User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "admin",
		LastName:  "admin",
		Active:    1,
	})

	handler := http.HandlerFunc(testApp.SubscribeToPlan)
	handler.ServeHTTP(rr, req)

	testApp.Wait.Wait()

	if rr.Code != http.StatusSeeOther {
		t.Error("expected status code of statusseeother, but got", rr.Code)
	}
}
