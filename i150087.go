package golangexamples

func ConcatSlice(sliceToConcat []byte) string {

	s := string(sliceToConcat)

	return s

}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {

	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] = sliceToEncrypt[i] + 3

	}

	return sliceToEncrypt

}
