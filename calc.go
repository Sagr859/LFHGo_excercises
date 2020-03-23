package main

import "fmt"

func main() {
	var op byte = ' '
	fmt.Println("|||||||| BINARY CALCULATOR SG19 |||||||\n Enter 2 integers and operator to calculation\n")
	var v1, v2 float64
	fmt.Printf("Enter Value: ")
	n, err := fmt.Scanf("%v\n", &v1) //to prevent new line error while enetering data
	if err != nil || n != 1 {
		fmt.Println(n, err)
		return
	}

	fmt.Printf("Enter Operator: ")
	n1, err1 := fmt.Scanf("%c\n", &op)
	if err1 != nil || n1 != 1 {
		fmt.Println(n1, err1)
		return
	}

	fmt.Printf("Enter Value: ")
	n2, err2 := fmt.Scanf("%v\n", &v2)
	if err != nil || n != 1 {
		fmt.Println(n2, err2)
		return
	}

	switch {
	case op == '+':
		v1 += v2
	case op == '-':
		v1 -= v2
	case op == '*':
		v1 *= v2
	case op == '/':
		v1 /= v2
	default:
		fmt.Println("Invalid operator")
	}

	fmt.Println("Calculated value is ", v1)
	fmt.Println("Calculator Terminated")
}
