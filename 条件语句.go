package main 
import "fmt"
func main() {
	
	// func1()
	// func2()
	// func3()
	func4()
}


func func1() {
	 
	 var a int = 100
	 var b int = 200

	 if a == 100{
	 	if b== 200{
	 		fmt.Printf("a 的值为 100  b 的值为200 \n")
	 	}
	 }

}

func func2() {
	
	 var a int 
	 var b int
	 fmt.Printf("请输入密码:\n")
	 fmt.Scan(&a)
	 if a == 408972897{
	 	fmt.Printf("请再次输入密码\n")
	 	fmt.Scan(&b)
	 	if b == a{
	 		fmt.Printf("密码正确\n");
	 	}else{
	 		fmt.Printf("二次输入的密码不一致\n")
	 	}

	 }else{
	 	fmt.Printf("你输入的密码不正确\n")
	 }

}

func func3() {
	var grabe string = "B"
	var marks int = 90

	switch marks{
		case 90: grabe = "A"
		case 80: grabe = "B"
		case 50,60,70: grabe = "C"
		default : grabe = "D"
	}

	switch{
		case grabe == "A" :
		 fmt.Printf("优秀！\n")
		 case grabe == "B" :
		 fmt.Printf("良好！\n")
		 case grabe == "C" :
		 fmt.Printf("及格！\n")
		 case grabe == "D":
		 fmt.Printf("不及格！\n")
		default :
			fmt.Printf("差！\n")
	}

	fmt.Printf("你的成绩是:%s\n",grabe)
}

func func4() {

	/*
       type switch 来判读某个 interface 变量中实际存储的变量类型

       switch x.(type){
			case type:
				statement(s)
       }

	*/


       var a = 123

       switch i :=a.(type){
             case nil:
             	fmt.Printf("a 的类型: %T\n",i)
             case int:
             	fmt.Printf("a 的类型: int\n")
             case float64:
             	fmt.Printf("a 的类型: float64\n")
             case func(int) float64:
             	fmt.Printf("a 的类型: func(int) float64\n")
             case bool, string:
             	fmt.Printf("a 的类型: bool, string\n")
             default:
             	fmt.Printf("未知类型")

       }
}