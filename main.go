package main

import (
	"fmt"
	"log"

	"github.com/alkamalp/crm-golang/modules/actors"
	"github.com/alkamalp/crm-golang/modules/customers"
	"github.com/alkamalp/crm-golang/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql()

	//check connection
	checkdb, err := dbCrud.DB()
	if err != nil {
		log.Fatal(err)
	}

	//ping to database
	errconn := checkdb.Ping()
	if err != nil {
		log.Fatal(errconn)
	}

	fmt.Println("database connected..!")

	actorHandler := actors.NewRouter(dbCrud)
	actorHandler.Handle(router)

	customerHandler := customers.NewRouter(dbCrud)
	customerHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
