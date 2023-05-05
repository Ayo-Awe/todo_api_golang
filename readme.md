# 🚀 Simple Golang CRUD API using Fiber and GORM 📝

This is a simple Golang CRUD API built with Fiber and GORM. The API allows you to perform CRUD (Create, Read, Update, Delete) operations on a list of Todos. It connects to a MySQL database and provides the following endpoints:

- `GET /api/todos` - Get all todos <br>
  Example Response:

  ```json
  {
    "status": "success",
    "todos": [
      {
        "id": 3,
        "title": "Walk my dog",
        "description": "",
        "completed": true,
        "createdAt": "2023-05-04T21:12:26.648+01:00",
        "updatedAt": "2023-05-04T21:12:26.648+01:00"
      },
      {
        "id": 7,
        "title": "Make Cheese",
        "description": "Make some good tasting cheese",
        "completed": false,
        "createdAt": "2023-05-04T22:43:33.143+01:00",
        "updatedAt": "2023-05-04T22:43:33.143+01:00"
      }
    ]
  }
  ```

- `GET /api/todos/:id` - Get a single todo by ID<br> Example Response

  ```json
  {
    "status": "success",
    "todo": {
      "id": 3,
      "title": "Walk my dog",
      "description": "",
      "completed": false,
      "createdAt": "2023-05-04T21:12:26.648+01:00",
      "updatedAt": "2023-05-04T21:12:26.648+01:00"
    }
  }
  ```

- `POST /api/todos/:id` - Create a todo with a title and description<br>
  Example Request Body

  ```json
  {
    "title": "Go to the gym",
    "description": ""
  }
  ```

  <br>
    Example Response

  ```json
  {
    "status": "success",
    "todo": {
      "id": 5,
      "title": "Go to the gym",
      "description": "",
      "completed": false,
      "createdAt": "2023-05-04T21:12:26.648+01:00",
      "updatedAt": "2023-05-04T21:12:26.648+01:00"
    }
  }
  ```

- `PATCH /api/todos/:id` - Edit the title or description of a todo
  <br>Example Request Body

  ```json
  {
    "title": "Take a course",
    "description": "Take a codewithmosh course on NodeJS"
  }
  ```

  Example Response

  ```json
  {
    "status": "success",
    "todo": {
      "id": 3,
      "title": "Take a course",
      "description": "Take a codewithmosh course on NodeJS",
      "completed": false,
      "createdAt": "2023-05-04T21:12:26.648+01:00",
      "updatedAt": "2023-05-04T21:12:36.648+01:00"
    }
  }
  ```

- `PUT /api/todos/:id/status` - Update the completed status of a todo
  <br> Example Request Body

  ```json
  {
    "completed": true
  }
  ```

  <br>Example Response

  ```json
  {
    "status": "success",
    "todo": {
      "id": 5,
      "title": "Go to the gym",
      "description": "",
      "completed": true,
      "createdAt": "2023-05-04T21:12:26.648+01:00",
      "updatedAt": "2023-05-04T21:12:26.648+01:00"
    }
  }
  ```

- `DELETE /api/todos/:id` - Delete a todo by ID
  <br>Example Response

  ```json
  {
    "status": "success",
    "deleted": 2
  }
  ```

## 🛠️ Installation

1. Clone the repository: `git clone https://github.com/Ayo-Awe/todo_api_golang.git`
2. Navigate to the project directory: `todo_api_golang`
3. Install the dependencies: `go get ./...`
4. Set up the MySQL database and update the `.env` file with your database credentials. Depending on when you're reading this, there could be no .env file, you can edit the dsn in the database.go file
5. Run the server: `go run app.go`

## 🛡️ Authentication

This API does not require authentication.

## 🔑 Authorization

This API does not require authorization.

## 📦 Dependencies

This API uses the following dependencies:

- [Fiber](https://github.com/gofiber/fiber/) - Web framework for Golang
- [GORM](https://gorm.io/) - ORM library for Golang
- [MySQL](https://www.mysql.com/) - Open-source relational database management system

## 🚀 Deployment

This API is currently has a live deployment on [Render](https://render.com/).
Here is the API base url `https://go-todo-api.onrender.com/api/`

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check out the [issues page](https://github.com/Ayo-Awe/todo_api_golang/issues).

## 📄 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). See the [LICENSE](LICENSE) file for more information.
