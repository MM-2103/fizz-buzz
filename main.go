package main


func main() {
	for i := 1; i < 101; i++ {
		if i % 5 == 0 && i % 10 != 0  {
			println("Fizz")
		} else if i % 10 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
	return
}
