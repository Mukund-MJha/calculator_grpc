package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	calculator "josh/calculator_grpc/calculatorpb"
	"log"
)

func main() {

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculator.NewCalculatorServiceClient(cc)
	fmt.Print("Enter the method you want to try: ")
	var method string
	fmt.Scan(&method)

	switch method {
	case "Sum":
		var x, y int
		fmt.Print("Enter the 2 numbers you want to add: ")
		fmt.Scan(&x)
		fmt.Scan(&y)
		Add(c, x, y)
	case "Prime":
		var num int
		fmt.Print("Enter the prime number: ")
		fmt.Scan(&num)
		Prime(c, num)
		//case "Avg":
		//	fmt.Print("Enter the numbers:")
		//	x := input([]int{}, nil)
		//	ComputeAverage(c, x)
	}
}

//func input(x []int, err error) []int {
//	if err != nil {
//		return x
//	}
//	var d int
//	n, err := fmt.Scanf("%d", &d)
//	if n == 1 {
//		x = append(x, d)
//	}
//	return input(x, err)
//}

func Add(c calculator.CalculatorServiceClient, x int, y int) {

	fmt.Println("Making gRPC call for sum")

	req := calculator.SumRequest{
		Num1: float64(x),
		Num2: float64(y),
	}

	resp, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling greet grpc unary call: %v", err)
	}

	fmt.Println("Sum of the numbers is : ", resp.Sum)
}

func Prime(c calculator.CalculatorServiceClient, num int) {

	fmt.Println("Making gRPC call for prime")

	req := calculator.PrimeNumbersRequest{
		Limit: int64(num),
	}

	resp, err := c.PrimeNumbers(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling greet grpc unary call: %v", err)
	}
	log.Println("Response From Server, Prime Number : ")
	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while receving server stream : %v", err)
		}

		fmt.Print(" ", msg.GetPrimeNum())
	}
}
