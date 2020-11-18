package util

//IsLetter returns if a character is a letter.
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

//IsWhitespace returns if a character is whitespace or not.
func IsWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
