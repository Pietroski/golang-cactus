package project_root

import (
	"fmt"
	"github.com/Pietroski/golang-notification/errors"
	"os"
	"testing"
)

const (
	ProjectRootDirectoryName = "golang-VivaRealCrawler"
	ProjectRootDirectory = "/home/pietroski/go/src/github.com/Pietroski/golang-VivaRealCrawler/"
)

func TestInitializeProjectRoot(t *testing.T) {
	fmt.Println("TestInitializeProjectRoot")
	ProjectRoot.InitializeProjectRoot(ProjectRootDirectoryName)
	fmt.Println(ProjectRoot.GetProjectRootDirectory())
	fmt.Println()
}

func GETSETProjectRootDirectoryNameTest() {
	ProjectRoot.SetProjectRootDirectoryName(ProjectRootDirectoryName)
	prdn := ProjectRoot.GetProjectRootDirectoryName()
	if prdn != ProjectRootDirectoryName {
		errors.Error.AssertionTestError(0, ProjectRootDirectoryName, prdn)
	}
}

func TestSetProjectRootDirectoryName(t *testing.T) {
	fmt.Println("TestSetProjectRootDirectoryName")
	GETSETProjectRootDirectoryNameTest()
	fmt.Println()
}

func TestGetProjectRootDirectoryName(t *testing.T) {
	fmt.Println("TestGetProjectRootDirectoryName")
	GETSETProjectRootDirectoryNameTest()
	fmt.Println()
}

func GETSETCurrentWorkingDirectoryTest() {
	ProjectRoot.SetCurrentWorkingDirectory()
	pwd := ProjectRoot.GetCurrentWorkingDirectory()
	wd, err := os.Getwd()
	errors.Error.CheckFatalError(err)
	if pwd != wd {
		errors.Error.AssertionTestError(0, wd, pwd)
	}
}

func TestSetCurrentWorkingDirectory(t *testing.T) {
	fmt.Println("TestSetCurrentWorkingDirectory")
	GETSETCurrentWorkingDirectoryTest()
	fmt.Println()
}

func TestGetCurrentWorkingDirectory(t *testing.T) {
	fmt.Println("TestGetCurrentWorkingDirectory")
	GETSETCurrentWorkingDirectoryTest()
	fmt.Println()
}

func FindAndRetrieveProjectRootDirectoryTest() {
	dp, ok := ProjectRoot.FindAndRetrieveProjectRootDirectory(ProjectRootDirectoryName)
	if !ok {
		errors.Error.AssertionTestError(0, true, ok)
	}

	if ProjectRootDirectory != dp {
		errors.Error.AssertionTestError(1, ProjectRootDirectory, dp)
	}
}

func TestFindAndRetrieveProjectRootDirectory(t *testing.T) {
	fmt.Println("TestFindAndRetrieveProjectRootDirectory")
	FindAndRetrieveProjectRootDirectoryTest()
	fmt.Println()
}

func GETSETProjectRootDirectoryTest() {
	ProjectRoot.SetProjectRootDirectory()
	dp := ProjectRoot.GetProjectRootDirectory()

	if ProjectRootDirectory != dp {
		errors.Error.AssertionTestError(1, ProjectRootDirectory, dp)
	}
}

func TestGetProjectRootDirectory(t *testing.T) {
	fmt.Println("TestGetProjectRootDirectory")
	GETSETProjectRootDirectoryTest()
	fmt.Println()
}

func TestSetProjectRootDirectory(t *testing.T) {
	fmt.Println("TestSetProjectRootDirectory")
	GETSETProjectRootDirectoryTest()
	fmt.Println()
}

func GETSETProjectRootDirectoryTestFail() {
	ProjectRoot.SetProjectRootDirectoryName("goland-VivaRealCrawler")
	ProjectRoot.FindAndRetrieveProjectRootDirectory("goland-VivaRealCrawler")
	ProjectRoot.SetProjectRootDirectory()
	dp := ProjectRoot.GetProjectRootDirectory()

	if ProjectRootDirectory != dp {
		errors.Error.AssertionTestError(1, ProjectRootDirectory, dp)
	}
}

func TestGetProjectRootDirectoryFail(t *testing.T) {
	fmt.Println("TestGetProjectRootDirectory")
	GETSETProjectRootDirectoryTestFail()
	fmt.Println()
}

func TestSetProjectRootDirectoryFail(t *testing.T) {
	fmt.Println("TestSetProjectRootDirectory")
	GETSETProjectRootDirectoryTestFail()
	fmt.Println()
}
