package main
import "testing"

type testpair2 struct {
c  Shape
perival float64
areaval float64
}
var tests2 = []testpair2{
{Circle{r:5},31.400000000000001,78.5},
{Rectangle{l:3,b:4},14,12},
}
func TestPeri(t *testing.T) {
for _, pair := range tests2 {
v:=pair.c.Peri();
w:=pair.c.Area();

/*Below test for testing perimeter function*/
if v != pair.perival {
t.Error(
"For", pair.perival,
"expected", pair.perival,
"got", v,
)
}
/*Below test for testing area function*/
if w != pair.areaval {
t.Error(
"For", pair.areaval,
"expected", pair.areaval,
"got", w,
)
}


}
}
