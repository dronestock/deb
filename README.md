# Deb
[![编译状态](https://github.ruijc.com:20443/api/badges/dronestock/deb/status.svg)](https://github.ruijc.com:20443/dronestock/deb)
[![Golang质量](https://goreportcard.com/badge/github.com/dronestock/deb)](https://goreportcard.com/report/github.com/dronestock/deb)
![版本](https://img.shields.io/github/go-mod/go-version/dronestock/deb)
![仓库大小](https://img.shields.io/github/repo-size/dronestock/deb)
![最后提交](https://img.shields.io/github/last-commit/dronestock/deb)
![授权协议](https://img.shields.io/github/license/dronestock/deb)
![语言个数](https://img.shields.io/github/languages/count/dronestock/deb)
![最佳语言](https://img.shields.io/github/languages/top/dronestock/deb)
![星星个数](https://img.shields.io/github/stars/dronestock/deb?style=social)

持续集成系统Deb插件，用于快速生成Debain/Ubuntu等类系统安装包

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
- name: 上传到腾讯云
  image: ccr.ccs.tencentyun.com/dronestock/deb
  settings:
    secret_id: xxx
    secret_key: xxx
```

更多使用教程，请参考[文档](https://www.dronestock.tech/plugin/stock/drone)

## 交流

![微信群](https://www.dronestock.tech/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 插件列表

- [Git](https://www.dronestock.tech/plugin/stock/git) 使用Git推送和拉取代码
- [Maven](https://www.dronestock.tech/plugin/stock/maven) Maven编译、打包、测试以及发布到仓库
- [Protobuf](https://www.dronestock.tech/plugin/stock/protobuf) Protobuf编译、静态检查以及高级功能
- [Docker](https://www.dronestock.tech/plugin/stock/docker) Docker编译、打包以及发布到镜像仓库
- [Node](https://www.dronestock.tech/plugin/stock/node) Node编译、打包以及发布到仓库
- [Cos](https://www.dronestock.tech/plugin/stock/cos) 腾讯云对象存储基本配置、文件上传等
- [Mcu](https://www.dronestock.tech/plugin/stock/mcu) 各种模块依赖文件修改
- [Apisix](https://www.dronestock.tech/plugin/stock/apisix) Apisix网关插件
- [Ftp](https://www.dronestock.tech/plugin/stock/ftp) Ftp文件插件

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢

[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://www.jetbrains.com/?from=dronestock/deb)
