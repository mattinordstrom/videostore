# Retro Video Store - VHS 
    
## Setup
### Stop current postgres:
```sh
sudo systemctl stop postgresql
```
or
```sh
sudo systemctl stop postgresql
sudo systemctl disable postgresql
```

### Create DB:


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

## Start DB and Run
```sh
./start.sh
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

___

  
##  
##  
## Run tests:  
```sh
$ go test ./... -v
```  
  

##  
##  
## Run lint:  
```sh  
$ golangci-lint run -v
```  
  

##  
##  
## Precommit hook:  
```sh
$ pip3 install pre-commit
$ pre-commit install
```  
  
##  
#### Run precommit manually    
```sh
$ pre-commit run --all-files
```  