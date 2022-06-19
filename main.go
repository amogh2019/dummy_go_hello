package main

import (
	"fmt"

	"github.com/amogh2019/dummy_go_math/calculationutil"
	"github.com/amogh2019/dummy_go_math/gcdutil"
	"github.com/amogh2019/dummy_go_math/geometryutil"
	v2Cal "github.com/amogh2019/dummy_go_math/v2/calculationutil"
	"github.com/ddadumitrescu/hellomod"
)

func main() {

	fmt.Println("Hello in Angrezi")

	hellomod.Salut()
	hellomod.SayHello()

	fmt.Println(geometryutil.RectangleArea(2, 3))
	fmt.Println(calculationutil.SubBFromA(10, 3))
	fmt.Println(v2Cal.AddNums(10, 10, 10, 10, 10))
	fmt.Printf("gcd of 45 and 27  is %d\n", gcdutil.FindGcd(45, 27))
}
