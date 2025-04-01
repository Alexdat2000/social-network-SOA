package utils

import (
	"errors"
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
