#编译调试

makefile中已经写好编译脚本，先前端编译，生成前端dist文件，然后编译后端，生成generate文件，最后修改相应的环境配置（config/config.ini）
## 脚本编译
~~~ bin
make build-vue
make build-go
make run
~~~

## 源代码编译
~~~ bin
npm install
npm run build
npm run serve

go generate cmd/gomanager/main.go
go build -o bin/gomanager cmd/gomanager/main.go
./bin/gomanager --config=./config/config.ini
~~~

## docker
ko配置
~~~ bin
go install github.com/google/ko@latest
~~~
设置docker的repo
~~~ bin
export PATH=$PATH:$(GOBIN)
export GOPATH="$(go env GOPATH)"
export KO_DOCKER_REPO='gcr.io/my-project'
ko login -p 密码 -u 用户名 gcr.io
~~~
镜像制作 (前提是执行go generate生成statik文件)
~~~
ko publish ./cmd/gomanager --tags=release_1.0.0 --base-import-path
~~~