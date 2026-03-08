package data

// ───────────────────────────────────────────────────────────────────────────
// Profile
// ───────────────────────────────────────────────────────────────────────────

type Profile struct {
	Name      string
	Alias     string
	Role      string
	Motto     string
	Location  string
	Email     string
	GitHub    string
	LinkedIn  string
	WakaTime  string
	About     []string
	Orgs      []string
}

var Me = Profile{
	Name:     "Denis Arsyatya",
	Alias:    "WaxArsatia",
	Role:     "Backend Developer · Network Engineer · Linux SysAdmin",
	Motto:    `"In a world of variables, be a constant."`,
	Location: "Indonesia",
	Email:    "waxarsatia@gmail.com",
	GitHub:   "github.com/WaxArsatia",
	LinkedIn: "linkedin.com/in/denis-arsyatya",
	WakaTime: "wakatime.com/@WaxArsatia",
	Orgs:     []string{"ForumLinuxIndonesia"},
	About: []string{
		"Backend — Designing and building robust REST APIs, microservices, and server-side logic",
		"Systems — Low-level systems programming with Rust; SSH-over-WebSocket, performance-critical tools",
		"Go — Building concurrent, high-performance services and CLI tools with Go",
		"Networking — Network design, configuration, and infrastructure management",
		"Linux — System administration, automation, shell scripting, and active member of ForumLinuxIndonesia",
		"JS/TS Ecosystem — Deep expertise in TypeScript, Node.js, React, and Next.js",
		"Always exploring new languages, tools, and best practices to sharpen the craft",
	},
}

// ───────────────────────────────────────────────────────────────────────────
// Tech Stack
// ───────────────────────────────────────────────────────────────────────────

type SkillGroup struct {
	Category string
	Skills   []string
}

var TechStack = []SkillGroup{
	{
		Category: "Languages",
		Skills:   []string{"TypeScript", "Go", "Rust", "JavaScript", "Node.js"},
	},
	{
		Category: "Frameworks & Libraries",
		Skills:   []string{"React", "Next.js", "Express", "Bubble Tea / Wish"},
	},
	{
		Category: "Infrastructure & DevOps",
		Skills:   []string{"Linux", "Bash / Shell", "Docker", "Nginx", "Git / GitHub"},
	},
	{
		Category: "Databases",
		Skills:   []string{"PostgreSQL", "MongoDB", "MySQL", "Redis"},
	},
	{
		Category: "Also Familiar With",
		Skills:   []string{"Python", "Java", "PHP"},
	},
}

// ───────────────────────────────────────────────────────────────────────────
// Projects
// ───────────────────────────────────────────────────────────────────────────

type Project struct {
	Name        string
	Description string
	Lang        string
	URL         string
	Stars       int
	Tags        []string
}

var Projects = []Project{
	{
		Name:        "SSHWS-Rust",
		Description: "SSH over WebSocket tunnel — bridges SSH connections through WebSocket, enabling SSH in environments where only HTTP/S is allowed.",
		Lang:        "Rust",
		URL:         "github.com/WaxArsatia/SSHWS-Rust",
		Stars:       1,
		Tags:        []string{"ssh", "websocket", "networking", "rust"},
	},
	{
		Name:        "pseudocode-converter",
		Description: "AI-powered tool that converts pseudocode (Indonesian academic format) into runnable C++ code, built with Next.js and OpenAI.",
		Lang:        "TypeScript",
		URL:         "github.com/WaxArsatia/pseudocode-converter",
		Stars:       0,
		Tags:        []string{"ai", "nextjs", "typescript", "education"},
	},
	{
		Name:        "next-overmath",
		Description: "Interactive math problem-solving platform built with Next.js, featuring step-by-step solutions and LaTeX rendering.",
		Lang:        "TypeScript",
		URL:         "github.com/WaxArsatia/next-overmath",
		Stars:       0,
		Tags:        []string{"math", "nextjs", "typescript", "education"},
	},
	{
		Name:        "auctores-discord-bot",
		Description: "Feature-rich Discord bot for community management with moderation, utilities, and automated workflows.",
		Lang:        "TypeScript",
		URL:         "github.com/WaxArsatia/auctores-discord-bot",
		Stars:       0,
		Tags:        []string{"discord", "bot", "typescript", "automation"},
	},
	{
		Name:        "whoami.ssh",
		Description: "This SSH portfolio app! An interactive terminal portfolio served over SSH, built with charmbracelet/wish and Bubble Tea.",
		Lang:        "Go",
		URL:         "github.com/WaxArsatia/whoami.ssh",
		Stars:       0,
		Tags:        []string{"go", "ssh", "tui", "portfolio", "bubbletea"},
	},
}
