package cli

import "regexp"

func parseTemplate(tokens map[string]string, buffer string) string {
	re := regexp.MustCompile("(@\\w*)+")

	return re.ReplaceAllStringFunc(buffer, func(s string) string {
		return tokens[s]
	})
}
