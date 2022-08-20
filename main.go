package main

import (
	"MyGelSpider/controllers"
	"fmt"
	"time"
)

func main() {
	var choice string

	for {

		fmt.Println("欢迎来到 Gelbooru 下载器。 下载整页内的所有图片请按1；下载特定ID的图片请按2； 退出请按3。")
		fmt.Print("请输入你的选择:")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			fmt.Println("*您选择了下载页面内所有图片*")
			controllers.ControllerForCase1()
		case "2":
			fmt.Println("这是下载特定ID的图片的具体业务")
			controllers.ControllerForCase2()
		case "3":
			fmt.Println("程序退出")
			return
		default:
			fmt.Println("你输入的有误，请重新输入。")
		}

		time.Sleep(time.Second / 2)
	}
}
