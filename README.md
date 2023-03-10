# NETFLAKES

Netflakes is an assessment test from [busha](https://busha.co) to consume and expose movie apis from `[https://swapi.dev/api](https://swapi.dev/api)` . The project is built using the ports and adapters pattern. This enables the components of the application to easily be replaced given that such components adhere to the specified interfaces.

### Endpoints

> BASE_URL: [https://netflakes.herokuapp.com](https://netflakes.herokuapp.com/api/movies)

**Fetch Movies - GET**

`/api/movies`

**Add Comment - POST**

`/api/movies/:movie_id/add-comment`

Request Body Sample

> {
> "body”: “This is a sample comment for testing purposes only….”,
> “created_by”: “Teej4y”
> }

**Get Movie Comments - GET**

`/api/movies/:movie_id/comments`

**Get Movie Characters - GET**

`/api/movies/:movie_id/characters?sort_by=<height | gender | name>,order=<asc | desc>,filter_by_gender=<male | female>`

### POSTMAN COLLECTION

[https://www.postman.com/cloudy-meteor-121927/workspace/busha-test/collection/11505765-553fb34e-317a-4cf7-ab28-9a13c4087173?action=share&creator=11505765](https://www.postman.com/cloudy-meteor-121927/workspace/busha-test/collection/11505765-553fb34e-317a-4cf7-ab28-9a13c4087173?action=share&creator=11505765)

**RUNNING THE PROJECT**

To run the project on docker in a local environment, you would need to set the environment variables below as part of the docker env when creating the container using the `-e` flag.
If the project, is ran locally, then an `app.env` file is needed to store the environment variables.

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=teej4y
DB_PASSWORD=password
DB_NAME=netflakes
DB_URL="postgresql://root:secret@localhost:5432/netflakes?sslmode=disable"
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=password
```

### IMPROVEMENTS

[ ] More test cases needs to be written to cover all or most of the functionalites
[ ] Performance Metrics and evaluations could be added
