package kata

func CountSheeps(numbers []bool) (count int) {
	for _, value := range numbers {
		if value {
			count++
		}
	}
	return count
}
