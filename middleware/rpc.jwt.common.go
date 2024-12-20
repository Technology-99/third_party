package middleware

import (
	"context"
	"fmt"
	"github.com/Technology-99/third_party/cache_key"
	"github.com/Technology-99/third_party/commKey"
	"github.com/Technology-99/third_party/consts"
	utils "github.com/Technology-99/third_party/cryptography"
	"github.com/Technology-99/third_party/response"
	"github.com/Technology-99/third_party/sony"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// note: 基于grpc的中间件，实现jwt校验

func (SvcCtx *MiddleContext) JwtUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	result := new(Resp)

	ctx = context.WithValue(ctx, CtxFullMethod, info.FullMethod)
	result.Path = info.FullMethod

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

	// note: token校验
	Authorization := tempMD.Get(commKey.HeaderAuthorization)

	logc.Infof(ctx, "打印这个authorization: %v", Authorization)

	if len(Authorization) < 1 {
		result.Code = response.AUTHORIZATION_NOT_FOUND
		result.Msg = response.StatusText(response.AUTHORIZATION_NOT_FOUND)
		return result, nil
	}

	if len(Authorization[0]) <= 7 {
		result.Code = response.AUTHORIZATION_NOT_FOUND
		result.Msg = response.StatusText(response.AUTHORIZATION_NOT_FOUND)
		return result, nil
	}

	token := GetAccessToken(Authorization)

	// note: 获取redis中的key
	AccessKey := tempMD.Get(commKey.HeaderXAccessKeyFor)
	if len(AccessKey) < 1 {
		result.Code = response.ACCESSKEY_NOT_FOUND
		result.Msg = response.StatusText(response.ACCESSKEY_NOT_FOUND)
		//resultMsg, _ := ReturnProtoMsg(*result, info.FullMethod)
		//logx.Infof("能运行到这里吗")
		return result, nil
	}

	key := fmt.Sprintf(cache_key.ACCESS_TOKEN_KEY, consts.SHIELD_RPC_SERVICE_NAME, AccessKey[0])

	AccessPublicKey, err := SvcCtx.Redis.Get(key)
	if err != nil {
		result.Code = response.ACCESSKEY_NOT_FOUND
		result.Msg = response.StatusText(response.ACCESSKEY_NOT_FOUND)
		return result, nil
	}

	if len(AccessPublicKey) <= 10 {
		result.Code = response.ACCESSKEY_NOT_FOUND
		result.Msg = response.StatusText(response.ACCESSKEY_NOT_FOUND)
		return result, nil
	}

	// note: 判断是否需要校验jwt
	claims, err := utils.ParseToken(token, AccessPublicKey)
	if err != nil {
		result.Code = response.ACCESS_TOKEN_INVALID
		result.Msg = response.StatusText(response.ACCESS_TOKEN_INVALID)
		return result, nil
	}

	if err = claims.Valid(); err != nil {
		result.Code = response.ACCESS_TOKEN_INVALID
		result.Msg = response.StatusText(response.ACCESS_TOKEN_INVALID)
		return result, nil
	}

	logc.Infof(ctx, "打印解码出来的内容: %v+", claims)
	ctx = context.WithValue(ctx, "AccessKey", AccessKey[0])
	ctx = context.WithValue(ctx, "UserId", claims.UserId)
	ctx = context.WithValue(ctx, "TenantID", claims.TenantID)

	return handler(ctx, req)
}

func (SvcCtx *MiddleContext) JwtStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (resp interface{}, err error) {
	result := new(Resp)

	ctx := ss.Context()

	ctx = context.WithValue(ctx, CtxFullMethod, info.FullMethod)
	result.Path = info.FullMethod

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

	// note: token校验
	Authorization := tempMD.Get(commKey.HeaderAuthorization)

	logc.Infof(ctx, "打印这个authorization: %v", Authorization)

	if len(Authorization) < 1 {
		result.Code = response.AUTHORIZATION_NOT_FOUND
		result.Msg = response.StatusText(response.AUTHORIZATION_NOT_FOUND)
		return result, nil
	}

	if len(Authorization[0]) <= 7 {
		result.Code = response.AUTHORIZATION_NOT_FOUND
		result.Msg = response.StatusText(response.AUTHORIZATION_NOT_FOUND)
		return result, nil
	}

	token := GetAccessToken(Authorization)

	// note: 获取redis中的key
	AccessKey := tempMD.Get(commKey.HeaderXAccessKeyFor)
	if len(AccessKey) < 1 {
		result.Code = response.ACCESSKEY_NOT_FOUND
		result.Msg = response.StatusText(response.ACCESSKEY_NOT_FOUND)
		//resultMsg, _ := ReturnProtoMsg(*result, info.FullMethod)
		//logx.Infof("能运行到这里吗")
		return result, nil
	}

	key := fmt.Sprintf(cache_key.ACCESS_TOKEN_KEY, consts.SHIELD_RPC_SERVICE_NAME, AccessKey[0])

	AccessPublicKey, err := SvcCtx.Redis.Get(key)
	if err != nil {
		result.Code = response.ACCESSKEY_NOT_FOUND
		result.Msg = response.StatusText(response.ACCESSKEY_NOT_FOUND)
		return result, nil
	}

	if len(AccessPublicKey) <= 10 {
		result.Code = response.ACCESSKEY_NOT_FOUND
		result.Msg = response.StatusText(response.ACCESSKEY_NOT_FOUND)
		return result, nil
	}

	// note: 判断是否需要校验jwt
	claims, err := utils.ParseToken(token, AccessPublicKey)
	if err != nil {
		result.Code = response.ACCESS_TOKEN_INVALID
		result.Msg = response.StatusText(response.ACCESS_TOKEN_INVALID)
		return result, nil
	}

	if err = claims.Valid(); err != nil {
		result.Code = response.ACCESS_TOKEN_INVALID
		result.Msg = response.StatusText(response.ACCESS_TOKEN_INVALID)
		return result, nil
	}

	logc.Infof(ctx, "打印解码出来的内容: %v+", claims)
	ctx = context.WithValue(ctx, "AccessKey", AccessKey[0])
	ctx = context.WithValue(ctx, "UserId", claims.UserId)
	ctx = context.WithValue(ctx, "TenantID", claims.TenantID)

	return handler(srv, ss), nil
}

func GetAccessToken(headAuthorization []string) string {
	return strings.TrimPrefix(headAuthorization[0], "Bearer ")
}
