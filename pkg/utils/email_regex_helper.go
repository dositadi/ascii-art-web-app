package utils

import "regexp"

func IsEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z][a-z0-9_\-\.]*@[a-z]+\.[a-z]{2,}`)

	return re.MatchString(email)
}
