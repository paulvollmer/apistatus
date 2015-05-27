package apistatus

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testserver() httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `H=LLO W=RLD`)
	}))
	return *ts
}

// should throw error for invalid URL
func TestPackage(t *testing.T) {
	status := ApiStatus{}

	_, err := status.Check("")
	if err == nil {
		t.Error("check", err)
	}

	_, err = status.Check("foo://")
	if err == nil {
		t.Error("check", err)
	}

	_, err = status.Check("http://www.sample.com")
	if err != nil {
		t.Error("check", err)
	}

	_, err = status.Check("https://www.sample.com")
	if err != nil {
		t.Error("check", err)
	}
}

func TestOnline_true(t *testing.T) {
	ts := testserver()
	defer ts.Close()
	status := ApiStatus{}
	status.Check(ts.URL)
	if status.Online != true {
		t.Error("Online is not true")
	}
}

func TestOnline_message(t *testing.T) {
	ts := testserver()
	defer ts.Close()
	status := ApiStatus{}
	status.Check(ts.URL)
	if status.Message != "OK" {
		t.Error("Message is not ok")
	}
}

func TestOnline_statusCode(t *testing.T) {
	ts := testserver()
	defer ts.Close()
	status := ApiStatus{}
	status.Check(ts.URL)
	if status.StatusCode != 200 {
		t.Error("StatusCode is not 200")
	}
}

func TestOffline_false(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://nodomain:1234")
	if status.Online != false {
		t.Error("Online is not false")
	}
}

func TestOffline_message(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://nodomain:1234")
	if status.Message != "failed" {
		t.Error("Message is not undefined")
	}
}

func TestOffline_statusCode(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://nodomain:1234")
	if status.StatusCode != 0 {
		t.Error("StatusCode is not 0")
	}
}

// TODO: func TestRedirect(t *testing.T) {}

func TestError_ClientError(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/404")
	if status.Category != "Client Error" {
		t.Error("Category is not Client Error")
	}
}

func TestError_ServerError(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/505")
	if status.Category != "Server Error" {
		t.Error("Category is not Server Error")
	}
}

func TestError_NotFound(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/404")
	if status.Message != "Not Found" {
		t.Error("Message is not Not Found")
	}
}

func TestError_404(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/404")
	if status.StatusCode != 404 {
		t.Error("statuscode is not 404")
	}
}

// TODO: Error 404 redirect

func TestStatusCode_123(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/123")
	if status.StatusCode != 123 {
		t.Error("statuscode is not 123")
	}
}

func TestStatusCode_777_NonStandardCategory(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/777")
	if status.Category != "Non-standard Category" {
		t.Error("Category is not Non-standard Category")
	}
}

// func TestStatusCode_123_NonStandardCategory(t *testing.T) {
// 	status := ApiStatus{}
// 	status.Check("http://httpbin.org/status/123")
// 	if status.Message != "Non-standard Category" {
// 		t.Error("Message is not Non-standard Category")
// 	}
// }

func TestStatusCode_500_NonStandardCategory(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://httpbin.org/status/500")
	if status.Message != "Internal Server Error" {
		t.Error("Message is not Internal Server Error")
	}
}

// TODO: HAR Requests

func TestLatency(t *testing.T) {
	status := ApiStatus{}
	status.Check("http://mockbin.com")
	if status.Latency < 0.0 {
		t.Error("Latency Error")
	}
}
