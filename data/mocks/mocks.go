package mocks

import (
	"bytes"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/project_root"
	"github.com/Pietroski/golang-notification/errors"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	ProjectRootDirectoryName = "golang-VivaRealCrawler"
	IODirectoryPathFromRootPath = "data/out/"

	InitialPage = "../../../data/mocks/initial_page_mock.html"
	MiddlePage = "../../../data/mocks/middle_page_mock.html"
	FinalPage = "../../../data/mocks/final_page_mock.html"
)

var (
	Mocker IMock = &mocker{}
)

type IMock interface {
	MockProjectRoot()
	MockFileGenerator()

	MockHTTPResponse(fileToMock string) *http.Response

	MockGoqueryDocument()
}

type mocker struct {
	//
}

func (m *mocker) MockProjectRoot() {
	project_root.ProjectRoot.InitializeProjectRoot(ProjectRootDirectoryName)
}

func (m *mocker) MockFileGenerator() {
	//file_manager.FileManager.SetOutputDirectory(IODirectoryPathFromRootPath)
}

func (m *mocker) MockHTTPResponse(fileToMock string) *http.Response {
	HTMLFile, err := ioutil.ReadFile(fileToMock)
	errors.Error.CheckFatalError(err)

	mockedBody := io.NopCloser(bytes.NewReader(HTMLFile))

	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body: mockedBody,
	}
	return httpResponse
}

func (m *mocker) MockGoqueryDocument() {
	//
}
