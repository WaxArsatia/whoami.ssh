package tui

import "github.com/charmbracelet/lipgloss"

// ───────────────────────────────────────────────────────────────────────────
// Ayu Dark palette
// ───────────────────────────────────────────────────────────────────────────

const (
	colBg       = lipgloss.Color("#0b0e14") // background
	colSurface  = lipgloss.Color("#131721") // slightly lifted surface
	colOverlay  = lipgloss.Color("#1e232b") // normal black
	colMuted    = lipgloss.Color("#3d424d") // mid-tone
	colSubtle   = lipgloss.Color("#686868") // bright black
	colText     = lipgloss.Color("#bfbdb6") // foreground
	colTextDim  = lipgloss.Color("#8a9199") // dimmed foreground
	colCyan     = lipgloss.Color("#95e6cb") // bright cyan
	colBlue     = lipgloss.Color("#59c2ff") // bright blue
	colPurple   = lipgloss.Color("#d2a6ff") // bright magenta
	colPink     = lipgloss.Color("#f07178") // bright red
	colGreen    = lipgloss.Color("#aad94c") // bright green
	colYellow   = lipgloss.Color("#ffb454") // bright yellow
	colOrange   = lipgloss.Color("#f9af4f") // normal yellow (orange in Ayu)
	colTeal     = lipgloss.Color("#90e1c6") // normal cyan
	colWhite    = lipgloss.Color("#ffffff") // bright white
)

// ───────────────────────────────────────────────────────────────────────────
// Base styles
// ───────────────────────────────────────────────────────────────────────────

var (
	// Full-screen wrapper
	AppStyle = lipgloss.NewStyle().
		Background(colBg).
		Foreground(colText)

	// ── Tab bar ──────────────────────────────────────────────────────────────

	TabBarStyle = lipgloss.NewStyle().
		Background(colSurface).
		Padding(0, 1)

	TabStyle = lipgloss.NewStyle().
		Foreground(colSubtle).
		Padding(0, 2)

	ActiveTabStyle = lipgloss.NewStyle().
		Foreground(colCyan).
		Bold(true).
		Padding(0, 2).
		Underline(true)

	TabDividerStyle = lipgloss.NewStyle().
		Foreground(colMuted)

	// ── Status bar ───────────────────────────────────────────────────────────

	StatusBarStyle = lipgloss.NewStyle().
		Background(colSurface).
		Foreground(colSubtle).
		Padding(0, 1)

	StatusKeyStyle = lipgloss.NewStyle().
		Foreground(colCyan).
		Background(colOverlay).
		Padding(0, 1).
		Margin(0, 1, 0, 0)

	StatusDescStyle = lipgloss.NewStyle().
		Foreground(colSubtle)

	// ── Content area ─────────────────────────────────────────────────────────

	ContentStyle = lipgloss.NewStyle().
		Padding(1, 3)

	// ── Section headings ─────────────────────────────────────────────────────

	SectionTitleStyle = lipgloss.NewStyle().
		Foreground(colCyan).
		Bold(true).
		MarginBottom(1)

	SectionSubtitleStyle = lipgloss.NewStyle().
		Foreground(colSubtle).
		Italic(true)

	// ── Banners & highlights ─────────────────────────────────────────────────

	BannerStyle = lipgloss.NewStyle().
		Foreground(colCyan).
		Bold(true)

	HighlightStyle = lipgloss.NewStyle().
		Foreground(colCyan)

	AccentStyle = lipgloss.NewStyle().
		Foreground(colPink)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(colGreen)

	WarningStyle = lipgloss.NewStyle().
		Foreground(colYellow)

	DimStyle = lipgloss.NewStyle().
		Foreground(colSubtle)

	BoldStyle = lipgloss.NewStyle().
		Foreground(colText).
		Bold(true)

	// ── Boxes & borders ──────────────────────────────────────────────────────

	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colOverlay).
		Padding(1, 2)

	HighlightBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colCyan).
		Padding(1, 2)


	// ── Project card ─────────────────────────────────────────────────────────

	ProjectTitleStyle = lipgloss.NewStyle().
		Foreground(colBlue).
		Bold(true)

	ProjectLangStyle = lipgloss.NewStyle().
		Foreground(colOrange).
		Italic(true)

	ProjectURLStyle = lipgloss.NewStyle().
		Foreground(colSubtle).
		Italic(true)

	TagStyle = lipgloss.NewStyle().
		Foreground(colPurple).
		Background(colOverlay).
		Padding(0, 1).
		Margin(0, 1, 0, 0)

	// ── Home-specific ─────────────────────────────────────────────────────────

	NameStyle = lipgloss.NewStyle().
		Foreground(colWhite).
		Bold(true)

	RoleStyle = lipgloss.NewStyle().
		Foreground(colTeal)

	MottoStyle = lipgloss.NewStyle().
		Foreground(colSubtle).
		Italic(true)

	PromptStyle = lipgloss.NewStyle().
		Foreground(colCyan)
)

// LangColor returns a color for a programming language tag.
func LangColor(lang string) lipgloss.Color {
	switch lang {
	case "Go":
		return colTeal
	case "Rust":
		return colOrange
	case "TypeScript", "JavaScript":
		return colYellow
	default:
		return colBlue
	}
}
