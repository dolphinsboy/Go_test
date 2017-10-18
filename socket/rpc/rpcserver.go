package main

import (
	"errors"
	"net/rpc"
	"net"
	"fmt"
	"os"
	"net/http"
)

type Args struct {
	A,B int
}

type Qutoient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith)Multiply(args *Args, reply *int)error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith)Divide(args *Args, quo *Qutoient)error {
	if args.B == 0{
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil{
		fmt.Fprintf(os.Stderr, "listen error %s", err.Error())
		os.Exit(1)
	}
	fmt.Println("Start serve")
	http.Serve(l, nil)
}