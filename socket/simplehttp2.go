package main
import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	//用于解析地址和端口号
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	//用于创建连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil{
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}