package main

import (
	"time"

	"github.com/eisandbar/ytlive/app/adder"
	"github.com/eisandbar/ytlive/app/checker"
	"github.com/eisandbar/ytlive/app/getter"
	"github.com/eisandbar/ytlive/app/server"
	"github.com/eisandbar/ytlive/app/store"
	"github.com/eisandbar/ytlive/app/youtubeapi"
)

const (
	CONN_PORT = "3000"
)

func main() {
	// // Get new list of streams
	service := youtubeapi.NewService()

	pgStore := store.NewPGStore()

	checker.NewChecker(&pgStore, service, time.Minute*30)

	// adder.NewAdder(&pgStore, service, true, time.Minute*30, 8*time.Hour)
	adder.NewAdder(&pgStore, service, false, time.Minute*30, 8*time.Hour)

	getter.NewGetter(&pgStore)

	myServer := server.NewServer(&pgStore)
	server.StartServer(myServer, CONN_PORT)

}
