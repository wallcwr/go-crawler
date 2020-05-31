package writer

import (
	"io/ioutil"
	"log"
	"os"
)

//创建目录
func MakeDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 保存文件
func Save(file []byte, fileName string)  error {
	err := ioutil.WriteFile(fileName, file, os.ModePerm)
	if err != nil {
		log.Fatal("write file error: ", err)
		return err
	}
	log.Println("save successfully")
	return nil
}