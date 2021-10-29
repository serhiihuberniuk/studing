//отже задача,
//в тебе на гугл диску має бути якась exel таблиця
//треба написати скіпр, який іде на гугл диск (конфіги для конекшину читає з файла в yaml форматі)
//і дістає всі exel і пропонує користувачу вибрати,
//користувач вибирає файл і скріпт у табличці в консолі виводить вміст exel файлу
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"studing/tasks/task_16/configs"
	"studing/tasks/task_16/parser"
	"studing/tasks/task_16/printer"
	"studing/tasks/task_16/provider"
	"studing/tasks/task_16/scanner"
)

func main() {
	ctx := context.Background()

	config, err := configs.NewConfig("./tasks/task_16/configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	googleDriveProvider, err := provider.NewGoogleDriveProvider(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	excelFiles, err := googleDriveProvider.GetExcelFilesNames(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(excelFiles) == 0 {
		log.Fatal("no excel files found")
	}

	for _, fileName := range excelFiles {
		fmt.Println(fileName)
	}

	fmt.Println("Enter name of file you want to read: ")
	fileName := scanner.ScanTerminal()

	tempFileName := config.DownloadPath + fileName
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = os.Remove(tempFileName); err != nil {
			log.Fatal(err)
		}
	}()

	err = googleDriveProvider.DownloadFile(ctx, fileName, tempFile)
	if err != nil {
		log.Fatal(err)
	}

	myParser := parser.NewExcelParser()
	myPrinter := printer.NewExcelPrinter(myParser)

	if err := myPrinter.PrintExcelFile(ctx, tempFileName); err != nil {
		log.Fatal(err)
	}
}
