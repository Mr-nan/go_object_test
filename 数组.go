package main 
import "fmt"
func main() {
	
	/*
   		声明数组
   		var variable_name [size] variable_type
	*/

   	/*数组balance 长度为10 类型为floa32*/
   	var balance [10] float32

   	/*数组初始化*/
   	var balance1 = [5] float32 {1000.0,2.0,3.0,4.0,50.0}
   	 fmt.Printf("balance1=%f\n",balance1[3])

   	/*
   	 初始化数组中{}中的元素个数不能大于【】中的数字
   	 如果忽略【】中的数字不设置大小，go 语言会根据元素的个数来设置数组的大小
   	*/

   	 var balance2 = [...]float32{10.0,11,3.0}

   	 /*数组赋值*/
   	 balance2[2] = 50.0
   	 fmt.Printf("balance2[2]=%f\n",balance[2])

   	 /*
   	   多维数组
   	   var variable_name [size][size]...[size] variable_type
   	   var threedim [5][10][4] int

   	 */



   	   /*
   	    二维数组
   	    var variable_name [x][y] variable_type
   	   */

   	    var a = [3][4]int{
   	    	{0,1,2,3},
   	    	{4,5,6,7},
   	    	{8,9,10,11}}

   	    for i := 0; i < len(a); i++ {
   	    	for j:=0;j<len(a[i]);j++{
   	    		fmt.Printf("a[%d][%d] = %d\n",i,j,a[i][j])
   	    	}
   	    }


   	    /*
   	      向函数传递数组

		// 形参设定数组大小
		func funcName(param [10]int){
		
		}

		// 形参未设定数组大小
		func funcName(param []int){
	
		}

   	    */

		var arr = []int{1,2,3,10,5}
		var avg float32
		avg = getAverage(arr,5)
		fmt.Printf("平均值:%f\n",avg)

		/*
		  浮点精度计算
		*/

		 a1:=1.69
		 b1:=1.7
		 c1:= a1*b1 //结果应该是2.873
		 fmt.Println(c1) // 输出的是2.8729999999999998

		 a2:=1690
		 b2:=1700
		 c2:=a2 *b2
		 fmt.Println(c2)
		 fmt.Println(float64(c2)/1000000)


}


func getAverage(arr []int,size int) float32 {
	var i , sum int
	var avg float32

	for i= 0; i < size; i++ {
		sum+=arr[i]
	}
	avg = float32(sum) / float32(size)
    return avg
}











