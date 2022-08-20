package models

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

type WriteCounter struct {
	Compeleted uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Compeleted += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {

	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %d B complete", wc.Compeleted)

}

//fileUrl := "https://img3.gelbooru.com//images/20/da/20da0a88bb6f5c2f87cba0c1960be71d.gif"

func DownloadFile(DirPath string, fileName string, url string) error {

	filePath := DirPath + "/" + fileName
	out, err := os.Create(filePath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	var Size int64
	if resp.Status == "200 OK" {
		contentLength := path.Base(resp.Header.Get("Content-Length"))
		if contentLength != "" {
			parse, err := strconv.ParseInt(contentLength, 10, 64)
			if err != nil {
				return err
			}
			Size = parse
		}
		fmt.Printf("FileName = %s \nSize = %d B\n", fileName, Size)
	} else {
		fmt.Println("连接图床网站过程中出现问题, 错误信息：", err)
		return err
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	out.Close()

	fmt.Print("\n")
	err = os.Rename(filePath+".tmp", filePath)
	if err != nil {
		return err
	}

	return nil
}
