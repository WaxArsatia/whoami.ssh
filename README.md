# whoami.ssh

> An interactive SSH portfolio — visit it right from your terminal.

## Live

|         |                                                          |
| ------- | -------------------------------------------------------- |
| **Web** | [https://whoami.denis.my.id](https://whoami.denis.my.id) |
| **SSH** | `ssh whoami.denis.my.id`                                 |

```bash
ssh whoami.denis.my.id
```

## Features

- Built with [charmbracelet/wish](https://github.com/charmbracelet/wish) — SSH server framework
- Beautiful TUI powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea) + [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Ayu Dark color theme
- 5 interactive sections: Home · About · Skills · Projects · Contact
- Viewport scrolling, keyboard navigation, responsive layout

## Navigation

| Key                    | Action                     |
| ---------------------- | -------------------------- |
| `tab` / `shift+tab`    | Cycle through sections     |
| `0` – `4`              | Jump directly to a section |
| `↑` / `↓` or `j` / `k` | Scroll content             |
| `g` / `G`              | Scroll to top / bottom     |
| `q` / `ctrl+c`         | Quit                       |

Section keys: `0` home · `1` about · `2` skills · `3` projects · `4` contact

## Project Structure

```
whoami.ssh/
├── main.go               # SSH server entry point (wish + bubbletea middleware)
├── go.mod
├── frontend/
│   └── index.html        # Static web landing page (Tailwind + Alpine.js)
├── internal/
│   ├── data/
│   │   └── profile.go    # Profile, TechStack, and Projects data
│   └── tui/
│       ├── app.go        # Root model, tab bar, status bar, key handling
│       ├── home.go       # Home / landing view
│       ├── about.go      # About section (viewport)
│       ├── skills.go     # Skills section (viewport)
│       ├── projects.go   # Projects section (viewport)
│       ├── contact.go    # Contact section (viewport)
│       └── styles.go     # Ayu Dark palette + shared lipgloss styles
└── whoami.ssh            # Compiled binary (generated after build)
```

## Running Locally

```bash
# Requires Go
go run .

# Connect (in another terminal) — default port is 2222 locally
ssh -p 2222 localhost
```

> The SSH host key is auto-generated at `.ssh/id_ed25519` on first run.

To run on the standard SSH port 22 (like production):

```bash
sudo go run . --port 22
ssh localhost
```

## Docker

```dockerfile
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o whoami.ssh .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/whoami.ssh .
EXPOSE 22
CMD ["./whoami.ssh", "--port", "22"]
```

```bash
docker build -t whoami-sh .
# run on standard port 22
docker run -p 22:22 whoami-sh
```

## Stack

- **Go** — runtime
- **[wish](https://github.com/charmbracelet/wish)** — SSH server framework
- **[bubbletea](https://github.com/charmbracelet/bubbletea)** — TUI framework (MVU)
- **[lipgloss](https://github.com/charmbracelet/lipgloss)** — terminal styling
- **[bubbles](https://github.com/charmbracelet/bubbles)** — viewport component
- **[log](https://github.com/charmbracelet/log)** — structured logging

## License

MIT — Denis Arsyatya (WaxArsatia)
