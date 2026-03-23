package tranformer

import (
	"slices"
	"strings"
)

func (at AsciiTransform) FormatAsciiWords(asciiWords [][][]string) string {
	var ascii strings.Builder

	for i := 0; i < len(asciiWords); i++ {
		current := asciiWords[i]
		if slices.Compare(current[0], []string{""}) == 0 {
			ascii.WriteString("\n")
			continue
		}
		for j := 0; j < 8; j++ {
			for k := 0; k < len(current); k++ {
				if k != len(current)-1 {
					ascii.WriteString(current[k][j])
				} else {
					ascii.WriteString(current[k][j] + "\n")
				}
			}
		}
		ascii.WriteString("\n")
	}
	return ascii.String()
}
