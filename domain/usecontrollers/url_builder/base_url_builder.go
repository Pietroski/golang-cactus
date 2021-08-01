package url_builder

import (
	"fmt"
)

const (
	CuritibaBaseURL = "https://www.vivareal.com.br/venda/parana/curitiba"
	SaoPauloBaseURL = "https://www.vivareal.com.br/venda/sp/sao-paulo"
)

var (
	baseURLSlice = []string{
		CuritibaBaseURL, SaoPauloBaseURL,
	}

	baseSearchSlice = []string{
		"",
		"",
	}
)

func (u *SURLBuilder) buildSearchURLs() map[string][]string {
	baseSearchURLsBuilderSlice := u.baseSearchURLsBuilder(baseURLSlice, baseSearchSlice)

	var apartments []string
	var houses []string

	for i, v := range baseSearchURLsBuilderSlice {
		if i % 2 == 0 {
			apartments = append(apartments, v)
		} else {
			houses = append(houses, v)
		}
	}

	var baseSearchURLsMap = map[string][]string{
		"apartments": apartments,
		"houses": houses,
	}

	return baseSearchURLsMap
}

func (u *SURLBuilder) baseSearchURLsBuilder(baseURLSliceParameter []string, baseSearchSliceParameter []string) []string {
	baseSearchURLSliceSize := u.CountURLsNumber(baseURLSliceParameter, baseSearchSliceParameter)
	u.SetURLsNumber(baseSearchURLSliceSize)
	var baseSearchURLSlice = make([]string, baseSearchURLSliceSize)

	sliceIndex := 0
	for _, base := range baseURLSlice {
		for _, searchBase := range baseSearchSlice {
			baseURL := fmt.Sprintf("%v/%v", base, searchBase)
			baseSearchURLSlice[sliceIndex] = baseURL
			sliceIndex++
		}
	}

	return baseSearchURLSlice
}

func (u *SURLBuilder) CountURLsNumber(URLsSlices ...[]string) int {
	var counter = 1

	for _, URLSlice := range URLsSlices {
		counter *= len(URLSlice)
	}

	return counter
}