# Pwnbox-Starter

> Pwnbox-Starter is a cross-platform Go CLI utility that automates CTF and penetration testing workflows. It instantly scaffolds organized workspace directories, auto-discovers OpenVPN profiles, launches predefined TMUX hacking environments, and generates interactive Markdown report templates. Perfect for HackTheBox, TryHackMe, and Bug Bounties.



<img width="692" height="650" alt="Screenshot 2026-06-30 041626" src="https://github.com/user-attachments/assets/d824de6c-a1ae-46a0-8845-3f7f9e3511b3" />




<img width="717" height="516" alt="Screenshot 2026-06-30 042255" src="https://github.com/user-attachments/assets/ceb758cb-c4be-4381-9085-ede4be8c672c" />




<img width="1905" height="747" alt="Screenshot 2026-06-30 040319" src="https://github.com/user-attachments/assets/48826f59-95c1-4a4f-bacc-c4fcbe4729a6" />




<img width="1901" height="943" alt="Screenshot 2026-06-30 040334" src="https://github.com/user-attachments/assets/ec2bfb22-fa6f-48a0-b639-ebb37a71915c" />



---

## The Motivation

In the world of penetration testing and CTFs, the first 10 minutes of every engagement are always identical: manually creating `/scans` and `/exploits` directories, finding and connecting to the correct VPN profile, splitting terminal windows into multiple panes, and creating a blank Markdown template for notes. 

Pwnbox-Starter was engineered out of the developer's pure necessity—and constructive laziness—to entirely eliminate this repetitive friction. By completely automating the setup phase, this tool ensures you can jump straight into enumeration and exploitation the moment a machine boots up.

---

## Architecture & Flow

<img width="1453" height="8192" alt="Pwnbox Starter VPN Workflow-2026-06-29-201236" src="https://github.com/user-attachments/assets/a3b8229f-51c0-4462-b0a8-188a5df46795" />


---

## Features

Pwnbox-Starter is designed to eliminate the repetitive setup phase of penetration testing so you can immediately focus on enumeration and exploitation. 

- **Intelligent Workspace Scaffolding:** Automatically generates a structured project directory (e.g., `~/CTF/htb/machine_name`) containing organized subfolders for `/scans`, `/exploits`, `/loot`, and `/credentials`.
- **VPN Auto-Discovery & Orchestration:** The utility automatically scans your local directories (including cross-environment WSL detection) for `.ovpn` profiles. It copies the config and establishes the VPN connection in a detached, dedicated pop-out window.
- **Graphical File Selection:** Type `browse` directly in the CLI to launch a native OS file-picker dialog (supports Windows Forms, macOS osascript, and Linux zenity) for effortless VPN profile selection.
- **TMUX Integration:** Instantly spins up a customized 4-window terminal multiplexer session tailored for hacking (`recon`, `exploit`, `shell`, `root`). If your connection drops, your shells and scans remain persistent in the background.
- **Interactive Report Generation:** Prompts the user for target metadata (OS, Difficulty, Initial Notes) and automatically builds a pre-formatted `report.md` template for documenting vulnerabilities, flags, and write-ups.
- **Environment Requirements:** Because this is a pwning utility relying heavily on Linux security tooling, it requires a Unix environment. It runs flawlessly on a dedicated **Linux OS (Kali/Parrot)** or inside a **Virtual Machine (VMware / VirtualBox)**. For Windows and macOS users, it runs perfectly via **WSL (Windows Subsystem for Linux)** or your hypervisor of choice.

---

## Usage & Commands

To keep this document clean, all installation instructions, cross-compilation guides, and CLI flag references have been moved to a dedicated documentation file.

For full documentation on how to build, install, and use `pwnbox-starter`, see the **[COMMANDS Reference Guide](./COMMANDS.md)**.
