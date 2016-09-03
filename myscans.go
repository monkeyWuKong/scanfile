package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	//"strings"
)

func isDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

//对字符串进行SHA1哈希
func GetFileSHA1(path string) (string, error) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		return "", err
	}

	hash := sha1.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func main() {
	var jumpstr string
	var rootpath string
	var a bool
	var b error
	var hashstr string

	RstName := "Resul.txt"
	DstFile, err := os.Create(RstName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer DstFile.Close()

	//get root path
	fmt.Println("Enter the rootpath and string you want to ignore(eg  D:/gowork \\.exe ) :")
	fmt.Scanf("%s %s", &rootpath, &jumpstr)
	if isDirExists(rootpath) {
		fmt.Println(rootpath)
	} else {
		fmt.Println("rootpath is not exist.")
		return
	}

	//ignore some file or path
	//fmt.Println("Enter string you want to ignore:")
	//fmt.Scanf("%s", &jumpstr)
	fmt.Println(jumpstr)
	//string judge
	_, err = regexp.Compile(jumpstr)
	if err != nil {
		fmt.Println("string is illegal")
		return
	}

	filepath.Walk(rootpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}

			if f.IsDir() {
				return nil
			}

			//jump as you need
			a, b = regexp.MatchString(jumpstr, path)
			if a && nil == b {
				//fmt.Println("jump -->  ", path)
				return nil
			}

			//get hash
			hashstr, b = GetFileSHA1(path)
			//get file size
			FileSize := strconv.FormatInt(f.Size(), 10)
			result := f.Name() + "," + hashstr + "," + FileSize + "\r\n"
			//write line
			println(path)
			DstFile.WriteString(result)
			return nil
		})
}
