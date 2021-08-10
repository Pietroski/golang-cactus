package crawler

import (
	"github.com/Pietroski/golang-VivaRealCrawler/internal/services/config/file_manager"
	"github.com/PuerkitoBio/goquery"
)

type ICrawler interface {
	//
}

type crawler struct {
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

func NewCrawler(stdout string, fm file_manager.IFileManager) *crawler {
	fm.SetOutputDirectory(stdout)

	cs := &crawler{
		OutputFileClient: fm,
		OutputFilePath:   stdout,
	}

	return cs
}

