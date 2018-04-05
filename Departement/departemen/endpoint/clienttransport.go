package endpoint

import (
	"context"
	"time"

	svc "projectHotel/hotel/departemen/server"

	pb "projectHotel/hotel/departemen/grpc"

	util "projectHotel/hotel/util/grpc"
	disc "projectHotel/hotel/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.DepartemenService"
)

func NewGRPCDepartemenClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.DepartemenService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addDepartemenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddDepartemenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addDepartemenEp = retry
	}

	var readDepartemenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDepartemenEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDepartemenEp = retry
	}

	var updateDepartemenEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateDepartemen, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateDepartemenEp = retry
	}

	var readDepartemenByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDepartemenByNama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDepartemenByNamaEp = retry
	}

	var readDepartemenByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDepartemenByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDepartemenByKeteranganEp = retry
	}
	return DepartemenEndpoint{AddDepartemenEndpoint: addDepartemenEp,
		ReadDepartemenEndpoint: readDepartemenEp, UpdateDepartemenEndpoint: updateDepartemenEp,
		ReadDepartemenByNamaEndpoint:       readDepartemenByNamaEp,
		ReadDepartemenByKeteranganEndpoint: readDepartemenByKeteranganEp}, nil
}

func encodeAddDepartemenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Departemen)
	return &pb.AddDepartemenReq{
		IdDepartemen:   req.IdDepartemen,
		NamaDepartemen: req.NamaDepartemen,
		Status:         req.Status,
		Keterangan:     req.Keterangan,
	}, nil
}

func encodeReadDepartemenRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateDepartemenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Departemen)
	return &pb.UpdateDepartemenReq{
		IdDepartemen:   req.IdDepartemen,
		NamaDepartemen: req.NamaDepartemen,
		Status:         req.Status,
		Keterangan:     req.Keterangan,
	}, nil
}

func encodeReadDepartemenByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Departemen)
	return &pb.ReadDepartemenByNamaReq{NamaDepartemen: req.NamaDepartemen}, nil
}

func encodeReadDepartemenByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Departemen)
	return &pb.ReadDepartemenByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func decodeDepartemenResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadDepartemenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDepartemenResp)
	var rsp svc.Departemens

	for _, v := range resp.AllDepartemen {
		itm := svc.Departemen{
			IdDepartemen:   v.IdDepartemen,
			NamaDepartemen: v.NamaDepartemen,
			Status:         v.Status,
			Keterangan:     v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadDepartemenbyNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDepartemenByNamaResp)
	return svc.Departemen{
		IdDepartemen:   resp.IdDepartemen,
		NamaDepartemen: resp.NamaDepartemen,
		Status:         resp.Status,
		Keterangan:     resp.Keterangan,
	}, nil
}

func decodeReadDepartemenByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDepartemenByKeteranganResp)
	var rsp svc.Departemens

	for _, v := range resp.KetDepartemen {
		itm := svc.Departemen{
			IdDepartemen:   v.IdDepartemen,
			NamaDepartemen: v.NamaDepartemen,
			Status:         v.Status,
			Keterangan:     v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddDepartemenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddDepartemen",
		encodeAddDepartemenRequest,
		decodeDepartemenResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddDepartemen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddDepartemen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDepartemenEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDepartemen",
		encodeReadDepartemenRequest,
		decodeReadDepartemenResponse,
		pb.ReadDepartemenResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDepartemen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDepartemen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateDepartemen(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateDepartemen",
		encodeUpdateDepartemenRequest,
		decodeDepartemenResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateDepartemen")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateDepartemen",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDepartemenByNama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDepartemenByNama",
		encodeReadDepartemenByNamaRequest,
		decodeReadDepartemenbyNamaRespones,
		pb.ReadDepartemenByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDepartemenByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDepartemenByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDepartemenByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDepartemenByKeterangan",
		encodeReadDepartemenByKeteranganRequest,
		decodeReadDepartemenByKeteranganResponse,
		pb.ReadDepartemenByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDepartemenByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDepartemenByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
