FROM scratch
COPY ./build/app1 /
COPY ./config/config.dev.yaml /
ENTRYPOINT ["/app1"]


# docker build -t admin-server:latest .     #构建docker镜像
# docker run -d -p 8082:8082 --name admin admin-server:latest   #运行docker容器