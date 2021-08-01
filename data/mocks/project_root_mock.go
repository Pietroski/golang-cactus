package mocks

import (
	"github.com/Pietroski/golang-VivaRealCrawler/domain/project_root"
)

var (
	ProjectRootMocker IProjectRootMock = &SProjectRootMocker{}
)

type IProjectRootMock interface {
	MockProjectRoot()
}

type SProjectRootMocker struct {
	//
}

func (prm *SProjectRootMocker) MockProjectRoot() {
	project_root.ProjectRoot.InitializeProjectRoot(ProjectRootDirectoryName)
}
