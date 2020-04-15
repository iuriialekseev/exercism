package collatzconjecture

func CollatzConjecture(i int) (int, error) {
	count := 0

	for {
		switch {
		case i%2 == 0:
			i /= 2
		case i == 1:
			return count, nil
		default:
			i = i*3 + 1
		}
		count++
	}
}
