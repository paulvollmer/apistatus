package apistatus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	CategoryInformationalMin = 100
	CategoryInformationalMax = 199
	CategorySuccessMin       = 200
	CategorySuccessMax       = 299
	CategoryRedirectMin      = 300
	CategoryRedirectMax      = 399
	CategoryClientErrorMin   = 400
	CategoryClientErrorMax   = 499
	CategoryServerErrorMin   = 500
	CategoryServerErrorMax   = 599
)

type ApiStatus struct {
	Online     bool    `json:"online"`
	Latency    float64 `json:"latency"`
	StatusCode int     `json:"statusCode"`
	Category   string  `json:"category"`
	Message    string  `json:"message"`
}

// send request to the server
func (a *ApiStatus) Check(theURL string) (int, error) {
	if theURL == "" {
		return 0, errors.New("URL string is empty")
	} else {
		uri, err := url.Parse(theURL)
		if err != nil {
			return 0, err
		}

		if uri.Scheme == "http" || uri.Scheme == "https" {
			// request...
			t1 := time.Now()
			resp, err := http.Get(theURL)
			if err != nil {
				// fmt.Println("something went wrong", err)
				a.Online = false
				a.StatusCode = 0
				a.Category = "request failed"
				a.Message = "failed"
				return 0, nil
			}
			// fmt.Printf("%# v", resp)
			t2 := time.Now()
			d := t2.Sub(t1)

			a.Online = true
			a.Latency = d.Seconds()
			a.StatusCode = resp.StatusCode
			a.Category = a.CategoryText(resp.StatusCode)
			a.Message = a.StatusText(resp.StatusCode)
			// log.Println("<== Check ready to ship... statuscode is ", a.StatusCode)
			return a.StatusCode, nil
		} else {
			return 0, errors.New("url scheme '" + uri.Scheme + "' not supported")
		}
	}
}

func (a *ApiStatus) GetJSON() string {
	jsonData, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error", err)
	}
	// fmt.Println(string(jsonData))
	return string(jsonData)
}

func (a *ApiStatus) CategoryText(statusCode int) string {
	if statusCode >= CategoryInformationalMin && statusCode <= CategoryInformationalMax {
		return "Informational"
	} else if statusCode >= CategorySuccessMin && statusCode <= CategorySuccessMax {
		return "Success"
	} else if statusCode >= CategoryRedirectMin && statusCode <= CategoryRedirectMax {
		return "Redirect"
	} else if statusCode >= CategoryClientErrorMin && statusCode <= CategoryClientErrorMax {
		return "Client Error"
	} else if statusCode >= CategoryServerErrorMin && statusCode <= CategoryServerErrorMax {
		return "Server Error"
	} else {
		return "Non-standard Category"
	}
}

// extend the status text
func (a *ApiStatus) StatusText(statusCode int) string {
	switch statusCode {
	case 100, 101,
		200, 201, 202, 203, 204, 205, 206,
		300, 301, 302, 303, 304, 305, 307,
		400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418,
		500, 501, 502, 503, 504, 505,
		428, 429, 431, 511:
		return http.StatusText(statusCode)
	case 102:
		return "Processing"
	case 207:
		return "Multi-Status"
	case 226:
		return "IM Used"
	case 308:
		return "Permanent Redirect"
	case 422:
		return "Unprocessable Entity"
	case 423:
		return "Locked"
	case 424:
		return "Failed Dependency"
	case 426:
		return "Upgrade Required"
	case 451:
		return "Unavailable For Legal Reasons"
	case 506:
		return "Variant Also Negotiates"
	case 507:
		return "Insufficient Storage"
	default:
		return "Non-standard Status"
	}
}
