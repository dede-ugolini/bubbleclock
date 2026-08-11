package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

type model struct {
	now        time.Time
	width      int
	height     int
	idxFont    int
	idxColors  int
	use24hours bool
	showDate   bool
	showHelp   bool
	blink      bool
	bold       bool
	ready      bool
	quitting   bool
}

type tickMsg time.Time

var fonts = []string{
	"5lineoblique",
	"acrobatic   ",
	"alligator   ",
	"alligator2  ",
	"avatar      ",
	"banner      ",
	"banner3-D   ",
	"banner3     ",
	"basic       ",
	"bell        ",
	"big         ",
	"bigchief    ",
	"bulbhead    ",
	"chunky      ",
	"colossal    ",
	"computer    ",
	"contessa    ",
	"cosmic      ",
	"cricket     ",
	"doh         ",
	"doom        ",
	"dotmatrix   ",
	"drpepper    ",
	"eftifont    ",
	"epic        ",
	"fender      ",
	"fuzzy       ",
	"gothic      ",
	"graffiti    ",
	"hollywood   ",
	"jazmine     ",
	"larry3d     ",
	"nancyj-fancy",
	"o8          ",
	"poison      ",
	"puffy       ",
	"rev         ",
	"roman       ",
	"rounded     ",
	"script      ",
	"slant       ",
	"starwars    ",
	"univers     ",
}

var palette = []lipgloss.Color{
	lipgloss.Color("42"),  // green
	lipgloss.Color("214"), // orange
	lipgloss.Color("81"),  // cyan
	lipgloss.Color("212"), // pink
	lipgloss.Color("15"),  // white
}

func (m model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tickMsg:
		m.now = time.Time(msg)
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.quitting = true
			return m, tea.Quit

		case "s":
			m.idxFont = (m.idxFont + 1) % len(fonts)
			return m, nil

		case "S":
			m.idxFont = (max(0, m.idxFont-1)) % len(fonts)
			return m, nil

		case "c":
			m.idxColors = (m.idxColors + 1) % len(palette)
			return m, nil

		case "b":
			m.bold = !m.bold
			return m, nil

		case "h":
			m.showHelp = !m.showHelp
			return m, nil

		case "t":
			m.use24hours = !m.use24hours
			return m, nil

		case "d":
			m.showDate = !m.showDate
			return m, nil

		case "B":
			m.blink = !m.blink
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting || !m.ready {
		return ""
	}

	font := strings.Trim(fonts[m.idxFont], " ")
	help := ""

	if m.showHelp {
		help = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).
			Render(fmt.Sprintf("q quit	 c color	t 12/24h	d date	h show help  b bold	 B blink	 s font (%s)", font))
	}

	date := ""
	if m.showDate {
		dateLine := m.now.Format("Monday, January 2, 2006")
		date = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render(dateLine)
	}

	clock, suffix := m.formatTime()
	str := figure.NewFigure(clock+" "+suffix, font, false).Slicify()
	strClock := lipgloss.NewStyle().Foreground(palette[m.idxColors]).Bold(m.bold).Render(strings.Join(str, "\n"))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, strClock+"\n\n"+date+"\n"+help)
}

func (m model) formatTime() (clockStr, suffix string) {
	if m.use24hours {
		return m.now.Format("15:04:05"), ""
	}

	colonLit := true
	if m.blink {
		colonLit = m.now.Second()%2 == 0
	}

	colon := " "
	if colonLit {
		colon = ":"
	}
	h := m.now.Hour() % 12
	if h == 0 {
		h = 12
	}
	suffix = "AM"
	if m.now.Hour() >= 12 {
		suffix = "PM"
	}
	return fmt.Sprintf("%02d %s %02d %s %02d", h, colon, m.now.Minute(), colon, m.now.Second()), suffix
}

func initialModel() model {
	return model{
		now:        time.Now(),
		width:      0,
		height:     0,
		idxFont:    0,
		idxColors:  0,
		use24hours: false,
		showDate:   true,
		showHelp:   true,
		blink:      false,
		bold:       false,
		ready:      false,
		quitting:   false,
	}
}

func main() {
	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
