package main

import "fmt"
func main(){
for i:=0;i<10;i++ {
fmt.Println(Fib(i))
}
}
/*Fibonacci Recursive function*/
func Fib(x int) int{
  if x==0{
    return 0
  } else if x==1{
    return 1
  }
return Fib(x-1)+Fib(x-2)
}
