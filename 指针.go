package main 
import "fmt"
func main() {


	/* 
		一个指针变量指向了一个值的内存地址
		指针声明的格式： var var_name *var_type
	*/

		var ip *int 			/* 指向整型 */
		var fp *float32 		/* 指向浮点型 */

		var a int = 10
		var b float32 = 10.11

		ip = &a  /* 指针变量的存储地址*/
		fp = &b

		fmt.Printf("a变量的地址是：%x\nb变量的地址是：%x\n",&a,&b)

		// 指针变量的存储地址
		fmt.Printf("ip变量存储的指针地址%x\nfp变量存储的指针地址:%x\n",ip,fp)

		// 指针访问值
		fmt.Printf("*ip变量的值:%d\n*fp变量的值:%f\n",*ip,*fp)

		/*
		 空指针，当一个指针被定义后没有分配到任何变量时，它的值为nil
		 nil指针也称为空指针
         nil在概念上和其他语言的null、None、nil、NULL 一样，都是代零值的空值

		空指针判断

		if(ptr !=nil) if(ptr == nil)

		*/

		func1()
		func2()

}

func func1() {
	/*
		指针数组
	*/

	a :=[]int {10,11,12}
	var ptr [3]*int

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < len(ptr); i++ {
		fmt.Printf("ptr[%d] = %d\n",i,*ptr[i])
	}
}

func func2() {
	/*
      指向指针的指针
      **ptr
	*/

      var a int= 10
      var ip *int = &a
      var ptr **int = &ip
      
      fmt.Printf("变量a的值%d\n",a)
      fmt.Printf("指针变量 *ip的值%d\n",*ip)
      fmt.Printf("指向指针的指针变量 **ptr的值%d\n",**ptr)
     
}