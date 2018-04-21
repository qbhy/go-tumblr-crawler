package main

import (
	"github.com/qbhy/go-utils"
	"path"
	"fmt"
	"strings"
	"tumblr-crawler/downloader"
	"runtime"
	io "io/ioutil"
)

func ParseSites(filename string) []string {
	data, _ := io.ReadFile(filename)
	wrapSites := string(data)
	wrapSites = strings.Replace(wrapSites, "\t", ",", -1)
	wrapSites = strings.Replace(wrapSites, "\n", ",", -1)
	wrapSites = strings.Replace(wrapSites, "\r", ",", -1)
	wrapSites = strings.Replace(wrapSites, " ", ",", -1)

	return strings.Split(wrapSites, ",")
}

func main() {

	config := downloader.NewConfig()
	currentPath := utils.CurrentPath()
	sites := []string{}
	var proxies downloader.ProxyConfig

	// 获取代理配置
	proxyPath := path.Join(currentPath, "proxies.json")
	if exists, _ := utils.PathExists(proxyPath); exists {
		proxies = downloader.ProxyConfig{}
		config.Load(proxyPath, &proxies)
		fmt.Println(proxies)
	}

	// 获取站点配置
	sitesPath := path.Join(currentPath, "sites.txt")
	if exists, _ := utils.PathExists(sitesPath); exists {
		sites = ParseSites(sitesPath)
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
