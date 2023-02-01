// server.go
package main

import (
	"net"
	"log"
  pb "route/proto" // 透過 `proto` 生成的 Server Stub
	"context" // 新增
  "google.golang.org/protobuf/proto" // 新增
	"google.golang.org/grpc"  // 新增 grpc library
)

// 在這裡，因為在這個 routeGuideServer interface 中有一個 mustEmbedUnimplementedRouteGuideServer()，所以必須要有這個東西，是用來實現向上兼容的。
type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer

	// 接著我們先做 Server 的假 DB 部分。
	// 在這個 routeGuideServer 中，加上這一行 DB 的 type 定義。
	features []*pb.Feature
	// 我們把這個 DB 叫做 features，裡面是有個一些 feature 的 Slice。
}

func dbServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{
				Name: "東京鐵塔",
				Location: &pb.Point {
					Latitude: 353931000,
					Longitude: 139444400,
				},
			},
			{
				Name: "淺草寺",
				Location: &pb.Point {
					Latitude: 357147651,
					Longitude: 139794466,
				},
			},
			{
				Name: "晴空塔",
				Location: &pb.Point {
					Latitude: 357100670,
					Longitude: 139808511,
				},
			},
		},
	}
}

func (s *routeGuideServer) GetFeature(cxt context.Context, point *pb.Point) (*pb.Feature, error){
	for _,feature := range s.features{
		if proto.Equal(feature.Location, point){
			return feature, nil
		}
	}
	return nil, nil
}

func main() {
	// 生成一個listener
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener a the address")
	}

	// server
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, dbServer())
	log.Fatalln(grpcServer.Serve(lis))
}