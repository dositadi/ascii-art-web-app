package tranformer

import (
	"bufio"
	"fmt"
	"os"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (at *AsciiTransform) ReadAsciiFromFont(rn rune, banner string) ([]string, *m.Error) {
	var output []string
	fontPath := ""

	switch banner {
	case "standard":
		fontPath = "fonts/standard.txt"
	case "shadow":
		fontPath = "fonts/shadow.txt"
	case "thinkertoy":
		fontPath = "fonts/thinker_toy.txt"
	}

	fmt.Println("in read: ", string(rn), banner)

	file, err := os.Open(fontPath)
	if err != nil {
		return nil, &m.Error{
			Error:   h.PROCESS_TEXT_ERR,
			Details: h.PROCESS_TEXT_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}

	calcEndline := func(startLine int) int {
		return startLine + 9
	}

	scanner := bufio.NewScanner(file)
	startLine := at.CalculateStartLine(rn)
	currentLine := 0
	endLine := calcEndline(startLine)
	fmt.Println(startLine, endLine)

	for scanner.Scan() {
		if currentLine >= startLine && currentLine <= endLine {
			line := scanner.Text()

			output = append(output, line)
		} else if startLine < currentLine || currentLine > endLine {
			break
		}
		currentLine = currentLine + 1
	}

	if err2 := scanner.Err(); err2 != nil {
		return nil, &m.Error{
			Error:   h.PROCESS_TEXT_ERR,
			Details: h.PROCESS_TEXT_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return output, nil
}
