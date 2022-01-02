# WOL
一个简单的 Wake On Lan 唤醒程序

## 配置
```yaml
port: 8080
nic: "eno1"     # 指定设备的网卡（一台设备多网卡的情况）
url: "/powerOn" # 如果是暴露在公网请设置复杂的 url 值
macAddress: "14-C9-F1-0D-FC-36" # 需要唤醒设备的 MAC 地址
```

## 绑定 IOS 快捷指令
![image](https://mikuac.com/images/github_wol_preview.jpg)

## 交叉编译

Mac 下编译 Linux 和 Windows 64位可执行程序
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

Linux 下编译 Mac 和 Windows 64位可执行程序
```shell
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

Windows 下编译 Mac 和 Linux 64位可执行程序
```shell
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```
