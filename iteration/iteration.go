package iteration

func Repeat(s string, times int) string {

	var result string
	for i := 0; i < times; i++ {
		result += s
	}

	return result
}
