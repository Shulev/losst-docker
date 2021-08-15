## ___Install:___
```
mkdir -p /opt/pgadmin
```
```
mkdir -p /opt/pgadmin/{postgres-data,pgadmin-data}
```

## ___Usage:___
```
Clone the repo
```
```
docker-compose up -d
```


## ___Example:___
```
Open pgAdmin4 with http://localhost:5454 in the browser
```
```
psql -h localhost -p 5432 -U postgres
```

