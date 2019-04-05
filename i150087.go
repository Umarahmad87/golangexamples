package golangexamples

func ConcatSlice(sliceToConcat []byte) string {

	//s := string(sliceToConcat)
	s1 := ""
	for i:=0;i<len(sliceToConcat);i++ {

		s1 += string(sliceToConcat[i])
		if (i!=len(sliceToConcat)-1){
			s1 += "-"}
	}

	return s1

}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {

	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] = sliceToEncrypt[i] + 3

	}

	return sliceToEncrypt

}
