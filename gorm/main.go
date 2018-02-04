package main

import (
    "fmt"
    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
    start := time.Now()  
    
    db, err := gorm.Open("mysql", "root:PASSWORD@tcp(URL:3306)/DATABASE?charset=utf8&parseTime=True")
    if err != nil {
      panic("failed to connect database")
    }
    defer db.Close()
  
    //   db.AutoMigrate(&Product{})

    // Create
    //   db.Create(&Product{Code: "L1212", Price: 1000})
  
    var product []Product
    db.Find(&product)
   
    fmt.Println(product[0])
    
	elapsed := time.Since(start)
    
    elapsedTimeStr := elapsed.String()
    fmt.Printf("Query took %s", elapsed)

	return events.APIGatewayProxyResponse{Body: elapsedTimeStr, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
} 
