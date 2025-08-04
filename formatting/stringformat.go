package formatting

import "strings"

// For input "first.last", this function returns "First Last".
func FormatName(name string) string {
	if name == "" {
		return "Unknown"
	}

	parts := strings.Split(name, ".")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, " ")
}
