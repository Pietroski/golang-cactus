package flux_control_and_sequence_manager

import (
	"fmt"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/client"
	crawler2 "github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/crawler"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/file_manager"
	"io/ioutil"
	"log"
	"runtime"
	"sync"
	"time"
)

const (
	GetMethod = "GET"
)

var (
	CrawlerController ICrawlerController = &SCrawlerController{}

	wg sync.WaitGroup
)

type ICrawlerController interface {
	DistributeCrawlers(baseSearchURLMap map[string][]string, optionsURLMap map[string]string, URLsNumber int)
	CheckMaximumProcesses(wishedGoroutinesNumber int) bool
	CrawlerFlux(URL string, index int,  baseURL string, optionsURL string)
}

type SCrawlerController struct {
	//
}

func (cc *SCrawlerController) DistributeCrawlers(
	baseSearchURLMap map[string][]string, optionsURLMap map[string]string, URLsNumber int,
) {
	cmp := cc.CheckMaximumProcesses(URLsNumber)
	i := 0
	if cmp {
		for keyType, baseSearchURLSlice := range baseSearchURLMap {
			for _, baseSearchURL := range baseSearchURLSlice {
				optionsSearchURL := optionsURLMap[keyType]
				initialURL := baseSearchURL + optionsSearchURL

				wg.Add(1)
				go func(URL string, index int, baseURL string, optionsURL string) {
					defer wg.Done()
					cc.CrawlerFlux(URL, index, baseURL, optionsURL)
				}(initialURL, i, baseSearchURL, optionsSearchURL)
				i++
			}
		}
		wg.Wait()
	}
}

func (cc *SCrawlerController) CheckMaximumProcesses(wishedGoroutinesNumber int) bool {
	mp := runtime.NumCPU()
	if wishedGoroutinesNumber <= mp {
		return true
	}
	return false
}

// CrawlerFlux controls the flux of a single crawler
func (cc *SCrawlerController) CrawlerFlux(URL string, index int,  baseURL string, optionsURL string) {
	client.CrawlerClient.NewCrawlerClient()
	client.CrawlerClient.SetDefaultCrawlerRequester(GetMethod, "https://www.movingimoveis.com.br/venda-pr-curitiba/apartamentos/page/2", nil)

	filename := fmt.Sprintf("viva_real_%v", index)

	outDir := file_manager.FileManager.GetOutputDirectory()
	response := client.CrawlerClient.CrawlerDo()

	for _, cookie := range response.Cookies() {
		fmt.Println(cookie)
	}

	//log.Fatal("Technical stop")

	fmt.Println("--------------------------- INIT HERE --------------------------")
	bs, _ := ioutil.ReadAll(response.Body)
	s := string(bs)
	fmt.Println(s)

	fmt.Println("--------------------------- END HERE --------------------------")
	log.Fatal("Technical stop")

	crawler := crawler2.NewCrawler(outDir)
	crawler.SetURLParameters(baseURL, "", optionsURL)
	crawler.SetDocumentToProcess(response)
	crawler.WriteBuildingsListFileToCrawl(filename)

	for crawler.CheckAndSetNextPageQuery() {
		crawler.SetNextDocumentToProcess()
		crawler.WriteBuildingsListFileToCrawl(filename)
		time.Sleep(time.Second * 8)
	}

	fmt.Println(fmt.Sprintf("Finished process %v ->", index))
}

func SetBuildingsList() {
	//
}

func GetBuildingsList() {
	//
}

func CrawOnBuildingsList() {
	//
}

func LoopOnCrawler() {
	//
}
