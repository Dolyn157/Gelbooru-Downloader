package controllers

import (
	"MyGelSpider/model"
	"MyGelSpider/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ControllerForCase1() {
	var tags string
	var startPid int
	var maxpage int

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入要下载图片的tags：")
	input, _ := inputReader.ReadString('\n')
	tags = strings.TrimSpace(input)

	var isCorrect bool = false
	var typed1 string

	for isCorrect == false {
		fmt.Print("请输入下载的起始页码：")
		fmt.Scanln(&typed1)
		if utils.OnlyContainsDigits(typed1) {
			startPid, _ = strconv.Atoi(typed1)
			isCorrect = true
		} else {
			fmt.Println("你输入的内容包含数字以外的字符，请重新输入。")
		}
	}

	isCorrect = false
	var typed2 string

	for isCorrect == false {
		fmt.Print("从起始页码开始往后要下载几页：")
		fmt.Scanln(&typed2)
		if utils.OnlyContainsDigits(typed2) {
			i, _ := strconv.Atoi(typed2)
			if i > 0 {
				maxpage = i
				isCorrect = true
			}
		} else {
			fmt.Println("你输入的内容包含数字以外的字符，且必须大于零，请重新输入。")
		}
	}

	dirDst := "./Pages"
	os.Mkdir(dirDst, os.ModePerm)

	fmt.Println("开始下载：")

	for i := startPid; i < startPid+maxpage; i++ {
		ApiURI := fmt.Sprint("https://gelbooru.com/index.php?page=dapi&s=post&q=index&limit=21", "&tags=", tags, "&pid=", i, "&json=1")
		//判断一下网络链接通不通
		fmt.Println("当前正在处理的页面：", ApiURI)
		var apiData models.ApiData
		apiData.GetFileAndDownload(ApiURI, dirDst)

	}
	fmt.Print("本轮下载结束。", tags, startPid, maxpage, "\n\n")

}

func ControllerForCase2() {
	var picID int

	var isCorrect bool = false
	var typed1 string

	for isCorrect == false {
		fmt.Print("请输入要下载的图片ID：")
		fmt.Scanln(&typed1)
		if utils.OnlyContainsDigits(typed1) {
			picID, _ = strconv.Atoi(typed1)
			isCorrect = true
		} else {
			fmt.Println("你输入的内容包含数字以外的字符，请重新输入。")
		}
	}

	ApiURI := fmt.Sprint("https://gelbooru.com/index.php?page=dapi&s=post&q=index", "&id=", picID, "&json=1")
	dirDst := "./SinglePictures"
	os.Mkdir(dirDst, os.ModePerm)

	fmt.Println("开始下载：", ApiURI)

	var apiData models.ApiData
	apiData.GetFileAndDownload(ApiURI, dirDst)

}
