package main

func main() {
	var p1 Person
	p1.int = 1001
	p1.string = "Jiten"
	p1.bool = true
}

type Item struct {
	int
	string
}

type Person struct {
	int
	string
	bool
}
