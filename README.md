# video store  
  
## Setup
docker build -t videostore-postgres-db ./  
  
docker run -d --name videostore-postgresdb-container -p 5432:5432 videostore-postgres-db  
  
## Troubleshoot: Stop postgresql service
sudo service postgresql stop


