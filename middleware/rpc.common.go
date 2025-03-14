package middleware

import (
	"context"
	"errors"
	"github.com/Technology-99/third_party/commKey"
	"github.com/Technology-99/third_party/response"
	"github.com/Technology-99/third_party/sony"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// note: 基于grpc的中间件，实现读取metadata中的信息映射到context中

func StreamHeaderParseInterceptor() grpc.StreamServerInterceptor {
	return func(svr any, stream grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		ctx := stream.Context()

		result := &Resp{
			Code: response.SUCCESS,
			Msg:  "ok",
			Path: info.FullMethod,
		}

		ctx = context.WithValue(ctx, CtxFullMethod, info.FullMethod)

		// note: metadata中尝试获取requestId, 如果不存在就生成一个
		tempMD, isExist := metadata.FromIncomingContext(ctx)
		if !isExist {
			result.Code = response.METADATA_NOT_FOUND
			result.Msg = response.StatusText(response.METADATA_NOT_FOUND)
			return errors.New(response.StatusText(response.METADATA_NOT_FOUND))
		}

		requestId := tempMD.Get(commKey.HeaderXRequestIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxRequestID, requestId[0])
			result.RequestID = requestId[0]
		} else {
			tempRequestId := sony.NextId()
			ctx = context.WithValue(ctx, CtxRequestID, tempRequestId)
			result.RequestID = tempRequestId
		}

		//note: 读取metadata中的信息
		xTenantIDFor := tempMD.Get(commKey.HeaderXTenantIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxTenantId, xTenantIDFor[0])
		} else {
			result.Code = response.METADATA_NOT_FOUND
			result.Msg = response.StatusText(response.METADATA_NOT_FOUND)
			return errors.New(response.StatusText(response.METADATA_NOT_FOUND))
		}

		xDomainIdFor := tempMD.Get(commKey.HeaderXDomainIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxDomainId, xDomainIdFor[0])
		} else {
			result.Code = response.METADATA_NOT_FOUND
			result.Msg = response.StatusText(response.METADATA_NOT_FOUND)
			return errors.New(response.StatusText(response.METADATA_NOT_FOUND))
		}

		return handler(svr, stream)
	}
}

func UnaryHeaderParseInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (any, error) {

		result := &Resp{
			Code: response.SUCCESS,
			Msg:  "ok",
			Path: info.FullMethod,
		}

		ctx = context.WithValue(ctx, CtxFullMethod, info.FullMethod)

		// note: metadata中尝试获取requestId, 如果不存在就生成一个
		tempMD, isExist := metadata.FromIncomingContext(ctx)
		if !isExist {
			result.Code = response.METADATA_NOT_FOUND
			result.Msg = response.StatusText(response.METADATA_NOT_FOUND)
			return result, nil
		}

		requestId := tempMD.Get(commKey.HeaderXRequestIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxRequestID, requestId[0])
			result.RequestID = requestId[0]
		} else {
			tempRequestId := sony.NextId()
			ctx = context.WithValue(ctx, CtxRequestID, tempRequestId)
			result.RequestID = tempRequestId
		}

		//note: 读取metadata中的信息
		xTenantIDFor := tempMD.Get(commKey.HeaderXTenantIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxTenantId, xTenantIDFor[0])
		} else {
			result.Code = response.METADATA_NOT_FOUND
			result.Msg = response.StatusText(response.METADATA_NOT_FOUND)
			return result, nil
		}

		xDomainIdFor := tempMD.Get(commKey.HeaderXDomainIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxDomainId, xDomainIdFor[0])
		} else {
			result.Code = response.METADATA_NOT_FOUND
			result.Msg = response.StatusText(response.METADATA_NOT_FOUND)
			return result, nil
		}

		return handler(ctx, req)
	}
}
