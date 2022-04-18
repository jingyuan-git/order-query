# Order - Query - Product 

## 1. Project Introduction
    
   - The order query project based on vue and gin, which separates the front and rear of the full stack. The project provides the function of search orders and filter orders, and can display the information related to the amount.

## 2. How to run

-  Required

   - PostgreSQL database
   - gin
   - vue
   - IDE recommendation: VSCode
  

    ```bash
    # clone the project
    git clone https://github.com/jingyuan-git/order-query.git

    # create an order database in PostgreSQL
    # and than will write database-related information into the configuration
    # in `server/conf/app.ini`
    ``` 

### 2.1 server project

- conf

    ```bash
    You should modify `server/conf/app.ini`

    [server]
    ; debug or release
    RunMode = debug
    ; Host = localhost
    HttpPort = 8000
    ...

    [database]
    Type = postgres
    User = 
    Password = 
    Host = 127.0.0.1
    Port = 5432
    Name = 
    TimeZone = Australia/Melbourne
    DataPath = ../data/
    ...
    ```

- build and run

    ``` bash
    cd server

    # use go mod And install the go dependency package
    go mod tidy

    # Compile 
    go build -o server main.go (windows the compile command is go build -o server.exe main.go )

    # Run binary
    ./server (windows The run command is server.exe)
    ```

- Project information and existing API

    ```
    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /api/v1/orders            --> server/routers/v1.GetOrders (6 handlers)

    Listening port is 8000
    ```

### 2.2 web project

- Config
    ```
    You can config in `web-v2/public/config.js`, for server api.

    
    const BaseConfig = {
        ServerApiUrl: "http://localhost:8000",
    }
    ```

- Project setup
    ```
    cd web-v2
    yarn install
    ```

- Compiles and hot-reloads for development
    ```
    yarn serve
    ```

- Compiles and minifies for production
    ```
    yarn build
    ```

- Lints and fixes files
    ```
    yarn lint
    ```

- Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

# 3. Project Display

![image](https://raw.githubusercontent.com/jingyuan-git/order-query/2a36e00c31dcdbe18a55ea4308049e924a105a0e/data/DisplayInterface.png)