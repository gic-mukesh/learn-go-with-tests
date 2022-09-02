package iteration

func repeat(character string, iteration int) string {

	var repeated string

	for i := 0; i < iteration; i++ {
		repeated += character
	}

	return repeated
}
