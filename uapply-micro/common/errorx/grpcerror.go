package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcErrorToHttp(err error) error {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return NewError(e.Message(), NotFound)
			case codes.InvalidArgument:
				return NewError(e.Message(), ParamInvalid)
			case codes.Internal:
				return NewSysError()
			case codes.AlreadyExists:
				return NewError(e.Message(), Exist)
			default:
				return NewSysError()
			}
		}
		return NewSysError()
	}
	return NewSysError()
}
