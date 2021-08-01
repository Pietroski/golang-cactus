package main

import (
	"fmt"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/project_root"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/file_manager"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/flux_control_and_sequence_manager"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/url_builder"
)

const (
	ProjectRootDirectoryName = "golang-VivaRealCrawler"
	IODirectoryPathFromRootPath = "data/out/"
)

var (
	baseURLsMap map[string][]string
	optionsURLsMap map[string]string
	totalURLs int
)

// init initializes controllers and client configurations and/or instances
func init() {
	project_root.ProjectRoot.InitializeProjectRoot(ProjectRootDirectoryName)
	file_manager.FileManager.SetOutputDirectory(IODirectoryPathFromRootPath)

	baseURLsMap, optionsURLsMap, totalURLs = url_builder.URLBuilder.BuildURLs()
}

// main calls the flux controllers
func main() {
	delete(baseURLsMap, "houses")
	delete(optionsURLsMap, "houses")
	newBase := []string{baseURLsMap["apartments"][0]}
	totalURLs = 1
	fmt.Println(baseURLsMap, optionsURLsMap, totalURLs, newBase)
	flux_control_and_sequence_manager.CrawlerController.DistributeCrawlers(baseURLsMap, optionsURLsMap, totalURLs)
}
