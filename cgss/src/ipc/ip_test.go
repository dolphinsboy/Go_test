package ipc

import (
	"testing"
	"fmt"
)

type EchoServer struct {

}

//实现接口的方法
//http://www.cnblogs.com/simplelovecs/p/5359520.html
func (server *EchoServer) Handle(request, params string) *Response {

	return &Response{"OK", "ECHO:"+request+" ~ "+params}
}

func (server *EchoServer) Name() string{
	return "EchoServer"
}

func TestIpc(t *testing.T){
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	fmt.Println(client2)
	fmt.Println(client1)

	resq1,_ := client1.Call("From client1", "Test")
	resq2,_ := client2.Call("From client2", "Test")

	if resq1.Body != "ECHO:From client1 ~ Test" || resq2.Body != "ECHO:From client2 ~ Test" {
		t.Error("IpcClient.Call failed.resq1:",resq1.Body, "resq2:", resq2.Body)
	}

	client1.Close()
	client2.Close()

}

