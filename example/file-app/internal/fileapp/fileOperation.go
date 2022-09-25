package fileapp

// package外から参照したいものの名前の先頭を大文字にするというのは，
// 関数名だけでなく，変数名や構造体名にも適用される．
import (
	"fmt"
	"os"
)

// write file
// 存在しない場合は、新規作成。
// 存在する場合はからにして書き込む
func WriteFile(filename string, str string) {
	// f, err := os.Create("write.txt")
	// 書き込み
	f, err := os.Create("write.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := []byte(str) // <= byte制限はないよね?
	// 16バイト制限はどこからきた? 解決
	// https: //github.com/Runacy/distributed-service-go/issues/1
	if _, err := f.Write(data); err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}

}
