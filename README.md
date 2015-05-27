# apistatus
implementation of https://github.com/Mashape/apistatus in go.

### Install

go get github.com/paulvollmer/apistatus

### Usage

    status := Apistatus{}
    data, err := status.Check("http://github.com")
    if err != nil {
      panic(err)
    }
    fmt.Println(data)

### TODO's
[ ] HAR

### MIT license

Copyright (c) 2015, Mashape (https://www.mashape.com/)
