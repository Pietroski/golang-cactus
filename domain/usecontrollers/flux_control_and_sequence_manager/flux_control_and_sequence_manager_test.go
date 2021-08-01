package flux_control_and_sequence_manager

import (
	"fmt"
	"testing"
)

func TestCheckMaximumProcesses(t *testing.T) {
	fmt.Println("TestCheckMaximumProcesses")
	CrawlerController.CheckMaximumProcesses(4)
	fmt.Println()
}
