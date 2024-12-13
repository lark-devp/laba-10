package main

import (
	"flag"
	"log"
	"net/http"
	"web-10/internal/query/api"
	"web-10/internal/query/config"
	"web-10/internal/query/provider"
	"web-10/internal/query/usecase"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "../../configs/query_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase("", prv)
	srv := api.NewServer(cfg.IP, cfg.Port, use)

	log.Printf("Сервер запущен на %s\n", srv.Address)
	log.Fatal(http.ListenAndServe(srv.Address, srv.Router))
}
