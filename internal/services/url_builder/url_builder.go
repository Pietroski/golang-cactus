package url_builder

import "github.com/Pietroski/golang-VivaRealCrawler/internal/models/url"

type Assembler interface {
	//
}

type urlBuilder struct {
	BaseSearchURL string
	BaseURLPaths  []url.Paths
	URLsNumber    int
}

func NewURLBuilder(
	baseSearchUrl string,
	baseUrlPaths []url.Paths,
) Assembler {
	return &urlBuilder{
		BaseSearchURL: baseSearchUrl,
		BaseURLPaths:  baseUrlPaths,
		URLsNumber:    0,
	}
}
