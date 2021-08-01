package mocks

import (
	"io"
	"strings"
)

const (
	Method = "GET"
	URL    = "https://www.example.com"
)

var (
	ClientMock IClientMock = &SClientMock{}
)

type IClientMock interface {
	MockIOReader() io.Reader
}

type SClientMock struct{}

func (scm *SClientMock) MockIOReader() io.Reader {
	str := "This is a io.Reader Mock"
	r := strings.NewReader(str)

	return r
}
