


/*
    函数定义格式
    func function_name([parameter list])[return_types]{
			函数体
    }

    func: 函数由func开始声明

    function_name:函数名称，函数名和参数列表一起构成了函数签名

    parameter list: 参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。
    参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。

    return_type:返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下return_types不是必须的

    函数体 :函数定义的代码集合

*/

package main 

import "fmt"
import "math"
   
func main() {
	
    var a int = 10
    var b int = 20
    var ret int

    ret = max(a,b)
    fmt.Printf("%d和%d的最大值是:%d\n",a,b,ret)

    c,d :=swap("哈哈","呵呵")
    fmt.Println(c,d)

    fmt.Printf("交换前 a=%d,b=%d\n",a,b)
    swap2(a,b)
    fmt.Printf("交换后 a=%d,b=%d\n",a,b)
   
    fmt.Printf("======引用传递======\n")
    fmt.Printf("交换前 a=%d,b=%d\n",a,b)
    swap3(&a,&b)
    fmt.Printf("交换后 a=%d,b=%d\n",a,b)

    /*
      go 语言可以很灵活的创建函数，并且作为值使用
      以下实例中我们在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数math.sqrt()
    */

     // 声明函数变量
      getSquareRoot :=func (x float64)float64 {
      	return math.Sqrt(x)
      }

      // 使用函数
      fmt.Println(getSquareRoot(9))

      /*
         函数闭包
         go 语言支持匿名函数，可作为闭包。
         匿名函数是一个“内联”语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

      */

       // nextNumber 为一个函数，函数i 为 0
         nextNumber :=getSequence()

         // 调用 nextNumber 函数， i变量自增 1并且返回

         fmt.Println(nextNumber())
         fmt.Println(nextNumber())
         fmt.Println(nextNumber())

         // 创建新的函数 nextNumber1,并查看结果
         nextNumber1 :=getSequence()
         fmt.Println(nextNumber1())
         fmt.Println(nextNumber1())
       
       /*
			函数方法
			go语言中同时有函数和方法，一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类似的一个值或者一个指针。
			所有给定类型的方法属于该类型的方法集。

	        func (variable_name variable_date_type) function_name() [return_type]{
					函数体
	        }
       */

	    var c1 Circle
	    c1.radius = 10.00
	    fmt.Println("圆的面积为:",c1.getArea())

}

/*定义结构体*/
type Circle struct{
	radius float64
}

/*该method 属于Circle类型对象的方法*/
func (c Circle) getArea() float64 {
	// c.radius 即为Circle类型对象中的属性
	return 3.14 * c.radius *c.radius
}

/*匿名函数*/
func getSequence() func() int {
	 i:=0
	 return func() int {
	 	i+=1
	 	return i
	 }
	
}

/*
  返回两个数的最大值
*/
func max(num1,num2 int) int{
	
	var result int

	if(num1 > num2){
		result = num1
	}else{
		result = num2
	}
	return result

}

/*
   返回多个值
*/
func swap(x,y string)(string,string) {
	
	return y,x
}


/*
  值传递-在调用函数时将实际参数复制一份传递到函数中，这样在函数中对参数进行修改，将不影响到实际参数
*/

 func swap2(a,b int){
 	 var temp int
 	 temp = a
 	 a = b
 	 b = temp
 	 fmt.Printf("交换了a=%d:b=%d\n",a,b)
 }

 /*
   引用传递-在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
 */
func swap3(a *int , b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Printf("引用传递交换了a=%d:b=%d\n",*a,*b)
}












