package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-res", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"post-search-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "2022-04-27"},
		{key: "end", value: "2022-04-30"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2022-04-27"},
		{key: "end", value: "2022-04-30"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		// {key: "csrf_token", value: nosurf.Token()}
		{key: "first_name", value: "Jhonatan"},
		{key: "last_name", value: "Teodoro"},
		{key: "email", value: "dasdas@dasda.com"},
		{key: "phone", value: "352236235"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {

			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Error(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf(
					"For %s. expected %d but got %d",
					test.name, test.expectedStatusCode, resp.StatusCode,
				)
			}

		} else {
			// fmt.Println("Hey ")
			params := url.Values{}
			for _, item := range test.params {
				params.Add(item.key, item.value)
			}
			// client := ts.Client()
			resp, err := ts.Client().PostForm(ts.URL+test.url, params)
			if err != nil {
				t.Log(err)
				t.Error(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf(
					"For %s. expected %d but got %d",
					test.name, test.expectedStatusCode, resp.StatusCode,
				)
			}

		}
	}
}
