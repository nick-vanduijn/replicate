package replicate

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const Version = "0.1.0"

func PrintHeader(msg string) {
	fmt.Println("┌─────────────────────────────────────┐")
	fmt.Printf("│  🚀 %-29s │\n", msg)
	fmt.Println("└─────────────────────────────────────┘")
}

func PrintStep(status string, msg string, timeStr string) {
	fmt.Printf("  %s %s...         %s\n", status, msg, timeStr)
}

func PrintProgress(msg string, progress int) {
	bar := strings.Repeat("█", progress/10) + strings.Repeat("░", 10-progress/10)
	fmt.Printf("  ↻ %s...          [%s] %d%%\n", msg, bar, progress)
}

func PrintSummary() {
	fmt.Println()
	fmt.Printf("Environment: %s\n", runtime.GOOS)
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Region: %s\n", "local")
}

func timeAndPrintStep(status string, msg string, start time.Time) {
	elapsed := time.Since(start)
	PrintStep(status, msg, fmt.Sprintf("%.1fs", elapsed.Seconds()))
}

func CheckHomebrew() (int, error) {
	PrintHeader("Homebrew Check Started")

	start := time.Now()
	if _, err := exec.LookPath("brew"); err != nil {
		timeAndPrintStep("⨯", "Checking for Homebrew", start)
		PrintSummary()
		return 1, fmt.Errorf("brew not found in PATH")
	}
	timeAndPrintStep("✓", "Checking for Homebrew", start)

	start = time.Now()
	out, err := exec.Command("brew", "--version").CombinedOutput()
	if err != nil {
		timeAndPrintStep("⨯", "Running version check", start)
		PrintSummary()
		return 2, fmt.Errorf("failed to run `brew --version`: %v", err)
	}
	timeAndPrintStep("✓", "Running version check", start)

	fmt.Println(string(out))
	PrintSummary()
	return 0, nil
}

func RunInstallScript() (int, error) {
	PrintHeader("Homebrew Install Started")

	start := time.Now()
	if runtime.GOOS != "darwin" {
		timeAndPrintStep("⨯", "Checking OS compatibility", start)
		PrintSummary()
		return 1, fmt.Errorf("homebrew installation is only supported on macOS")
	}
	timeAndPrintStep("✓", "Checking OS compatibility", start)

	start = time.Now()
	exePath, err := os.Executable()
	var scriptPath string
	if err != nil {
		return 2, fmt.Errorf("failed to get executable path: %v", err)
	}
	scriptPath = filepath.Join(filepath.Dir(exePath), "install_homebrew.sh")

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		scriptPath = "./scripts/install_homebrew.sh"
	}

	if _, err := os.Stat(scriptPath); err != nil {
		timeAndPrintStep("⨯", "Locating install script", start)
		PrintSummary()
		return 2, fmt.Errorf("install script not found: %s", scriptPath)
	}
	timeAndPrintStep("✓", "Locating install script", start)

	start = time.Now()
	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		timeAndPrintStep("⨯", "Running install script", start)
		PrintSummary()
		return 3, fmt.Errorf("install script failed: %v", err)
	}

	timeAndPrintStep("✓", "Running install script", start)

	PrintSummary()
	return 0, nil
}