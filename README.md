# goproc

**goproc** is a simple, modular CLI tool written in Go for displaying basic system information. 
Ideal for learning Go, exploring system stats, showcasing Go development practices and practical CLI tooling.

## Features

- 🖥️ Hostname and CPU info
- 💾 Memory usage (current, total, system)
- 📀 Disk usage (total, used, free, percentage)
- ⏱️ OS info and system uptime (coming soon)

## Usage

Run the CLI from the project root:

```bash
go run main.go

Example output:

goproc - Simple System Info
===========================
Hostname: YOUR-HOSTNAME
CPU Cores: 12

Memory Info:
  Allocated: 0.15 MB
  Total Allocated: 0.15 MB
  System Memory: 6.09 MB

Disk Usage (/):
  Total: 499.55 GB
  Used:  486.99 GB
  Free:  12.57 GB
  Used Percent: 97.48%

Goals
This project is designed to:

Practice writing modular Go code

Build familiarity with Git and GitHub

Serve as a clean, presentable CLI project for job applications
