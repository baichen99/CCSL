# CCSL 中国手语与聋人研究中心

## 环境变量说明

- `CCSL_ENV`: 可取`dev`和`prod`，分别对应开发环境和生产环境

## 开发环境准备

> 以下`Makefile`中的脚本均为 Shell 脚本, Windows 用户可使用`WSL`

- 创建开发数据库

```bash
$ make devdb
```

- 导入数据库（SQL 文件）

- 下载依赖

```bash
$ make download
```

> 在下载的时候可能会出现因为被墙无法下载依赖的情况，可以使用 proxychains 之类的终端代理软件

- 启动后端

```bash
$ make backend
```

- 启动前端

```bash
$ make frontend
```

## 部署

- 部署前端

```bash
$ make deploy-frontend
```

- 部署后端

```bash
$ make deploy-backend
# ssh to server
$
```

## 语料库

- 国家通用手语对比语料库：UniversalContrast
- 上海手语动词语料库: ShanghaiVerb
- 专有名词语料库: ProperNoun
- 手语语篇数据库: Text
- 手语研究文献数据库
- 手语研究术语库
