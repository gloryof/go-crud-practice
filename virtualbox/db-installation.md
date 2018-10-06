## PostgreSQLをインストール

```
# yum -y localinstall https://download.postgresql.org/pub/repos/yum/10/redhat/rhel-7-x86_64/pgdg-centos10-10-2.noarch.rpm
# yum -y install postgresql10-server
# /usr/pgsql-10/bin/postgresql-10-setup initdb
# systemctl enable postgresql-10
# systemctl start postgresql-10
```

## firewalldの設定
```
# firewall-cmd --permanent --new-service=postgres

# firewall-cmd --permanent --service=postgres --set-short=postgres
# firewall-cmd --permanent --service=postgres --set-description=postgres
# firewall-cmd --permanent --service=postgres --add-port=5432/tcp

# firewall-cmd --permanent --add-service=postgres --zone=public

# firewall-cmd --reload
```

## PostgreSQLの設定変更
```
$ diff -u /var/lib/pgsql/10/data/postgresql.conf{.orig,}
--- /var/lib/pgsql/10/data/postgresql.conf.orig	2018-09-24 12:01:51.974214881 +0900
+++ /var/lib/pgsql/10/data/postgresql.conf	2018-09-24 12:02:49.419361469 +0900
@@ -56,7 +56,7 @@
 
 # - Connection Settings -
 
-#listen_addresses = 'localhost'		# what IP address(es) to listen on;
+listen_addresses = '*'		# what IP address(es) to listen on;
 					# comma-separated list of addresses;
 					# defaults to 'localhost'; use '*' for all
 					# (change requires restart)
@@ -110,7 +110,7 @@
 
 # - Memory -
 
-shared_buffers = 128MB			# min 128kB
+shared_buffers = 512MB			# min 128kB
 					# (change requires restart)
 #huge_pages = try			# on, off, or try
 					# (change requires restart)
```
```
$ diff /var/lib/pgsql/10/data/pg_hba.conf{,.orig}
83,84d82
< host	all		all		192.168.1.0/24		md5
```
```
# systemctl restart postgresql-10
```

## DBユーザの追加

```
# su - postgres
$ psql
```
ログインしたら`/env/db/ddl/create-database.sql`を実行。
