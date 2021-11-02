//отже задача,
//в тебе на гугл диску має бути якась exel таблиця
//треба написати скіпр, який іде на гугл диск (конфіги для конекшину читає з файла в yaml форматі)
//і дістає всі exel і пропонує користувачу вибрати,
//користувач вибирає файл і скріпт у табличці в консолі виводить вміст exel файлу
package main

import (
	"context"
	"log"

	"studing/tasks/task_16/configs"
	"studing/tasks/task_16/printer"
	"studing/tasks/task_16/provider/google_provider"
	"studing/tasks/task_16/service"

	"studing/tasks/task_16/scanner"
)

func main() {
	ctx := context.Background()

	config, err := configs.NewConfig("./tasks/task_16/configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	googleProvider, err := google_provider.NewGoogleProvider(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	driveProvider, err := google_provider.NewGoogleDriveProvider(ctx, googleProvider)
	if err != nil {
		log.Fatal(err)
	}

	sheetProvider, err := google_provider.NewGoogleSheetsProvider(ctx, googleProvider)
	if err != nil {
		log.Fatal(err)
	}

	p := printer.NewExcelPrinter(sheetProvider)
	s := scanner.NewTerminalScanner()
	myService := service.New(s, p)

	if err := myService.OpenExcelFile(ctx, driveProvider); err != nil {
		log.Fatal(err)
	}
}
