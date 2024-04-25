package main

import (
	"errors"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

type Body struct {
	Name string `json:"name"`
}

func add(numero1 int, numero2 int) (int, error) {
	result := numero1 + numero2
	return result, nil
}

func division(numero1 int, numero2 int) (int, error) {
	if numero2 == 0 {
		return 0, errors.New("no se puede dividir por 0")
	}
	result := numero1 / numero2
	return result, nil
}

func main() {
	// engine := gin.New()

	// engine.POST("/test", func(context *gin.Context) {
	// 	body := Body{}

	// 	if err := context.BindJSON(&body); err != nil {
	// 		context.AbortWithError(http.StatusBadRequest, err)
	// 		return
	// 	}

	// 	fmt.Println(body)
	// 	context.JSON(http.StatusAccepted, &body)
	// })

	// engine.Run(":3000")
}
