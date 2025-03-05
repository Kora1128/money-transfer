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


## Locking Strategy and Implementations

1. Grained Lock on Individual Accounts during Transfer
    - Instead of using a global lock on all transfers , we lock the two accounts involved in a transfer.
    - This improves concurrency cuz we can allow multiple transfers happening at the same time.

2. Implemented Ordered Locking to prevent deadlock
    - When two users send money to each other at same time, it may lock in different orders causing deadlock. So to avoid that we compare user_ids and do a ordered lock.

3. RWMutex during Concurrent Reads
    - `Get account balance` would be used in different places including exposing in a API and also during transfer to fetch the current balance. So implemented read-only lock so that multiple goroutines can read them concurrently.


## Sample Curls for testing

Account Create
```
curl --location 'http://localhost:8080/account/create' \
--header 'Content-Type: application/json' \
--data '{
    "user_id" : "kora",
    "balance": 50
}'
```

Get account Balance
```
curl --location 'http://localhost:8080/account/balance?user_id=kora' \
--data ''
```

Transfer money

```
curl --location 'http://localhost:8080/transfer' \
--header 'Content-Type: application/json' \
--data '{
    "from_user_id" : "kora",
    "to_user_id" : "kowshik",
    "amount": 50
}'
```