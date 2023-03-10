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
sudo service postgresql stop
```

```sh
docker ps -a
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
PUT localhost:3000/rental/:rentalid/return

## Output PDF (reciept)
GET localhost:3000/rental/receipt/:rentalid

## Frontend
```sh
npm install -g http-server
cd frontend
http-server
```
Go to http://127.0.0.1:8080 in browser