
/*
  go 语言运算符
  运算符用于在程序运行时执行数学或逻辑运算。
  go 语言内置的运算符有：

  1、算术运算符 （+ - * / % ++ --）
  2、关系运算符  (== != > < >= <=)
  3、逻辑运算符  (&& || !)
  4、位运算符    (& | ^ << >>)
  5、赋值运算符  (= += -= *= /= %= <<= >>= &= ^= |=)
  6、其他运算符  (& *)

*/


package main 
import "fmt"
func main() {

	// fun1()
	// fun2()
	// fun3()
	// fun4()
	// fun5()
	fun6()
	
}

func fun1() {

	var a int  = 21
	var b int = 10
	var c int

	c = a+b;
	fmt.Printf("第1行-c的值为%d\n",c)
	c = a-b;
	fmt.Printf("第2行-c的值为%d\n",c)
	c = a*b;
	fmt.Printf("第3行-c的值为%d\n",c)
	c = a/b;
	fmt.Printf("第4行-c的值为%d\n",c)
	c = a%b;
	fmt.Printf("第5行-c的值为%d\n",c)
	a++
	fmt.Printf("第6行-a的值为%d\n",a)
	b--
	fmt.Printf("第7行-b的值为%d\n",b)
}

func fun2() {

	var a int = 21
	var b int = 10

	if(a == b){
		println("a 等于 b")
	}else{
		println("a 不等于 b")
	}

	if(a>b){
		println("a 大于 b ")
	}else{
		println("a 小于 b")
	}

	a=5
	b=20

	if(a<=b){
		println("a 小于等于 b")
	}

	if(b>=a){
		println("b 大于等于 a")
	}
}

func fun3() {
	var a bool = true
	var b bool = false
	if(a && b){
		fmt.Printf("第1行 条件为 true\n")
	}

	if(a || b){
		fmt.Printf("第2行 条件为 false\n")
	}

	a = false
	b = true

	if(a && b){
		fmt.Printf("第3行 条件为 true\n")
	}else{
		fmt.Printf("第3行 条件为 false\n")
	}
	if(!(a && b)){
		fmt.Printf("第4行 条件为 true\n")
	}
}

func fun4() {
	/*
      A = 0011 1100
      B = 0000 1101

      A & B = 0000 1100
      A | B = 0011 1101
      A ^ B = 0011 0001
	*/

      var a uint = 60 	// 60 = 0011 1100
      var b uint = 13  	// 13 = 0000 1101
      var c uint = 0

      c = a & b 		
      fmt.Printf("第1行c=%d\n",c) //	12 = 0000 1100

      c = a | b 		
      fmt.Printf("第2行c=%d\n",c) //	61 = 0011 1101

      c=a^b
      fmt.Printf("第3行c=%d\n",c) // 49 = 0011 0001

      c=a<<2
      fmt.Printf("第4行c=%d\n",c)  // 240 = 1111 000, 左移就是乘以2 n次方， 60 * 2的2次方 = 60 *4 = 240

      c=a>>2
      fmt.Printf("第5行c=%d\n",c) // 15 = 0000 1111, 右移就是除以2  n次方，  60 / 2的2次方 = 60/ 4 = 15
}

func fun5() {

	var a int = 21
	var c int

	c = a 
	fmt.Printf("第1行 = 运算符 c的值为%d \n",c)
	c += a 
	fmt.Printf("第2行 += 运算符 c的值为%d \n",c)
	c -= a 
	fmt.Printf("第3行 -= 运算符 c的值为%d \n",c)
	c *= a 
	fmt.Printf("第4行 *= 运算符 c的值为%d \n",c)
	c /= a 
	fmt.Printf("第5行 /= 运算符 c的值为%d \n",c)
	c %= a 
	fmt.Printf("第6行 = 运算符 c的值为%d \n",c)

	c = 40
	c <<= 2 
	fmt.Printf("第7行 <<=2 运算符 c的值为%d \n",c)
	c >>= 2
	fmt.Printf("第8行 >>=2 运算符 c的值为%d \n",c)
	c &= 2
	fmt.Printf("第9行 &= 运算符 c的值为%d \n",c)
	c ^= 2
	fmt.Printf("第10行 ^= 运算符 c的值为%d \n",c)
	c |= 2
	fmt.Printf("第11行 |= 运算符 c的值为%d \n",c)

	/*
		第1行 = 运算符 c的值为21
		第2行 += 运算符 c的值为42
		第3行 -= 运算符 c的值为21
		第4行 *= 运算符 c的值为441
		第5行 /= 运算符 c的值为21
		第6行 = 运算符 c的值为0
		第7行 <<=2 运算符 c的值为160
		第8行 >>=2 运算符 c的值为40
		第9行 &= 运算符 c的值为0
		第10行 ^= 运算符 c的值为2
		第11行 |= 运算符 c的值为2
	*/

}

func fun6() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/*运算符实例*/
	fmt.Printf("第1行 a变量类型为 = %T\n",a)
	fmt.Printf("第2行 b变量类型为 = %T\n",b)
	fmt.Printf("第3行 c变量类型为 = %T\n",c)

	/* & 和 *运算符实例*/
	ptr = &a
	fmt.Printf("a的值为 %d\n",a)
	fmt.Printf("prt 为 %d\n",*ptr)
	fmt.Printf("a的地址为 %p\n",&a)
	fmt.Printf("ptr的地址为 %p\n",ptr)
}