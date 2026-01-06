# instalacion docker
> descargar: imagen docker pull postgres:16
> levantar contenedor: docker run --name som-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5431:5432 -d postgres:16

### librerias gorm

>go get -u gorm.io/gorm
>go get -u gorm.io/driver/postgres
>go get -u github.com/google/uuid