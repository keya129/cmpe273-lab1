package main

import "fmt"
/*This defines interface*/
type Shape interface {
Peri() float64
Area() float64
}
type Rectangle struct {
  l,b float64
}
type Circle struct{
  r float64
}
func(r Rectangle) Area() float64{
  return r.l*r.b;
}
func(c Circle) Area() float64{
  return 3.14*c.r*c.r
}
func Printarea(s Shape){
fmt.Println("Area :")
fmt.Println(s.Area())
}

func(r Rectangle) Peri() float64{
  return 2*(r.l+r.b)
}
func(c Circle) Peri() float64{
  return 3.14*2*c.r
}
func Printperi(s Shape){
fmt.Println("Perimeter :")
fmt.Println(s.Peri())
}

func main() {
c:=Circle{r:5}
r:=Rectangle{l:3,b:4}
Printperi(c)
Printperi(r)
Printarea(c)
Printarea(r)

}
