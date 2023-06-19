package main

import "glog"

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

	gLogger.Info("\n \033[36m Executing query of: %s", "SELECT * FROM USUARIOS WHERE ID = 2")

}
