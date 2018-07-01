CREATE USER "crud-user" WITH PASSWORD 'crud-user';
CREATE DATABASE "go-crud" WITH OWNER = "crud-user" ENCODING = "UTF-8";
GRANT ALL PRIVILEGES ON DATABASE "go-crud" TO "crud-user";