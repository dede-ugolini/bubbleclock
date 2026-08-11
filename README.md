# BubbleClock

A terminal clock that renders the current time as ASCII art using [go-figure](https://github.com/common-nighthawk/go-figure), built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## Usage

```
go run .
```

## Keybindings

| Key | Action                                   |
| --- | ---------------------------------------- |
| `q` | Quit                                     |
| `s` | Next font                                |
| `S` | Previous font                            |
| `c` | Next color                               |
| `b` | Toggle bold                              |
| `B` | Toggle blinking colon                    |
| `t` | Toggle 12h / 24h format                  |
| `d` | Toggle date display                      |
| `h` | Toggle help bar                          |

## Features

- Rotates through a collection of ASCII art fonts (44 built-in)
- Color cycling through a built-in palette
- Optional date line (`Monday, January 2, 2006`)
- 12/24-hour time format switching
- Blinking colon effect
- Bold text rendering
- Centered, full-terminal layout
