package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// clOut is the output writer for all banner/display output.
var clOut io.Writer = os.Stdout

// ANSI color codes — disabled automatically when not a terminal or NO_COLOR=1
var (
	clReset  = "\033[0m"
	clBold   = "\033[1m"
	clCyan   = "\033[1;36m"
	clGreen  = "\033[1;32m"
	clYellow = "\033[1;33m"
	clRed    = "\033[1;31m"
	clDim    = "\033[2m"
)

func init() {
	if os.Getenv("NO_COLOR") == "1" || !isTerminal() {
		clReset, clBold, clCyan, clGreen, clYellow, clRed, clDim = "", "", "", "", "", "", ""
	}
}

// isTerminal reports whether stdout is a terminal.
func isTerminal() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func printInfo(msg string)    { fmt.Printf("%s[*]%s %s\n", clCyan, clReset, msg) }
func printSuccess(msg string) { fmt.Printf("%s[+]%s %s\n", clGreen, clReset, msg) }
func printWarning(msg string) { fmt.Printf("%s[!]%s %s\n", clYellow, clReset, msg) }
func printError(msg string)   { fmt.Fprintf(os.Stderr, "%s[-]%s %s\n", clRed, clReset, msg) }
func printHeader(msg string)  { fmt.Printf("\n%s%s%s\n", clBold, msg, clReset) }
func printDivider()           { fmt.Printf("%s%s%s\n", clDim, "------------------------------------------------------------", clReset) }

// generateUUID produces a random RFC 4122-style UUID.
func generateUUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// browseFile opens a native file dialog (Windows, macOS, or Linux) to select an .ovpn file.
func browseFile() string {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "-Sta", "-NoProfile", "-Command", `
Add-Type -AssemblyName System.Windows.Forms;
$f = New-Object System.Windows.Forms.OpenFileDialog;
$f.Filter = "OpenVPN Config (*.ovpn)|*.ovpn|All Files (*.*)|*.*";
$f.Title = "Select VPN Config";
$f.ShowHelp = $true;
$f.Multiselect = $false;
[void]$f.ShowDialog();
if ($f.FileName) { Write-Output $f.FileName }
`)
		out, err := cmd.Output()
		if err == nil {
			return strings.TrimSpace(string(out))
		}
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("osascript", "-e", `POSIX path of (choose file with prompt "Select VPN Config" of type {"ovpn"})`)
		out, err := cmd.Output()
		if err == nil {
			return strings.TrimSpace(string(out))
		}
	} else {
		cmd := exec.Command("zenity", "--file-selection", "--file-filter=*.ovpn", "--title=Select VPN Config")
		out, err := cmd.Output()
		if err == nil {
			return strings.TrimSpace(string(out))
		}
	}
	return ""
}
