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

- JWT : JSON Web Token

- to generate ssl 

- to create tls certificate 

- step -1 : to create private key and csr (certificate signing request)

```openssl req -new -subj "/C=US/ST=Utah/CN=localhost" -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr```

- step-2 : with csr get crt (certificate)
  
```openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt```

- once you have two files, change in your main.go file. Instead of r.Run use r.RUNTLS
  
  -log.Fatal(r.RunTLS(":"+port, "security/localhost.crt", "security/localhost.key"))