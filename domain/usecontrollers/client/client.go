package client

import (
	"fmt"
	"github.com/Pietroski/golang-notification/errors"
	"io"
	"net/http"
)

const (
	USER_AGENT = "Mozilla/5.0 (X11; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0"
)

var (
	CrawlerClient ICrawlerClient = &SCrawlerClient{}
)

type ICrawlerClient interface {
	NewCrawlerClient()
	NewCrawlerRequester(method string, URL string, body io.Reader)
	GetCrawlerClientHTTPRequest() *http.Request
	SetDefaultHeaders()
	SetDefaultCrawlerRequester(method string, URL string, body io.Reader)
	CrawlerDo() *http.Response
}

type SCrawlerClient struct {
	client  *http.Client
	request *http.Request
}

// NewCrawlerClient instantiates a new crawler client
// with default redirect checking.
func (ccs *SCrawlerClient) NewCrawlerClient() {
	ccs.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect ->", req.URL.Host, req, via)
			return http.ErrUseLastResponse
		},
	}
}

// NewCrawlerRequester instantiates a new HTTP request for the crawler client
func (ccs *SCrawlerClient) NewCrawlerRequester(method string, URL string, body io.Reader) {
	var err error
	ccs.request, err = http.NewRequest(method, URL, body)
	errors.Error.CheckFatalError(err)
}

// GetCrawlerClientHTTPRequest gets and returns the pointer to the instance of the crawler requester
func (ccs *SCrawlerClient) GetCrawlerClientHTTPRequest() *http.Request {
	request := ccs.request
	return request
}

// SetDefaultHeaders sets the default header
func (ccs *SCrawlerClient) SetDefaultHeaders() {
	ccs.request.Header.Add("User-Agent", USER_AGENT)
}

// SetDefaultCrawlerRequester initializes the crawler requester and sets its default headers
func (ccs *SCrawlerClient) SetDefaultCrawlerRequester(method string, URL string, body io.Reader) {
	ccs.NewCrawlerRequester(method, URL, body)
	ccs.SetDefaultHeaders()
}

// CrawlerDo fires the request previous configured
func (ccs *SCrawlerClient) CrawlerDo() *http.Response {
	response, err := ccs.client.Do(ccs.request)
	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
	}
	errors.Error.CheckFatalError(err)

	return response
}

func (ccs *SCrawlerClient) responseMiddleware(resp *http.Response) *http.Response {
	//window := js.Global().Get("window")
	//fmt.Println(window)
	//resp.Cookies()
	//
	return resp
}
