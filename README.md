#Fidibo Interview 
This is a Task which is assigned to me during interview process for Fidibo company.

###How to Run?
To run the Project:
* run the `create_db.sql` scrip in `adapter/store/sql` folder to create our database to store Users.
* run `docker compose up` and enjoy the project.

###How to Send Request
There is three endpoints:
* `/register`: creates a new user to be able to use main functionality of the project, which is calling and API from Fidibo.
* `/login`: this end point generate a valid jwt token for already registered users. you must put the token inside your header for calling the `/search/book` endpoint.
* `/search/book`: this endpoint does the main work, it sends a request to Fidibo's search service and parses the result. it also has a caching mechanism to reduce time of repetitive requests. it stores the received data for `10 min` in redis cache. NOTE that if you haven't a valid jwt token, you can't use this endpoint.  