package main

import "fmt"

func main() {

	s:="anvfd"
	by:=[]byte(s)
	for i:=range by{
		fmt.Printf("%T\n",int(by[i]+1))
	}
	//fmt.Println(by)
	//for i:=range by{
	//	by[i]=by[i]-'a'
	//	fmt.Println(by[i])
	//	//fmt.Println(string(69))
	//}
	//fmt.Println(by)
	//fmt.Println(string(by))

	plain:="hello"
	key:="234137"
	fmt.Println(Jm(plain,key))
}

func Jm(plaintext string, key string)  string{
	byPlain:=[]byte(plaintext)
	bykey:=[]byte(key)
	var encode []int
	encode=make([]int,len(key))


	for i:=0;i<len(key);i++{
		encode[i]= int(bykey[i] - '0')
		//fmt.Printf("%V\n",encode[i])
	}
	//fmt.Printf("%T",encode[])
	for i:=0;i<len(byPlain);i++{
		if i&1==0{
			byPlain[i]= byte('a'+(int(byPlain[i]-'a') + encode[i%len(plaintext)])%26)
		}else {
			byPlain[i]= byte('a'+(int(byPlain[i]-'a') - encode[i%len(plaintext)])%26)
		}
	}
	return string(byPlain)
}