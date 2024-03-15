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
Returns an auth token to use to auth against the other endpoints
Public and openee

GET /discover
- [x] List all other users
- [x] Sorted by other user's distance from the user
- [x] Add distance from the user to the returned payload
- [x] Authenticate with token 
- [x] EXTEND: Add a filter on age and gender
- [ ] EXTEND: Exclude users already swiped on
- [ ] BONUS: attractiveness model

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



Fundamentals:
- data modelling
- tdd
- debugging
- symmetric encryption -> jwt
- choosing a framework
- trunk dev branch
- conventional commits