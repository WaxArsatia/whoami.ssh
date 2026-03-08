package tui

import "github.com/charmbracelet/lipgloss"

// ───────────────────────────────────────────────────────────────────────────
// Ayu Dark palette
// ───────────────────────────────────────────────────────────────────────────

const (
	colBg      = lipgloss.Color("#0b0e14")
	colSurface = lipgloss.Color("#131721")
	colOverlay = lipgloss.Color("#1e232b")
	colMuted   = lipgloss.Color("#3d424d")
	colSubtle  = lipgloss.Color("#686868")
	colText    = lipgloss.Color("#bfbdb6")
	colTextDim = lipgloss.Color("#8a9199")
	colCyan    = lipgloss.Color("#95e6cb")
	colBlue    = lipgloss.Color("#59c2ff")
	colPurple  = lipgloss.Color("#d2a6ff")
	colPink    = lipgloss.Color("#f07178")
	colGreen   = lipgloss.Color("#aad94c")
	colYellow  = lipgloss.Color("#ffb454")
	colOrange  = lipgloss.Color("#f9af4f")
	colTeal    = lipgloss.Color("#90e1c6")
	colWhite   = lipgloss.Color("#ffffff")
)

// ───────────────────────────────────────────────────────────────────────────
// Per-session styles
// ───────────────────────────────────────────────────────────────────────────

type Styles struct {
	r *lipgloss.Renderer

	App lipgloss.Style

	TabBar     lipgloss.Style
	Tab        lipgloss.Style
	ActiveTab  lipgloss.Style
	TabDivider lipgloss.Style

	StatusBar  lipgloss.Style
	StatusKey  lipgloss.Style
	StatusDesc lipgloss.Style

	Content lipgloss.Style

	SectionTitle    lipgloss.Style
	SectionSubtitle lipgloss.Style

	Banner    lipgloss.Style
	Highlight lipgloss.Style
	Accent    lipgloss.Style

	Success lipgloss.Style
	Warning lipgloss.Style
	Dim     lipgloss.Style
	Bold    lipgloss.Style

	Box          lipgloss.Style
	HighlightBox lipgloss.Style

	ProjectTitle lipgloss.Style
	ProjectLang  lipgloss.Style
	ProjectURL   lipgloss.Style
	Tag          lipgloss.Style

	Name   lipgloss.Style
	Role   lipgloss.Style
	Motto  lipgloss.Style
	Prompt lipgloss.Style
}

func NewStyles(r *lipgloss.Renderer) Styles {
	n := r.NewStyle

	return Styles{
		r: r,

		App: n().Background(colBg).Foreground(colText),

		// ── Tab bar ──────────────────────────────────────────────────────────
		TabBar:     n().Background(colSurface).Padding(0, 1),
		Tab:        n().Foreground(colSubtle).Padding(0, 2),
		ActiveTab:  n().Foreground(colCyan).Bold(true).Padding(0, 2).Underline(true),
		TabDivider: n().Foreground(colMuted),

		// ── Status bar ───────────────────────────────────────────────────────
		StatusBar:  n().Background(colSurface).Foreground(colSubtle).Padding(0, 1),
		StatusKey:  n().Foreground(colCyan).Background(colOverlay).Padding(0, 1).Margin(0, 1, 0, 0),
		StatusDesc: n().Foreground(colSubtle),

		// ── Content area ────────────────────────────────────────────────────
		Content: n().Padding(1, 3),

		// ── Section headings ─────────────────────────────────────────────────
		SectionTitle:    n().Foreground(colCyan).Bold(true).MarginBottom(1),
		SectionSubtitle: n().Foreground(colSubtle).Italic(true),

		// ── Banners & highlights ─────────────────────────────────────────────
		Banner:    n().Foreground(colCyan).Bold(true),
		Highlight: n().Foreground(colCyan),
		Accent:    n().Foreground(colPink),

		// ── States ───────────────────────────────────────────────────────────
		Success: n().Foreground(colGreen),
		Warning: n().Foreground(colYellow),
		Dim:     n().Foreground(colSubtle),
		Bold:    n().Foreground(colText).Bold(true),

		// ── Boxes & borders ──────────────────────────────────────────────────
		Box:          n().Border(lipgloss.RoundedBorder()).BorderForeground(colOverlay).Padding(1, 2),
		HighlightBox: n().Border(lipgloss.RoundedBorder()).BorderForeground(colCyan).Padding(1, 2),

		// ── Project card ─────────────────────────────────────────────────────
		ProjectTitle: n().Foreground(colBlue).Bold(true),
		ProjectLang:  n().Foreground(colOrange).Italic(true),
		ProjectURL:   n().Foreground(colSubtle).Italic(true),
		Tag:          n().Foreground(colPurple).Background(colOverlay).Padding(0, 1).Margin(0, 1, 0, 0),

		// ── Home-specific ─────────────────────────────────────────────────────
		Name:   n().Foreground(colWhite).Bold(true),
		Role:   n().Foreground(colTeal),
		Motto:  n().Foreground(colSubtle).Italic(true),
		Prompt: n().Foreground(colCyan),
	}
}

func (s Styles) New() lipgloss.Style { return s.r.NewStyle() }

// ───────────────────────────────────────────────────────────────────────────
// Helpers
// ───────────────────────────────────────────────────────────────────────────

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
