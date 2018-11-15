
package main
import "fmt"
import "unsafe"

const(
  a="abc"
  b=len(a)
  c=unsafe.Sizeof(a)
)

func main() {
	const length int = 10
	const width int = 5
	var area int
	const a1,b1,c1 = 1,false,"str"
	area = length * width;
    fmt.Printf("面积为:%d",area)
    println()
    println(a1,b1,c1)
    println(a,b,c)

}