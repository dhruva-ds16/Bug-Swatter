package Validator

import (
	"regexp"
)

var emailRegexp = regexp.MustCompile("(?i)" + "^[a-z0-9!#$%&'*+/=?^_`{|}~.-]+" + "@" + "[a-z0-9-]+(\\.[a-z0-9-]+)*$") //MustCompile arguments are in this order 1. Makes it case insensitive 2. Validates local part 3. Validates domain part

func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}
