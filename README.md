# Concurrent Money Transfer System

## Overview

This is a concurrent money transfer system built in Go to handle money tarnsfers between users with built in protections.

## Features

- Creation of Users/accounts with starting balance.
- Make Transfer to other users
- Atomic Money Tranfers
- Prevention of overdrafts
- Concurrently safe account operations

## Pre-requisites

Go 1.20 or higher

## Setup and Installation

1. Clone the repository
2. Run `go mod tidy` to download all dependencies
3. Run the application with `go run cmd/main.go`

## API Endpoints

1. Create Account

URL: `/account/create`
Method: `POST`
Request Body
```
{
  "user_id": "kowshik",
  "balance": 25.50
}
```

Response Body
```
{"status":"account created","user_id":"kora"}
```

2. Get Balance

URL: `/account/balance?user_id=kowshik`
Method: `GET`

Response Body
```
{"balance":25.50}
```

3. Transfer

URL: `/transfer`
Method: `POST`

Request Body
```
{
    "from_user_id" : "kowshik",
    "to_user_id" : "ram",
    "amount": 15.45
}
```

Response Body
```
{"status":"transfer successful"}
```
