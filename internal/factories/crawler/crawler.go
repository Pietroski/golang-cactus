package crawler

var (
	Factory ICrawlerFactory = &crawlerFactory{}
)

type ICrawlerFactory interface {
	//
}

type crawlerFactory struct {
	//
}
