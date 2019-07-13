## 第三方组件下载

go get -u github.com/jeffallen/mqtt



## 编译

### windows

go build -o mosquitto_sub.exe main.go



### linux

go build -o mosquitto_sub main.go



### arm 交叉编译 (linux 环境)

(1). go env: 查看当前的 GOARCH 值 (记住, 交叉编译完成后, 需要还原)

(2). 切换为arm环境

export GOARCH=arm

(3). 编译

go build -o mosquitto_sub main.go

(4). 切换为原来的环境

export GOARCH=第(1)步中的环境

如:  export GOARCH=amd64



## 使用

1. 使用 mosquitto_sub -help 查看命令参数