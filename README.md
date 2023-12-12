# go-bookstore-oauth-api
Authoritation api

## Cassandra

Cleint https://github.com/gocql/gocql


In case anyone wants a quick way to run Cassandra with docker, here is how you do it:

Create the container with ports mapped as if you installed the binary on your system:

```
docker run --name cassandra -d -p 7199:7199 -p 7000:7000 -p 9042:9042 -p 9160:9160 -p7001:7001 cassandra:3.11
```
You will use `localhost` and default port (9042) to connect from your app.

Create the Keyspace:
```
docker exec -it cassandra  cqlsh -e "CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1}"
```

Create the Table:
```
docker exec -it cassandra cqlsh -e "USE oauth; CREATE TABLE access_tokens(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);"
```

Run commands in cqlsh like this:
```
docker exec -it cassandra cqlsh
```

Stop the service like this:
```
docker stop cassandra
```

Start it again like this:
```
docker start cassandra
```
