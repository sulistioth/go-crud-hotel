package server

import "context"

type Status int32
type CreatedBy string
type UpdatedBy string

const (
	//ServiceID is dispatch service ID
	ServiceID           = "departemen.hotel.id"
	OnAdd     Status    = 1
	OnAdd2    CreatedBy = "Admin"
	OnAdd3    UpdatedBy = "Admin"
)

type Departemen struct {
	IdDepartemen   string
	NamaDepartemen string
	Status         string
	Keterangan     string
}
type Departemens []Departemen

// interface sebagai parameter
type ReadWriter interface {
	AddDepartemen(Departemen) error
	ReadDepartemen() (Departemens, error)
	UpdateDepartemen(Departemen) error
	ReadDepartemenByNama(string) (Departemen, error)
	ReadDepartemenByKeterangan(string) (Departemens, error)
}

//interface sebagai nilai return
type DepartemenService interface {
	AddDepartemenService(context.Context, Departemen) error
	ReadDepartemenService(context.Context) (Departemens, error)
	UpdateDepartemenService(context.Context, Departemen) error
	ReadDepartemenByNamaService(context.Context, string) (Departemen, error)
	ReadDepartemenByKeteranganService(context.Context, string) (Departemens, error)
}
