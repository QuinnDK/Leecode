package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
	go语言自带的输入格式无法实现空格输入
	*/
	//var name string
	//
	//
	//fmt.Scanf("%s",&name)
	//fmt.Print(name)

	/*
	实现空格输入
	*/
	//inputReader := bufio.NewReader(os.Stdin)	// 使用了自动类型推导，不需要var关键字声明
		inputReader := bufio.NewReader(os.Stdin)

		fmt.Println("Please input your name: ")

		//input, err = inputReader.ReadString('\n')
		input, err := inputReader.ReadString('\n')

		if err == nil {
			fmt.Printf("Your name is: %s\n", input)
		}

	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))


		s := "yes字符串"
		for _,b := range []byte(s){
			fmt.Printf("%X ",b)
		}

}
