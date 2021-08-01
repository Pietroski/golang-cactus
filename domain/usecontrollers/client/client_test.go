package client

import (
	"fmt"
	"github.com/Pietroski/golang-VivaRealCrawler/data/mocks"
	"github.com/Pietroski/golang-notification/errors"
	"github.com/Pietroski/golang-pietroski-reflections/converters"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

const (
	Method = "GET"
	URL    = "https://www.example.com"
)

var (
	ClientMock IClientMock = &SClientMock{}
)

type IClientMock interface {
	ClientMocker() *SCrawlerClient
}

type SClientMock struct{}

func (scm *SClientMock) ClientMocker() *SCrawlerClient {
	var ClientMock = SCrawlerClient{}
	CrawlerClient.SetDefaultCrawlerRequester(Method, URL, nil)

	return &ClientMock
}

func init() {
	ClientMock.ClientMocker()
}

func TestNewCrawlerClient(t *testing.T) {
	fmt.Println("TestNewCrawlerClient")

	var CrawlerClientToTest = SCrawlerClient{}
	CrawlerClientToTest.NewCrawlerClient()

	cctc := CrawlerClientToTest.client
	cctcType := reflect.TypeOf(cctc)
	expectedType := reflect.TypeOf(&http.Client{})

	if cctcType != expectedType {
		errors.Error.AssertionTestError(1, expectedType, cctcType)
	}

	cctcCheckRedirectType := reflect.TypeOf(cctc.CheckRedirect)
	expectedCCTCCRType := reflect.TypeOf(
		//TODO: Mock this function return type and itself if possible
		func(req *http.Request, via []*http.Request) error { return fmt.Errorf("mocked error") },
	)

	if cctcCheckRedirectType != expectedCCTCCRType {
		errors.Error.AssertionTestError(2, expectedCCTCCRType, cctcCheckRedirectType)
	}

	fmt.Println()
}

type NewCrawlerRequesterTestInputStruct struct {
	Method string
	URL    string
	Body   io.Reader
}

var bodyType = reflect.TypeOf(strings.NewReader("Body Type"))

type NewCrawlerRequesterTestOutputStruct struct {
	Method string
	URL    string
	Body   io.Reader
}

type NewCrawlerRequesterTestStruct struct {
	input  NewCrawlerRequesterTestInputStruct
	output NewCrawlerRequesterTestOutputStruct
}

var NewCrawlerRequesterTestMap = map[int]NewCrawlerRequesterTestStruct{
	0: {
		input: NewCrawlerRequesterTestInputStruct{
			Method: "GET",
			URL:    "https://www.example.com",
			Body:   mocks.ClientMock.MockIOReader(),
		},
		output: NewCrawlerRequesterTestOutputStruct{
			Method: "GET",
			URL:    "https://www.example.com",
			Body:   mocks.ClientMock.MockIOReader(),
		},
	},
	1: {
		input: NewCrawlerRequesterTestInputStruct{
			Method: "POST",
			URL:    "https://www.another-example.com",
			Body:   mocks.ClientMock.MockIOReader(),
		},
		output: NewCrawlerRequesterTestOutputStruct{
			Method: "POST",
			URL:    "https://www.another-example.com",
			Body:   mocks.ClientMock.MockIOReader(),
		},
	},
}

var NewCrawlerRequesterTestSlice = converters.MapToSlice(NewCrawlerRequesterTestMap)

func CrawlerClientHTTPRequestTest() {
	NewCrawlerRequesterTestSliceSize := NewCrawlerRequesterTestSlice.Len()
	for index := 0; index < NewCrawlerRequesterTestSliceSize; index++ {
		v := NewCrawlerRequesterTestSlice.Index(index)
		input := v.FieldByName("input")
		inputMethod := input.FieldByName("Method").String()
		inputURL := input.FieldByName("URL").String()
		inputBody := strings.NewReader(input.FieldByName("Body").String())

		output := v.FieldByName("output")
		outputMethod := output.FieldByName("Method").String()
		outputURL := output.FieldByName("URL").String()
		outputBody := output.FieldByName("Body")
		//outputBodyValue := output.FieldByName("Body").Elem().Elem().Field(0).String()
		//.Type() // .Elem().Type() //.Elem().Type() // .Field(0).Type()
		//newOutputBody := io.NopCloser(strings.NewReader(outputBodyValue))

		CrawlerClient.NewCrawlerRequester(inputMethod, inputURL, inputBody)
		request := CrawlerClient.GetCrawlerClientHTTPRequest()

		if request.Method != outputMethod {
			errors.Error.AssertionTestError(index, outputMethod, request.Method)
		}
		if request.URL.String() != outputURL {
			errors.Error.AssertionTestError(index, outputURL, request.URL)
		}
		// reflect.TypeOf(request.Body) == reflect.TypeOf(newOutputBody)
		if reflect.DeepEqual(request.Body, outputBody)  {
			errors.Error.AssertionTestError(index, request.Body, outputBody)
		}
	}
}

func TestNewCrawlerRequester(t *testing.T) {
	fmt.Println("TestNewCrawlerRequester")
	CrawlerClientHTTPRequestTest()
	fmt.Println()
}

func TestGetCrawlerClientHTTPRequest(t *testing.T) {
	fmt.Println("TestGetCrawlerClientHTTPRequest")
	CrawlerClientHTTPRequestTest()
	fmt.Println()
}

type SetDefaultHeadersStruct struct {
	USER_AGENT string
}

var SetDefaultHeadersMap = map[int]SetDefaultHeadersStruct{
	0: {
		USER_AGENT: "Mozilla/5.0 (X11; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0",
	},
}

func SetDefaultHeadersTest() {
	CrawlerClient.SetDefaultHeaders()

	HTTPRequest := CrawlerClient.GetCrawlerClientHTTPRequest()
	HTTPRequestHeader := HTTPRequest.Header

	for k, v := range SetDefaultHeadersMap {
		USER_AGENT := v.USER_AGENT
		ua := HTTPRequestHeader.Get("User-Agent")
		if USER_AGENT != HTTPRequestHeader.Get("User-Agent") {
			errors.Error.AssertionTestError(k, USER_AGENT, ua)
		}
	}
}

func TestSetDefaultHeaders(t *testing.T) {
	fmt.Println("TestSetDefaultHeaders")
	SetDefaultHeadersTest()
	fmt.Println()
}

func SetDefaultCrawlerRequesterTest() {
	NewCrawlerRequesterTestSliceSize := NewCrawlerRequesterTestSlice.Len()
	for index := 0; index < NewCrawlerRequesterTestSliceSize; index++ {
		v := NewCrawlerRequesterTestSlice.Index(index)
		input := v.FieldByName("input")
		inputMethod := input.FieldByName("Method").String()
		inputURL := input.FieldByName("URL").String()
		inputBody := strings.NewReader(input.FieldByName("Body").String())

		output := v.FieldByName("output")
		outputMethod := output.FieldByName("Method").String()
		outputURL := output.FieldByName("URL").String()
		outputBody := output.FieldByName("Body")
		//outputBodyValue := output.FieldByName("Body").Elem().Elem().Field(0).String()
		//.Type() // .Elem().Type() //.Elem().Type() // .Field(0).Type()
		//newOutputBody := io.NopCloser(strings.NewReader(outputBodyValue))

		CrawlerClient.SetDefaultCrawlerRequester(inputMethod, inputURL, inputBody)
		request := CrawlerClient.GetCrawlerClientHTTPRequest()

		if request.Method != outputMethod {
			errors.Error.AssertionTestError(index, outputMethod, request.Method)
		}
		if request.URL.String() != outputURL {
			errors.Error.AssertionTestError(index, outputURL, request.URL)
		}
		// reflect.TypeOf(request.Body) == reflect.TypeOf(newOutputBody)
		if reflect.DeepEqual(request.Body, outputBody)  {
			errors.Error.AssertionTestError(index, request.Body, outputBody)
		}
	}
}

func TestSetDefaultCrawlerRequester(t *testing.T) {
	fmt.Println("TestSetDefaultCrawlerRequester")
	SetDefaultCrawlerRequesterTest()
	fmt.Println()
}

func CrawlerDoTest() {
	CrawlerClient.NewCrawlerClient()
	CrawlerClient.SetDefaultCrawlerRequester("GET", "https://vivareal.com.br", nil)

	CrawlerClient.CrawlerDo()
}

func TestCrawlerDo(t *testing.T) {
	fmt.Println("TestCrawlerDo")
	CrawlerDoTest()
	fmt.Println()
}