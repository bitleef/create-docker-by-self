package main

import (
	"fmt"
	"os"
	"strings"
)

/*
* 采用带参数执行对应的函数，每次输入的参数有且只能有一个，相应的参数如下：
* 1. UTS_Namespace 隔离nodename和domainname两个系统标识
* 2. IPC_Namespace 隔离System V IPC和POSIX message queues
* 3. PID_Namespace 隔离进程ID
* 4. Mount_Namespace 隔离各个进程看到的挂载点视图
* 5.
 */
var CallFunctionName string

const (
	UTS_NAMESPACE   = "UTS_Namespace"
	IPC_NAMESPACE   = "IPC_Namespace"
	PID_NAMESPACE   = "PID_Namespace"
	MOUNT_NAMESPACE = "Mount_Namespace"
)

func main() {
	var err error
	for idx, arg := range os.Args {
		if idx == 2 {
			fmt.Println("too many parameters:", os.Args)
			return
		} else if idx == 1 {
			CallFunctionName = strings.Replace(arg, "-", "_", -1)
			fmt.Println("call function is:", CallFunctionName)
		}
	}

	switch CallFunctionName {
	case UTS_NAMESPACE:
		err = UTS_Namespace()
		if err != nil {
			fmt.Printf("%s failed to run, err is %v", CallFunctionName, err)
		}
	case IPC_NAMESPACE:
		err = IPC_Namespace()
		if err != nil {
			fmt.Printf("%s failed to run, err is %v", CallFunctionName, err)
		}
	case PID_NAMESPACE:
		err = PID_Namespace()
		if err != nil {
			fmt.Printf("%s failed to run, err is %v", CallFunctionName, err)
		}
	case MOUNT_NAMESPACE:
		err = Mount_Namespace()
		if err != nil {
			fmt.Printf("%s failed to run, err is %v", CallFunctionName, err)
		}
	default:
		fmt.Println("call function is NULL, please input parameter")
	}
	if err != nil {
		fmt.Println(CallFunctionName, "failed to run")
	}
}
