package main
import ("testing";"time")

type testpair struct {
input int
output int
}
var tests = []testpair{
{ 1, 1 },
{ 2, 2},
{10,10},

}
/*Fibonacci function to test sleep function*/
func TestSleep(t *testing.T) {
for _, pair := range tests {
t1 := time.Now()

Sleep(pair.input)
t2 := time.Now()
v:=t2.Sub(t1);
k:=int(v.Nanoseconds())
k=k/1000000000
if k != pair.output  {
t.Error(
"For", pair.input,
"expected", pair.output,
"got", k,
)
}


}
}
