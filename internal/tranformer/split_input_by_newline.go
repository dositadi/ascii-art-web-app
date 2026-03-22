package tranformer

import (
	"bufio"
	"fmt"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (at *AsciiTransform) SplitInputByNewline(input string) ([]string, *m.Error) {
	var output []string

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		output = append(output, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, &m.Error{
			Error:   h.PROCESS_TEXT_ERR,
			Details: h.PROCESS_TEXT_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}
	fmt.Println(output)
	return output, nil
}
