package cli

import "regexp"

func parseTemplate(tokens map[string]string, buffer string) string {
	if len(tokens) == 0 {
		return buffer
	}

	re := regexp.MustCompile("(@\\w*)+")

	return re.ReplaceAllStringFunc(buffer, func(s string) string {
		return tokens[s]
	})
}
