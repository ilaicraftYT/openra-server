package protocol

func MakeArray(count int, f func(int) byte) []byte {
	result := make([]byte, count)
	for i := 0; i < count; i++ {
		result[i] = f(i)
	}
	return result
}
