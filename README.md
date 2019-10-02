# CCSL 中国手语与聋人研究中心

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

> 在下载的时候可能会出现因为被墙无法下载依赖的情况，可以使用 proxychains 之类的终端代理软件(或者设置GOPROXY)

- 启动后端开发

```bash
$ make backend
```

- 启动前端开发

```bash
$ make frontend
```

## 部署

- 部署前端

```bash
$ make deploy-frontend
```

- 部署后端
替换${version_number}为版本号数字
```bash
$ make deploy-backend # 输入Api版本号
# ssh到服务器
$ load
$ export CCSL_VERSION=${version_number}
$ cd api
$ docker-compose down
$ docker-compose up -d
```
## 备注

### 环境变量说明
- `CCSL_ENV`: 可取`dev`和`prod`，分别对应开发环境和生产环境
- `POSTGRES_USER`: 数据库用户名
- `POSTGRES_PASSWORD`: 数据库密码
- `POSTGRES_DB`: 数据库名

### 语料库名称

- 国家通用手语对比语料库：Lexical Database for Chinese National Sign Language (Lexical Database)
- 上海手语动词语料库: Corpus for Shanghai Sign Language Verb(Verb Corpus)
- 专有名词语料库: Corpus for Proper Nouns in CSL(Proper Nouns Corpus)
- 手语语篇数据库: Chinese Sign Language Corpus for Sign Texts(Text Corpus)
- 手语研究文献数据库: Literature Database for Sign Language Research(Literature Database)
- 手语研究术语库: Database for Technical Terms in Sign Linguistics(Terms Database)
