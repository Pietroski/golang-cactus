package file_manager

import (
	"bufio"
	"github.com/Pietroski/golang-VivaRealCrawler/domain/project_root"
	"github.com/Pietroski/golang-notification/errors"
	"os"
	"path/filepath"
	"sync"
)

var (
	FileManager IFileManager = &SFileManager{}
	wg sync.WaitGroup
)

// IFileManager is the FileManager package interface
type IFileManager interface {
	SetOutputDirectory(string)
	GetOutputDirectory() string
	WriteToFile(filename string, content string)
	ReadFromFile(filename string) <-chan string
	GetNumberOfLinesFromFile(filename string) int64
}
// SFileManager is the FileManager package struct
type SFileManager struct {
	outputDirectory string
}

// GetOutputDirectory gets the output directory
// that the files will be written in
func (fm *SFileManager) GetOutputDirectory() string {
	outputPath := fm.outputDirectory

	return outputPath
}

// SetOutputDirectory sets the output directory
// that the files will be written in
func (fm *SFileManager) SetOutputDirectory(IODirectoryPathFromRootPath string) {
	rootPath := project_root.ProjectRoot.GetProjectRootDirectory()
	outputPath := filepath.Join(rootPath, IODirectoryPathFromRootPath)

	fm.outputDirectory = outputPath
}

// WriteToFile writes to the chosen file the given content
func (fm *SFileManager) WriteToFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	errors.Error.CheckFatalError(err)

	_, err = file.WriteString(content)
	errors.Error.CheckFatalError(err)

	closingErr := file.Close()
	errors.Error.CheckFatalError(closingErr, "Error closing the file")
}

// ReadFromFile reads from the chosen file its content
func (fm *SFileManager) ReadFromFile(filename string) <-chan string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	errors.Error.CheckFatalError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	err = scanner.Err()
	errors.Error.CheckFatalError(err)

	fileNumberOfLines := fm.GetNumberOfLinesFromFile(filename) + 1
	lines := make(chan string, fileNumberOfLines)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()
	wg.Wait()
	close(lines)

	return lines
}

// GetNumberOfLinesFromFile return the numbers of line from an specific file
func (fm *SFileManager) GetNumberOfLinesFromFile(filename string) int64 {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	errors.Error.CheckFatalError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	err = scanner.Err()
	errors.Error.CheckFatalError(err)

	var lineCount int64
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
