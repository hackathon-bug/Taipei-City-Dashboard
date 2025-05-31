// Package app initiates the Gin server and connects to the postgreSQL database
/*
Developed By Taipei Urban Intelligence Center 2023-2024

// Lead Developer:  Igor Ho (Full Stack Engineer)
// Systems & Auth: Ann Shih (Systems Engineer)
// Data Pipelines:  Iima Yu (Data Scientist)
// Design and UX: Roy Lin (Prev. Consultant), Chu Chen (Researcher)
// Testing: Jack Huang (Data Scientist), Ian Huang (Data Analysis Intern)
*/
package app

import (
	"TaipeiCityDashboardBE/app/cache"
	"TaipeiCityDashboardBE/app/initial"
	"TaipeiCityDashboardBE/app/middleware"
	"TaipeiCityDashboardBE/app/models"
	"TaipeiCityDashboardBE/app/routes"
	"TaipeiCityDashboardBE/global"
	"TaipeiCityDashboardBE/logs"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// app.go is the main entry point for this application.
// initiates configures the postgreSQL, Redis, Gin router and starts the server.

// StartApplication initiates the main backend application, including the Gin router, postgreSQL, and Redis.
func StartApplication() {
	// 1. Connect to postgreSQL and Redis
	models.ConnectToDatabases("MANAGER", "DASHBOARD")
	cache.ConnectToRedis()

	// 2. Initiate default Gin router with logger and recovery middleware
	routes.Router = gin.Default()

	// 2.5 init redis
	users, _, _, err := models.GetAllUsers(0,0,"","","","")
	if err != nil{
		panic(err)
	}
	const blackListSet = "user:blacklist"
    const whiteListSet = "user:whitelist"
	for _, user := range users {
        userIDStr := fmt.Sprintf("%d", user.ID)

		if err := cache.Redis.Del(blackListSet, whiteListSet).Err(); err != nil {
			panic("delete black list and white list from redis error:" + err.Error())
		}
        if user.IsBlacked != nil && *user.IsBlacked {
            if err := cache.Redis.SAdd(blackListSet, userIDStr).Err(); err != nil {
                panic("add black list to redis error:" + err.Error())
            }
        }
        if user.IsWhitelist != nil && *user.IsWhitelist {
            if err := cache.Redis.SAdd(whiteListSet, userIDStr).Err(); err != nil {
                panic("add white list to redis error:" + err.Error())
            }
        }
    }


	// 3. Add common middlewares that need to run on all routes
	routes.Router.Use(middleware.AddCommonHeaders)
	// routes.Router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://tuic.gov.taipei"},
	// 	AllowMethods:     []string{"GET"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	// 4. Configure routes and routing groups (./router.go)
	routes.ConfigureRoutes()

	// 5. Configure http server
	addr := global.GinAddr

	err = endless.ListenAndServe(addr, routes.Router)
	if err != nil {
		logs.Warn(err)
	}
	logs.FInfo("Server on %v stopped", addr)

	// If the server stops, close the database connections
	models.CloseConnects("MANAGER", "DASHBOARD")
	cache.CloseConnect()
}

func MigrateManagerSchema() {
	models.ConnectToDatabases("MANAGER")
	models.MigrateManagerSchema()
	initial.InitDashboardManager()
	models.CloseConnects("MANAGER")
}

func InsertDashbaordSampleData() {
	models.ConnectToDatabases("DASHBOARD")
	initial.InitSampleCityData()
	models.CloseConnects("DASHBOARD")
}
