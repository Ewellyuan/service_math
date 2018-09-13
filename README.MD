### 本实例基于 go1.10.3

* 监听端口：8080

* 安装依赖
```
RUN go get -d -v ./...
RUN go install -v ./...
```
* 本项目有 2个 branch，需要分别切换到对应版本打包
```
git checkout service_math_v1
git checkout service_math_v2
```

* 本地运行
```
go run main.go
```

* 测试请求 v1
```
curl -X POST -d '{"x":120,"y":99}' http://localhost:8080/sum
curl -X POST -d '{"x":119,"y":100}' http://localhost:8080/mul
```

* 测试请求 v2
```
curl -X POST -d '{"x":120,"y":99}' http://localhost:8080/sum
curl -X POST -d '{"x":119,"y":100}' http://localhost:8080/mul
curl -X POST -d '{"x":119,"y":100}' http://localhost:8080/dec
```
