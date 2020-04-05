## Assignment ##

- Authenticate user using Basic Auth (username and password)
- Upon successful authentication, create a JWT token for user
- Store the token in database and cache to check validity
- Send the token to user as HTTP Cookie
- All further requests from user would have this HTTP cookie which can be verified for validity
- Once user logs out, delete the token from cache and database and clear the cookie on user device as well


## Configuration Parameters / Env variables ##
- HTTP_PORT (Server port; default:"3000")
- DATABASE_USER (Postgres Database username; required)
- DATABASE_PASSWORD (Postgres Database password; required)
- DATABASE_HOST (Postgres Database password; required)
- DATABASE_PORT (Postgres Database port; required)
- DATABASE_NAME (Postgres Database name; required)
- REDIS_ADDRESS (Redis address; required)
- REDIS_PASSWORD (Redis password)
- REDIS_DB (Redis database version; default:"0")
- JWT_SIGNING_SECRET (required)
- JWT_EXPIRY_DURATION (default:"30m")


## To create Postgres database ## 
$ psql   
postgres=# CREATE USER jio_user WITH PASSWORD 'jio_pass';  
postgres=# GRANT ALL PRIVILEGES ON DATABASE "jio_db" to jio_user;   
postgres=# CREATE DATABASE jio_db;     
postgres=# GRANT ALL PRIVILEGES ON DATABASE "jio_db" to jio_user;  


## How to run ## 
$ JWT_SIGNING_SECRET=abcdefgh REDIS_ADDRESS="localhost:6379" DATABASE_HOST=localhost DATABASE_PORT=5432 DATABASE_USER=jio_user DATABASE_PASSWORD=jio_pass DATABASE_NAME=jio_db go run main.go 

There are 3 endpoints accessible from browser: 
- http://localhost:3000/login (Asks user to enter credentials)
- http://localhost:3000/secure (Shows whether user is logged in or not)
- http://localhost:3000/logout (To logout the user)
 
Flow:  
Open http://localhost:3000/login in browser and enter credentials as test : test   
User would get redirected to http://localhost:3000/secure upon successful authentication   
After logging out, user would be shown "Not logged in" message on secure page.   


## TODO ## 
- Add Dockerfile
- Add test cases