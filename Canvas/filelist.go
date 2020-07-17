package Canvas

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ostype = os.Getenv("GOOS") // 获取系统类型
)

var listfile []string //获取文件列表

func Listfunc(path string, f os.FileInfo, err error) error {
	//用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
	ok := strings.HasSuffix(path, ".jsc")
	if ok {
		listfile = append(listfile, path) //将目录push到listfile []string中
	}
	//fmt.Println(ostype) // print ostype
	//fmt.Println(path) //list the file
	return nil
}

func GetFileList(path string) []string {
	//var strRet string
	err := filepath.Walk(path, Listfunc) //

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return listfile
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
