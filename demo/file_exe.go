package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)



func main() {

	//readFileFromByIOUtil()

	///writeFileFromByIOUtil()

	useBuff()
}

func useBuff()  {

	var s string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("请输入内容")
	s,_ = reader.ReadString('\n')
	fmt.Printf("你输入的内容%s",s)
}


func writeFileFromByIOUtil()  {

	name := "bjtu"
	err := ioutil.WriteFile("./exe.txt",[]byte(name),0666)
	if err != nil{
		fmt.Println("write file failed")
		return
	}
}

func writeFileFromBuffer()  {

	fileObj ,err := os.OpenFile("./exe.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)

	if err != nil{
		fmt.Printf("open file failed,err:%v",err)
		return
	}

	defer fileObj.Close()

	writer := bufio.NewWriter(fileObj)

	writer.WriteString("hello ,北京")//将数据写入缓存
	writer.Flush()//将缓存的内容写入文件
}

func writeFile()  {

	fileObj ,err := os.OpenFile("./exe.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)

	if err != nil{
		fmt.Printf("open file failed,err:%v",err)
		return
	}

	fileObj.WriteString("你好，世界 ，hello world")


}

func readFileFromByIOUtil()  {

	ret,err := ioutil.ReadFile("/Users/bytedance/dev/code/demo_code/demo/Demo.go")
	if err != nil {
		fmt.Println("read file failed")
	}
	if err == io.EOF{
		fmt.Println("end")
	}
	fmt.Println(string(ret))
}

func readFileFromBuffer()  {

	fileObj, err := os.Open("/Users/bytedance/dev/code/demo_code/demo/Demo.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
	}

	//关闭文件
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)

	for  {
		line,err := reader.ReadString('\n')
		if err != nil{
			fmt.Printf("read line failed,err:%v\n", err)
			return
		}
		if err == io.EOF{
			return
		}

		fmt.Println(line)

	}

}


func readFile() {
	fileObj, err := os.Open("./Demo.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
	}

	//关闭文件
	defer fileObj.Close()

	//读文件
	var temp [128]byte
	for {
		n, err := fileObj.Read(temp[:])
		if err == io.EOF {
			//读到文件末尾
			fmt.Println("end")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed ,err:%v\n", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(temp[:n]))
		if n < 0 {
			return
		}

	}
}
