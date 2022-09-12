FROM postgres:latest
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=aims_enrollment_db
COPY ./init.sql /docker-entrypoint-initdb.d/init.sql