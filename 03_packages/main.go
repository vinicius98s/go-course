package main

import (
	"fmt"
	"math"

	"github.com/vinicius98s/go-course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.3))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("suiciniv"))
}
