package iteration

func Repeat(r string) string {
	var repated string
	for i := 0; i < 5; i++ {
		repated = repated + r
	}

	return repated
}
