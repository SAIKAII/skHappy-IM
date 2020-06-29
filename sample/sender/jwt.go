package main

import "context"

type JWTSt struct {
	JWTString string
}

func (j *JWTSt) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"jwt": j.JWTString}, nil
}

func (j *JWTSt) RequireTransportSecurity() bool {
	return false
}
