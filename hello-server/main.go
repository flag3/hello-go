package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "strconv"
)

type jsonData struct {
  Number int
  String string 
  Bool   bool 
}

var value int = 0

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/hello", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World.\n")
  })
  e.GET("/json", jsonHandler)
  e.POST("/post", postHandler)
  e.GET("/hello/:name", helloHandler)
  e.GET("/ping", pingHandler)
  e.GET("/incremental", incrementalHandler)
  e.GET("/fizzbuzz", fizzbuzzHandler)
  e.GET("/json", jsonHandler)
  e.Start(":8080")
}

func jsonHandler(c echo.Context) error {
  res := jsonData{
    Number: 10,
    String: "hoge", 
    Bool: false,
  }
  return c.JSON(http.StatusOK, &res)
}

func postHandler(c echo.Context) error {
  var data jsonData

  if err := c.Bind(&data); err != nil {
    return c.JSON(http.StatusBadRequest, data)
  }
  return c.JSON(http.StatusOK, data)
}

func helloHandler(c echo.Context) error {
  name := c.Param("name")
  return c.String(http.StatusOK, "Hello, "+name+".\n")
}

func pingHandler(c echo.Context) error {
  return c.String(http.StatusOK, "pong\n")
}

func incrementalHandler(c echo.Context) error {
  value++
  return c.String(http.StatusOK, strconv.Itoa(value))
}

func fizzbuzzHandler(c echo.Context) error {
  str := "1"
  if count := c.QueryParam("count"); count == "" {
    return c.String(http.StatusOK, "30")
  } else {
    n, err := strconv.Atoi(count)
    if err != nil {
      return c.String(http.StatusBadRequest, "Bad Request")
    } else {
      for i := 2; i <= n; i++ {
        if i % 15 == 0 {
          str += "\nFizzBuzz"
        } else if i % 3 == 0 {
          str += "\nFizz"
        } else if i % 5 == 0 {
          str += "\nBuzz"
        } else {
          str += "\n" + strconv.Itoa(i)
        }
      }
      return c.String(http.StatusOK, str)
    }
  }
}
