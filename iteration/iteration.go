package iterations

func Repeat(character string, repetitions int) string {
	var text string

	for i := 0; i < repetitions; i++ {
		text += character
	}

	return text
}