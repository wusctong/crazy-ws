# 使用Java 8及以上版本的基础镜像（CraftSocketSocketProxy要求Java 8+）
FROM openjdk:8-jre-slim

# 设置工作目录
WORKDIR /app

# 将本地的JAR文件复制到容器中（确保文件名与你实际的JAR一致）
COPY CraftSocketProxy1.0.1.jar /app/craftsocketproxy.jar

# 暴露端口（Render会动态分配端口，这里只是提示作用）
EXPOSE 8080

# 启动命令（关键：使用Render的动态端口$PORT，替换为你的MC服务器信息）
CMD ["sh", "-c", "java -jar /app/craftsocketproxy.jar --s \
  -host play.onecube.fr \
  -port 25565 \
  -proxy $PORT \
  -path /boost"]