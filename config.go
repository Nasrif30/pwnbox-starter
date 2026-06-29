package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

const version = "2.0"

// printBanner renders PWNSTRT as compact word-art (68 cols, 4px wide letters, 2-char cells).
func printBanner() {
	// Fill chars per letter: P=PW  W=WN  N=NB  S=ST  T=TR  R=RT  T2=TM  gap=2sp
	// Each letter block is 8 chars (4 cells × 2 chars). Total: 7×8 + 6×2 = 68 cols.
	rows := []struct{ c, t string }{
		{"\033[2;36m", ` pwnbox-starter | pwn . start . root . shell . kali . htb`},
		// row 0: top bars / caps
		{"\033[1;36m", `PWPWPW    WN    WN  NB    NB    STSTST  TRTRTRTR  RTRTRT    TMTMTMTM`},
		// row 1: upper body
		{"\033[1;36m", `PW    PW  WN    WN  NBNB  NB  ST          TRTR    RT    RT    TMTM  `},
		// row 2: upper body (repeated for weight)
		{"\033[0;36m", `PW    PW  WN    WN  NBNB  NB  ST          TRTR    RT    RT    TMTM  `},
		// row 3: belly / crossbar
		{"\033[1;32m", `PWPWPW    WNWNWNWN  NB  NBNB    STST      TRTR    RTRTRT      TMTM  `},
		// row 4: lower body
		{"\033[1;32m", `PW          WNWN    NB  NBNB        ST    TRTR    RT  RT        TMTM  `},
		// row 5: lower body
		{"\033[0;32m", `PW          WNWN    NB    NB          ST  TRTR    RT    RT      TMTM  `},
		// row 6: base / foot
		{"\033[0;32m", `PW                  NB    NB  STSTST      TRTR    RT    RT      TMTM  `},
		{"\033[2;32m", ` ctf . vpn . tmux . enum . scan . loot . report . notes . creds`},
	}

	fmt.Println()
	for _, r := range rows {
		fmt.Printf("%s%s\033[0m\n", r.c, r.t)
	}
	fmt.Printf("\033[2m------------------------------------------------------------\033[0m\n")
	fmt.Printf(" pwnbox-starter | github.com/Nasrif30\n")
	fmt.Printf("\033[2m------------------------------------------------------------\033[0m\n\n")
}





var (
	homeDir    string
	baseDir    string
	configPath string

	optPlatform   string
	optMachine    string
	optTargetIP   string
	optVPNURL     string
	optConnect    bool
	optTmux       bool
	optReport     bool
	optDisconnect bool
	optStatus     bool
	optList       bool
	optArchive    bool
	optRecent     bool
	optLog        bool
	optDryRun     bool
	optConfig     string
	optVersion    bool
)

// knownPlatforms is the list shown as hints in interactive mode.
var knownPlatforms = []string{
	"htb", "thm", "pg", "vulnhub",
	"bugbounty", "bb", "ctf", "pwn",
	"tryhackme", "hackthebox", "offsec",
	"pwnedlabs", "portswigger", "rootme",
}

func initPaths() {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		homeDir = "/root"
	}
	baseDir = filepath.Join(homeDir, "CTF")
	configPath = filepath.Join(homeDir, ".pwnbox.conf")
}

func parseFlags() {
	flag.StringVar(&optPlatform, "p", "", "Platform (e.g. htb, thm, pg, bugbounty)")
	flag.StringVar(&optMachine, "m", "", "Machine / target name")
	flag.StringVar(&optTargetIP, "i", "", "Target IP address")
	flag.StringVar(&optVPNURL, "v", "", "URL to download .ovpn config")
	flag.BoolVar(&optConnect, "c", false, "Connect to VPN after setup")
	flag.BoolVar(&optTmux, "t", false, "Launch organised tmux session")
	flag.BoolVar(&optReport, "r", false, "Generate report.md")
	flag.BoolVar(&optDisconnect, "d", false, "Disconnect active VPN session")
	flag.BoolVar(&optStatus, "status", false, "Show workspace status (read-only)")
	flag.BoolVar(&optList, "l", false, "List active sessions and workspaces")
	flag.BoolVar(&optArchive, "archive", false, "Archive current workspace to ~/CTF/archive/")
	flag.BoolVar(&optRecent, "recent", false, "Show recent projects")
	flag.BoolVar(&optLog, "log", false, "Enable session logging to ~/CTF/logs/")
	flag.BoolVar(&optDryRun, "dry-run", false, "Simulate all operations without creating files")
	flag.StringVar(&optConfig, "config", "", "Path to alternate config file")
	flag.BoolVar(&optVersion, "version", false, "Print version and exit")

	flag.Usage = printUsage
	flag.Parse()
}

func printUsage() {
	fmt.Printf("%spwnbox-starter%s — CTF & Pentest workspace manager\n\n", clBold, clReset)
	fmt.Printf("%sUSAGE%s\n", clCyan, clReset)
	printDivider()
	fmt.Printf("  pwnbox-starter [flags]\n\n")
	fmt.Printf("%sEXAMPLES%s\n", clCyan, clReset)
	printDivider()
	fmt.Printf("  pwnbox-starter                                       (interactive)\n")
	fmt.Printf("  pwnbox-starter -p htb       -m lame    -i 10.10.10.3\n")
	fmt.Printf("  pwnbox-starter -p thm       -m anthem  -i 10.10.x.x  -t\n")
	fmt.Printf("  pwnbox-starter -p pg        -m monitor -i 192.168.x.x -v https://vpn.ovpn -c -t\n")
	fmt.Printf("  pwnbox-starter -p bugbounty -m tesla   -i 10.0.0.1\n")
	fmt.Printf("  pwnbox-starter -p ctf       -m pwn2024\n")
	fmt.Printf("  pwnbox-starter --status\n")
	fmt.Printf("  pwnbox-starter --archive\n")
	fmt.Printf("  pwnbox-starter --recent\n")
	fmt.Printf("  pwnbox-starter -d\n\n")
	fmt.Printf("%sPLATFORMS%s\n", clCyan, clReset)
	printDivider()
	fmt.Printf("  %s\n\n", strings.Join(knownPlatforms, "  "))
	fmt.Printf("%sFLAGS%s\n", clCyan, clReset)
	printDivider()
	flag.PrintDefaults()
	fmt.Println()
}

func getWorkspacePath() string {
	return filepath.Join(baseDir, optPlatform, optMachine)
}

// checkDependencies warns on missing optional tools.
// Only errors on tools required for the specific action being performed.
func checkDependencies() {
	type dep struct {
		name     string
		optional bool
		onlyIf   bool // only check if this condition is true
	}

	deps := []dep{
		{"curl", false, optVPNURL != ""},
		{"wget", true, optVPNURL != ""},
		{"openvpn", false, optConnect || optDisconnect},
		{"tmux", false, optTmux},
		{"vim", true, false},
		{"tree", true, false},
		{"tar", true, optArchive},
	}

	allOk := true
	for _, d := range deps {
		// Skip check if condition not applicable
		if !d.optional && !d.onlyIf {
			continue
		}
		_, err := exec.LookPath(d.name)
		if err != nil {
			if !d.optional {
				printError(fmt.Sprintf("%s is required for this action but not installed", d.name))
				allOk = false
			} else if d.onlyIf {
				printWarning(fmt.Sprintf("%s not installed (optional)", d.name))
			}
		}
	}
	if allOk {
		printSuccess("Dependencies OK")
	}
}

// interactiveMode guides the user through setup with menus and simple prompts.
func interactiveMode() {
	reader := bufio.NewReader(os.Stdin)

	// ── Platform picker ──────────────────────────────────────────
	fmt.Printf("\n%s  Select Platform:%s\n", clBold, clReset)
	printDivider()
	for i, p := range knownPlatforms {
		fmt.Printf("  %s[%2d]%s  %s\n", clCyan, i+1, clReset, p)
	}
	fmt.Printf("  %s[  ]%s  other (type it)\n\n", clDim, clReset)

	platInput := readLine(reader, "  Choice [1]: ")
	if platInput == "" {
		platInput = "1"
	}
	idx := 0
	_, err := fmt.Sscanf(platInput, "%d", &idx)
	if err == nil && idx >= 1 && idx <= len(knownPlatforms) {
		optPlatform = knownPlatforms[idx-1]
	} else {
		// User typed a custom platform name
		optPlatform = strings.ToLower(strings.TrimSpace(platInput))
		if !isValidPlatformStr(optPlatform) {
			printWarning("Invalid platform name. Using 'htb'.")
			optPlatform = "htb"
		}
	}
	printSuccess(fmt.Sprintf("Platform: %s", optPlatform))

	// ── Machine / target name ────────────────────────────────────
	fmt.Printf("\n%s  Target / Machine name:%s\n", clBold, clReset)
	for {
		optMachine = readLine(reader, "  Name: ")
		optMachine = strings.ToLower(strings.TrimSpace(optMachine))
		if optMachine == "" {
			printWarning("Machine name is required.")
			continue
		}
		if !isValidMachineStr(optMachine) {
			printWarning("Use lowercase letters, numbers, hyphens only (e.g. lame, blue-team).")
			continue
		}
		break
	}
	printSuccess(fmt.Sprintf("Machine: %s", optMachine))

	// ── Target IP (optional) ─────────────────────────────────────
	fmt.Printf("\n%s  Target IP%s %s(optional — press Enter to skip)%s\n", clBold, clReset, clDim, clReset)
	optTargetIP = readLine(reader, "  IP: ")
	optTargetIP = strings.TrimSpace(optTargetIP)

	// ── Feature toggles ──────────────────────────────────────────
	fmt.Printf("\n%s  Features:%s\n", clBold, clReset)
	printDivider()
	optConnect = menuYN(reader, "Connect VPN?    ")
	if optConnect {
		ovpnFiles := findOVPNFiles()
		if len(ovpnFiles) > 0 {
			fmt.Printf("\n%s  Found local VPN configs:%s\n", clDim, clReset)
			for i, f := range ovpnFiles {
				displayPath := f
				if strings.HasPrefix(f, homeDir) {
					displayPath = "~" + f[len(homeDir):]
				}
				fmt.Printf("  %s[%2d]%s  %s\n", clCyan, i+1, clReset, displayPath)
			}
		}

		fmt.Printf("\n%s  VPN config URL or local path%s\n  %s(choose number, type 'browse', paste path, or press Enter)%s\n", clBold, clReset, clDim, clReset)
		optVPNURL = strings.TrimSpace(readLine(reader, "  Input: "))

		if strings.ToLower(optVPNURL) == "browse" || strings.ToLower(optVPNURL) == "b" {
			printInfo("Opening file browser...")
			selected := browseFile()
			if selected != "" {
				optVPNURL = selected
				printSuccess("Selected: " + optVPNURL)
			} else {
				printWarning("No file selected.")
				optVPNURL = ""
			}
		} else {
			idx := 0
			if _, err := fmt.Sscanf(optVPNURL, "%d", &idx); err == nil && idx >= 1 && idx <= len(ovpnFiles) {
				optVPNURL = ovpnFiles[idx-1]
			} else if strings.HasPrefix(optVPNURL, "~/") {
				optVPNURL = filepath.Join(homeDir, optVPNURL[2:])
			}
		}
	}
	optTmux   = menuYN(reader, "Launch TMUX?    ")
	optReport = menuYN(reader, "Generate report?")
	fmt.Println()
}

// findOVPNFiles searches for .ovpn files in ~/Downloads, current directory, and Windows Downloads if in WSL.
func findOVPNFiles() []string {
	var files []string
	searchDirs := []string{
		filepath.Join(homeDir, "Downloads"),
		".",
	}

	// Add WSL Windows Downloads folder
	if runtime.GOOS == "linux" {
		if b, err := os.ReadFile("/proc/version"); err == nil && strings.Contains(strings.ToLower(string(b)), "microsoft") {
			cmd := exec.Command("cmd.exe", "/c", "echo %USERPROFILE%")
			out, err := cmd.Output()
			if err == nil {
				winPath := strings.TrimSpace(string(out))
				winPath = strings.ReplaceAll(winPath, "\\", "/")
				if len(winPath) >= 2 && winPath[1] == ':' {
					drive := strings.ToLower(string(winPath[0]))
					wslDownloads := fmt.Sprintf("/mnt/%s%s/Downloads", drive, winPath[2:])
					searchDirs = append(searchDirs, wslDownloads)
				}
			}
		}
	}

	for _, dir := range searchDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".ovpn") {
				files = append(files, filepath.Join(dir, e.Name()))
			}
		}
	}
	return files
}

// readLine reads a single line from stdin with a displayed prompt.
func readLine(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	val, _ := r.ReadString('\n')
	return strings.TrimSpace(val)
}

// menuYN prints a styled inline yes/no prompt and returns true for 'y'/'Y'/1.
func menuYN(r *bufio.Reader, label string) bool {
	fmt.Printf("  %s%-18s%s %s[y/N]%s: ", clBold, label, clReset, clDim, clReset)
	val, _ := r.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(val)) == "y"
}


// prompt prints a labelled prompt with an optional default shown in brackets.
func prompt(r *bufio.Reader, label, def string) string {
	if def != "" {
		fmt.Printf("  %s%s%s [%s%s%s]: ", clBold, label, clReset, clDim, def, clReset)
	} else {
		fmt.Printf("  %s%s%s: ", clBold, label, clReset)
	}
	val, _ := r.ReadString('\n')
	val = strings.TrimSpace(val)
	if val == "" {
		return def
	}
	return val
}

// promptValidated keeps re-prompting until the value passes the validator.
// Typing 'help', '-h', or '--help' at any prompt shows usage and exits.
func promptValidated(r *bufio.Reader, label, def string, validate func(string) bool) string {
	for {
		val := prompt(r, label, def)
		switch strings.TrimSpace(val) {
		case "help", "-h", "--help", "-help":
			fmt.Println()
			printUsage()
			os.Exit(0)
		}
		if val == "" {
			printWarning(label + " is required.")
			continue
		}
		if !validate(val) {
			printWarning(fmt.Sprintf("'%s' is not a valid value. Use lowercase letters, numbers, hyphens only.", val))
			continue
		}
		return val
	}
}

// promptYN asks a yes/no question and returns true for 'y'.
func promptYN(r *bufio.Reader, label string) bool {
	fmt.Printf("  %s%s%s [y/N]: ", clBold, label, clReset)
	val, _ := r.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(val)) == "y"
}

func isValidPlatformStr(s string) bool {
	ok, _ := regexp.MatchString(`^[a-z0-9]+$`, s)
	return ok
}

func isValidMachineStr(s string) bool {
	ok, _ := regexp.MatchString(`^[a-z0-9][a-z0-9-]*$`, s)
	return ok
}

// validateInput checks platform and machine name formats.
func validateInput() error {
	if !isValidPlatformStr(optPlatform) {
		return fmt.Errorf("invalid platform '%s': lowercase letters and numbers only (e.g. htb, thm, bugbounty)", optPlatform)
	}
	if !isValidMachineStr(optMachine) {
		return fmt.Errorf("invalid machine '%s': lowercase letters, numbers, hyphens (e.g. lame, blue-team)", optMachine)
	}
	if optTargetIP != "" {
		ok, _ := regexp.MatchString(`^[a-zA-Z0-9.\-:]+$`, optTargetIP)
		if !ok {
			return fmt.Errorf("invalid IP/Target address '%s'", optTargetIP)
		}
	}
	return nil
}
