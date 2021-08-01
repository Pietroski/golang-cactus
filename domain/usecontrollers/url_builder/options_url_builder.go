package url_builder

import (
	"github.com/Pietroski/golang-notification/errors"
	"strings"
)

const (
	rawOptionsType = "#tipos="
)

var (
	optionsApartmentSearchSlice = []string{
		"apartamento_residencial",
		"cobertura_residencial",
		"flat_residencial",
		"kitnet_residencial",
	}
	optionsHouseSearchSlice = []string{
		"casa_residencial",
		"condominio_residencial",
		"sobrado_residencial",
	}

	optionsMap = map[string][]string{
		"apartments": optionsApartmentSearchSlice,
		"houses": optionsHouseSearchSlice,
	}
)

func (u *SURLBuilder) buildOptionsURL() map[string]string {
	optionsMapKeyQuantity := len(optionsMap)
	var _optionsMap = make(map[string]string, optionsMapKeyQuantity)

	for keyValue, optionsSlice := range optionsMap {
		var optionsType strings.Builder
		s := u.optionsURLsBuilder(&optionsType, rawOptionsType, optionsSlice)
		_optionsMap[keyValue] = s
	}

	return _optionsMap
}

func (u *SURLBuilder) optionsURLsBuilder(
	optionsType *strings.Builder, baseOptionString string, optionsSlice []string,
) string {
	_, err := optionsType.WriteString(baseOptionString)
	errors.Error.CheckFatalError(err)

	concatenatedOptionsSlice := strings.Join(optionsSlice, ",")

	_, err = optionsType.WriteString(concatenatedOptionsSlice)
	errors.Error.CheckFatalError(err)

	optionsTypeString := optionsType.String()
	return optionsTypeString
}
