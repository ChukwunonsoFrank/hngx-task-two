# HNGX Backend Task(Stage 2)

👩‍💻 HNGX Backend Task 1 - A CRUD API that handles a person(user) resource

## Getting Started

#### Clone the repo:

```bash
git clone https://github.com/ChukwunonsoFrank/hngx-task-two.git
cd hngx-task-two
```

#### Set environment variables:

```bash
cp .env.example .env
```
Ensure you have PostgreSQL installed on your local machine to be able to spin up a test database for the setup.

####  Start the server:

```bash
go build && ./hng-task-two
```
This will spin up the local server on port 8080 and log a similar message as shown below.

```bash
2023/09/13 14:45:46 Server running on port 8000
```

## Request/Response Formats

#### GET request

```bash
GET /api/{user_id} HTTP/1.1
```
```json
    {
        "id": {user_id},
        "name": "example_name"
    }
```

#### POST request

```bash
POST /api/{user_id} HTTP/1.1

{
    "name": "example_name"
}
```
```json
    {
        "id": {user_id},
        "name": "example_name"
    }
```

#### PUT request

```bash
PUT /api/{user_id} HTTP/1.1

{
    "name": "new_example_name"
}
```
```json
    {
        "id": {user_id},
        "name": "new_example_name"
    }
```

#### DELETE request

```bash
DELETE /api/{user_id} HTTP/1.1
```
```json
    {
        "id": {user_id},
        "name": "example_name"
    }
```

## Sample Usage

As shown below, we are making a GET request with HTTPie API client but you can make use of any API testing tool available to you.

#### Sample GET request
![Sample GET request]()

