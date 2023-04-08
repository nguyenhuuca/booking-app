# booking-app
Booking 


#### Context Diagram
---
![Alt text](doc/image/Booking-Context.png)


#### Container Diagram
---
![Alt text](doc/image/Booking-Container.png)

#### Component Diagram
----
![Alt text](doc/image/Booking-component.png)

#### Tructure project
---
```
├───booking-service  
│   ├───api   // handler incoming request
│   ├───config
│   ├───db     
│   │   └───migrations
│   ├───dto
│   ├───logic     
│   ├───repo
│   └───storage  // maping entity
├───doc
│   ├───image
│   └───planuml
└───webapp
    ├───assets
    ├───css
    └───js


```

#### Set up & Installation
---
##### Installation:
- Golang >= 1.20
- PostgresSQL.
- Creata database with name bookingdb
- Dowload all package
    ```
    cd booking-service
    go mod tidy
    ```
- Run migration data, change database server info on your machine.
    ```
    migrate -database postgres://{user}:{pass}@{host}/bookingdb?sslmode=disable -path db/migrations up
    ```


##### Running Appp on LocalHost:
- Change file api/.env.example to api/.env and configure database info for postgres sql
```shell
DB_USER=postgres
```

