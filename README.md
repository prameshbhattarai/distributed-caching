# Distributed Caching with Groupcache

## Description

Embedded distributed caching within multiple servers using the groupcache library.

## Setup

### Install Golang

To set up the development environment, ensure that you have Go installed. You can download and install Go from the [official Go website](https://golang.org/dl/).

## Deployment

### Starting the Application
 To start the application, use the following command:

```bash
make start
```
### Stopping the Application
 To stop the application and free up the used ports, use the following command:

```bash
make stop
```

### Clean the build files
 To clear all build files, use the following command:

```bash
make clean
```

## Testing  

### Fetching Data and Caching  
  You can test the caching mechanism using curl commands.

  First Request (Fetch from DB and Cache)
```bash
curl --location 'http://localhost:8081?key=apple'
```
This is the first request, so the data will be fetched from the database, store it in the cache, and then returned in the response.

  Subsequent Requests (Fetch from Cache)

```bash
curl --location 'http://localhost:8082?key=apple'
curl --location 'http://localhost:8083?key=apple'
curl --location 'http://localhost:8084?key=apple'
``` 
These requests will return the response from the cache, provided the cache has not expired.
