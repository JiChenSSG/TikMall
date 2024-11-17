// Code generated by Kitex v0.11.3. DO NOT EDIT.

package authservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	auth "github.com/jichenssg/tikmall/gen/kitex_gen/auth"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"DeliverToken": kitex.NewMethodInfo(
		deliverTokenHandler,
		newDeliverTokenArgs,
		newDeliverTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RefreshToken": kitex.NewMethodInfo(
		refreshTokenHandler,
		newRefreshTokenArgs,
		newRefreshTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"VerifyToken": kitex.NewMethodInfo(
		verifyTokenHandler,
		newVerifyTokenArgs,
		newVerifyTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteToken": kitex.NewMethodInfo(
		deleteTokenHandler,
		newDeleteTokenArgs,
		newDeleteTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	authServiceServiceInfo                = NewServiceInfo()
	authServiceServiceInfoForClient       = NewServiceInfoForClient()
	authServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func deliverTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.DeliverTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).DeliverToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeliverTokenArgs:
		success, err := handler.(auth.AuthService).DeliverToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeliverTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeliverTokenArgs() interface{} {
	return &DeliverTokenArgs{}
}

func newDeliverTokenResult() interface{} {
	return &DeliverTokenResult{}
}

type DeliverTokenArgs struct {
	Req *auth.DeliverTokenReq
}

func (p *DeliverTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.DeliverTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeliverTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeliverTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeliverTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeliverTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.DeliverTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeliverTokenArgs_Req_DEFAULT *auth.DeliverTokenReq

func (p *DeliverTokenArgs) GetReq() *auth.DeliverTokenReq {
	if !p.IsSetReq() {
		return DeliverTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeliverTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeliverTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeliverTokenResult struct {
	Success *auth.DeliveryResp
}

var DeliverTokenResult_Success_DEFAULT *auth.DeliveryResp

func (p *DeliverTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.DeliveryResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeliverTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeliverTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeliverTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeliverTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.DeliveryResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeliverTokenResult) GetSuccess() *auth.DeliveryResp {
	if !p.IsSetSuccess() {
		return DeliverTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeliverTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.DeliveryResp)
}

func (p *DeliverTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeliverTokenResult) GetResult() interface{} {
	return p.Success
}

func refreshTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RefreshTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).RefreshToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RefreshTokenArgs:
		success, err := handler.(auth.AuthService).RefreshToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefreshTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRefreshTokenArgs() interface{} {
	return &RefreshTokenArgs{}
}

func newRefreshTokenResult() interface{} {
	return &RefreshTokenResult{}
}

type RefreshTokenArgs struct {
	Req *auth.RefreshTokenReq
}

func (p *RefreshTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RefreshTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefreshTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefreshTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefreshTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefreshTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.RefreshTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefreshTokenArgs_Req_DEFAULT *auth.RefreshTokenReq

func (p *RefreshTokenArgs) GetReq() *auth.RefreshTokenReq {
	if !p.IsSetReq() {
		return RefreshTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefreshTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefreshTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefreshTokenResult struct {
	Success *auth.RefreshResp
}

var RefreshTokenResult_Success_DEFAULT *auth.RefreshResp

func (p *RefreshTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.RefreshResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefreshTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefreshTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefreshTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefreshTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.RefreshResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefreshTokenResult) GetSuccess() *auth.RefreshResp {
	if !p.IsSetSuccess() {
		return RefreshTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefreshTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.RefreshResp)
}

func (p *RefreshTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefreshTokenResult) GetResult() interface{} {
	return p.Success
}

func verifyTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.VerifyTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).VerifyToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *VerifyTokenArgs:
		success, err := handler.(auth.AuthService).VerifyToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*VerifyTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newVerifyTokenArgs() interface{} {
	return &VerifyTokenArgs{}
}

func newVerifyTokenResult() interface{} {
	return &VerifyTokenResult{}
}

type VerifyTokenArgs struct {
	Req *auth.VerifyTokenReq
}

func (p *VerifyTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.VerifyTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *VerifyTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *VerifyTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *VerifyTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *VerifyTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.VerifyTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var VerifyTokenArgs_Req_DEFAULT *auth.VerifyTokenReq

func (p *VerifyTokenArgs) GetReq() *auth.VerifyTokenReq {
	if !p.IsSetReq() {
		return VerifyTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *VerifyTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VerifyTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type VerifyTokenResult struct {
	Success *auth.VerifyResp
}

var VerifyTokenResult_Success_DEFAULT *auth.VerifyResp

func (p *VerifyTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.VerifyResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *VerifyTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *VerifyTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *VerifyTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *VerifyTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.VerifyResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *VerifyTokenResult) GetSuccess() *auth.VerifyResp {
	if !p.IsSetSuccess() {
		return VerifyTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *VerifyTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.VerifyResp)
}

func (p *VerifyTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VerifyTokenResult) GetResult() interface{} {
	return p.Success
}

func deleteTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.DeleteTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).DeleteToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteTokenArgs:
		success, err := handler.(auth.AuthService).DeleteToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteTokenArgs() interface{} {
	return &DeleteTokenArgs{}
}

func newDeleteTokenResult() interface{} {
	return &DeleteTokenResult{}
}

type DeleteTokenArgs struct {
	Req *auth.DeleteTokenReq
}

func (p *DeleteTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.DeleteTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.DeleteTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteTokenArgs_Req_DEFAULT *auth.DeleteTokenReq

func (p *DeleteTokenArgs) GetReq() *auth.DeleteTokenReq {
	if !p.IsSetReq() {
		return DeleteTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteTokenResult struct {
	Success *auth.DeleteTokenResp
}

var DeleteTokenResult_Success_DEFAULT *auth.DeleteTokenResp

func (p *DeleteTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.DeleteTokenResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.DeleteTokenResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteTokenResult) GetSuccess() *auth.DeleteTokenResp {
	if !p.IsSetSuccess() {
		return DeleteTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.DeleteTokenResp)
}

func (p *DeleteTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteTokenResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DeliverToken(ctx context.Context, Req *auth.DeliverTokenReq) (r *auth.DeliveryResp, err error) {
	var _args DeliverTokenArgs
	_args.Req = Req
	var _result DeliverTokenResult
	if err = p.c.Call(ctx, "DeliverToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefreshToken(ctx context.Context, Req *auth.RefreshTokenReq) (r *auth.RefreshResp, err error) {
	var _args RefreshTokenArgs
	_args.Req = Req
	var _result RefreshTokenResult
	if err = p.c.Call(ctx, "RefreshToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyToken(ctx context.Context, Req *auth.VerifyTokenReq) (r *auth.VerifyResp, err error) {
	var _args VerifyTokenArgs
	_args.Req = Req
	var _result VerifyTokenResult
	if err = p.c.Call(ctx, "VerifyToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteToken(ctx context.Context, Req *auth.DeleteTokenReq) (r *auth.DeleteTokenResp, err error) {
	var _args DeleteTokenArgs
	_args.Req = Req
	var _result DeleteTokenResult
	if err = p.c.Call(ctx, "DeleteToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
