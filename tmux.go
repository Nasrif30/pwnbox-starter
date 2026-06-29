package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func launchTmux(ws string) {
	if optDryRun {
		printInfo("[DRY-RUN] Would launch TMUX session")
		return
	}

	sessionName := fmt.Sprintf("pwnbox-%s-%s", optPlatform, optMachine)

	err := exec.Command("tmux", "has-session", "-t", sessionName).Run()
	if err == nil {
		printInfo("Tmux session already exists.")
		printSuccess("Attach using: tmux attach -t " + sessionName)
		return
	}

	exec.Command("tmux", "new-session", "-d", "-s", sessionName, "-c", ws).Run()
	exec.Command("tmux", "rename-window", "-t", sessionName+":0", "shell").Run()
	exec.Command("tmux", "new-window", "-t", sessionName, "-n", "nmap", "-c", filepath.Join(ws, "scans")).Run()
	exec.Command("tmux", "new-window", "-t", sessionName, "-n", "gobuster", "-c", filepath.Join(ws, "exploits")).Run()
	exec.Command("tmux", "new-window", "-t", sessionName, "-n", "notes", "-c", filepath.Join(ws, "notes")).Run()
	exec.Command("tmux", "select-window", "-t", sessionName+":0").Run()

	fmt.Printf("\n%s%s%s\n", clCyan, "=========================================", clReset)
	fmt.Printf("%s  TMUX Session Ready%s\n", clBold, clReset)
	fmt.Printf("  Attach: %stmux attach -t %s%s\n", clGreen, sessionName, clReset)
	fmt.Printf("%s%s%s\n\n", clCyan, "=========================================", clReset)
}
