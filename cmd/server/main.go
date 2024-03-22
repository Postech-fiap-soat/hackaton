package main

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"hackaton/internal/app/point_record"
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
	repository := point_record.NewPointRecordRepository(db)
	registerPoint := point_record.NewRegisterPointUseCase(repository)
	httpHandler := point_record.NewHttpHandler(registerPoint)
	router := bunrouter.New(bunrouter.Use(reqlog.NewMiddleware()))
	router.WithGroup("/api/v1", func(apiV1Routes *bunrouter.Group) {
		apiV1Routes.POST("/register-point", httpHandler.RegisterPoint)
		apiV1Routes.GET("/daily-report", httpHandler.GetRegistersDay)
	})
	log.Fatalf(http.ListenAndServe(":8001", router).Error())
}
