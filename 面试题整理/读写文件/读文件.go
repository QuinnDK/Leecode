package 读写文件

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
//利用ioutil.ReadFile直接从文件读取到[]byte中
func Read0() string {
	f,err:=ioutil.ReadFile("test" )
	if err!=nil{
		fmt.Println("read fail",err)
	}
	return string(f)
}
//先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
func Read1() string {
	f,err:=os.Open("test")
	if err!=nil{
		fmt.Println("read fail",err)
		return ""
	}
	//把file 读取到缓冲区
	defer f.Close()
	var chunk []byte
	buf:=make([]byte,1024)

	for{
		//从file读取到buf中
		n,err:=f.Read(buf)
		if err!=nil&&err!=io.EOF{
			fmt.Println("read buf fail",err)
			return ""
		}
		//说明读取结束
		if n==0{
			break
		}
		chunk=append(chunk,buf[:n]...)
	}
	return string(chunk)
}

//先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte
func Read2() (string) {
	fi, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	var chunks []byte

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		//fmt.Println(string(buf))
		chunks = append(chunks, buf...)
	}
	return string(chunks)
	//fmt.Println(string(chunks))
}
//读取到file中，再利用ioutil将file直接读取到[]byte中
func Read3() string {
	f,err:=os.Open("test")
	if err!=nil{
		fmt.Println("read file fail",err)
		return ""
	}
	defer f.Close()
	fd,err:=ioutil.ReadAll(f)
	if err!=nil{
		fmt.Println("read to fd fail",err)
		return ""
	}
	return string(fd)
}