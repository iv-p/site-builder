package templateloader

// Template holds data for a single template file
type Template struct {
	Template string
	Scripts  map[string]bool
	CSS      string
}
