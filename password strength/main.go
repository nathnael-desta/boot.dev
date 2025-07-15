package main

func isValidPassword(password string) bool {
	if len(password) < 4 || len(password) > 12 {
		return false
	}
	hasDigit := false
	hasUpper := false
	for i := 0; i < len(password); i++ {
		if 48 <= int(password[i]) && int(password[i]) <= 57 {
			hasDigit = true
			continue
		}
		if 65 <= int(password[i]) && int(password[i]) <= 90 {
			hasUpper = true
		}
	}
	return hasDigit && hasUpper
}
