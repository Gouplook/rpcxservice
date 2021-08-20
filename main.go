/**
 * @Author: yinjinlin
 * @File:  main
 * @Description:
 * @Date: 2021/8/18 下午4:27
 */

package  main

import (
	"context"
	"flag"
	"fmt"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)


var (
	addr = flag.String("addr","localhost:8972","servic address")
)

type Arith struct {

}

func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	fmt.Println("Mul C=", reply.C)
	return nil
}


func (a *Arith) Add (ctx context.Context,args example.Args, reply *example.Reply) error {
	reply.C = args.A + args.B
	fmt.Println("Add C=", reply.C)
	return nil
}

// server
func main(){
	flag.Parse()
	s := server.NewServer()

	s.RegisterName("Arith", new(Arith),"")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}


	fmt.Println("service ....")
}