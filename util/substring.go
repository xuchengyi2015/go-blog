package util

func SubString(str string, begin, length int) string {
	//fmt.Printf("substring=%s\n", str)
	rs := []rune(str)
	lth := len(rs)

	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	//fmt.Printf("begin=%d,end=%d,lth=%d\n", begin, end, lth)
	return string(rs[begin:end])
}
