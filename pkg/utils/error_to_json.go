package utils

import (
	"encoding/json"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func ErrorToJson(err m.Error) []byte {
	output, _ := json.Marshal(err)

	return output
}
