# Go Fiber Template Project

A starter template for building RESTful APIs with Go Fiber framework and GORM ORM, featuring environment configuration, SQLite database, and auto-migrations.

## Features
- Fiber web framework
- GORM ORM with SQLite support
- Environment variable configuration
- Auto database migrations
- Structured project layout (MVC-like)
- Request logging middleware

````
## Project Structure
├── controller/ # Request handlers
│ └── hello.go
├── model/ # Data models
│ └── user.go
├── migration/ # Database migrations
│ └── migrate.go
├── route/ # Route definitions
│ └── router.go
├── main.go # Application entry point
└── .env.example # Environment template
````

## Prerequisites
- Go 1.16+
- SQLite (for database)

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/glendecado/template.git
   cd template

2. **Set up environment**
    cp .env.example .env

3. **Install dependencies**
    go mod tidy 

4. **Run the application**
    go run main.go


## Models and migration

1. **Open file model*


2. **Define your model structure with GORM tags**
````
// model/model.go
package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}


type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:200;not null"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null;index"` // Foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
}

//you can add more
````

3. **Open migration/migrate.go**

4. **Add your new model to the Migrate slice:**
````
// migration/migrate.go
package migration

import (
	"github.com/glendecado/template/model"
)

var Migrate = []interface{}{
	&model.User{},
	&model.Post{}, // Add new model here
}
````

## Controller and route

1. **Add new File in Controller Folder**
    ex. controller/postController.go

2. **you can add new functions make sure to put "package controller" **

````
package controller

import "github.com/gofiber/fiber/v2"

//make sure the first letter is capitalized
func AddPost(c *fiber.Ctx) error {
	return c.SendString("added")
}
````

3. **put the controller to route**

````
package route

import (
	"github.com/glendecado/template/controller"
	"github.com/gofiber/fiber/v2"
)

func Run(app *fiber.App) {

	//sample route
	app.Get("/", controller.Hello)

    //addnew here
    app.Get("/", controller.AddUser)

    //you can also use
    //app.Get()
    //app.Post()
    //app.Destroy()
    //app.Update()
}

````
