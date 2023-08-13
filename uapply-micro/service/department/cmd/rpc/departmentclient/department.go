// Code generated by goctl. DO NOT EDIT!
// Source: department.proto

package departmentclient

import (
	"context"

	"uapply-micro/service/department/cmd/rpc/department"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreInfoRsp  = department.CreInfoRsp
	CreReq      = department.CreReq
	DelReq      = department.DelReq
	DepInfoRsp  = department.DepInfoRsp
	DepReq      = department.DepReq
	DepsInfoRsp = department.DepsInfoRsp
	EachDep     = department.EachDep
	Empty       = department.Empty
	GetRegRsp   = department.GetRegRsp
	MarkInfo    = department.MarkInfo
	RbdReq      = department.RbdReq
	RbdRsp      = department.RbdRsp
	RegData     = department.RegData
	RegInfo     = department.RegInfo
	RegReq      = department.RegReq
	UdpReq      = department.UdpReq
	UserUid     = department.UserUid

	Department interface {
		CreateDep(ctx context.Context, in *CreReq, opts ...grpc.CallOption) (*CreInfoRsp, error)
		GetDep(ctx context.Context, in *DepReq, opts ...grpc.CallOption) (*DepInfoRsp, error)
		GetDeps(ctx context.Context, in *DepReq, opts ...grpc.CallOption) (*DepsInfoRsp, error)
		UpdDep(ctx context.Context, in *UdpReq, opts ...grpc.CallOption) (*DepInfoRsp, error)
		DelDep(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DepInfoRsp, error)
		Register(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*Empty, error)
		GetRegister(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*GetRegRsp, error)
		GetRegByDepIds(ctx context.Context, in *RbdReq, opts ...grpc.CallOption) (*RbdRsp, error)
		GetRegisterMsg(ctx context.Context, in *UserUid, opts ...grpc.CallOption) (*RegInfo, error)
		MarkUser(ctx context.Context, in *MarkInfo, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultDepartment struct {
		cli zrpc.Client
	}
)

func NewDepartment(cli zrpc.Client) Department {
	return &defaultDepartment{
		cli: cli,
	}
}

func (m *defaultDepartment) CreateDep(ctx context.Context, in *CreReq, opts ...grpc.CallOption) (*CreInfoRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.CreateDep(ctx, in, opts...)
}

func (m *defaultDepartment) GetDep(ctx context.Context, in *DepReq, opts ...grpc.CallOption) (*DepInfoRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.GetDep(ctx, in, opts...)
}

func (m *defaultDepartment) GetDeps(ctx context.Context, in *DepReq, opts ...grpc.CallOption) (*DepsInfoRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.GetDeps(ctx, in, opts...)
}

func (m *defaultDepartment) UpdDep(ctx context.Context, in *UdpReq, opts ...grpc.CallOption) (*DepInfoRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.UpdDep(ctx, in, opts...)
}

func (m *defaultDepartment) DelDep(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DepInfoRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.DelDep(ctx, in, opts...)
}

func (m *defaultDepartment) Register(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*Empty, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultDepartment) GetRegister(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*GetRegRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.GetRegister(ctx, in, opts...)
}

func (m *defaultDepartment) GetRegByDepIds(ctx context.Context, in *RbdReq, opts ...grpc.CallOption) (*RbdRsp, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.GetRegByDepIds(ctx, in, opts...)
}

func (m *defaultDepartment) GetRegisterMsg(ctx context.Context, in *UserUid, opts ...grpc.CallOption) (*RegInfo, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.GetRegisterMsg(ctx, in, opts...)
}

func (m *defaultDepartment) MarkUser(ctx context.Context, in *MarkInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := department.NewDepartmentClient(m.cli.Conn())
	return client.MarkUser(ctx, in, opts...)
}
