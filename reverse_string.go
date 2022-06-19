package reverse_string

func ReverseString(input string) (output string) {
	for _, s := range input {
		defer func(symbol rune) {
			output += string(symbol)
		}(s)
	}
	return output
}
