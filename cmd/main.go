package main

import "github.com/Vygos/glog"

type Person struct {
	Name string
	Age  int
}

func main() {
	gLogger := glog.NewGLogger("Golang")

	person := Person{"John Doe", 30}

	gLogger.Info("Starting application on port %s ", ":8080")
	gLogger.Info("Receiving event of payload: %+v", person)
	gLogger.Warn("Connection has to much opened connections total: %d", 10)

	gLogger.Info("Executing query of: %s", glog.Cyan, "SELECT * FROM USUARIOS WHERE ID = 2")
	gLogger.Warn("Connection has to much opened connections total: %d", 10)

}
