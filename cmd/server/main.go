package main

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"hackaton/internal/app/point_record"
	"hackaton/internal/app/user"
	"hackaton/internal/config"
	"hackaton/internal/infra"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inicializando app")
	LoadAPP(cfg)
}

func LoadAPP(cfg *config.Config) {
	db, err := infra.GetDBConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	pointRecordRepository := point_record.NewPointRecordRepository(db)
	userRepository := user.NewUserRepository(db)
	sender := point_record.NewPointRecordSender(cfg)
	registerPoint := point_record.NewRegisterPointUseCase(pointRecordRepository, sender, userRepository)
	pointRecordHandler := point_record.NewHttpHandler(registerPoint)
	userUsecase := user.NewUserUseCase(userRepository, cfg)
	userHandler := user.NewHttpHandler(userUsecase)
	router := bunrouter.New(bunrouter.Use(reqlog.NewMiddleware()))
	router.WithGroup("/api/v1", func(apiV1Routes *bunrouter.Group) {
		apiV1Routes.POST("/register-point", pointRecordHandler.RegisterPoint)
		apiV1Routes.GET("/daily-report", pointRecordHandler.GetRegistersDay)
		apiV1Routes.GET("/monthly-report", pointRecordHandler.GetMonthlyReport)
		apiV1Routes.POST("/login", userHandler.Login)
	})
	log.Fatalf(http.ListenAndServe(":8001", router).Error())
}
