[![Build Status](https://travis-ci.org/rcarmstrong/go-bamboo.svg?branch=master)](https://travis-ci.org/rcarmstrong/go-bamboo) [![GoDoc](https://godoc.org/github.com/rcarmstrong/go-bamboo?status.svg)](https://godoc.org/github.com/rcarmstrong/go-bamboo)

# Go-Bamboo
## Go client library for communicating with an Atlassian Bamboo CI Server API.

go-github requires Go version 1.7 or greater.


**Atlassian's Bamboo API documentation refers to endpoints as 'services'. This client library was modeled after go-github which also uses the term 'service' to describe an interface that implements all methods that interact with a certain endpoint group (the plan service implements plan related methods). As to not confuse the two, this documentation will refer to an API 'service' as a resource, i.e. go-bamboo exposes multiple services to communicate with the Bamboo API resources.**

## Table of Contents ##
- [Usage](#usage)
    * [Authenticaiton](#authenticaiton)
- [Bamboo Rest API Documentation](#bamboo-rest-api-documentation)
- [Permissions](#permissions)
    * [Project Plan Permissions](#sproject-plan-permissions)



## Usage ##
```go
import bamboo "github.com/rcarmstrong/go-bamboo"
```

### Authenticaiton ###
At the moment, go-bamboo only supports simple credentials for authentication

```go
bambooClient := bamboo.NewSimpleClient(nil, "myUsername", "myPassword")

// Optionally set a different connection URL for the bamboo client.
// Defaults to "http://localhost:8085/rest/api/latest/"
bambooClient.SetURL("https://my.bambooserver.com:8085/")
```

You may optionally pass in your own http client, replacing the nil above, to be used as the go-bamboo http client.

## Bamboo Rest API Documentation ##
Atlassian Bamboo's Rest API documentation can be frustrating at time in how much it lacks in detail. With this project, I hope to save you from some of that frustration. The API documentation can be found [here](https://docs.atlassian.com/atlassian-bamboo/REST/6.2.5/) for those who are curious, with a more detailed but incomplete doc living [here.](https://developer.atlassian.com/server/bamboo/bamboo-rest-resources/)

## Permissions ##
Bamboo allows an admin to set access control on resources such as projects and plans. Most things have five levels of access:
- View
- Edit
- Build
- Clone
- Admin

The expected strings for these permissions are defined as the constants ReadPermission, WritePermission, BuildPermission, ClonePermission, and AdminPermission. Read and Write are the same as View and Edit, the names just differ from the UI to the API.

### Project Plan Permissions ###

Project plan permissions refers to the permissions a plan inherited for the project for a specific set of users, groups, or roles. The ProjectPlan service exposes the addition, removal, and changing of these permissions. Individual users, groups and the Logged In Users role can be given permission to view(read)/edit(write)/build/clone/administer(admin) the project's plans. Only the Anonymous Users role is restricted to only being able to have view permission.

## Cloning a plan ##

Returns general info about the API.

-- TODO --
example

