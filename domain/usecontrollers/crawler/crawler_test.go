package crawler

import (
	"fmt"
	"github.com/Pietroski/golang-VivaRealCrawler/data/mocks"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/file_manager"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/project_root"
	"github.com/Pietroski/golang-notification/errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"reflect"
	"testing"
)

const (
	InitialPage = "../../../data/mocks/html_mocks/initial_page_mock.html"
	MiddlePage  = "../../../data/mocks/html_mocks/middle_page_mock.html"
	FinalPage   = "../../../data/mocks/html_mocks/final_page_mock.html"

	ProjectRootDirectoryName = "golang-VivaRealCrawler"
	stdout = "domain/usecontrollers/crawler/test_results/"
)

var (
	TCrawler *SCrawler
)

func init() {
	project_root.ProjectRoot.InitializeProjectRoot(ProjectRootDirectoryName)
	file_manager.FileManager.SetOutputDirectory(stdout)

	outDir := file_manager.FileManager.GetOutputDirectory()
	fmt.Println("Output directory ->", outDir)

	TCrawler = NewCrawler(outDir)
	mockedHTTPResponse := mocks.Mocker.MockHTTPResponse(InitialPage)
	TCrawler.SetDocumentToProcess(mockedHTTPResponse)
}

// NewCrawler tests
type TNewCrawlerStruct struct {
	input  string
	output *SCrawler
}

var TNewCrawlerMap = map[int]TNewCrawlerStruct{
	0: {
		input:  "/domain/usecontrollers/crawler",
		output: &SCrawler{},
	},
	1: {
		input:  "/domain/usecontrollers/crawler",
		output: &SCrawler{},
	},
}

func NewCrawlerTest() {
	for tn, t := range TNewCrawlerMap {
		i := t.input
		o := t.output

		crawler := NewCrawler(i)

		outputType := reflect.TypeOf(o)
		crawlerType := reflect.TypeOf(crawler)

		if outputType != crawlerType {
			errors.Error.AssertionTestError(tn, outputType, crawlerType)
		}
	}
}

func TestNewCrawler(t *testing.T) {
	fmt.Println("TestNewCrawler")
	NewCrawlerTest()
	fmt.Println()
}

// SetDocumentToProcess tests
type TSetDocumentToProcessStruct struct {
	input    *http.Response
	expected reflect.Type
}

var TSetDocumentToProcessMap = map[int]TSetDocumentToProcessStruct{
	0: {
		input:    mocks.Mocker.MockHTTPResponse(InitialPage),
		expected: reflect.TypeOf(&goquery.Document{}),
	},
	1: {
		input:    mocks.Mocker.MockHTTPResponse(MiddlePage),
		expected: reflect.TypeOf(&goquery.Document{}),
	},
	2: {
		input:    mocks.Mocker.MockHTTPResponse(FinalPage),
		expected: reflect.TypeOf(&goquery.Document{}),
	},
}

func SetDocumentToProcessTest() {
	for testNumber, test := range TSetDocumentToProcessMap {
		i := test.input
		o := test.expected

		c := &SCrawler{
			DocumentToProcess: &goquery.Document{},
		}
		returnType := reflect.TypeOf(c.DocumentToProcess)

		c.SetDocumentToProcess(i)

		if c.DocumentToProcess == nil && returnType != o {
			log.Panicf("Failed on test number -> %v", testNumber)
		}
	}
}

func TestSetDocumentToProcess(t *testing.T) {
	fmt.Println("TestSetDocumentToProcess")
	SetDocumentToProcessTest()
	fmt.Println()
}

// GetDocumentToProcess tests
type TGetDocumentToProcessStruct struct {
	input  *http.Response
	output *goquery.Document
}

var TGetDocumentToProcessMap = map[int]TGetDocumentToProcessStruct{
	0: {
		input:    mocks.Mocker.MockHTTPResponse(InitialPage),
		output: &goquery.Document{},
	},
	1: {
		input:    mocks.Mocker.MockHTTPResponse(MiddlePage),
		output: &goquery.Document{},
	},
	2: {
		input:    mocks.Mocker.MockHTTPResponse(FinalPage),
		output: &goquery.Document{},
	},
}

func GetDocumentToProcessTest() {
	for tn, t := range TGetDocumentToProcessMap {
		i := t.input
		o := t.output

		c := &SCrawler{
			DocumentToProcess: &goquery.Document{},
		}
		c.SetDocumentToProcess(i)

		returnType := reflect.TypeOf(c.GetDocumentToProcess())
		outputType := reflect.TypeOf(o)

		if outputType != returnType {
			errors.Error.AssertionTestError(tn, outputType, returnType)
		}

	}
}

func TestGetDocumentToProcess(t *testing.T) {
	fmt.Println("TestGetDocumentToProcess")
	GetDocumentToProcessTest()
	fmt.Println()
}

// WriteBuildingsListFileToCrawl tests
type TWriteBuildingsListFileToCrawlStruct struct {
	input string
}

var TWriteBuildingsListFileToCrawlMap = map[int]TWriteBuildingsListFileToCrawlStruct{
	0: {
		"write_test",
	},
}

func WriteBuildingsListFileToCrawlTest() {
	for tn, t := range TWriteBuildingsListFileToCrawlMap {
		i := t.input

		fmt.Println()
		TCrawler.WriteBuildingsListFileToCrawl(i)
		fmt.Println("Test number ->", tn)
	}
}

func TestWriteBuildingsListFileToCrawl(t *testing.T) {
	fmt.Println("TestWriteBuildingsListFileToCrawl")
	WriteBuildingsListFileToCrawlTest()
	fmt.Println()
}

// transformNextPageQuery Tests
type transformNextPageQueryStruct struct {
	input  string
	output string
}

var transformNextPageQueryMap = map[int]transformNextPageQueryStruct{
	0: {
		input:  "#page=1",
		output: "?page=1",
	},
}

func transformNextPageQueryTest() {
	for tn, t := range transformNextPageQueryMap {
		i := t.input
		o := t.output

		fro := TCrawler.transformNextPageQuery(i)

		if fro != o {
			errors.Error.AssertionTestError(tn, o, fro)
		}
	}
}

func TestTransformNextPageQuery(t *testing.T) {
	fmt.Println("TestTransformNextPageQuery")
	transformNextPageQueryTest()
	fmt.Println()
}
