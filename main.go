package main

import (
	"fmt"
	"math"
	// "errors"
	// "time"
)

func main() {
	var num int
	fmt.Println("input the number, I will caculate the prime numbers for you:")
	fmt.Scanln(&num)
	prime(num)
}

func prime(n int){
	var b = make ([]bool, n+1)
	b[1]=false
	b[2]=true

	for i:=3; i<=n; i+=2{
		b[i]=true
	}
	for i:=4; i<=n; i+=2{
		b[i]=false
	}
	up:=int (math.Sqrt(float64(n)))
	for i:=3;i<=up;i+=2{
		if b[i]{
			for j:=2*i; j<=n; j+=i{
				b[j]=false
			}
		}
	}

	var result []int 
	for i :=1; i<=n; i++{
		if b[i]{
			result=append(result,i)
		}
	}

	fmt.Println("final prime array is ", result, len(result))
}