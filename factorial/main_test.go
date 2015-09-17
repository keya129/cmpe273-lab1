package main
import "testing"

type testpair struct {
value int
fibval int
}
var tests = []testpair{
{ 0, 0 },
{ 1, 1 },
{ 2, 1 },
{ 3, 2 },
{ 4, 3 },
{ 5, 5 },
{ 6, 8 },
{ 7, 13 },
{ 8, 21 },
{ 9, 34 },

}
/*Fibonacci function to test recursive function*/

func TestFib(t *testing.T) {
for _, pair := range tests {
v := Fib(pair.value)
if v != pair.fibval {
t.Error(
"For", pair.value,
"expected", pair.fibval,
"got", v,
)
}


}
}
