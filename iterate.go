package iteration

func Iterate(statement string) string {
	var output string

	for i := 0; i < 3; i++ {
		output += statement
	}

	return output
}
