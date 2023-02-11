package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "strconv"
  "encoding/json"
  "io/ioutil"
)

type jsonData struct {
  Number int    `json:"number,omitempty"`
  String string `json:"string,omitempty"`
  Bool   bool   `json:"bool,omitempty"`
}

type applicationJsonData struct {
  Right int `json:"right"`
  Left  int `json:"left"`
}

type applicationJsonDataAnswer struct {
  Answer int `json:"answer"`
}

type applicationJsonDataError struct {
  Error string  `json:"error"`
}

type Student struct {
  Number int `json:"student_number"`
  Name string `json:"name"`
}

type Class struct {
  Number int `json:"class_number"`
  Students []Student `json:"students"`
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
  e.POST("application/json", applicationJsonHandler)
  e.GET("/students/:class/:studentNumber", studentsHandler)
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
  if err := c.Bind(data); err != nil {
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

func applicationJsonHandler(c echo.Context) error {
  var data applicationJsonData

  if err := c.Bind(&data); err != nil {
    var dataError applicationJsonDataError
    dataError.Error = "Bad Request"
    return c.JSON(http.StatusBadRequest, dataError)
  }
  var dataAnswer applicationJsonDataAnswer
  dataAnswer.Answer = data.Right + data.Left
  return c.JSON(http.StatusOK, dataAnswer)
}

func studentsHandler(c echo.Context) error {
  bytes, err := ioutil.ReadFile("./students.json")
  if err != nil {
    return c.String(http.StatusBadRequest, "Bad Request")
  }

  var students []Class
  if err := json.Unmarshal(bytes, &students); err != nil {
    return c.String(http.StatusBadRequest, "Bad Request")
  }

  class := c.Param("class")
  studentNumber := c.Param("studentNumber")
  n, err1 := strconv.Atoi(class)
  m, err2 := strconv.Atoi(studentNumber)

  if (err1 != nil || err2 != nil) {
    return c.String(http.StatusBadRequest, "Bad Request")
  } else if (n > len(students)) || (m > len(students[n-1].Students)) {
    var dataError applicationJsonDataError
    dataError.Error = "Student Not Found"
    return c.JSON(http.StatusNotFound, dataError)
  } else {
    return c.JSON(http.StatusOK, students[n-1].Students[m-1])
  }
}
