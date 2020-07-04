#!/bin/bash

DB_DIR=$(cd $(dirname $0) && pwd)
cd $DB_DIR

mysql -uroot -ppassword -e "DROP DATABASE IF EXISTS isubata; CREATE DATABASE isubata;"
mysql -uroot -ppassword isubata < ./isubata.sql
