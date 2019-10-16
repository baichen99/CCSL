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

> 在部署之前确保能够使用ccsl@ccsl.shu.edu.cn命令直接ssh连接到服务器，需要开启VPN

- 部署前端

```bash
$ make deploy-frontend
```

- 部署后端

```bash
$ make deploy-backend # 输入Api版本号，上一个版本号在configs/.version文件中，版本号依次递增即可
```
* ssh 到服务器
```bash
$ export CCSL_VERSION=${版本号}
$ cd api && docker-compose up -d
```
## 备注

### 配置说明
- App：应用相关配置
    - Domain：部署域名
    - CORS：允许跨域资源请求的域名
    - AllowCookie：是否允许Cookie

- Log：日志
    - Level：日志级别，可为fatal, error, warn, info, debug, disable

- Listener：监听域名和端口设置

- JWT：JSON Web Token相关的配置
    - ExpireHours：JWT密钥有效时间（小时）
    - PrivateKey和PublicKey：JWT密钥，其中签名算法采用ES512算法，公钥和私钥的生成方法如下
    - Debug：是否显示debug日志

```bash
$ openssl ecparam -genkey -name secp521r1 -noout -out ecdsa-p521-private.pem
# 拷贝ecdsa-p521-private.pem文件内的内容到PrivateKey
$ openssl ec -in ecdsa-p521-private.pem -pubout -out ecdsa-p521-public.pem
# 拷贝ecdsa-p521-public.pem文件内的内容到PublicKey
```

- Postgresql：PG数据库相关的设置
    - Server：数据库服务器地址
    - Port：数据库端口
    - User：数据库用户
    - Password：数据库密码
    - Database：数据库名称

- File：上传文件相关设置
    - Dir：上传文件存储目录



### 环境变量说明
- `CCSL_ENV`: 可取`dev`和`prod`，分别对应开发环境和生产环境，通过此环境变量读取配置文件
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
