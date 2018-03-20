[![Build Status](https://travis-ci.org/rcarmstrong/go-bamboo.svg?branch=master)](https://travis-ci.org/rcarmstrong/go-bamboo) [![GoDoc](https://godoc.org/github.com/rcarmstrong/go-bamboo?status.svg)](https://godoc.org/github.com/rcarmstrong/go-bamboo)

# Go-Bamboo
## go-bamboo is a Go client library for accessing the Atlassian Bamboo CI API v3.

go-github requires Go version 1.7 or greater.


**As to not confuse the ideas of an API service and a service within the code used to communicate with that API service; the code uses the term resource to describe the API service. I.e. go-bamboo exposes multiple services to communicate with the API resources.**

## Table of Contents ##
- [Usage](#usage)
    * [Authenticaiton](#authenticaiton)
- [General-Information](#general-information)
- [Server-Information](#server-information)
- [Project-Service](#project-service)
- [Plan-Service](#plan-service)
- [Result-Service](#result-service)
- [Chart-Service](#chart-service)
- [Queue-Service](#queue-service)
- [Export-Service](#export-service)
- [Clone-Service](#clone-service)
- [Dependency-Service](#dependency-service)
- [Elastic-Configuration-Service](#elastic-configuration-service)
- [Reindex-Service](#reindex-service)
- [Current-User-Service](#current-user-service)


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

You may optionally pass in your own http client, replaces the nil above, to be used as the go-bamboo http client.

## General-Information ##

Returns general info about the API.

-- TODO --
example

## Server-Information ##

Returns general information about the Bamboo server.
Implemented via the InfoService.

-- TODO --
example

## Project-Service ##

Returns information about configured projects and information about individual projects.
Some features implemented.

-- TODO --
example

## Plan-Service ##

Returns information about configured plans and information about individual plans.
Some features implemented.

-- TODO --
example

## Result-Service ##

-- TODO --
implement
example

## Chart-Service ##

-- TODO --
implement
example

## Queue-Service ##

-- TODO --
implement
example

## Export-Service ##

-- TODO --
implement
example

## Clone-Service ##

-- TODO --
implement
example

## Dependency-Service ##

-- TODO --
implement
example

## Elastic-Configuration-Service ##

-- TODO --
implement
example

## Reindex-Service ##

-- TODO --
implement
example

## Current-User-Service ##

-- TODO --
implement
example