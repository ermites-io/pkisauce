//go:generate protoc -I. --go-grpc_out=require_unimplemented_servers=false:.  --go_out=. --go_opt=paths=source_relative svc.proto

package svc
