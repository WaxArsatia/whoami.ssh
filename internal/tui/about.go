package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/WaxArsatia/whoami.ssh/internal/data"
)

type aboutView struct {
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func newAboutView(w, h int) aboutView {
	contentH := h - tabBarHeight - statusBarHeight
	if contentH < 1 {
		contentH = 24
	}
	vp := viewport.New(w, contentH)
	vp.SetContent(buildAboutContent(w))
	return aboutView{
		viewport: vp,
		ready:    w > 0 && h > 0,
		width:    w,
		height:   h,
	}
}

func (v aboutView) Init() tea.Cmd { return nil }

func (v aboutView) Update(msg tea.Msg) (aboutView, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		v.width = msg.Width
		v.height = msg.Height
		contentH := msg.Height - tabBarHeight - statusBarHeight
		if contentH < 1 {
			contentH = 1
		}
		v.viewport.Width = msg.Width
		v.viewport.Height = contentH
		v.viewport.SetContent(buildAboutContent(msg.Width))
		v.ready = true
	}
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v aboutView) View() string {
	return v.viewport.View()
}

func buildAboutContent(w int) string {
	p := data.Me
	var sb strings.Builder

	// ── Header ──────────────────────────────────────────────────────────────
	sb.WriteString("\n")
	sb.WriteString(SectionTitleStyle.Render("  $ cat profile.toml") + "\n\n")

	// Profile card
	cardW := minInt(w-6, 72)
	card := buildProfileCard(p, cardW)
	sb.WriteString(card + "\n\n")

	// ── About entries ────────────────────────────────────────────────────────
	sb.WriteString(SectionTitleStyle.Render("  # highlights") + "\n\n")
	for i, item := range p.About {
		icon := getAboutIcon(i)
		line := fmt.Sprintf("  %s  %s\n", icon, item)
		sb.WriteString(line)
	}

	sb.WriteString("\n")
	return sb.String()
}

func buildProfileCard(p data.Profile, w int) string {
	rows := []struct{ key, val string }{
		{"name", p.Name},
		{"alias", p.Alias},
		{"role", p.Role},
		{"location", "Indonesia"},
		{"email", p.Email},
		{"github", p.GitHub},
		{"linkedin", p.LinkedIn},
		{"wakatime", p.WakaTime},
		{"motto", p.Motto},
	}

	var lines []string
	for _, r := range rows {
		keyStr := lipgloss.NewStyle().Foreground(colCyan).Width(12).Render(r.key)
		sep := DimStyle.Render(" = ")
		valStr := lipgloss.NewStyle().Foreground(colText).Render(r.val)
		lines = append(lines, fmt.Sprintf("  %s%s%s", keyStr, sep, valStr))
	}

	box := BoxStyle.Width(w).Render(strings.Join(lines, "\n"))
	return box
}

func getAboutIcon(i int) string {
	icons := []string{"🔧", "🦀", "🐹", "🌐", "🐧", "📦", "🚀"}
	if i < len(icons) {
		return AccentStyle.Render(icons[i])
	}
	return AccentStyle.Render("•")
}
