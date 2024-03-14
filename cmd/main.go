package main

import (
	"github.com/yourscarent/cars_manager/internal/app"
	"github.com/yourscarent/cars_manager/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
