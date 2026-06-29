package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	initPaths()
	parseFlags()

	// --version
	if optVersion {
		fmt.Printf("pwnbox-starter v%s\n", version)
		os.Exit(0)
	}

	// -d disconnect does not show banner
	if optDisconnect {
		disconnectVPN()
		return
	}

	printBanner()
	checkDependencies()
	fmt.Println()

	// No flags → interactive mode
	if flag.NFlag() == 0 {
		interactiveMode()
	}

	if optPlatform == "" || optMachine == "" {
		printError("Platform (-p) and Machine (-m) are required.")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	if err := validateInput(); err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	workspace := getWorkspacePath()

	// Configuration summary
	fmt.Println()
	printHeader("Configuration Summary")
	printDivider()
	fmt.Printf("  %-14s %s\n", "Platform:", optPlatform)
	fmt.Printf("  %-14s %s\n", "Machine:", optMachine)
	fmt.Printf("  %-14s %s\n", "Workspace:", workspace)
	if optTargetIP != "" {
		fmt.Printf("  %-14s %s\n", "Target IP:", optTargetIP)
	}
	if optVPNURL != "" {
		fmt.Printf("  %-14s %s\n", "VPN URL:", optVPNURL)
	}
	fmt.Printf("  %-14s %t\n", "Connect VPN:", optConnect)
	fmt.Printf("  %-14s %t\n", "TMUX:", optTmux)
	fmt.Printf("  %-14s %t\n", "Report:", optReport)
	fmt.Printf("  %-14s %t\n", "Dry Run:", optDryRun)
	printDivider()
	fmt.Println()

	setupWorkspace(workspace)

	if optVPNURL != "" {
		downloadVPN()
	}

	if optConnect {
		connectVPN()
	}

	if optTmux {
		launchTmux(workspace)
	}

	if optReport {
		generateReport(workspace)
	}

	fmt.Println()
	printHeader("Completed Successfully")
	printDivider()
	fmt.Printf("  %-14s %s\n", "Workspace:", workspace)
	fmt.Printf("  %-14s %s/notes/machine.md\n", "Notes:", workspace)
	if optReport {
		fmt.Printf("  %-14s %s/report.md\n", "Report:", workspace)
	}
	if optTmux {
		fmt.Printf("  %-14s pwnbox-%s-%s\n", "TMUX Session:", optPlatform, optMachine)
	}
	printDivider()
}
