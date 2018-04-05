package server

import (
	"context"
)

//langkah ke-5
type departemen struct {
	writer ReadWriter
}

func NewDepartemen(writer ReadWriter) DepartemenService {
	return &departemen{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (t *departemen) AddDepartemenService(ctx context.Context, departemen Departemen) error {
	//fmt.Println("customer")
	err := t.writer.AddDepartemen(departemen)
	if err != nil {
		return err
	}

	return nil
}

func (c *departemen) ReadDepartemenService(ctx context.Context) (Departemens, error) {
	dep, err := c.writer.ReadDepartemen()
	//fmt.Println("customer", cus)
	if err != nil {
		return dep, err
	}
	return dep, nil
}

func (t *departemen) UpdateDepartemenService(ctx context.Context, dep Departemen) error {
	err := t.writer.UpdateDepartemen(dep)
	if err != nil {
		return err
	}
	return nil
}

func (t *departemen) ReadDepartemenByNamaService(ctx context.Context, nama string) (Departemen, error) {
	dep, err := t.writer.ReadDepartemenByNama(nama)
	//fmt.Println("customer:", cus)
	if err != nil {
		return dep, err
	}
	return dep, nil
}

func (t *departemen) ReadDepartemenByKeteranganService(ctx context.Context, keterangan string) (Departemens, error) {
	dep, err := t.writer.ReadDepartemenByKeterangan(keterangan)
	if err != nil {
		return dep, err
	}
	return dep, nil
}
