package main

import "fmt"

func main() {

	var ch byte
	var str string

	//字符
	//1.单引号
	//2.字符，往往都是只有一个字符  除了 \n \t 转义字符

	ch = 'a'
	fmt.Println(ch)
	//字符串
	//1.双引号
	//2.字符串可以有一个或多个字符组成
	//3.字符串都是隐藏了一个结束符  '\0'
	str = "a"
	fmt.Println(str)

	str = "hello go"
	fmt.Println(str[0],str[1])
	fmt.Printf("%c,%c\n",str[0],str[1])

	
}
