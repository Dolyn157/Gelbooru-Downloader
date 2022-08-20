package models

import (
	"MyGelSpider/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiData struct {
	Post       []post
	Attributes attributes `json:"@attributes"`
}

type attributes struct {
	Count int `json:"count"`
}

type post struct {
	FileUrl string `json:"file_url"`
	Source  string `json:"Source"`
}

func (a *ApiData) GetFileAndDownload(ApiURI string, DirPath string) {

	//加载配置文件，拼接URL字符串
	resp, err := http.Get(ApiURI)
	if err != nil {
		fmt.Println("连接图床网站过程中出现问题, 错误信息：", err)
		return
	}

	// Read the Response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Unmarshall

	json.Unmarshal(data, &a)

	if a.Attributes.Count == 0 {
		fmt.Println("在您指定的 tags 或页面当中找不到图片")
		return
	}

	//遍历当前页面中的每一个图片

	for i, _ := range a.Post {

		fmt.Printf("正在处理第 %d 个文件 File URL = %s", i+1, a.Post[i].FileUrl)
		fileName := utils.GetNameFromURI(a.Post[i].FileUrl)

		err := DownloadFile(DirPath, fileName, a.Post[i].FileUrl)
		if err != nil {
			fmt.Printf("下载过程发送错误，错误信息： %v \n\n", err)
		}
		fmt.Print("下载完成\n\n")
	}

}
