package matematika

func SumAll(num ...int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}
