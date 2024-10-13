package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/defrell01/goodcom_client/internal/comments"
	"github.com/defrell01/goodcom_client/internal/config"
	"github.com/defrell01/goodcom_client/internal/files"
)

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

	fmt.Println("Извлеченные комментарии:")
	for _, filePath := range filesList {
		ext := filepath.Ext(filePath)
		commentsList, err := comments.ExtractCommentsFromFile(filePath, ext)
		if err != nil {
			log.Printf("Ошибка извлечения комментариев из файла %s: %v", filePath, err)
			continue
		}

		if len(commentsList) > 0 {
			fmt.Printf("Комментарии из файла %s:\n", filePath)
			for _, comment := range commentsList {
				fmt.Println(comment)
			}
			fmt.Println()
		} else {
			fmt.Printf("Нет комментариев в файле %s\n", filePath)
		}
	}
}
