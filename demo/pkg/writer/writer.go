package writer

import (
	"io/ioutil"
	"log"
	"os"
)

func MakeDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Save(file []byte, fileName string)  error {
	err := ioutil.WriteFile(fileName, file, os.ModePerm)
	if err != nil {
		log.Fatal("write file error: ", err)
		return err
	}
	log.Println("save successfully")
	return nil
}