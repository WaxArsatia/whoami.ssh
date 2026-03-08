package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/WaxArsatia/whoami.ssh/internal/data"
)

type projectsView struct {
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func newProjectsView(w, h int) projectsView {
	contentH := h - tabBarHeight - statusBarHeight
	if contentH < 1 {
		contentH = 24
	}
	vp := viewport.New(w, contentH)
	vp.SetContent(buildProjectsContent(w))
	return projectsView{
		viewport: vp,
		ready:    w > 0 && h > 0,
		width:    w,
		height:   h,
	}
}

func (v projectsView) Init() tea.Cmd { return nil }

func (v projectsView) Update(msg tea.Msg) (projectsView, tea.Cmd) {
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
		v.viewport.SetContent(buildProjectsContent(msg.Width))
		v.ready = true
	}
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v projectsView) View() string {
	return v.viewport.View()
}

func buildProjectsContent(w int) string {
	var sb strings.Builder
	cardW := minInt(w-4, 78)

	sb.WriteString("\n")
	sb.WriteString(SectionTitleStyle.Render("  $ ls -la ~/projects/") + "\n\n")

	for _, proj := range data.Projects {
		card := buildProjectCard(proj, cardW)
		sb.WriteString(card + "\n\n")
	}

	return sb.String()
}

func buildProjectCard(p data.Project, w int) string {
	// Language badge
	lang := lipgloss.NewStyle().
		Foreground(LangColor(p.Lang)).
		Background(colOverlay).
		Padding(0, 1).
		Render(p.Lang)

	titleRow := ProjectTitleStyle.Render(p.Name) + "  " + lang

	// Description
	desc := lipgloss.NewStyle().Foreground(colTextDim).Render(p.Description)

	// Tags
	var tagParts []string
	for _, tag := range p.Tags {
		tagParts = append(tagParts, TagStyle.Render("#"+tag))
	}
	tags := strings.Join(tagParts, "")

	// URL
	url := ProjectURLStyle.Render(p.URL)

	content := titleRow + "\n" + desc + "\n\n" + tags + "\n" + url

	return BoxStyle.Width(w).Render(content)
}
