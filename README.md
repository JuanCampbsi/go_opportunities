![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)
![Test](https://github.com/go-telegram-bot-api/telegram-bot-api/actions/workflows/test.yml/badge.svg)
<div style="width:100%; display: flex; align-items: center;">
  <h1>Api Go REST - GIN 
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="60" width="70" style="margin-bottom: -15px; z-index: -10; margin-left: 1.25rem"/>
  </h1> 
</div>


### 🛠  Description   

</br>

This project aims to create a robust REST API using the Go/GIN programming language and to document the endpoints using Swagger.

The objective is to adhere to design patterns and to modularize functions, separating the configuration logic. For this purpose, a configuration folder has been created which contains files for logging, database connection settings, and SQLite initialization.

A custom logger has been developed for various data returns and endpoint information, such as success, error, etc.

Additionally, a custom data validation mechanism has been created without using external libraries, to validate data in the JSON structure.


## Swagger

</br>

<p align="center">
  <kbd>
 <img width="800" style="border-radius: 10px" height="400" src="https://github.com/JuanCampbsi/Preview_README/blob/86f37a264c34d108e5e1f52e9acc8c144fa81a12/assets/go-oportunities-swaggers.png" alt="Intro"> 
  </kbd>
  </br>
</p>

</br>


### ⌨ Installation
To use it, you need to clone the repository, install the dependencies and run the project.

```bash
# Open terminal/cmd and then Clone this repository
$ git clone https://github.com/JuanCampbsi/go-opportunities.git

# Access project folder in terminal/cmd
$ cd go-opportunities

# Install the dependencies
$ go mod tidy

# Run the application in development mode
$ go run main.go                           

```

</br>	

### ⌨ Stack of technologies and libraries

-   [Golang](https://go.dev/doc/) - version 1.20
-   [GIN](https://github.com/gin-gonic/gin) - version 1.9.0
-   [GORM](https://gorm.io/gorm ) - version 1.25.1
-   [SqLite](https://gorm.io/docs/connecting_to_the_database.html) - version 1.5.1
-   [Swaggo](https://github.com/swaggo/swag) - version 1.16.1
 
</br>

👨‍💻 **Author** 💻

Developed by [_Juan Campos_](https://www.linkedin.com/in/juancampos-ferreira/)