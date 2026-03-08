package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/WaxArsatia/whoami.ssh/internal/data"
)

type projectsView struct {
	viewport viewport.Model
	ready    bool
	width    int
	height   int
	styles   Styles
}

func newProjectsView(w, h int, st Styles) projectsView {
	contentH := h - tabBarHeight - statusBarHeight
	if contentH < 1 {
		contentH = 24
	}
	vp := viewport.New(w, contentH)
	vp.SetContent(buildProjectsContent(w, st))
	return projectsView{
		viewport: vp,
		ready:    w > 0 && h > 0,
		width:    w,
		height:   h,
		styles:   st,
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
		v.viewport.SetContent(buildProjectsContent(msg.Width, v.styles))
		v.ready = true
	}
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v projectsView) View() string {
	return v.viewport.View()
}

func buildProjectsContent(w int, st Styles) string {
	var sb strings.Builder
	cardW := minInt(w-4, 78)

	sb.WriteString("\n")
	sb.WriteString(st.SectionTitle.Render("  $ ls -la ~/projects/") + "\n\n")

	for _, proj := range data.Projects {
		card := buildProjectCard(proj, cardW, st)
		sb.WriteString(card + "\n\n")
	}

	return sb.String()
}

func buildProjectCard(p data.Project, w int, st Styles) string {
	lang := st.New().
		Foreground(LangColor(p.Lang)).
		Background(colOverlay).
		Padding(0, 1).
		Render(p.Lang)

	titleRow := st.ProjectTitle.Render(p.Name) + "  " + lang

	desc := st.New().Foreground(colTextDim).Render(p.Description)

	var tagParts []string
	for _, tag := range p.Tags {
		tagParts = append(tagParts, st.Tag.Render("#"+tag))
	}
	tags := strings.Join(tagParts, "")

	url := st.ProjectURL.Render(p.URL)

	content := titleRow + "\n" + desc + "\n\n" + tags + "\n" + url

	return st.Box.Width(w).Render(content)
}
