package main

import (
	"gin-orm-prac/database"
	"gin-orm-prac/routers"
)

func main() {
	database.InitModel();
	
	router := routers.InitRouter();

	router.Run();
}