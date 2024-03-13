# Muzz Task


## API

Started off with Gin for pure simplicity
Echo has jwt auth built in
Gin would need a plugin appleboy/gin-jwt but looks verbose


### Endpoints


GET /user/create
Create a random user
Use a faker type library to generate fake names etc
Public and open

POST /login
Payload of a user id? or just GET token?
Returns an auth token to use to auth against the other endpoints
Public and open

GET /discover
List all other users
Sorted by other user's distance from the user
Add distance from the user to the returned payload
Exclude users already swiped on
Add a filter on age and gender
Authenticate with token

POST /swipe
Payload of their user id and yes/no verdict
Compare that verdict to the their verdict of the user, if one exists
If a match, return matched payload
Authenticate with token



## Database & Models

Use GORM

```

Table users {
  id integer [primary key]
  name varchar
  age integer
  latlong varchar
  gender varchar
  created timestamp
}

Table swipes {
  swiper integer
  swipee integer
  interested bool
  created timestamp
}

Ref: users.id < swipes.swiper
Ref: users.id < swipes.swipee

```


Token authentication

Keep it basic. Symmetric encryption using a private secret
This is basically what JWT is
https://github.com/golang-jwt/jwt
