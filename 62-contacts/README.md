- to run postgres sql

```docker run -d -p 55432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=contactsbd postgres```

- to run pgadmin

```docker run -d -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=jitenp@outlook.com -e PGADMIN_DEFAULT_PASSWORD=root dpage/pgadmin4```

- to delete a container

```docker rm -f containerid```