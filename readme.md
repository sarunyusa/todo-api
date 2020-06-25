#Todo API

###How to run
1. Use `make api-docker-image` to create docker image of API
2. `make start-compose` to start API with database
3. `make log-compose` to view log
4. `make stop-compose` and `make restart-compose` to stop and restart API and database
5. The API will available on `http://0.0.0.0:8080`

###API

####`GET` `/`
Health check

####`POST` `/todo`
Create todo

example request
```json
{
	"topic":"test 3",
	"detail":"test detail 3",
	"due_date":"2020-06-30T18:30:00.000Z"
}
```

example response
```json
{
    "code": 200,
    "data": {
        "topic": "test 3",
        "detail": "test detail 3",
        "due_date": "2020-06-30T18:30:00Z",
        "id": "brqheo83u9lina4h78vg",
        "is_done": false,
        "create_at": "2020-06-25T21:30:09.2943343Z",
        "update_at": "2020-06-25T21:30:09.2943343Z"
    }
}
```

####`PUT` `/todo/{id}`
Edit todo of the {id}

example request
```json
{
	"topic":"test 3",
	"detail":"test detail 3",
	"due_date":"2020-06-30T18:30:00.000Z"
}
```

example response
```json
{
    "code": 200,
    "data": {
        "topic": "test 3",
        "detail": "test detail 3",
        "due_date": "2020-06-30T18:30:00Z",
        "id": "brqheo83u9lina4h78vg",
        "is_done": false,
        "create_at": "2020-06-25T21:30:09.2943343Z",
        "update_at": "2020-06-25T21:30:09.2943343Z"
    }
}
```

####`DELETE` `/todo/{id}`
Delete todo of the {id} - There are no request and response body

####`PUT` `/todo/{id}/done`
Set todo of the {id} to done - There are no request and response body

####`GET` `/todo/{id}`
Get todo by {id}
```json
{
    "code": 200,
    "data": {
        "topic": "test 3",
        "detail": "test detail 3",
        "due_date": "2020-06-30T18:30:00Z",
        "id": "brqheo83u9lina4h78vg",
        "is_done": false,
        "create_at": "2020-06-25T21:30:09.2943343Z",
        "update_at": "2020-06-25T21:30:09.2943343Z"
    }
}
```

####`GET` `/todo/`
Get todo that not done yet
```json
{
    "code": 200,
    "data": [
      {
        "topic": "test 3",
        "detail": "test detail 3",
        "due_date": "2020-06-30T18:30:00Z",
        "id": "brqheo83u9lina4h78vg",
        "is_done": false,
        "create_at": "2020-06-25T21:30:09.2943343Z",
        "update_at": "2020-06-25T21:30:09.2943343Z"
      },
      {
        "topic": "test 4",
        "detail": "test detail 4",
        "due_date": "2020-06-30T18:30:00Z",
        "id": "brqheo83u9lina4h78vg",
        "is_done": false,
        "create_at": "2020-06-25T21:30:09.2943343Z",
        "update_at": "2020-06-25T21:30:09.2943343Z"
      }
    ]
}
```
