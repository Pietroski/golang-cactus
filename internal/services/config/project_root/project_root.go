package project_root

import (
	"fmt"
	"github.com/Pietroski/golang-notification/errors"
	"os"
	"strings"
)

var (
	ProjectRoot IProjectRoot = &SProjectRoot{}
)

type IProjectRoot interface {
	InitializeProjectRoot(projectRootDirectoryName string)
	SetProjectRootDirectoryName(projectRootDirectoryName string)
	GetProjectRootDirectoryName() string
	SetCurrentWorkingDirectory()
	GetCurrentWorkingDirectory() string

	FindAndRetrieveProjectRootDirectory(projectRootDirectoryName string) (directoryPath string, ok bool)

	SetProjectRootDirectory()
	GetProjectRootDirectory() string
}

type SProjectRoot struct {
	projectRootDirectoryName string
	projectRootDirectory     string
	currentWorkingDirectory  string
}

// InitializeProjectRoot finds and establishes the project's root directory
// Obs: the Project Name shall be different from its parent folders
func (pr *SProjectRoot) InitializeProjectRoot(projectRootDirectoryName string) {
	pr.SetProjectRootDirectoryName(projectRootDirectoryName)
	pr.SetCurrentWorkingDirectory()
	pr.SetProjectRootDirectory()
}

// SetProjectRootDirectoryName sets the current root project's name
func (pr *SProjectRoot) SetProjectRootDirectoryName(projectRootDirectoryName string) {
	pr.projectRootDirectoryName = projectRootDirectoryName
}

// GetProjectRootDirectoryName gets the current root project's name
func (pr *SProjectRoot) GetProjectRootDirectoryName() string {
	return pr.projectRootDirectoryName
}

// GetCurrentWorkingDirectory gets the current working directory
func (pr *SProjectRoot) GetCurrentWorkingDirectory() string {
	currentWorkingDirectory := pr.currentWorkingDirectory

	return currentWorkingDirectory
}

// SetCurrentWorkingDirectory sets the current working directory
func (pr *SProjectRoot) SetCurrentWorkingDirectory() {
	currentWorkingDirectory, err := os.Getwd()
	errors.Error.CheckPanicError(err)

	pr.currentWorkingDirectory = currentWorkingDirectory
}

// GetProjectRootDirectory gets the project root directory
func (pr *SProjectRoot) GetProjectRootDirectory() string {
	projectRootDirectory := pr.projectRootDirectory

	return projectRootDirectory
}

// SetProjectRootDirectory sets the current root directory
func (pr *SProjectRoot) SetProjectRootDirectory() {
	projectRootDirectoryName := pr.GetProjectRootDirectoryName()
	projectRootDirectory, ok := pr.FindAndRetrieveProjectRootDirectory(projectRootDirectoryName)
	if !ok {
		fmt.Println("Project root directory not found!!")
		return
	}

	pr.projectRootDirectory = projectRootDirectory
}

// FindAndRetrieveProjectRootDirectory finds and returns the current root directory
func (pr *SProjectRoot) FindAndRetrieveProjectRootDirectory(
	projectRootDirectoryName string,
) (
	directoryPath string,
	ok bool,
) {
	pwd := pr.GetCurrentWorkingDirectory()

	projectRootDirectoryNameStringSize := len(projectRootDirectoryName)
	pwdIndex := strings.Index(pwd, projectRootDirectoryName)
	if pwdIndex == -1 {
		return "", false
	}

	limitIndex := pwdIndex + projectRootDirectoryNameStringSize
	rootDirectoryPath := pwd[:limitIndex] + "/"

	return rootDirectoryPath, true
}
