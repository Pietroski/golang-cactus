package mocks

import "github.com/Pietroski/golang-VivaRealCrawler/domain/usecontrollers/file_manager"

var (
	FileManagerMocker IFileManagerMock = &SFileManager{}
)

type IFileManagerMock interface {
	MockOutputDirectory(mockedOutputDirectory string)
}

type SFileManager struct {
	//
}

func (fmm *SFileManager) MockOutputDirectory(mockedOutputDirectory string) {
	file_manager.FileManager.SetOutputDirectory(mockedOutputDirectory)
}
