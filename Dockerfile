FROM postgres
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB videostore
COPY videostore.sql /docker-entrypoint-initdb.d/