# Retro Video Store - VHS & DVD
  
## Setup
### Create DB:
```sh
sudo service postgresql stop
```

```sh
docker build -t videostore-postgres-db ./  
```  

```sh
docker run -d --name videostore-postgresdb-container -p 5432:5432 videostore-postgres-db  
```

### Migrate DB:
Creates table **rentals** with columns. 
```sh
go run migrate/migratedb.go
```

### Run DB:
```sh
docker start <container_id>
```

## Run
```sh
go run main.go
```

## Make requests
### Post new rental:
POST localhost:3000/rental  
{ "VideoName": "Die hard", "Customer": "John Smith" }  

### List rentals:
GET localhost:3000/rentals?customer=John%20Smith

### Return rental:
PUT localhost:3000/rental/:id/return

## Output PDF (WIP)
receipt.pdf