# Project `wife.storage`

Simple RESTful service for Wife project.

## Description

There are available hierarchy of entities in service:

* User 1
  * Friends
    * User ...
  * Tasks
    * Task ...
* User 2
  * Friends
    * User ...
  * Tasks
    * Task ...
* ...

## Stack

* [Go](https://golang.org/) 
* [PostgreSQL](https://www.postgresql.org/)

## API

Data is available with `Authorization` header in request ([Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication)). Validating is made by request:

* `POST host:port/login` - authentication try

### Users

Format `JSON`:

	{
		"login": "user",
		"password": "password",
		"name": "James May"
	}

Requests:

* `GET host:port/user/{id}` - returns single user data
* `GET host:port/users` - returns array for all users data
* `POST host:port/user` - creates user
