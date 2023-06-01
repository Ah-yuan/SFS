package main

import(
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

type Action struct{
	JobName	string
	Exec	string
	Para	int
	Start	int
	Id	int
}


func GetTrace(path string)([]Action, int){
	file, err := os.Open(path)
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan(){
		txtlines = append(txtlines, scanner.Text())
	}
	trace := []Action{}
	var s []string
	var i int
	var f int
	var id int
	var newAction Action
	var num int = 0
	for _,eachline := range txtlines{
		s = strings.Split(eachline," ")
		i, _ = strconv.Atoi(s[2])
		f, _ = strconv.Atoi(s[3])
		id, _ = strconv.Atoi(s[4])
		newAction = Action{s[0],s[1],i,f*9,id}
		trace = append(trace, newAction)
		num += 1
	}
	return trace, num
}

#这段代码定义了一个结构体 Action，它有5个属性：JobName，Exec，Para，Start和Id。然后定义了一个函数 GetTrace，该函数接受一个文件路径字符串，并返回两个值：一个类型为 []Action 的切片以及一个 int 类型的计数器。
#在函数里面，代码打开了指定路径的文件并读取其中的内容，将其按行分隔成字符串，并将每一行作为一个元素添加到名为 txtlines 切片中。接下来，定义了一些变量用于处理每一行内容，包括 s、i、f 和 id。然后使用一个循环遍历
#txtlines 中的每一行，对每一行进行分割并将其转换为适当的数据类型，最终将 Action 结构体的实例添加到 trace 切片中。最后返回 trace 和 num 两个值。
#这段代码的功能是从一个文件中读取文本，将其解析为 Action 结构体实例组成的切片，并返回该切片以及切片中 Action 实例的数量。
