package tui

import (
"fmt"
"strings"

"github.com/charmbracelet/lipgloss"
)

// asciiArt is the whoami.ssh full-size banner (requires >= 55 cols)
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
}

func newHomeView(w, h int) homeView {
return homeView{width: w, height: h}
}

func (v homeView) View() string {
w := v.width

// Banner
var banner string
if w >= 55 {
bannerStyle := lipgloss.NewStyle().Foreground(colCyan).Bold(true)
suffix := lipgloss.NewStyle().Foreground(colBlue).Bold(true).Render(".sh")
banner = bannerStyle.Render(asciiArt) + "\n" +
lipgloss.NewStyle().MarginLeft(44).Render(suffix)
} else {
banner = lipgloss.NewStyle().Foreground(colCyan).Bold(true).Render("whoami.ssh")
}

// Identity block
nameStr := NameStyle.Render("Denis Arsyatya") +
DimStyle.Render("  //  ") +
DimStyle.Render("WaxArsatia")
roleStr := RoleStyle.Render("\u2b21  Backend Developer \u00b7 Network Engineer \u00b7 Linux SysAdmin")
mottoStr := MottoStyle.Render(`"In a world of variables, be a constant."`)

// Stats
statsRow := buildStatsRow()

// Dividers
divW := minInt(w-6, 70)
divider := DimStyle.Render(strings.Repeat("\u2500", divW))

// Nav hint
navHint := buildNavHint()

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

func buildNavHint() string {
type kv struct{ key, desc string }
entries := []kv{
{"tab / shift+tab", "switch section"},
{"1-4", "jump to view"},
{"\u2191\u2193 / j k", "scroll"},
{"q / ctrl+c", "quit"},
}
var parts []string
for _, e := range entries {
key := StatusKeyStyle.Render(e.key)
desc := StatusDescStyle.Render(e.desc)
parts = append(parts, key+" "+desc)
}
return "  " + strings.Join(parts, "  ")
}

func buildStatsRow() string {
items := []struct{ label, value string }{
{"location", "Indonesia"},
{"focus", "Reliability \u00b7 Scalability \u00b7 Systems"},
{"org", "ForumLinuxIndonesia"},
}
cols := make([]string, 0, len(items))
for _, item := range items {
col := fmt.Sprintf("%s\n%s",
DimStyle.Render("  "+item.label),
"  "+HighlightStyle.Render(item.value),
)
cols = append(cols, col)
}
return lipgloss.JoinHorizontal(lipgloss.Top,
lipgloss.NewStyle().Width(30).Render(cols[0]),
lipgloss.NewStyle().Width(38).Render(cols[1]),
lipgloss.NewStyle().Width(28).Render(cols[2]),
)
}

func minInt(a, b int) int {
if a < b {
return a
}
return b
}
