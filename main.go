package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

)

type IntService interface {
	Sum(context.Context, int, int) int
	Mul(context.Context, int, int) int
	Dec(context.Context, int, int) int
}

type IntServiceImpl struct {}
func (IntServiceImpl) Sum(_ context.Context, x, y int) int {
	return x+y
}
func (IntServiceImpl) Mul(_ context.Context, x, y int) int {
	return x*y
}
func (IntServiceImpl) Dec(_ context.Context, x, y int) int {
	return x-y
}

// 请求、响应消息体
type sumRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type sumResponse struct {
	S int `json:"s"`
}
// 提供给 NewServer 函数
func makeSumEndpoint(intS IntService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(sumRequest)
		S := intS.Sum(ctx, req.X, req.Y)
		return sumResponse{S}, nil
	}
}
func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// 请求、响应消息体
type mulRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type mulResponse struct {
	S int `json:"s"`
}
func makeMulEndpoint(intS IntService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(mulRequest)
		S := intS.Mul(ctx, req.X, req.Y)
		return mulResponse{S}, nil
	}
}
func decodeMulRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request mulRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// 请求、响应消息体
type decRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type decResponse struct {
	S int `json:"s"`
}
func makeDecEndpoint(intS IntService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(decRequest)
		S := intS.Dec(ctx, req.X, req.Y)
		return decResponse{S}, nil
	}
}
func decodeDecRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request decRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// 主入口
func main() {
	initS := IntServiceImpl{}

	sumHandler := httptransport.NewServer(
		makeSumEndpoint(initS),
		decodeSumRequest,
		encodeResponse,
	)

	mulHandler := httptransport.NewServer(
		makeMulEndpoint(initS),
		decodeMulRequest,
		encodeResponse,
	)

	decHandler := httptransport.NewServer(
		makeDecEndpoint(initS),
		decodeDecRequest,
		encodeResponse,
	)

	// 路由配置
	http.Handle("/sum", sumHandler)
	http.Handle("/mul", mulHandler)
	http.Handle("/dec", decHandler)

	// 日志
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -X POST -d '{"x":119,"y":99999999}' http://localhost:8080/sum
// curl -X POST -d '{"x":119,"y":100}' http://localhost:8080/mul
