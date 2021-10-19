//написати програму, яка питає шлях до папки,
//потім отримавши його - сканує всю папку з вложеностями і виводить дублікати файлів
//(давай припустимо так, що дублікатом являються файли з однаковими іменами)
package main

import (
	"context"
	"fmt"
	"log"

	"studing/tasks/task_15/checker"
	"studing/tasks/task_15/finder"
	"studing/tasks/task_15/scanner"
)

func main() {
	ctx := context.Background()
	ts := scanner.NewTerminalScanner()
	ds := scanner.NewDirectoryScanner()
	f := finder.NewFinder()
	dc := checker.NewDirectoryChecker(ds, f)

	path, err := ts.Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}

	duplicates, err := dc.CheckDirectoryForDuplicates(ctx, path)
	if err != nil {
		log.Fatal(err)
	}

	if len(duplicates) != 0 {
		fmt.Println(duplicates)
	} else {
		fmt.Println("no duplicates occurred")
	}
}
