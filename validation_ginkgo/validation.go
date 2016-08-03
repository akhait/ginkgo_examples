package validation_ginkgo

import "strings"

var ReservedChars = []rune{'^','@','#','$', '&'}

//Check is a username is longer than 8 characters and does not contain any reserved characters
//On success, returns a greeting string and a nil error
//On error, return a nil string and an error
func RegisterNewUsername(username string) string {
	if len(username) < 8 {
		return "Your username should be longer than 8 characters"
	}
	for _, s := range ReservedChars {
		if strings.ContainsRune(username, s) {
			return "Do not use special characters in your username"
		}
	}
	return "Hello, " + username
}