package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/WaxArsatia/whoami.ssh/internal/data"
)

type skillsView struct {
	viewport viewport.Model
	ready    bool
	width    int
	height   int
	styles   Styles
}

func newSkillsView(w, h int, st Styles) skillsView {
	contentH := h - tabBarHeight - statusBarHeight
	if contentH < 1 {
		contentH = 24
	}
	vp := viewport.New(w, contentH)
	vp.SetContent(buildSkillsContent(w, st))
	return skillsView{
		viewport: vp,
		ready:    w > 0 && h > 0,
		width:    w,
		height:   h,
		styles:   st,
	}
}

func (v skillsView) Init() tea.Cmd { return nil }

func (v skillsView) Update(msg tea.Msg) (skillsView, tea.Cmd) {
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
		v.viewport.SetContent(buildSkillsContent(msg.Width, v.styles))
		v.ready = true
	}
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v skillsView) View() string {
	return v.viewport.View()
}

func buildSkillsContent(w int, st Styles) string {
	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString(st.SectionTitle.Render("  $ cat tech-stack.json") + "\n\n")

	pillStyle := st.New().
		Foreground(colCyan).
		Background(colOverlay).
		Padding(0, 1).
		MarginRight(1)

	categoryStyle := st.New().
		Foreground(colSubtle).
		Bold(true)

	for i, group := range data.TechStack {
		heading := categoryStyle.Render("  " + group.Category)
		sb.WriteString(heading + "\n  ")

		lineUsed := 2
		for j, skill := range group.Skills {
			pill := pillStyle.Render(skill)
			pillW := lipgloss.Width(pill) + 1
			if j > 0 && lineUsed+pillW > w-4 {
				sb.WriteString("\n  ")
				lineUsed = 2
			}
			sb.WriteString(pill + " ")
			lineUsed += pillW
		}
		sb.WriteString("\n")

		if i < len(data.TechStack)-1 {
			sb.WriteString("\n")
		}
	}

	sb.WriteString("\n")
	return sb.String()
}
