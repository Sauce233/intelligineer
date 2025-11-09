#!/bin/bash
PORT=8091 \
JWT_KEY=cognitive \
PATH_RBAC=/code/smart-engi-backend/rbac.yaml \
MYSQL_USER=engineering \
MYSQL_PASS=engineering \
MYSQL_HOST=127.0.0.1 \
MYSQL_PORT=3306 \
MYSQL_DB=engineering \
REDIS_HOST=127.0.0.1 \
REDIS_PORT=6379 \
REDIS_PASS=Engagement \
REDIS_DB=9 \
PATH_PHOTO=/var/www/file/construction-photos/ \
PATH_CONTRACT=/var/www/file/contracts/ \
COOKIE_DOMAIN=axogc.net \
COOKIE_HTTPONLY=true \
go run .
