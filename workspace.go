package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Session struct {
	UUID         string `json:"uuid"`
	Platform     string `json:"platform"`
	Machine      string `json:"machine"`
	Workspace    string `json:"workspace"`
	Created      string `json:"created"`
	LastOpened   string `json:"last_opened"`
	TargetIP     string `json:"target_ip"`
	VPNStatus    string `json:"vpn_status"`
	ReportStatus string `json:"report_status"`
}

func setupWorkspace(ws string) {
	if optDryRun {
		printInfo("[DRY-RUN] Would create workspace: " + ws)
		return
	}

	dirs := []string{"scans", "exploits", "loot", "notes", "screenshots", "credentials"}
	for _, d := range dirs {
		path := filepath.Join(ws, d)
		err := os.MkdirAll(path, 0755)
		if err == nil {
			printSuccess("Created " + path)
		}
	}

	os.MkdirAll(filepath.Join(baseDir, "vpn"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "logs"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "archive"), 0755)

	notesFile := filepath.Join(ws, "notes", "machine.md")
	if _, err := os.Stat(notesFile); os.IsNotExist(err) {
		ip := optTargetIP
		if ip == "" {
			ip = "TBD"
		}
		content := fmt.Sprintf("# Machine: %s\n\n**Platform:** %s  \n**Target IP:** %s  \n**Date:** %s\n\n---\n\n## Overview\n\n## Recon\n\n## Enumeration\n\n## Exploitation\n\n## Privilege Escalation\n\n## Loot\n\n## Flags\n\n| Flag | Value |\n|------|-------|\n| User |       |\n| Root |       |\n\n## Lessons Learned\n\n## References\n", optMachine, optPlatform, ip, timeNow())
		os.WriteFile(notesFile, []byte(content), 0644)
		printSuccess("Created notes template: " + notesFile)
	}

	sessionFile := filepath.Join(ws, ".pwnbox-session")
	sess := Session{
		UUID:      generateUUID(),
		Platform:  optPlatform,
		Machine:   optMachine,
		Workspace: ws,
		Created:   time.Now().Format(time.RFC3339),
	}
	b, _ := json.MarshalIndent(sess, "", "  ")
	os.WriteFile(sessionFile, b, 0644)
	printSuccess("Created session metadata")
}

// timeNow returns a formatted date string for use in templates.
func timeNow() string {
	return time.Now().Format("2006-01-02")
}
