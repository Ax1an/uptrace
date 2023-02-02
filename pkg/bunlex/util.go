package bunlex

func IsWhitespace(c byte) bool {
	switch c {
	case ' ', '\t':
		return true
	default:
		return false
	}
}

func IsAlnum(c byte) bool {
	return IsAlpha(c) || IsDigit(c)
}

func IsAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
