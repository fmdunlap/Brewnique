docker run --name brewnique-postgres --rm -p 5432:5432 -e POSTGRES_PASSWORD=localdevdb -d -v ${PWD}/dev_db:/var/lib/postgresql/data postgres