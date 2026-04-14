# whoami.ssh

Interactive portfolio served over SSH, written in Go with Charmbracelet's TUI stack.

## Live

|         |                                            |
| ------- | ------------------------------------------ |
| **Web** | [https://denis.my.id](https://denis.my.id) |
| **SSH** | `ssh denis.my.id`                          |

```bash
ssh denis.my.id
```

## Overview

This repository contains:

- An SSH server that launches a terminal UI portfolio session.
- A multi-view TUI (Home, About, Skills, Projects, Contact).
- A static web landing page in `frontend/index.html`.
- An example Nginx config in `nginx.conf` for the web domain setup.

## Features

- SSH server using `wish` + Bubble Tea middleware.
- Keyboard controls: `tab` / `shift+tab` to switch views.
- Section shortcuts: `0` to `4` for Home, About, Skills, Projects, and Contact.
- Quit controls: `q` / `ctrl+c`.
- Scrollable content areas for About, Skills, Projects, and Contact views.
- Responsive layout updates on terminal resize.
- Ayu Dark-inspired styling via Lip Gloss.

## Tech Stack

- Go 1.25
- `github.com/charmbracelet/wish`
- `github.com/charmbracelet/bubbletea`
- `github.com/charmbracelet/bubbles`
- `github.com/charmbracelet/lipgloss`
- `github.com/charmbracelet/log`
- Frontend landing page: plain HTML with Tailwind Play CDN and Alpine.js CDN

## Setup and Run

Prerequisite: Go installed.

```bash
go run .
```

By default, the SSH server listens on `0.0.0.0:2222`.

Connect from another terminal:

```bash
ssh -p 2222 localhost
```

Optional flags:

```bash
go run . --host 127.0.0.1 --port 2222
```

The server is configured to use host key path `.ssh/id_ed25519`.

## Project Structure

```text
whoami.ssh/
├── main.go
├── go.mod
├── go.sum
├── nginx.conf
├── frontend/
│   └── index.html
└── internal/
    ├── data/
    │   └── profile.go
    └── tui/
        ├── app.go
        ├── home.go
        ├── about.go
        ├── skills.go
        ├── projects.go
        ├── contact.go
        └── styles.go
```
