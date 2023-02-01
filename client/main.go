package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "route/proto"
)

// 因為之後需要再開其他 api ，所以這裡先把 Unary 的相關處理放到函式 getFeat 裡面。
func getFeat (client pb.RouteGuideClient){
	feature,err := client.GetFeature(context.Background(),&pb.Point{
		Latitude: 353931000,
		Longitude: 139444400,
	})
	if err != nil{
		log.Fatalln()
	}
	fmt.Println(feature)
}


func main() {
	// Dial 是 grpc 提供的一個方法，會是一個dail 請求，第一個參數會說他要播向哪裡，之後是一些 option。
	conn,err := grpc.Dial("localhost:5000",grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// NewRouteGuideClient 是 grpc 自動生成在 Client stub 產生的一個方法，可以在 route_grpc.go 找到它的定義，不過總之他會接收一個連接的 Interface，並且回傳一個 RouteGuideClient 類型的東西。
	client := pb.NewRouteGuideClient(conn)
	getFeat(client)
	
	// feature,err := client.GetFeature(
	// 	context.Background(),
	// 	&pb.Point{
	// 		Latitude: 353931000,
	// 		Longitude: 139444400,
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(feature)
}