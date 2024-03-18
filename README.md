# Zumm Technical Exercise

## Background

**This application is deployed at www.lgrv.net.**

This is a technical exercise produced for a company whose product is a niche dating app. It is built with Go 1.22 and the Echo framework.

This is the first Go project I have worked on, having come from a Python background, which was really interesting as it meant I had to strip all language and framework away in my brain and rely on fundamentals.

I used TDD throughout which really helped me encode the requirements, get the app off the ground quickly (originally with Gin and SQLite) and provide a safety net for refactoring, especially in conjunction with the Go debugger `delve`


## Application

### Framework

Initially I began with Gin as it seemed to be the simplest way to get the app going and satisfy the test I'd written for the canary endpoint. However, further down the line, after having already implemented JWT 'manually', I refactored to use Echo, mostly because it had JWT middleware built-in

### Routes

This application offers the following HTTP endpoints:

- / GET
  - Canary/healthcheck simply returning HTTP 200
- /user/create GET
  - Create and return a random user profile
- /login POST
  - Authenticate as a user, returns a JWT token
- /discover GET
  - Get user profiles near you
  - Protected by JWT auth
- /swipe POST
  - Register your verdict on another user
  - Protected by JWT auth


### Database

Initially I began with SQLite, which I am a huge fan of, but when it came to dockerising and productionising the application I reached straight for Postgres, mostly because it's what I'm used to and I didn't want any surprises.

GORM seems to be the SQL Alchemy of the Go world and was an easy choice for the ORM.

### Authentication

The production systems I've worked on in the past have all used AWS IAM for authentication. However, I have a good understanding of encryption and reading the requirements it seemed to me to require symmetrical encryption and decryption with a private secret. JWT is a ubiquitous implementation of symmetrical encryption, and the Echo framework has JWT middleware built in.

## Roadmap

So much of this application is unfinished with regards to normal engineering practice. By way of example, the JWT secret is stored in plain text, and so are the user passwords!
