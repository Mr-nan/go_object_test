package main
import "fmt"
func main() {

	/*	

	iota, 特殊常量，可以认为是一个可以被编译器修改的常量。
	iota, 在const关键字出现时被重置为 0 （const 内部的第一行之前），const中每新增一行常量声明将使得iota计数一次，
	iota可理解为const语句块中的行索引。

	*/
	const (
			a = iota 	//0
			b			//1
			c 			//2
			d =	"ha"	//独立值，iota +=1
			e			//"ha" iota+=1
			f = 100  	// iota +=1
			g 			// 100 iota+=1
			h = iota	// 7,恢复计数
			i 			// 8
		)
		fmt.Println(a,b,c,d,e,f,g,h,i)  // 0 1 2 ha ha 100 100 7 8

		const(
				j=1<<iota    	// 	1<<0  1左移0位   1
				k=3<<iota		//  3<<1  3左移1位	6
				l				//  3<<2  3左移2位   12
				n				//	3<<3  3左移3位	24
			)

		fmt.Println(j,k,l,n)	// 1 6 12 24
}