package main

import (
	"context"
	"fmt"
	"time"

	cli "projectHotel/hotel/departemen/endpoint"
	opt "projectHotel/hotel/util/grpc"
	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCDepartemenClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Departemen
	//client.AddDepartemenService(context.Background(), svc.Departemen{IdDepartemen: "DP007", NamaDepartemen: "Kebersihan", Keterangan: "Jahat"})
	//fmt.Println("data berhasil ditambahkan")

	//List Departemen
	//deps, _ := client.ReadDepartemenService(context.Background())
	//fmt.Println("semua departemen:", deps)

	//Update Departemen
	//client.UpdateDepartemenService(context.Background(), svc.Departemen{NamaDepartemen: "Room Room", Status: "0", IdDepartemen: "DP002"})
	//fmt.Println("data berhasil diupdate")

	//Get Departemen By Nama
	//parameter := "Ja%"
	//depNama, _ := client.ReadDepartemenByNamaService(context.Background(), "Ja%")
	//fmt.Println("daftar departemen berdasarkan keterangan:", depNama)

	//GetByKet
	depKet, _ := client.ReadDepartemenByKeteranganService(context.Background(), "%ja%")
	fmt.Println("all departemen", depKet)

}
