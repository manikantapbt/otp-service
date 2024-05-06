package server

import (
	"connectrpc.com/connect"
	"context"
	v1 "otp-service/internal/gen/v1"
	"otp-service/internal/service"
)

type IOtpServer interface {
	GenerateOTP(ctx context.Context, req *connect.Request[v1.GenerateOTPRequest]) (*connect.Response[v1.GenerateOTPResponse], error)
}

func NewOtpService(service service.IOtpService) IOtpServer {
	return &otpService{Service: service}
}

type otpService struct {
	Service service.IOtpService
}

func (o *otpService) GenerateOTP(ctx context.Context,
	req *connect.Request[v1.GenerateOTPRequest]) (*connect.Response[v1.GenerateOTPResponse], error) {
	response := &v1.GenerateOTPResponse{}

	err := o.Service.SendMessage(req.Msg)
	if err != nil {
		response.Error = &v1.OtpError{
			Message:   err.Error(),
			ErrorCode: 1,
		}
	} else {
		response.IsSuccess = true
	}
	return connect.NewResponse(response), nil
}
