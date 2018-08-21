package main

import (
	"flag"
	"github.com/micro/go-micro"
	"fmt"
	"microservice_learning/protobuf/dbagent"
	"golang.org/x/net/context"
	"time"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/broker/nsq"
	"microservice_learning/protobuf/logagent"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/broker"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")

)

func main() {
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"http://127.0.0.1:2379"}
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs=[]string{"127.0.0.1:4152"}
	})
	server := micro.NewService(
		micro.Name("howie"),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
	)
	server.Init()
	db := dbagent.NewDbAgentServerService("howie", server.Client())
	sub:= micro.NewPublisher("server.log.data", server.Client())
	// Register handler
	sub.Publish(context.TODO(), &logagent.Log{Time:time.Now().Unix(),Error:"apiapiapi  ",Data:"api_agent启动成功",Filename:"apiapiapi",Line:"35",Method:"apiapiapi"})
	// Call the greeter
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		rsp, err := db.GetOneTestUser(context.TODO(), &dbagent.StringValue{Value:"1"})
		if err != nil {
			fmt.Println(err.Error())
		}else {
			fmt.Println(rsp.String())
		}
	}
}

/*flag.Parse()
fmt.Println("serv", *serv)
r := utils.NewResolver(*serv)
b := grpc.RoundRobin(r)

ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
if err != nil {
	panic(err)
}
fmt.Println("conn...")

ticker := time.NewTicker(1 * time.Second)
for range ticker.C {
	client := protobuf.NewDbAgentServerClient(conn)
	resp, err := client.GetOneTestUser(context.Background(),&protobuf.StringValue{Value:"1"})
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(resp)
	}
}*/