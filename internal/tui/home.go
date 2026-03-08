package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var asciiArt = strings.Join([]string{
	`██╗    ██╗██╗  ██╗ ██████╗  █████╗ ███╗   ███╗██╗`,
	`██║    ██║██║  ██║██╔═══██╗██╔══██╗████╗ ████║██║`,
	`██║ █╗ ██║███████║██║   ██║███████║██╔████╔██║██║`,
	`██║███╗██║██╔══██║██║   ██║██╔══██║██║╚██╔╝██║██║`,
	`╚███╔███╔╝██║  ██║╚██████╔╝██║  ██║██║ ╚═╝ ██║██║`,
	` ╚══╝╚══╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚═╝`,
}, "\n")

type homeView struct {
	width  int
	height int
	styles Styles
}

func newHomeView(w, h int, st Styles) homeView {
	return homeView{width: w, height: h, styles: st}
}

func (v homeView) View() string {
	w := v.width
	st := v.styles

	var banner string
	if w >= 55 {
		bannerStyle := st.New().Foreground(colCyan).Bold(true)
		suffix := st.New().Foreground(colBlue).Bold(true).Render(".ssh")
		banner = bannerStyle.Render(asciiArt) + "\n" +
			st.New().MarginLeft(44).Render(suffix)
	} else {
		banner = st.New().Foreground(colCyan).Bold(true).Render("whoami.ssh")
	}

	nameStr := st.Name.Render("Denis Arsyatya") +
		st.Dim.Render("  //  ") +
		st.Dim.Render("WaxArsatia")
	roleStr := st.Role.Render("⬡  Backend Developer · Network Engineer · Linux SysAdmin")
	mottoStr := st.Motto.Render(`"In a world of variables, be a constant."`)

	statsRow := buildStatsRow(st)

	divW := minInt(w-6, 70)
	divider := st.Dim.Render(strings.Repeat("─", divW))

	navHint := buildNavHint(st)

	indent := "  "
	sections := []string{
		"",
		banner,
		"",
		indent + nameStr,
		indent + roleStr,
		indent + mottoStr,
		"",
		indent + divider,
		"",
		statsRow,
		"",
		indent + divider,
		"",
		navHint,
		"",
	}

	content := strings.Join(sections, "\n")

	lines := strings.Count(content, "\n") + 1
	topPad := (v.height - lines) / 3
	if topPad < 1 {
		topPad = 1
	}

	return strings.Repeat("\n", topPad) + content
}

func buildNavHint(st Styles) string {
	type kv struct{ key, desc string }
	entries := []kv{
		{"tab / shift+tab", "switch section"},
		{"1-4", "jump to view"},
		{"↑↓ / j k", "scroll"},
		{"q / ctrl+c", "quit"},
	}
	var parts []string
	for _, e := range entries {
		key := st.StatusKey.Render(e.key)
		desc := st.StatusDesc.Render(e.desc)
		parts = append(parts, key+" "+desc)
	}
	return "  " + strings.Join(parts, "  ")
}

func buildStatsRow(st Styles) string {
	items := []struct{ label, value string }{
		{"location", "Indonesia"},
		{"focus", "Reliability · Scalability · Systems"},
		{"org", "ForumLinuxIndonesia"},
	}
	cols := make([]string, 0, len(items))
	for _, item := range items {
		col := fmt.Sprintf("%s\n%s",
			st.Dim.Render("  "+item.label),
			"  "+st.Highlight.Render(item.value),
		)
		cols = append(cols, col)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top,
		st.New().Width(30).Render(cols[0]),
		st.New().Width(38).Render(cols[1]),
		st.New().Width(28).Render(cols[2]),
	)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
