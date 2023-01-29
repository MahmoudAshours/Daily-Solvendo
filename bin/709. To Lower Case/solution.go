package main
func main(){
	print(toLowerCase("Hello"))
}

func toLowerCase(s string) string {
	bs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bs[i] = toLower(s[i])
	}
	return string(bs)
}

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + 32
	}
	return b
}

// Using builtin function

// func toLowerCase(s string) string {
//     return strings.ToLower(s)
// }