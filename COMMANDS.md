# pwnbox-starter — Command Reference

All commands for building, installing, and using `pwnbox-starter` across platforms.

---

## Build

### WSL / Kali Linux / Debian / Ubuntu

```bash
cd /mnt/d/Web/CTF/pwnbox_starter

# Build for Linux (native)
go build -o pwnbox-starter

# Make executable
chmod +x pwnbox-starter

# Optional: install system-wide
sudo cp pwnbox-starter /usr/local/bin/pwnbox-starter
```

### Windows (PowerShell)

```powershell
cd D:\Web\CTF\pwnbox_starter

# Build for Windows
go build -o pwnbox-starter.exe

# Run directly
.\pwnbox-starter.exe
```

### Cross-compile from Windows → Linux (for Kali)

```powershell
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o pwnbox-starter-linux
```

### Cross-compile from Linux → Windows

```bash
GOOS=windows GOARCH=amd64 go build -o pwnbox-starter.exe
```

### Cross-compile for macOS (from any platform)

```bash
GOOS=darwin GOARCH=amd64  go build -o pwnbox-starter-macos-x64
GOOS=darwin GOARCH=arm64  go build -o pwnbox-starter-macos-arm64   # Apple Silicon
```

### Cross-compile for ARM (Raspberry Pi / mobile)

```bash
GOOS=linux GOARCH=arm64 go build -o pwnbox-starter-arm64
GOOS=linux GOARCH=arm   GOARM=7 go build -o pwnbox-starter-armv7
```

---

## Usage

### Interactive Mode (no arguments)

```bash
./pwnbox-starter
```

### Common CLI Commands

```bash
# Basic workspace setup
./pwnbox-starter -p htb -m lame

# With target IP
./pwnbox-starter -p htb -m lame -i 10.10.10.3

# With tmux session
./pwnbox-starter -p htb -m lame -i 10.10.10.3 -t

# With VPN download + connect + tmux
./pwnbox-starter -p htb -m lame -i 10.10.10.3 -v https://example.com/htb.ovpn -c -t

# Generate report
./pwnbox-starter -p htb -m lame -r

# Full setup: IP + VPN + tmux + report
./pwnbox-starter -p htb -m lame -i 10.10.10.3 -v https://example.com/htb.ovpn -c -t -r

# Disconnect active VPN
./pwnbox-starter -d

# Dry run (simulate everything, create nothing)
./pwnbox-starter -p htb -m lame -i 10.10.10.3 -t --dry-run

# Show version
./pwnbox-starter --version

# Show help
./pwnbox-starter -h
```

### Platform Examples

```bash
# HackTheBox
./pwnbox-starter -p htb -m lame    -i 10.10.10.3
./pwnbox-starter -p htb -m blue    -i 10.10.10.40
./pwnbox-starter -p htb -m nibbles -i 10.10.10.75 -t

# TryHackMe
./pwnbox-starter -p thm -m anthem   -i 10.10.x.x -t
./pwnbox-starter -p thm -m basicpentesting -i 10.10.x.x

# Proving Grounds
./pwnbox-starter -p pg -m monitor   -i 192.168.x.x -v https://portal.offensive-security.com/me.ovpn -c -t

# VulnHub (no VPN needed)
./pwnbox-starter -p vulnhub -m kioptrix -i 192.168.1.100 -t
```

### Interactive Mode Features (NEW in v2.0)

When you run `pwnbox-starter` without arguments, it launches a highly interactive wizard with the following features:

- **VPN Auto-Discovery**: Automatically scans your `~/Downloads` (and Windows `Downloads` if running in WSL Kali) and lists local `.ovpn` files for one-tap selection.
- **GUI File Picker (`browse`)**: At the VPN prompt, type `browse` or `b` to launch a native graphical file picker (works on Windows, macOS, and Linux/WSL).
- **GUI VPN Pop-out**: On Windows and macOS, the VPN process automatically pops out into a brand new terminal window so your current terminal remains free for hacking. On Linux, it tries to launch a new terminal emulator window or falls back to background daemon mode.
- **Interactive Report Generator**: Prompts for Target OS, Difficulty, and Initial Notes to dynamically construct a detailed `report.md` template for you.
- **Target IP with Ports**: Supports IP addresses with ports (e.g., `10.10.10.10:32042`) and hostnames (e.g., `example.htb`).

---

## Flags Reference

| Flag | Type | Description |
|------|------|-------------|
| `-p` | string | Platform name (e.g. `htb`, `thm`, `pg`) |
| `-m` | string | Machine / box name (e.g. `lame`, `blue-team`) |
| `-i` | string | Target IP address (now supports `IP:Port` format) |
| `-v` | string | URL or Local Path to download/copy `.ovpn` config |
| `-c` | bool | Connect to VPN after setup (Pops out in a new window) |
| `-t` | bool | Launch organised tmux session |
| `-r` | bool | Launch interactive `report.md` generator |
| `-d` | bool | Disconnect active VPN session |
| `-l` | bool | List active sessions and workspaces |
| `-h` | bool | Show help |
| `--status` | bool | Show workspace status (read-only) |
| `--archive` | bool | Compress workspace into `~/CTF/archive/` |
| `--recent` | bool | Show recent projects |
| `--log` | bool | Enable session logging to `~/CTF/logs/` |
| `--dry-run` | bool | Simulate all operations without creating files |
| `--config` | string | Path to alternate config file |
| `--version` | bool | Print version and exit |

---

## TMUX — After Launch

```bash
# Attach to the session created by pwnbox-starter
tmux attach -t pwnbox-htb-lame

# List all tmux sessions
tmux ls

# Kill a specific session
tmux kill-session -t pwnbox-htb-lame

# Switch between windows inside tmux
Ctrl+b then 0   # shell
Ctrl+b then 1   # nmap
Ctrl+b then 2   # gobuster
Ctrl+b then 3   # notes
```

---

## Workspace Layout

```
~/CTF/
├── htb/
│   └── lame/
│       ├── scans/
│       ├── exploits/
│       ├── loot/
│       ├── notes/
│       │   └── machine.md
│       ├── screenshots/
│       ├── credentials/
│       ├── report.md
│       └── .pwnbox-session
├── vpn/
│   └── htb.ovpn
├── logs/
└── archive/
```

---

## VPN Management

```bash
# Download and connect in one command
./pwnbox-starter -p htb -m lame -v https://your.ovpn -c

# Connect to already-downloaded config
./pwnbox-starter -p htb -m lame -c

# Check tun0 after connecting
ip addr show tun0

# Disconnect all VPN sessions
./pwnbox-starter -d

# Manual OpenVPN (bypass pwnbox-starter)
sudo openvpn --config ~/CTF/vpn/htb.ovpn
```

---

## Development

```bash
# Run without building
go run .

# Format code
gofmt -w .

# Vet for issues
go vet ./...

# Tidy dependencies
go mod tidy

# Build all platforms at once
GOOS=linux   GOARCH=amd64 go build -o dist/pwnbox-starter-linux-amd64
GOOS=windows GOARCH=amd64 go build -o dist/pwnbox-starter-windows-amd64.exe
GOOS=darwin  GOARCH=amd64 go build -o dist/pwnbox-starter-darwin-amd64
GOOS=darwin  GOARCH=arm64 go build -o dist/pwnbox-starter-darwin-arm64
```

---

## Install System-Wide

### Linux / WSL Kali

```bash
sudo cp pwnbox-starter /usr/local/bin/
# Now usable from anywhere:
pwnbox-starter -p htb -m lame -i 10.10.10.3 -t
```

### Windows — Add to PATH

```powershell
# Copy to a directory already in your PATH, or add new entry:
Copy-Item .\pwnbox-starter.exe "$env:USERPROFILE\bin\pwnbox-starter.exe"
```

### macOS

```bash
sudo cp pwnbox-starter-macos-arm64 /usr/local/bin/pwnbox-starter
chmod +x /usr/local/bin/pwnbox-starter
```
