package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func downloadVPN() {
	if optDryRun {
		printInfo("[DRY-RUN] Would process VPN from: " + optVPNURL)
		return
	}
	dest := filepath.Join(baseDir, "vpn", optPlatform+".ovpn")
	if _, err := os.Stat(dest); err == nil {
		printInfo("VPN config already exists at " + dest)
		return
	}

	// If the user provided a local file path, copy it
	if _, err := os.Stat(optVPNURL); err == nil {
		printInfo("Copying local VPN config...")
		input, err := os.ReadFile(optVPNURL)
		if err != nil {
			printError("Failed to read local VPN file: " + err.Error())
			return
		}
		if err := os.WriteFile(dest, input, 0600); err != nil {
			printError("Failed to copy VPN file to " + dest)
		} else {
			printSuccess("VPN copied to " + dest)
		}
		return
	}

	// Otherwise, treat it as a URL and download it
	printInfo("Downloading VPN config...")
	cmd := exec.Command("curl", "-#", "-L", "-o", dest, optVPNURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		printError("Failed to download VPN config from " + optVPNURL)
	} else {
		printSuccess("VPN downloaded to " + dest)
	}
}

func connectVPN() {
	if optDryRun {
		printInfo("[DRY-RUN] Would connect VPN")
		return
	}
	cfg := filepath.Join(baseDir, "vpn", optPlatform+".ovpn")
	if _, err := os.Stat(cfg); os.IsNotExist(err) {
		printError("VPN config not found: " + cfg)
		return
	}

	printInfo("Connecting to VPN...")

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "openvpn", "--config", cfg)
		if err := cmd.Start(); err != nil {
			printError("Failed to start OpenVPN in new window.")
		} else {
			printSuccess("VPN launched in a new window.")
		}
		return
	} else if runtime.GOOS == "darwin" {
		script := fmt.Sprintf(`tell app "Terminal" to do script "sudo openvpn --config '%s'"`, cfg)
		cmd := exec.Command("osascript", "-e", script)
		if err := cmd.Start(); err != nil {
			printError("Failed to start OpenVPN in new window.")
		} else {
			printSuccess("VPN launched in a new window.")
		}
		return
	}

	// For Linux, try to launch a new terminal emulator window first
	termCmds := [][]string{
		{"x-terminal-emulator", "-e", fmt.Sprintf("sudo openvpn --config %s", cfg)},
		{"gnome-terminal", "--", "sudo", "openvpn", "--config", cfg},
		{"konsole", "-e", "sudo", "openvpn", "--config", cfg},
		{"xfce4-terminal", "-e", fmt.Sprintf("sudo openvpn --config %s", cfg)},
	}

	for _, tCmd := range termCmds {
		cmd := exec.Command(tCmd[0], tCmd[1:]...)
		if err := cmd.Start(); err == nil {
			printSuccess("VPN launched in a new " + tCmd[0] + " window.")
			return
		}
	}

	// Fallback to daemon mode if no terminal emulator is found or if headless
	pidFile := fmt.Sprintf("/tmp/openvpn-%s.pid", optPlatform)
	cmd := exec.Command("sudo", "openvpn", "--config", cfg, "--daemon", "--writepid", pidFile)
	if err := cmd.Run(); err != nil {
		printError("Failed to start OpenVPN. Check your privileges or configuration.")
		return
	}

	for i := 0; i < 15; i++ {
		out, _ := exec.Command("ip", "addr", "show", "tun0").Output()
		if strings.Contains(string(out), "inet") {
			printSuccess("VPN Connected on tun0")
			return
		}
		time.Sleep(1 * time.Second)
	}
	printError("VPN connection timed out. Could not find tun0.")
}

func disconnectVPN() {
	printInfo("Disconnecting VPN...")
	files, _ := filepath.Glob("/tmp/openvpn-*.pid")
	for _, f := range files {
		b, _ := os.ReadFile(f)
		pid := strings.TrimSpace(string(b))
		if pid != "" {
			exec.Command("sudo", "kill", pid).Run()
			os.Remove(f)
			printSuccess("Killed VPN process " + pid)
		}
	}
}
