package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"learn-proto/protobuff"
)

func main() {

	product := &protobuff.Products{
		Data: []*protobuff.Product{
			{
				Id:    2,
				Name:  "wowww sekalii",
				Price: 12313,
				Stock: 123,
				Category: &protobuff.Category{
					Id:   2,
					Name: "kontols",
				},
			},
		},
	}

	fmt.Println("products --> ", product)

	data, err := proto.Marshal(product)
	if err != nil {
		fmt.Println("error kntl --> ", err)
	}

	//compact binary wire format
	fmt.Println("proto marshall --> ", data)

	newProducts := &protobuff.Products{}
	err = proto.Unmarshal(data, newProducts)
	if err != nil {
		fmt.Println("err --> ", err)
	}

	for _, p := range newProducts.Data {
		fmt.Println("for loop --> ", p.Name)
	}

	for _, bjir := range newProducts.GetData() {
		fmt.Println("product.getName()", bjir.GetName())
	}
}
