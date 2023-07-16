package main

import (
	"fmt"
	"os"

	pb "github.com/mhdianrush/go-buffer-protocol-products/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var logger = logrus.New()

func main() {
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Println(err.Error())
	}
	logger.SetOutput(file)

	// encode
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Iphone 13 Pro Max",
				Price: 15000000,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Mobile Phone",
				},
			},
			{
				Id:    2,
				Name:  "Macbook Pro M1 2020",
				Price: 16000000,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Laptop",
				},
			},
		},
	}
	data, err := proto.Marshal(products)
	if err != nil {
		logger.Println(err)
	}
	// compact binary wire format
	fmt.Println(data)

	// decode
	decodeProducts := &pb.Products{}
	if err = proto.Unmarshal(data, decodeProducts); err != nil {
		logger.Println(err)
	}
	fmt.Println(decodeProducts)
}
