package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func generateReport(ws string) {
	if optDryRun {
		printInfo("[DRY-RUN] Would generate report")
		return
	}

	reportFile := filepath.Join(ws, "report.md")
	if _, err := os.Stat(reportFile); err == nil {
		printInfo("Report already exists. Skipping.")
		return
	}

	fmt.Printf("\n%s  [ Report Generator ]%s\n", clBold, clReset)
	printDivider()
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("  %sEnter Target OS%s (Linux/Windows, or press Enter to skip): ", clDim, clReset)
	targetOS, _ := reader.ReadString('\n')
	targetOS = strings.TrimSpace(targetOS)

	fmt.Printf("  %sEnter Difficulty%s (Easy/Medium/Hard, or press Enter to skip): ", clDim, clReset)
	difficulty, _ := reader.ReadString('\n')
	difficulty = strings.TrimSpace(difficulty)

	fmt.Printf("  %sEnter Initial Notes%s (Press Enter to skip): ", clDim, clReset)
	initialNotes, _ := reader.ReadString('\n')
	initialNotes = strings.TrimSpace(initialNotes)

	content := fmt.Sprintf("# Penetration Test Report - %s\n\n", optMachine)
	content += fmt.Sprintf("**Platform:** %s\n", optPlatform)
	content += fmt.Sprintf("**Target IP:** %s\n", optTargetIP)
	if targetOS != "" {
		content += fmt.Sprintf("**OS:** %s\n", targetOS)
	}
	if difficulty != "" {
		content += fmt.Sprintf("**Difficulty:** %s\n", difficulty)
	}
	content += fmt.Sprintf("**Date:** %s\n\n---\n\n", time.Now().Format("2006-01-02"))

	if initialNotes != "" {
		content += fmt.Sprintf("## Initial Notes\n%s\n\n", initialNotes)
	}

	content += "## Summary\n\n## Enumeration\n\n## Exploitation\n\n## Post-Exploitation\n"
	os.WriteFile(reportFile, []byte(content), 0644)
	printSuccess("Report saved: " + reportFile)
}
