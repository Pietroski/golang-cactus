package file_manager

import (
	"fmt"
	"github.com/Pietroski/golang-VivaRealCrawler/data/mocks"
	"testing"
)

const (
	IODirectoryPathFromRootPath = "data/out"
)

func TestProjectRootMocker(t *testing.T) {
	fmt.Println("TestProjectRootMocker")
	fmt.Println("file_manager package initialization mock")
	mocks.Mocker.MockProjectRoot()
	fmt.Println()
}

func TestSetOutputDirectory(t *testing.T) {
	fmt.Println("TestSetOutputDirectory")
	FileManager.SetOutputDirectory(IODirectoryPathFromRootPath)
	fmt.Println(FileManager.GetOutputDirectory())
	fmt.Println()
}

func TestGetOutputDirectory(t *testing.T) {
	fmt.Println("TestGetOutputDirectory")
	FileManager.SetOutputDirectory(IODirectoryPathFromRootPath)
	fmt.Println(FileManager.GetOutputDirectory())
	fmt.Println()
}

func TestWriteToFile(t *testing.T) {
	fmt.Println("TestWriteToFile")
	outDir := FileManager.GetOutputDirectory()
	outDirFinal := outDir + "/" + "test.txt"
	fmt.Println(outDirFinal)
	FileManager.WriteToFile(outDirFinal, "\nHello World!!")
	fmt.Println()
}

func TestReadFromFile(t *testing.T) {
	fmt.Println("TestReadFromFile")
	outDir := FileManager.GetOutputDirectory()
	outDirFinal := outDir + "/" + "test.txt"
	fmt.Println(outDirFinal)
	test := FileManager.ReadFromFile(outDirFinal)

	fmt.Println(test)

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for t := range test {
			i++
			fmt.Println(i, "->", t)
		}
	}()
	wg.Wait()

	fmt.Println()
}

func TestGenericStaticFunctionality(t *testing.T) {
	fmt.Println("TestGenericStaticFunctionality")
	od := FileManager.GetOutputDirectory()
	fmt.Println("Static FileManager output directory ->", od)
	fm := &SFileManager{}
	fm_od := fm.GetOutputDirectory()
	fmt.Println("Initial instance's FileManager output directory ->", fm_od)
	fm.SetOutputDirectory(IODirectoryPathFromRootPath)
	fm_od = fm.GetOutputDirectory()
	fmt.Println("Final instance's FileManager output directory ->", fm_od)
	fmt.Println()
}
