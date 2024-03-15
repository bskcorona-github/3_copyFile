package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ファイルをコピーする関数
func copyFile(sourcePath, destinationPath string) error {
	// ソースファイルを開く
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 宛先ファイルを作成またはオープン
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// ソースファイルの内容を宛先ファイルにコピー
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// コマンドライン引数を取得
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: go run main.go <source_file> <destination_file>")
		return
	}

	sourcePath := args[1]
	destinationPath := args[2]

	// ソースファイルが存在するか確認
	_, err := os.Stat(sourcePath)
	if err != nil {
		fmt.Println("ソースファイルが見つかりません:", err)
		return
	}

	// ディレクトリがなければ作成
	destinationDir := filepath.Dir(destinationPath)
	err = os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		fmt.Println("宛先ディレクトリの作成に失敗しました:", err)
		return
	}

	// ファイルをコピー
	err = copyFile(sourcePath, destinationPath)
	if err != nil {
		fmt.Println("ファイルのコピーに失敗しました:", err)
		return
	}
	fmt.Println("ファイルをコピーしました。")
}
