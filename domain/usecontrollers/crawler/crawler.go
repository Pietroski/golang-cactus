package crawler

import (
	"fmt"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/client"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/file_manager"
	"github.com/Pietroski/golang-notification/errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

// NewCrawler Return a new crawler instance
// attached to its own file manager instance
func NewCrawler(stdout string) *SCrawler {
	var fm file_manager.IFileManager
	fm = &file_manager.SFileManager{}
	fm.SetOutputDirectory(stdout)

	cs := &SCrawler{
		OutputFileClient: fm,
		OutputFilePath:   stdout,
	}

	return cs
}

// ICrawler crawler interface
type ICrawler interface {
	SetDocumentToProcess(response *http.Response)
	WriteBuildingsListFileToCrawl(filename string)
	CheckAndSetNextPageQuery() bool
	transformNextPageQuery(rawQuery string) string
	SetNextDocumentToProcess()
	CreateNextPageURL()
	SetURLParameters(baseURL string, queryURL string, optionsURL string)
	GetURLParameters() (baseURL string, queryURL string, optionsURL string)

	SetBaseURL(baseURL string)
	GetBaseURL() string
	SetQueryURL(query string)
	GetQueryURL() string
	SetOptionsURL(optionsParameter string)
	GetOptionsURL() string
	SetNextPageURL()
	GetNextPageURL()
}

// SCrawler crawler struct
type SCrawler struct {
	baseURL    string
	queryURL   string
	optionsURL string

	DocumentToProcess *goquery.Document

	BuildingsListURL []string
	CurrentPageURL   string
	NextPageURL      string

	OutputFileClient file_manager.IFileManager
	OutputFilePath   string
}

// SetDocumentToProcess set the current parsed document to be processed
func (c *SCrawler) SetDocumentToProcess(response *http.Response) {
	document, err := goquery.NewDocumentFromReader(response.Body)
	errors.Error.CheckPanicError(err)

	c.DocumentToProcess = document
}

// GetDocumentToProcess set the current parsed document to be processed
func (c *SCrawler) GetDocumentToProcess() *goquery.Document {
	return c.DocumentToProcess
}

// WriteBuildingsListFileToCrawl writes out to the
// specified file the list of buildings to crawl
func (c *SCrawler) WriteBuildingsListFileToCrawl(filename string) {
	document := c.GetDocumentToProcess()
	document.Find(".property-card__content-link").Each(func(i int, s *goquery.Selection) {
		value, ok := s.Attr("href")
		if !ok {
			fmt.Print("Attribute not found!!")
			return
		}

		stringLine := fmt.Sprintf("%v\n", value)

		c.OutputFileClient.WriteToFile(c.OutputFilePath+"/"+filename, stringLine)
		//fmt.Println("HERE ->", value)
	})
}

// CheckAndSetNextPageQuery Check if there is a next page to get and return a boolean. Also,
// if there is, it sets the existing URL to NextPage variable;
// if there is not, it sets an empty string to NextPage variable
func (c *SCrawler) CheckAndSetNextPageQuery() bool {
	possibleNextPage := c.DocumentToProcess.Find(".js-change-page").Last()
	query, ok := possibleNextPage.Attr("href")
	if !ok && query == "" {
		return false
	}

	query = c.transformNextPageQuery(query)
	c.SetQueryURL(query)
	return true
}

// transformNextPageQuery transforms the the URL query from
// the retrieved raw format to the actual needed format
func (c *SCrawler) transformNextPageQuery(rawQuery string) string {
	newQuery := strings.Replace(rawQuery, "#", "?", 1)
	return newQuery
}

// SetNextDocumentToProcess sets the next HTML page document to be processed
func (c *SCrawler) SetNextDocumentToProcess() {
	c.SetNextPageURL()
	newURL := c.GetNextPageURL()

	fmt.Println(newURL)

	client.CrawlerClient.SetDefaultCrawlerRequester("GET", newURL, nil)
	response := client.CrawlerClient.CrawlerDo()

	c.SetDocumentToProcess(response)

	closingError := response.Body.Close()
	errors.Error.CheckPanicError(closingError, "error closing response body connection")
}

// CreateNextPageURL creates the next HTML page URL
func (c *SCrawler) CreateNextPageURL() string {
	var newURL string

	if c.queryURL != "" {
		newURL = c.baseURL + c.queryURL + c.optionsURL
	} else {
		newURL = c.baseURL + "/" + c.queryURL + c.optionsURL
	}

	return newURL
}

// SetURLParameters sets baseURL, queryURL and optionsURL that compose the search URL
func (c *SCrawler) SetURLParameters(baseURL string, queryURL string, optionsURL string) {
	c.SetBaseURL(baseURL)
	c.SetQueryURL(queryURL)
	c.SetOptionsURL(optionsURL)
}

// GetURLParameters gets baseURL, queryURL and optionsURL that compose the search URL
func (c *SCrawler) GetURLParameters() (baseURL string, queryURL string, optionsURL string) {
	_baseURL := c.GetBaseURL()
	_queryURL := c.GetQueryURL()
	_optionsURL := c.GetOptionsURL()
	return _baseURL, _queryURL, _optionsURL
}

// SetBaseURL sets base search URL
func (c *SCrawler) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// GetBaseURL gets base search URL
func (c *SCrawler) GetBaseURL() string {
	return c.baseURL
}

// SetQueryURL sets query parameter URL
func (c *SCrawler) SetQueryURL(query string) {
	c.queryURL = query
}

// GetQueryURL gets query parameter URL
func (c *SCrawler) GetQueryURL() string {
	return c.queryURL
}

// SetOptionsURL sets URL search options
func (c *SCrawler) SetOptionsURL(optionsParameter string) {
	c.optionsURL = optionsParameter
}

// GetOptionsURL gets URL search options
func (c *SCrawler) GetOptionsURL() string {
	return c.optionsURL
}

// SetNextPageURL sets the next HTML page URL
func (c *SCrawler) SetNextPageURL() {
	c.NextPageURL = c.CreateNextPageURL()
}

// GetNextPageURL gets the next HTML page URL
func (c *SCrawler) GetNextPageURL() string {
	nextPageURL := c.NextPageURL
	return nextPageURL
}

func GetBuildingPage() {
	//
}

func GetBuildingInfo() {
	//
}
