package testhelpers

func TrimInput(input string) string {
	if len(input) > 20 {
		return input[:20] + "..."
	} else {
		return input
	}
}
