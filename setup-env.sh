#!/bin/bash

# DB
echo "DIALECT=$DIALECT" >> .env
echo "USER_NAME=$MYSQL_USER" >> .env
echo "PASSWORD=$PASSWORD" >> .env
echo "PROTOCOL=$PROTOCOL" >> .env
echo "DB_NAME=$DB_NAME" >> .env