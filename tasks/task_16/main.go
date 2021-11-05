//отже задача,
//в тебе на гугл диску має бути якась exel таблиця
//треба написати скіпр, який іде на гугл диск (конфіги для конекшину читає з файла в yaml форматі)
//і дістає всі exel і пропонує користувачу вибрати,
//користувач вибирає файл і скріпт у табличці в консолі виводить вміст exel файлу
package main

import (
	"context"
	"log"

	"studing/tasks/task_16/printer"
	"studing/tasks/task_16/provider/google_provider"
	"studing/tasks/task_16/service"

	"studing/tasks/task_16/scanner"
)

func main() {
	ctx := context.Background()

	googleProvider, err := google_provider.New(ctx, "./tasks/task_16/provider/google_provider/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	p := printer.NewExcelPrinter()
	s := scanner.NewTerminalScanner()
	myService := service.New(s, googleProvider.GoogleSheetsProvider, p)

	if err := myService.OpenExcelFile(ctx, googleProvider.GoogleDriveProvider); err != nil {
		log.Fatal(err)
	}
}
