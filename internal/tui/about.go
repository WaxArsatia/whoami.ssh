package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/WaxArsatia/whoami.ssh/internal/data"
)

type aboutView struct {
	viewport viewport.Model
	ready    bool
	width    int
	height   int
	styles   Styles
}

func newAboutView(w, h int, st Styles) aboutView {
	contentH := h - tabBarHeight - statusBarHeight
	if contentH < 1 {
		contentH = 24
	}
	vp := viewport.New(w, contentH)
	vp.SetContent(buildAboutContent(w, st))
	return aboutView{
		viewport: vp,
		ready:    w > 0 && h > 0,
		width:    w,
		height:   h,
		styles:   st,
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
		v.viewport.SetContent(buildAboutContent(msg.Width, v.styles))
		v.ready = true
	}
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v aboutView) View() string {
	return v.viewport.View()
}

func buildAboutContent(w int, st Styles) string {
	p := data.Me
	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString(st.SectionTitle.Render("  $ cat profile.toml") + "\n\n")

	cardW := minInt(w-6, 72)
	card := buildProfileCard(p, cardW, st)
	sb.WriteString(card + "\n\n")

	sb.WriteString(st.SectionTitle.Render("  # highlights") + "\n\n")
	for i, item := range p.About {
		icon := getAboutIcon(i, st)
		line := fmt.Sprintf("  %s  %s\n", icon, item)
		sb.WriteString(line)
	}

	sb.WriteString("\n")
	return sb.String()
}

func buildProfileCard(p data.Profile, w int, st Styles) string {
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
		keyStr := st.New().Foreground(colCyan).Width(12).Render(r.key)
		sep := st.Dim.Render(" = ")
		valStr := st.New().Foreground(colText).Render(r.val)
		lines = append(lines, fmt.Sprintf("  %s%s%s", keyStr, sep, valStr))
	}

	box := st.Box.Width(w).Render(strings.Join(lines, "\n"))
	return box
}

func getAboutIcon(i int, st Styles) string {
	icons := []string{"🔧", "🦀", "🐹", "🌐", "🐧", "📦", "🚀"}
	if i < len(icons) {
		return st.Accent.Render(icons[i])
	}
	return st.Accent.Render("•")
}
