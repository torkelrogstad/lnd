package build

import "strings"

var tags string

func Tags() []string {
	if strings.TrimSpace(tags) == "" {
		return nil
	}
	var out []string
	splitted := strings.Split(tags, " ")
	for _, elem := range splitted {
		if elem != "" {
			out = append(out, elem)
		}
	}
	return out
}
