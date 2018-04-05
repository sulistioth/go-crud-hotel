package endpoint

import (
	"context"

	svc "projectHotel/hotel/departemen/server"

	kit "github.com/go-kit/kit/endpoint"
)

type DepartemenEndpoint struct {
	AddDepartemenEndpoint              kit.Endpoint
	ReadDepartemenEndpoint             kit.Endpoint
	UpdateDepartemenEndpoint           kit.Endpoint
	ReadDepartemenByNamaEndpoint       kit.Endpoint
	ReadDepartemenByKeteranganEndpoint kit.Endpoint
}

func NewDepartemenEndpoint(service svc.DepartemenService) DepartemenEndpoint {
	addDepartemenEp := makeAddDepartemenEndpoint(service)
	readDepartemenEp := makeReadDepartemenEndpoint(service)
	updateDepartemenEp := makeUpdateDepartemenEndpoint(service)
	readDepartemenByNamaEp := makeReadDepartemenByNamaEndpoint(service)
	readDepartemenByKeteranganEp := makeReadDepartemenByKeteranganEndpoint(service)
	return DepartemenEndpoint{AddDepartemenEndpoint: addDepartemenEp,
		ReadDepartemenEndpoint:             readDepartemenEp,
		UpdateDepartemenEndpoint:           updateDepartemenEp,
		ReadDepartemenByNamaEndpoint:       readDepartemenByNamaEp,
		ReadDepartemenByKeteranganEndpoint: readDepartemenByKeteranganEp,
	}
}

func makeAddDepartemenEndpoint(service svc.DepartemenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Departemen)
		err := service.AddDepartemenService(ctx, req)
		return nil, err
	}
}

func makeReadDepartemenEndpoint(service svc.DepartemenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadDepartemenService(ctx)
		return result, err
	}
}

func makeUpdateDepartemenEndpoint(service svc.DepartemenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Departemen)
		err := service.UpdateDepartemenService(ctx, req)
		return nil, err
	}
}

func makeReadDepartemenByNamaEndpoint(service svc.DepartemenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Departemen)
		result, err := service.ReadDepartemenByNamaService(ctx, req.NamaDepartemen)
		return result, err
	}
}

func makeReadDepartemenByKeteranganEndpoint(service svc.DepartemenService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Departemen)
		result, err := service.ReadDepartemenByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
