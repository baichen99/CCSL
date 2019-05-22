# CCSL API Service

## 环境变量说明

* `CCSL_ENV`: 可取`dev`和`prod`，分别对应开发环境和生产环境

## 开发环境准备

> 以下`Makefile`中的脚本均为`Shell`命令, Windows 用户可使用`WSL`

* 启动开发数据库

```bash
$ make devdb
```

## 部署