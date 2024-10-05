package iteration

func Repeat(character string, times int) (repeated string) {
	// var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return
}
