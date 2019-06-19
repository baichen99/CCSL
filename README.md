# CCSL 中国手语与聋人研究中心

## 环境变量说明

* `CCSL_ENV`: 可取`dev`和`prod`，分别对应开发环境和生产环境

## 开发环境准备

> 以下`Makefile`中的脚本均为Shell脚本, Windows 用户可使用`WSL`

* 创建开发数据库

```bash
$ make devdb
```

* 导入数据库（SQL文件）

* 下载依赖

```bash
$ make download
```
> 在下载的时候可能会出现因为被墙无法下载依赖的情况，可以使用proxychains之类的终端代理软件

* 启动后端

```bash
$ make backend
```

* 启动前端

```bash
$ make frontend
```

## 部署

* 部署前端

```bash
$ make deploy-frontend
```

* 部署后端

```bash
$ make deploy-backend
# ssh to server
$ 
```