package endpoint

import (
	"context"

	scv "projectHotel/hotel/departemen/server"

	pb "projectHotel/hotel/departemen/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcDepartemenServer struct {
	addDepartemen              grpctransport.Handler
	readDepartemen             grpctransport.Handler
	updateDepartemen           grpctransport.Handler
	readDepartemenByNama       grpctransport.Handler
	readDepartemenByKeterangan grpctransport.Handler
}

func NewGRPCDepartemenServer(endpoints DepartemenEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.DepartemenServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcDepartemenServer{
		addDepartemen: grpctransport.NewServer(endpoints.AddDepartemenEndpoint,
			decodeAddDepartemenRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddDepartemen", logger)))...),
		readDepartemen: grpctransport.NewServer(endpoints.ReadDepartemenEndpoint,
			decodeReadDepartemenRequest,
			encodeReadDepartemenResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDepartemen", logger)))...),
		updateDepartemen: grpctransport.NewServer(endpoints.UpdateDepartemenEndpoint,
			decodeUpdateDepartemenRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateDepartemen", logger)))...),
		readDepartemenByNama: grpctransport.NewServer(endpoints.ReadDepartemenByNamaEndpoint,
			decodeReadDepartemenByNamaRequest,
			encodeReadDepartemenByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDepartemenByNama", logger)))...),
		readDepartemenByKeterangan: grpctransport.NewServer(endpoints.ReadDepartemenByKeteranganEndpoint,
			decodeReadDepartemenByKeteranganRequest,
			encodeReadDepartemenByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDepartemenByKeterangan", logger)))...),
	}
}

func decodeAddDepartemenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddDepartemenReq)
	return scv.Departemen{IdDepartemen: req.GetIdDepartemen(), NamaDepartemen: req.GetNamaDepartemen(),
		Status: req.GetStatus(), Keterangan: req.GetKeterangan()}, nil
}

func decodeReadDepartemenRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateDepartemenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateDepartemenReq)
	return scv.Departemen{IdDepartemen: req.IdDepartemen, NamaDepartemen: req.NamaDepartemen,
		Status: req.Status, Keterangan: req.Keterangan}, nil
}

func decodeReadDepartemenByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDepartemenByNamaReq)
	return scv.Departemen{NamaDepartemen: req.NamaDepartemen}, nil

}

func decodeReadDepartemenByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDepartemenByKeteranganReq)
	return scv.Departemen{Keterangan: req.Keterangan}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadDepartemenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Departemens)

	rsp := &pb.ReadDepartemenResp{}

	for _, v := range resp {
		itm := &pb.ReadDepartemenByNamaResp{
			IdDepartemen:   v.IdDepartemen,
			NamaDepartemen: v.NamaDepartemen,
			Status:         v.Status,
			Keterangan:     v.Keterangan,
		}
		rsp.AllDepartemen = append(rsp.AllDepartemen, itm)
	}
	return rsp, nil
}

func encodeReadDepartemenByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Departemen)
	return &pb.ReadDepartemenByNamaResp{IdDepartemen: resp.IdDepartemen, NamaDepartemen: resp.NamaDepartemen,
		Status: resp.Status, Keterangan: resp.Keterangan}, nil
}

func (s *grpcDepartemenServer) AddDepartemen(ctx oldcontext.Context, departemen *pb.AddDepartemenReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addDepartemen.ServeGRPC(ctx, departemen)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcDepartemenServer) ReadDepartemen(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadDepartemenResp, error) {
	_, resp, err := s.readDepartemen.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDepartemenResp), nil
}

func (s *grpcDepartemenServer) UpdateDepartemen(ctx oldcontext.Context, dep *pb.UpdateDepartemenReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateDepartemen.ServeGRPC(ctx, dep)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcDepartemenServer) ReadDepartemenByNama(ctx oldcontext.Context, nama *pb.ReadDepartemenByNamaReq) (*pb.ReadDepartemenByNamaResp, error) {
	_, resp, err := s.readDepartemenByNama.ServeGRPC(ctx, nama)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDepartemenByNamaResp), nil
}

func encodeReadDepartemenByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Departemens)
	rsp := &pb.ReadDepartemenByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadDepartemenByNamaResp{
			IdDepartemen:   v.IdDepartemen,
			NamaDepartemen: v.NamaDepartemen,
			Status:         v.Status,
			Keterangan:     v.Keterangan,
		}
		rsp.KetDepartemen = append(rsp.KetDepartemen, itm)
	}

	return rsp, nil
}

func (s *grpcDepartemenServer) ReadDepartemenByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadDepartemenByKeteranganReq) (*pb.ReadDepartemenByKeteranganResp, error) {
	_, resp, err := s.readDepartemenByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDepartemenByKeteranganResp), nil
}
