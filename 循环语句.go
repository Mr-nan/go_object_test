package main 
import "fmt"
func main() {
	
	func1()
	func2()
	func3()

}

func func1() {
	var b int = 15
	var a int
	numbers :=[6] int{1,2,3,5}

	/*for 循环*/
	for i := 0; i < 10; i++ {
		fmt.Printf("i 的值为：%d\n",i)
	}
    fmt.Printf("\n");
	for a<b {
		a++
		fmt.Printf("a 的值为：%d\n",a)
	}
    fmt.Printf("\n");
	for index,value:=range numbers{
		fmt.Printf("第%d位的值为:%d\n",index,value)
	}
	fmt.Printf("\n");
}

func func2() {
	
    for i := 0; i < 5; i++ {
    	if(i==2){
    		fmt.Printf("%d 等于 2，我要退出了\n",i)
    		break
    	}
    }
    fmt.Printf("\n");
    for i := 0; i < 10; i++ {
    	
    	if(i % 2 ==0){
    		fmt.Printf("%d 是偶数继续跑\n",i)
    		continue
    	}

    	fmt.Printf("看看我是谁%d\n",i)

    }
    fmt.Printf("\n");


}

func  func3() {
	/*
      goto 语句可以无条件的转移到过程中指定的行，
      goto语句通常与条件语句配合使用，可用来实现条件转移，构成循环，跳出循环具体功能
      但是，在结构化程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使得理解和调试都产生困难
	*/

      var a int = 10

      LOOP : for a<20 {
      	if( a == 15){
      		a = a+1
      		goto LOOP
      	}

      	fmt.Printf("a 的值为：%d \n",a)
      	a++
      }
}