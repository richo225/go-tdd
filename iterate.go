package iteration

func Iterate(statement string, count int) string {
	var output string

	for i := 0; i < count; i++ {
		output += statement
	}

	return output
}
