package main

import (
	"github.com/qbhy/go-utils"
	"path"
	"fmt"
	"tumblr-crawler/downloader"
	"runtime"
	config2 "tumblr-crawler/config"
)

func main() {

	configLoader := config2.NewConfig()
	currentPath := utils.CurrentPath()
	sites := []config2.SiteConfig{}

	var proxies config2.ProxyConfig

	// 获取代理配置
	proxyPath := path.Join(currentPath, "proxies.json")
	if exists, _ := utils.PathExists(proxyPath); exists {
		proxies = config2.ProxyConfig{}
		configLoader.Load(proxyPath, &proxies)
		fmt.Println(proxies)
	}

	// 获取站点配置
	sitesPath := path.Join(currentPath, "sites.json")
	if exists, _ := utils.PathExists(sitesPath); exists {
		configLoader.Load(sitesPath, &sites)
	}

	// 设置最大协程数
	maxProcesses := runtime.NumCPU() //获取cpu个数
	runtime.GOMAXPROCS(maxProcesses) //限制同时运行的goroutines数量

	fmt.Println("CPU数量:", maxProcesses)

	// 下面这个for循环的意义就是利用信道的阻塞，一直从信道里取数据，直到取得跟并发数一样的个数的数据，则视为所有goroutines完成。

	if len(sites) > 0 {
		for _, site := range sites {
			siteInstance := downloader.NewSite(site, proxies)
			downloader.WaitGroupInstance.Add(1)
			siteInstance.StartDownload()
		}

		downloader.WaitGroupInstance.Wait()
	} else {
		fmt.Println("没有配置站点")
	}

	fmt.Println("WE DONE!!!")
}
