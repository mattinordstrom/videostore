# video store  
  
## Setup
docker build -t videostore-postgres-db ./  
  
docker run -d --name videostore-postgresdb-container -p 5432:5432 videostore-postgres-db  
  
## Troubleshoot: Stop postgresql service
sudo service postgresql stop


## MIGRATE
go run migrate/migratedb.go

## POST RENTAL
POST localhost:3000/rental  
{ "VideoName": "Die hard", "Customer": "John Smith" }  

## GET RENTALS
GET localhost:3000/rentals