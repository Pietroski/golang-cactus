package url_builder

var (
	URLBuilder IURLBuilder = &SURLBuilder{}
)

type IURLBuilder interface {
	BuildURLs() (baseSearchURLMap map[string][]string, optionsURLMap map[string]string, URLsNumber int)
	SetBaseSearchURLSlice(baseSearchURLSlice map[string][]string)
	GetBaseSearchURLSlice() (baseSearchURLSlice map[string][]string)
	SetOptionsURL(optionsURL map[string]string)
	GetOptionsURL() map[string]string
	SetURLsNumber(n int)
	GetURLsNumber() int
}

type SURLBuilder struct {
	baseSearchURLSlice map[string][]string
	optionsURL         map[string]string
	URLsNumber         int
}

func (u *SURLBuilder) BuildURLs() (baseSearchURLMap map[string][]string, optionsURLMap map[string]string, URLsNumber int) {
	_baseSearchURLMap := u.buildSearchURLs()
	_optionsURLMap := u.buildOptionsURL()
	_URLsNumber := u.URLsNumber

	u.SetBaseSearchURLSlice(_baseSearchURLMap)
	u.SetOptionsURL(_optionsURLMap)

	return _baseSearchURLMap, _optionsURLMap, _URLsNumber
}

func (u *SURLBuilder) SetBaseSearchURLSlice(baseSearchURLSlice map[string][]string) {
	u.baseSearchURLSlice = baseSearchURLSlice
}

func (u *SURLBuilder) GetBaseSearchURLSlice() (baseSearchURLSlice map[string][]string) {
	return u.baseSearchURLSlice
}

func (u *SURLBuilder) SetOptionsURL(optionsURL map[string]string) {
	u.optionsURL = optionsURL
}

func (u *SURLBuilder) GetOptionsURL() map[string]string {
	return u.optionsURL
}

func (u *SURLBuilder) SetURLsNumber(n int) {
	u.URLsNumber = n
}

func (u *SURLBuilder) GetURLsNumber() int {
	return u.URLsNumber
}
