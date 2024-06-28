package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/turbovladimir/csv-parser/pkg/db"
	"log"
)

func main() {
	godotenv.Load(".env")

	var command string
	fmt.Println("Enter a command:")

	_, err := fmt.Scanln(&command)

	if err != nil {
		log.Fatal(err)
	}

	if command == "generate" {
		g := db.ModelGenerator{}
		g.Generate("./pkg/db/models")
	}

	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	//jakartaTime, _ := time.LoadLocation("Europe/Moscow")
	//scheduler := cron.New(cron.WithLocation(jakartaTime))
	//
	//// stop scheduler tepat sebelum fungsi berakhir
	//defer scheduler.Stop()
	//
	//// set task yang akan dijalankan scheduler
	//// gunakan crontab string untuk mengatur jadwal
	//scheduler.AddFunc("* * * * *", process)
	//
	//// start scheduler
	//go scheduler.Start()
	//
	//// trap SIGINT untuk trigger shutdown.
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//<-sig
}
