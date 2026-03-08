package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ───────────────────────────────────────────────────────────────────────────
// Layout constants
// ───────────────────────────────────────────────────────────────────────────

const (
	tabBarHeight    = 2
	statusBarHeight = 1
)

// ───────────────────────────────────────────────────────────────────────────
// Views
// ───────────────────────────────────────────────────────────────────────────

type viewID int

const (
	viewHome viewID = iota
	viewAbout
	viewSkills
	viewProjects
	viewContact
	viewCount
)

type tabDef struct {
	id    viewID
	label string
	key   string
}

var tabs = []tabDef{
	{viewHome, "  home  ", "0"},
	{viewAbout, "  about  ", "1"},
	{viewSkills, "  skills  ", "2"},
	{viewProjects, "  projects  ", "3"},
	{viewContact, "  contact  ", "4"},
}

// ───────────────────────────────────────────────────────────────────────────
// Root model
// ───────────────────────────────────────────────────────────────────────────

type Model struct {
	width    int
	height   int
	active   viewID
	styles   Styles
	home     homeView
	about    aboutView
	skills   skillsView
	projects projectsView
	contact  contactView
}

func New(w, h int, r *lipgloss.Renderer) Model {
	st := NewStyles(r)
	return Model{
		width:    w,
		height:   h,
		active:   viewHome,
		styles:   st,
		home:     newHomeView(w, h, st),
		about:    newAboutView(w, h, st),
		skills:   newSkillsView(w, h, st),
		projects: newProjectsView(w, h, st),
		contact:  newContactView(w, h, st),
	}
}


func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.about.Init(),
		m.skills.Init(),
		m.projects.Init(),
		m.contact.Init(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.home.width = msg.Width
		m.home.height = msg.Height - tabBarHeight - statusBarHeight

		var cmd tea.Cmd
		m.about, cmd = m.about.Update(msg)
		cmds = append(cmds, cmd)
		m.skills, cmd = m.skills.Update(msg)
		cmds = append(cmds, cmd)
		m.projects, cmd = m.projects.Update(msg)
		cmds = append(cmds, cmd)
		m.contact, cmd = m.contact.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab":
			m.active = (m.active + 1) % viewCount
			return m, nil
		case "shift+tab":
			m.active = (m.active - 1 + viewCount) % viewCount
			return m, nil

		case "0":
			m.active = viewHome
			return m, nil
		case "1":
			m.active = viewAbout
			return m, nil
		case "2":
			m.active = viewSkills
			return m, nil
		case "3":
			m.active = viewProjects
			return m, nil
		case "4":
			m.active = viewContact
			return m, nil
		}
	}

	var cmd tea.Cmd
	switch m.active {
	case viewAbout:
		m.about, cmd = m.about.Update(msg)
		cmds = append(cmds, cmd)
	case viewSkills:
		m.skills, cmd = m.skills.Update(msg)
		cmds = append(cmds, cmd)
	case viewProjects:
		m.projects, cmd = m.projects.Update(msg)
		cmds = append(cmds, cmd)
	case viewContact:
		m.contact, cmd = m.contact.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.renderTabBar() + "\n" + m.renderContent() + "\n" + m.renderStatusBar()
}


func (m Model) renderTabBar() string {
	var parts []string
	n := m.styles.New

	logo := n().
		Foreground(colCyan).
		Bold(true).
		Background(colOverlay).
		Padding(0, 2).
		Render("whoami.ssh")

	parts = append(parts, logo)
	parts = append(parts, m.styles.Dim.Background(colSurface).Render(" "))

	for _, tab := range tabs {
		if tab.id == viewHome {
			continue
		}
		var style lipgloss.Style
		if tab.id == m.active {
			style = m.styles.ActiveTab
		} else {
			style = m.styles.Tab
		}
		parts = append(parts, style.Background(colSurface).Render(tab.label))
	}

	bar := lipgloss.JoinHorizontal(lipgloss.Center, parts...)
	filled := n().
		Background(colSurface).
		Width(m.width).
		Render(bar)

	border := n().
		Foreground(colOverlay).
		Render(strings.Repeat("─", m.width))

	return filled + "\n" + border
}

func (m Model) renderContent() string {
	switch m.active {
	case viewHome:
		return m.home.View()
	case viewAbout:
		return m.about.View()
	case viewSkills:
		return m.skills.View()
	case viewProjects:
		return m.projects.View()
	case viewContact:
		return m.contact.View()
	}
	return ""
}

func (m Model) renderStatusBar() string {
	scrollPct := m.scrollPercent()

	left := fmt.Sprintf("  %s  %s",
		m.styles.Dim.Render("ssh"),
		m.styles.Highlight.Render("github.com/WaxArsatia/whoami.ssh"),
	)

	right := m.styles.Dim.Render(fmt.Sprintf("%s  ", scrollPct))

	pad := m.width - lipgloss.Width(left) - lipgloss.Width(right)
	if pad < 0 {
		pad = 0
	}

	bar := left + strings.Repeat(" ", pad) + right

	return m.styles.New().
		Background(colSurface).
		Foreground(colSubtle).
		Width(m.width).
		Render(bar)
}

func (m Model) scrollPercent() string {
	switch m.active {
	case viewAbout:
		if m.about.ready {
			return pct(m.about.viewport.ScrollPercent())
		}
	case viewSkills:
		if m.skills.ready {
			return pct(m.skills.viewport.ScrollPercent())
		}
	case viewProjects:
		if m.projects.ready {
			return pct(m.projects.viewport.ScrollPercent())
		}
	case viewContact:
		if m.contact.ready {
			return pct(m.contact.viewport.ScrollPercent())
		}
	}
	return ""
}

func pct(f float64) string {
	p := int(f * 100)
	if p == 0 {
		return "top"
	}
	if p >= 100 {
		return "bot"
	}
	return fmt.Sprintf("%d%%", p)
}
