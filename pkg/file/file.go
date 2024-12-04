package file

import (
	"fmt"
	"os"
)
type FileConfig struct{
	Filename string
}
func(f *FileConfig)ReadFile()[]byte{
	data,err := os.ReadFile(f.Filename)
	if err != nil{
		fmt.Println(err)
	}
	return data
}
func(f *FileConfig)WriteFile(){
	_ = f.ReadFile()
}

func NewFileConfig(f_name string)*FileConfig{
	return &FileConfig{
		Filename: f_name,
	}
}