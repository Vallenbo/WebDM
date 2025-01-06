# 使用官方 MySQL 镜像作为基础镜像
FROM mysql

# 设置 MySQL 环境变量
ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=mydb
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# 将本地的初始化 SQL 脚本拷贝到容器中的 /docker-entrypoint-initdb.d/ 目录下
COPY init.sql /docker-entrypoint-initdb.d/

# 暴露 3306 端口
EXPOSE 3306

# 启动 MySQL 服务
CMD ["mysqld"]