package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/defrell01/goodcom_client/internal/comments"
	"github.com/defrell01/goodcom_client/internal/config"
	"github.com/defrell01/goodcom_client/internal/files"
)

type FileComments struct {
	FilePath string
	Comments []string
	HasError bool
	Error    error
}

func main() {
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	filesList, err := files.ScanDirectory(cfg.Directory, cfg.Extensions)
	if err != nil {
		log.Fatalf("Ошибка сканирования директории: %v", err)
	}

	fmt.Println("Найденые файлы:")
	for _, file := range filesList {
		fmt.Println(file)
	}

	var wg sync.WaitGroup

	results := make(chan FileComments, len(filesList))

	for _, filePath := range filesList {
		wg.Add(1)

		go func(filePath string) {
			defer wg.Done()

			ext := filepath.Ext(filePath)
			commentsList, err := comments.ExtractCommentsFromFile(filePath, ext)

			if err != nil {
				results <- FileComments{FilePath: filePath, HasError: true, Error: err}
			} else {
				results <- FileComments{FilePath: filePath, Comments: commentsList}
			}
		}(filePath)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Извлеченные комментарии:")
	for result := range results {
		if result.HasError {
			log.Printf("Ошибка извлечения комментариев из файла %s: %v", result.FilePath, result.Error)
			continue
		}

		if len(result.Comments) > 0 {
			fmt.Printf("Комментарии из файла %s:\n", result.FilePath)
			for _, comment := range result.Comments {
				fmt.Println(comment)
			}
			fmt.Println()
		} else {
			fmt.Printf("Нет комментариев в файле %s\n", result.FilePath)
		}
	}

}
