package utils

import (
	"errors"
	"google.golang.org/grpc/codes"
	"net/http"
	"strconv"
	"strings"
)

func ParsePostId(postId string) (uint32, error) {
	id, err := strconv.Atoi(postId)
	if err != nil {
		return 0, err
	}
	if id < 0 {
		return 0, errors.New("invalid post id")
	}
	return uint32(id), nil
}

func ParsePostPrivate(isPrivate string) bool {
	if strings.ToLower(isPrivate) == "true" || strings.ToLower(isPrivate) == "private" {
		return true
	}
	return false
}

func GrpcCodeToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
