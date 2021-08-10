package main

import (
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/file_manager"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/project_root"
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/url_builder"
)

const (
	ProjectRootDirectoryName = "golang-VivaRealCrawler"
	IODirectoryPathFromRootPath = "data/out/"
)

var (
	urlBuilder url_builder.Assembler
)

// init initializes controllers and client configurations and/or instances
func init() {
	project_root.ProjectRoot.InitializeProjectRoot(ProjectRootDirectoryName)
	file_manager.FileManager.SetOutputDirectory(IODirectoryPathFromRootPath)

	// urlBuilder = url_builder.NewURLBuilder()
}

// main calls the flux controllers
func main() {
	//
}
