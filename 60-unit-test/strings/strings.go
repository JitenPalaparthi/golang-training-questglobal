package strings

func Reverse(str string) string {
	rstr := ""
	for _, v := range str {
		rstr = string(v) + rstr
	}
	return rstr
}

func ToUpper(str string) string {
	//ustr := ""
	bytes := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= 97 && str[i] <= 122 {
			bytes[i] = str[i] - 32
		} else {
			bytes[i] = str[i]
		}
	}
	return string(bytes)
}

func GetVovelCount(str string) uint {
	var count uint
	for _, v := range str {
		if string(v) == "A" || string(v) == "a" || string(v) == "E" || string(v) == "e" || string(v) == "I" || string(v) == "i" || string(v) == "O" || string(v) == "o" || string(v) == "U" || string(v) == "u" {
			count++
		}
	}
	return count
}
