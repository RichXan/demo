#!/bin/bash

# 等待 MySQL 启动
echo "Waiting for MySQL to be ready..."
for i in {1..30}; do
    if docker exec xblog-mysql mysqladmin ping -h"localhost" -p"root123" --silent > /dev/null 2>&1; then
        break
    fi
    echo -n "."
    sleep 1
done
echo ""

# 检查并显示字符集配置
echo "Checking MySQL character set configuration..."
docker exec xblog-mysql mysql -uroot -proot123 -e "SHOW VARIABLES LIKE '%character%'; SHOW VARIABLES LIKE '%collation%';"

# 创建数据库（如果不存在）
echo "Creating database..."
docker exec xblog-mysql mysql -uroot -proot123 -e "DROP DATABASE IF EXISTS x_micro_blog; CREATE DATABASE x_micro_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 初始化数据库
echo "Initializing database schema..."
docker exec -i xblog-mysql mysql -uroot -proot123 --default-character-set=utf8mb4 x_micro_blog < scripts/init.sql

# 验证表的字符集
echo "Verifying table character sets..."
docker exec xblog-mysql mysql -uroot -proot123 x_micro_blog -e "SELECT table_name, table_collation FROM information_schema.tables WHERE table_schema = 'x_micro_blog';"

echo "Database initialization completed!" 