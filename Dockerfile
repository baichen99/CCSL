FROM alpine
LABEL maintainer="Adrian Duan<adrianduan@icloud.com>"

RUN apk update \
    && apk --no-cache add tzdata ca-certificates postgresql-client \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \ 
    && echo "Asia/Shanghai" > /etc/timezone

ENV CCSL_ENV=prod
EXPOSE 8888
WORKDIR /app
COPY ccsl .
COPY locales ./locales
COPY configs ./configs

CMD [ "./ccsl" ]