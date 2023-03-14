- to run postgres sql

```docker run -d -p 55432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=contactsbd postgres```

- to run pgadmin

```docker run -d -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=jitenp@outlook.com -e PGADMIN_DEFAULT_PASSWORD=root dpage/pgadmin4```

- to delete a container

```docker rm -f containerid```

- to start mysql

```docker run -d --name mysql1 -p 53306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e MYSQL_DATABASE=contactsdb mysql```

- to start adminer (Webased GUI for databases like mysql,postgres,mongodb,oracle,MS Sql,elasticsearch)

```docker run -d --name=database-gui -p 8080:8080 adminer```