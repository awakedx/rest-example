# REST API server example 
## Using features:
- Echo go web frabework - https://echo.labstack.com
- DB postgres, pgx -  https://github.com/jackc/pgx
- Config based on viper - https://github.com/spf13/viper
- Postman queries inlcuded

## Run on local machine
Create .env file with fields
- DB_USER=youruser
- DB_PASSWORD=yourpassword
- DB_HOST=localhost
- DB_PORT=5432
- DB_NAME=yourdbname
- SECRET_ACCESS=yoursecret  
![imgstart](https://github.com/awakedx/rest-example/blob/master/readme/start.png)
## If you want run in docker
- Change DB_HOST=db(named service in docker-compose.yml)  
![dockerpng](https://github.com/awakedx/rest-example/blob/master/readme/dockerbuild.png)
## I deployed this on AWS EC2 in docker , u can test it through
[Swagger doc](http://13.61.25.183/swagger//index.html)
