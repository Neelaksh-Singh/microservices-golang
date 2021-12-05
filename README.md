# Microservices in Golang

### Standard microsevices in golang
Go into any folder and run `go run main.go`

### CRUD Operations

##### Create(POST)

`curl -v localhost:9090/ -d '{"name": "", "description": "" , "price" : _ , "sku": _ }' | jq`

##### Update(PUT)

`curl -v localhost:9090/{id_of_item} -XPUT -d '{"name": "tea", "description":" a nice cup of tea"}' | jq`

##### Read(GET)

`curl -v localhost:9090 | jq`

#### Viewing Swagger API's := `http://localhost:9090/products/docs`


