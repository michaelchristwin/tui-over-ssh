# Bubble Tea SSH Name Prompt

A small SSH server built with [Wish](https://github.com/charmbracelet/wish) and [Bubble Tea](https://github.com/charmbracelet/bubbletea) (v2). When a client connects over SSH, they're dropped into a terminal UI that asks for their name. The answer is written to a local file.

## How it works

1. The server listens on `localhost:3000` and accepts SSH connections.
2. Each session gets its own Bubble Tea program running in the alternate screen.
3. The user types their name into a text input (placeholder: `Jan Hus`).
4. Pressing **Enter** saves the input to `output.log` and ends the session. Pressing **Ctrl+C** quits without saving.

## Requirements

- A recent version of [Go](https://go.dev/dl/) (Bubble Tea v2 requires a modern toolchain; check the Bubble Tea repository for the minimum)
- An SSH client (`ssh` is preinstalled on macOS, Linux, and current Windows)

## Dependencies

| Package | Purpose |
| --- | --- |
| `charm.land/bubbletea/v2` | TUI framework |
| `charm.land/bubbles/v2` | Text input component |
| `charm.land/wish/v2` | SSH server and middleware |
| `charm.land/ssh` | SSH session types |
| `charm.land/log/v2` | Logging |

> All Charm libraries must use matching v2 import paths. Mixing them with the old `github.com/charmbracelet/...` packages causes type mismatch errors.

## Setup

```sh
# Initialize the module (skip if go.mod already exists)
go mod init your-module-name

# Download dependencies
go mod tidy
```

## Running

```sh
go run .
```

You should see a log line saying the SSH server has started. In another terminal, connect with:

```sh
ssh localhost -p 3000
```

If your SSH client complains about the host key changing between runs, remove the old entry with `ssh-keygen -R "[localhost]:3000"`.

## Output

When a user presses Enter, their input is written to `output.log` in the directory the server was started from. The file is **overwritten** on every submission, so it only ever holds the most recent name.

## Configuration

The host and port are constants at the top of `main.go`:

```go
const (
	host = "localhost"
	port = "3000"
)
```

Change `host` to `0.0.0.0` to accept connections from other machines. The host key is stored at `.ssh/id_ed25519`; Wish generates it automatically if it doesn't exist. Don't commit this file to version control.

## Project structure

```
.
├── main.go      # Server setup, SSH middleware, and the Bubble Tea model
├── go.mod
├── go.sum
└── .ssh/        # Generated host key (git-ignored)
```

## Notes

- The `activeterm` middleware rejects connections that don't request a PTY, since Bubble Tea needs one. Connect with a normal interactive `ssh` command rather than passing a remote command.
- There is no authentication configured, so anyone who can reach the port can connect. Add a Wish auth option before exposing this publicly.

## Stopping the server

Press **Ctrl+C** in the terminal running the server. It shuts down gracefully, waiting up to 30 seconds for active sessions to finish.
