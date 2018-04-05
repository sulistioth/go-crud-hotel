package endpoint

import (
	"context"
	"fmt"

	sv "projectHotel/hotel/departemen/server"
)

func (ce DepartemenEndpoint) AddDepartemenService(ctx context.Context, departemen sv.Departemen) error {
	_, err := ce.AddDepartemenEndpoint(ctx, departemen)
	return err
}

func (ce DepartemenEndpoint) ReadDepartemenService(ctx context.Context) (sv.Departemens, error) {
	resp, err := ce.ReadDepartemenEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Departemens), err
}

func (ce DepartemenEndpoint) UpdateDepartemenService(ctx context.Context, dep sv.Departemen) error {
	_, err := ce.UpdateDepartemenEndpoint(ctx, dep)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce DepartemenEndpoint) ReadDepartemenByNamaService(ctx context.Context, nama string) (sv.Departemen, error) {
	req := sv.Departemen{NamaDepartemen: nama}
	resp, err := ce.ReadDepartemenByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	dep := resp.(sv.Departemen)
	return dep, err
}

func (me DepartemenEndpoint) ReadDepartemenByKeteranganService(ctx context.Context, keterangan string) (sv.Departemens, error) {
	req := sv.Departemen{Keterangan: keterangan}
	fmt.Println(req)
	resp, err := me.ReadDepartemenByKeteranganEndpoint(ctx, req)
	fmt.Println("me resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	dep := resp.(sv.Departemens)
	return dep, err
}
