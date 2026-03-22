package tranformer

import (
	"fmt"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (at *AsciiTransform) ReadWords(input []string, banner string) ([][][]string, *m.Error) {
	var output [][][]string

	for i := 0; i < len(input); i++ {
		current := input[i]
		var wordsAsciiChar [][]string
		if strings.Compare(current, "") == 0 {
			wordsAsciiChar = append(wordsAsciiChar, []string{""})
			output = append(output, wordsAsciiChar)
			fmt.Println("Entered")
			continue
		}
		for _, rn := range current {
			fmt.Println(string(rn))
			asciiChar, err := at.ReadAsciiFromFont(rn, banner)
			if err != nil {
				return nil, &m.Error{
					Error:   h.PROCESS_TEXT_ERR,
					Details: h.PROCESS_TEXT_ERR_DETAIL,
					Code:    h.SERVER_ERR_CODE,
				}
			}
			wordsAsciiChar = append(wordsAsciiChar, asciiChar)
		}
		output = append(output, wordsAsciiChar)
	}
	return output, nil
}
