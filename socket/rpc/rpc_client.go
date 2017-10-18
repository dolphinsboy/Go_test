package main
import (
	"net/rpc"
	"fmt"
	"os"
)

type Args struct {
	A,B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.0:1234")
	fmt.Println("Start connect")
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect error %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected")

	args := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil{
		fmt.Fprintf(os.Stderr, "call error %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}