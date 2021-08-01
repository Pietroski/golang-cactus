package url_builder

import (
	"fmt"
	"testing"
)

var (
	BaseURLArray = [4]string{
		CuritibaBaseURL + "apartamento_residencial/",
		CuritibaBaseURL + "casa_residencial/",
		SaoPauloBaseURL + "apartamento_residencial/",
		SaoPauloBaseURL + "casa_residencial/",
	}
	optionsArray = map[string]string{
		"apartment": "#tipos=apartamento_residencial,cobertura_residencial,flat_residencial,kitnet_residencial",
		"house": "#tipos=casa_residencial,condominio_residencial,sobrado_residencial",
	}
)

type buildURLOptionsStruct struct {
	//
}

var (
	buildURLOptionsMap = map[int]buildURLOptionsStruct{
		//
	}
)

func TestBuildURLs(t *testing.T) {
	fmt.Println("TestBuildURLs")
	fro, _, _ := URLBuilder.BuildURLs()
	fmt.Println(fro)
}
