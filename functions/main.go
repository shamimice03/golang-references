package main

func buffer(s string) (int, error) {
	b := make([]string, ' ', 100)
	//print(b[0])
	print(len(b))
	for i := 0; i < 2; i++ {
		read(b, s)
		//print(b[0])
	}

	return len(b), nil
}

func read(buf []string, s string) int {
	buf = append(buf, s)
	return len(buf)
}

func main() {

	words := []string{"Hello", "World"}

	for _, value := range words {
		//fmt.Println(value)
		buffer(value)
	}
}
