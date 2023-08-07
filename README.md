# MongoAndGoBasedCRUDApi
CRUD application created in Go programming language and in Backend used MongoDB Nosql Database for Storing user data

Requirement to run Application is
1. you need to install Monogodb in your local system and mongod server should run
2. you need to give respective operation request through Postman, or any other Http cilent
3. GET request can run by url :-http://localhost:8080/user/get/3 where 3 is an example id of user
4. POST request can run by url :-http://localhost:8080/user/post
5. UPDATE request can run by url:-http://localhost:8080/user/update/3 here you need to give the field value in form ok key-value pair which you want to update
6. DELETE request can run by url:-http://localhost:8080/user/delete/3 user with id=3 will be deleted from database
