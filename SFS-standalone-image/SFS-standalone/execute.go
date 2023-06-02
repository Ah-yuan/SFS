package main

import(
	"time"
	"fmt"
	"log"
	"sync"
	"strconv"
	"os/exec"
)

type PidI struct{
	Pid	int
	Job	string
	N	int
	Id	int
	St	time.Time
	Credit	int
}

func Send(job Action, pids chan PidI){
	// Send just send request to receiver
	o := time.Now()
	//time.Sleep(time.Duration(job.Start)*time.Millisecond)
	new_pid := PidI{-10, job.JobName, job.Para,job.Id,o, -3}
	pids <- new_pid
}
func Execute(job PidI, p string, pids chan PidI, core string, queue chan PidI){
	// execute request and also update job direction
	var cmd *exec.Cmd
	start_time := job.St
	t1 := time.Now()
	if p == "N"{
		cmd = exec.Command("schedtool","-N","-a",core,"-e","python", "fib.py", strconv.Itoa(job.N),strconv.Itoa(job.Id))
	}else{
		//cmd = exec.Command("schedtool","-N","-a",core,"-e","python","fib.py", strconv.Itoa(job.N))
		cmd = exec.Command("schedtool","-F","-p","20","-a",core,"-e","python","fib.py", strconv.Itoa(job.N),strconv.Itoa(job.Id))
	}
	err := cmd.Start()
	if err != nil{
		log.Fatal("logs exec 1", err)
	}
	tw := time.Now()
	fmt.Println("logs wait time",tw.Sub(t1))
	//actions.m[job.Job] = cmd.Process.Pid
	//new_pid := PidI{0,job.Job,job.N,job.Id}
	pid := cmd.Process.Pid
	var new_pid PidI
	if cmd != nil{
		new_pid = PidI{pid,job.Job,job.N,job.Id,time.Now(), job.Credit}
	}else{
		new_pid = PidI{0,job.Job,job.N,job.Id,time.Now(), job.Credit}
	}
	//actions.Lock()
	//actions.m[job.Job] = new_pid
	//actions.Unlock()
	queue <- new_pid
	err = cmd.Wait()
	if err != nil{
		log.Fatal("exec 2",err)
	}
	t2 := time.Now()
	new_pid.Credit = -2
	pids <- new_pid
	fmt.Println(job.Job,t2.Sub(t1).Milliseconds())
	//cmd = exec.Command("kill", "-9", strconv.Itoa(pid))
	//err = cmd.Start()
	//if err != nil{
	//	log.Fatal("logs exec 3", err)
	//}
	//debug.FreeOSMemory()
        //log4sys.Warn("NumGoroutine:",runtime.NumGoroutine())
	fmt.Println("logs TIME: ",job.Job, t1.Sub(start_time), t2.Sub(start_time))
}

```
这段代码定义了一个名为Execute的函数，它有五个参数：job、p、pids、core和queue。该函数用于执行命令并更新作业方向。
在函数中，首先根据传入的参数p来决定执行哪个命令。当p等于"N"时，执行的是以下命令：
schedtool -N -a core -e python fib.py job.N job.Id
其中，job.N和job.Id是从任务job中获取的参数，用于作为fib.py脚本的输入参数。
当p不等于"N"时，执行的是以下命令：
schedtool -F -p 20 -a core -e python fib.py job.N job.Id
同样，job.N和job.Id也是从任务job中获取的参数。
执行命令后，会将返回的进程ID存储到pid变量中，并根据它创建一个新的PidI类型的结构体new_pid，并将其加入到队列queue中。接着，等待命令执行完成，并将new_pid发送到pids通道中。
在控制台输出一个字符串"logs TIME: "，后面跟着任务的ID（job.Job）、任务开始执行的时间（t1.Sub(start_time)）和任务结束执行的时间（t2.Sub(start_time)）。
其中，t1.Sub(start_time)表示从函数开始执行到任务开始执行的时间间隔，t2.Sub(start_time)表示从函数开始执行到任务结束执行的时间间隔。
```
func ExecuteNoChannel(wg *sync.WaitGroup, job Action, p string, pids chan PidI, start_time time.Time, cpuC string){
        defer wg.Done()
        //time.Sleep(time.Duration(job.Start) * time.Millisecond)
        t1 := time.Now()
        var cmd *exec.Cmd
        if p == "N"{
		cmd = exec.Command("schedtool","-N","-a",cpuC,"-e","python", job.Exec, strconv.Itoa(job.Para),strconv.Itoa(job.Id))
        }else{
                cmd = exec.Command("schedtool","-R","-p","20","-a","0x1","-e","python", job.Exec, strconv.Itoa(job.Para),strconv.Itoa(job.Id))
        }
	err := cmd.Start()
	if err != nil{
		log.Fatal("exec 1", err)
	}
	tw := time.Now()
        fmt.Println("logs wait time",tw.Sub(t1))
	//pid := cmd.Process.Pid
        //new_pid := PidI{cmd.Process.Pid,job.JobName}
        //pids <- new_pid
        //err := cmd.Wait()
        //if err != nil{
        //        log.Fatal(err)
        //}
	err = cmd.Wait()
        if err != nil{
                log.Fatal("exec 2",err)
        }

        t2 := time.Now()
        //pids <- new_pid
        fmt.Println(job.JobName,t2.Sub(t1).Milliseconds())
	fmt.Println("logs TIME: ",job.JobName, t1.Sub(start_time), t2.Sub(start_time))
}

