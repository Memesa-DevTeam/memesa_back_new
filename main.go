package main

import (
	"go.uber.org/fx"
	"memesa_go_backend/app"
)

func main() {
	fx.New(app.MemesaServices()).Run()
}
