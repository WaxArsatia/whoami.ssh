package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/WaxArsatia/whoami.ssh/internal/data"
)

type contactView struct {
	viewport viewport.Model
	ready    bool
	width    int
	height   int
	styles   Styles
}

func newContactView(w, h int, st Styles) contactView {
	contentH := h - tabBarHeight - statusBarHeight
	if contentH < 1 {
		contentH = 24
	}
	vp := viewport.New(w, contentH)
	vp.SetContent(buildContactContent(w, st))
	return contactView{
		viewport: vp,
		ready:    w > 0 && h > 0,
		width:    w,
		height:   h,
		styles:   st,
	}
}

func (v contactView) Init() tea.Cmd { return nil }

func (v contactView) Update(msg tea.Msg) (contactView, tea.Cmd) {
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
		v.viewport.SetContent(buildContactContent(msg.Width, v.styles))
		v.ready = true
	}
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v contactView) View() string {
	return v.viewport.View()
}

func buildContactContent(w int, st Styles) string {
	p := data.Me
	cardW := minInt(w-6, 60)

	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString(st.SectionTitle.Render("  $ curl -s contact.json | jq") + "\n\n")

	contacts := []struct {
		icon, label, value, color string
	}{
		{"📧", "email", p.Email, string(colCyan)},
		{"🐙", "github", p.GitHub, string(colBlue)},
		{"💼", "linkedin", p.LinkedIn, string(colBlue)},
		{"⏱️ ", "wakatime", p.WakaTime, string(colTeal)},
	}

	var lines []string
	for _, c := range contacts {
		icon := c.icon + " "
		label := st.New().Foreground(colSubtle).Width(12).Render(c.label)
		sep := st.Dim.Render(": ")
		val := st.New().Foreground(lipgloss.Color(c.color)).Render(c.value)
		lines = append(lines, fmt.Sprintf("  %s%s%s%s", icon, label, sep, val))
	}

	card := st.HighlightBox.Width(cardW).Render(strings.Join(lines, "\n"))
	sb.WriteString(card + "\n\n")

	sb.WriteString(st.SectionTitle.Render("  # organizations") + "\n\n")
	for _, org := range p.Orgs {
		sb.WriteString(fmt.Sprintf("  %s  %s\n",
			st.Accent.Render("◆"),
			st.New().Foreground(colText).Render(org),
		))
	}
	sb.WriteString("\n")

	divW := minInt(w-6, 60)
	sb.WriteString("  " + st.Dim.Render(strings.Repeat("─", divW)) + "\n\n")
	msg := st.Motto.Render(`"The best code is the code that runs reliably at 3 AM without waking anyone up."`)
	sb.WriteString("  " + msg + "\n\n")
	sb.WriteString("  " + st.Dim.Render("— Denis Arsyatya") + "\n\n")

	return sb.String()
}
