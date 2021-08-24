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
	"github.com/rcrowley/go-metrics"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)


var (
	addr = flag.String("addr","localhost:10002","service address")
	zkAddr = flag.String("zkAddr","localhost:2181","zookeeper address")
	basePath = flag.String("base","Users/uplook/youpin","prefix path")
)

type Arith struct {

}


// 为客户端提供服务
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
	// 创建一个服务实例。
	s := server.NewServer()

	// 注册中心zookeeper 插件必须在注册服务之前添加到Server中
	addRegistryPlugin(s)
	err := s.RegisterName("Arith",new(Arith), "")


	if err != nil {
		panic(err)
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}

	func (){
		defer s.Close()
	}()

	fmt.Println("service ....")
}

// 添加注册中心插件 ZooKeeper
func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath: *basePath,
		Metrics: metrics.NewRegistry(), // 用来更新服务TPS
		UpdateInterval: time.Minute,  // 更新时间
	}

	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}

	s.Plugins.Add(r)

}