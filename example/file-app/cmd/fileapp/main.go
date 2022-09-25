package main

import "github.com/Runacy/distributed-service-go/example/file-app/internal/fileapp"

func main() {
	fileapp.WriteFile("write.txt", "goで書いたよ！")
}
