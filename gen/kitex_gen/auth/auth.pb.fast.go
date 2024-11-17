// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package auth

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DeliverTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeliverTokenReq[number], err)
}

func (x *DeliverTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeliveryResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeliveryResp[number], err)
}

func (x *DeliveryResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AccessToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliveryResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RefreshTokenReq[number], err)
}

func (x *RefreshTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RefreshResp[number], err)
}

func (x *RefreshResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AccessToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VerifyTokenReq[number], err)
}

func (x *VerifyTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyTokenReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Method, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyTokenReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Path, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VerifyResp[number], err)
}

func (x *VerifyResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *VerifyResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Url, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Method, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteTokenReq[number], err)
}

func (x *DeleteTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteTokenResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *DeliverTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeliverTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *DeliveryResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeliveryResp) fastWriteField1(buf []byte) (offset int) {
	if x.AccessToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAccessToken())
	return offset
}

func (x *DeliveryResp) fastWriteField2(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRefreshToken())
	return offset
}

func (x *RefreshTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RefreshTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetRefreshToken())
	return offset
}

func (x *RefreshResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RefreshResp) fastWriteField1(buf []byte) (offset int) {
	if x.AccessToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAccessToken())
	return offset
}

func (x *RefreshResp) fastWriteField2(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRefreshToken())
	return offset
}

func (x *VerifyTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *VerifyTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *VerifyTokenReq) fastWriteField2(buf []byte) (offset int) {
	if x.Method == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMethod())
	return offset
}

func (x *VerifyTokenReq) fastWriteField3(buf []byte) (offset int) {
	if x.Path == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPath())
	return offset
}

func (x *VerifyResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *VerifyResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *VerifyResp) fastWriteField2(buf []byte) (offset int) {
	if x.Url == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUrl())
	return offset
}

func (x *VerifyResp) fastWriteField3(buf []byte) (offset int) {
	if x.Method == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetMethod())
	return offset
}

func (x *DeleteTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *DeleteTokenResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *DeliverTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeliverTokenReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *DeliveryResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeliveryResp) sizeField1() (n int) {
	if x.AccessToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAccessToken())
	return n
}

func (x *DeliveryResp) sizeField2() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRefreshToken())
	return n
}

func (x *RefreshTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RefreshTokenReq) sizeField1() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetRefreshToken())
	return n
}

func (x *RefreshResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RefreshResp) sizeField1() (n int) {
	if x.AccessToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAccessToken())
	return n
}

func (x *RefreshResp) sizeField2() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRefreshToken())
	return n
}

func (x *VerifyTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *VerifyTokenReq) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *VerifyTokenReq) sizeField2() (n int) {
	if x.Method == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMethod())
	return n
}

func (x *VerifyTokenReq) sizeField3() (n int) {
	if x.Path == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPath())
	return n
}

func (x *VerifyResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *VerifyResp) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *VerifyResp) sizeField2() (n int) {
	if x.Url == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUrl())
	return n
}

func (x *VerifyResp) sizeField3() (n int) {
	if x.Method == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetMethod())
	return n
}

func (x *DeleteTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteTokenReq) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *DeleteTokenResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_DeliverTokenReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_DeliveryResp = map[int32]string{
	1: "AccessToken",
	2: "RefreshToken",
}

var fieldIDToName_RefreshTokenReq = map[int32]string{
	1: "RefreshToken",
}

var fieldIDToName_RefreshResp = map[int32]string{
	1: "AccessToken",
	2: "RefreshToken",
}

var fieldIDToName_VerifyTokenReq = map[int32]string{
	1: "Token",
	2: "Method",
	3: "Path",
}

var fieldIDToName_VerifyResp = map[int32]string{
	1: "UserId",
	2: "Url",
	3: "Method",
}

var fieldIDToName_DeleteTokenReq = map[int32]string{
	1: "Token",
}

var fieldIDToName_DeleteTokenResp = map[int32]string{}
