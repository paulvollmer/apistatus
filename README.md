# apistatus [![Build Status](https://travis-ci.org/paulvollmer/apistatus.svg?branch=master)](https://travis-ci.org/paulvollmer/apistatus) ![License](https://img.shields.io/npm/l/apistatus.svg)
implementation of https://github.com/Mashape/apistatus in go.

### Install

go get github.com/paulvollmer/apistatus

### Usage

    ...
    status := Apistatus{}
    statusCode, err := status.Check("http://github.com")
    if err != nil {
      panic(err)
    }
    fmt.Println(statusCode)
    ...

### TODO's
[ ] HAR

### MIT license

Copyright (c) 2015, Mashape (https://www.mashape.com/)
