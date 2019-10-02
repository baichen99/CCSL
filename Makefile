.PHONY:build

# 构建Docker镜像
build:
	@make compile && read -p "Input API Version:" CCSL_VERSION && echo Deploying API Version $${CCSL_VERSION} now...&& git tag $${CCSL_VERSION} && git push origin --tags && echo $${CCSL_VERSION} > configs/.version  && docker build -t ccsl-api:$${CCSL_VERSION} . && docker save ccsl-api:$${CCSL_VERSION} > docker.tar

# 编译二进制Linux可执行文件
compile:
	@export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build

download:
	@go mod download

# 创建开发数据库
devdb:
	@docker run --name ccsl-pg -e POSTGRES_PASSWORD=password -e POSTGRES_USER=ccsl -e POSTGRES_DB=ccsl -p 0.0.0.0:5432:5432 -d postgres

# 清理环境
clean:
	@go mod tidy && rm -rf node_modules
# 运行后端开发
backend:
	@export CCSL_ENV=dev && go run main.go
# 运行前端开发
frontend:
	@npm run dev

deploy-backend:
	@make build && scp docker.tar ccsl@ccsl.shu.edu.cn:/home/ccsl/

deploy-frontend:
	@npm run build && scp -r dist/* ccsl@ccsl.shu.edu.cn:/home/ccsl/web/
