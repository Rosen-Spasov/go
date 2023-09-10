package main

import (
	"fmt"
	"net/http"

	"github.com/Rosen-Spasov/go/webServer/controllers"
	"github.com/Rosen-Spasov/go/webServer/models"
)

func main() {
	fmt.Println("Hello from webServer")
	models.GetUsers()
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
