# Pwnbox-Starter

> Pwnbox-Starter is a cross-platform Go CLI utility that automates CTF and penetration testing workflows. It instantly scaffolds organized workspace directories, auto-discovers OpenVPN profiles, launches predefined TMUX hacking environments, and generates interactive Markdown report templates. Perfect for HackTheBox, TryHackMe, and Bug Bounties.

![Banner Screenshot](./banner.png) *(Upload your terminal banner screenshot here)*

![Interactive Setup Screenshot](./setup.png) *(Upload your interactive setup screenshot here)*

![Report & VPN Screenshot](./report.png) *(Upload your VPN/Report screenshot here)*

![Final Workspace Setup](./workspace.png) *(Upload your final workspace/summary screenshot here)*

---

## The Motivation

In the world of penetration testing and CTFs, the first 10 minutes of every engagement are always identical: manually creating `/scans` and `/exploits` directories, finding and connecting to the correct VPN profile, splitting terminal windows into multiple panes, and creating a blank Markdown template for notes. 

Pwnbox-Starter was engineered out of the developer's pure necessity—and constructive laziness—to entirely eliminate this repetitive friction. By completely automating the setup phase, this tool ensures you can jump straight into enumeration and exploitation the moment a machine boots up.

---

## Architecture & Flow

<!-- PASTE YOUR MERMAID DIAGRAM IMAGE OR CODE HERE -->

---

## Features

Pwnbox-Starter is designed to eliminate the repetitive setup phase of penetration testing so you can immediately focus on enumeration and exploitation. 

- **Intelligent Workspace Scaffolding:** Automatically generates a structured project directory (e.g., `~/CTF/htb/machine_name`) containing organized subfolders for `/scans`, `/exploits`, `/loot`, and `/credentials`.
- **VPN Auto-Discovery & Orchestration:** The utility automatically scans your local directories (including cross-environment WSL detection) for `.ovpn` profiles. It copies the config and establishes the VPN connection in a detached, dedicated pop-out window.
- **Graphical File Selection:** Type `browse` directly in the CLI to launch a native OS file-picker dialog (supports Windows Forms, macOS osascript, and Linux zenity) for effortless VPN profile selection.
- **TMUX Integration:** Instantly spins up a customized 4-window terminal multiplexer session tailored for hacking (`recon`, `exploit`, `shell`, `root`). If your connection drops, your shells and scans remain persistent in the background.
- **Interactive Report Generation:** Prompts the user for target metadata (OS, Difficulty, Initial Notes) and automatically builds a pre-formatted `report.md` template for documenting vulnerabilities, flags, and write-ups.
- **Cross-Platform Native:** Written entirely in Go with zero external dependencies. Fully compatible across Windows (PowerShell), Linux (native & WSL), and macOS.

---

## Usage & Commands

To keep this document clean, all installation instructions, cross-compilation guides, and CLI flag references have been moved to a dedicated documentation file.

For full documentation on how to build, install, and use `pwnbox-starter`, see the **[COMMANDS Reference Guide](./COMMANDS.md)**.
