package controllers

import (
	"context"
	"fmt"
)

type key string

var UserIDKey key

func SetUserID(c context.Context, t string) context.Context {
	return context.WithValue(c, UserIDKey, t)
}

func GetUserID(ctx context.Context) (string, error) {
	v := ctx.Value(UserIDKey)

	user_id, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("user_id not found")
	}

	return user_id, nil
}
