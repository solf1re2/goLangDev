package problem

import "fmt"

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func Four() {
	fmt.Println("hi")
	biggest := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			res := i * j
			// fmt.Println(res)
			resString := fmt.Sprint(res)
			// fmt.Println(resString)
			resLength := len(resString)
			// fmt.Println(resLength)
			firsthalf := resString[:resLength/2]
			// fmt.Println(firsthalf)
			secondhalf := resString[resLength/2:]
			// fmt.Println(firsthalf + " - " + secondhalf)
			secondhalfReversed := Reverse(secondhalf)
			if firsthalf == secondhalfReversed {
				if res > biggest {
					fmt.Printf("%v (from %v and %v )", res, i, j)
					biggest = res
				}
			}
		}
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
