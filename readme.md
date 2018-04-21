# go-tumblr-crawler
Easily download all the photos/videos from tumblr blogs. 下载指定的 Tumblr 博客中的图片，视频。golang版本。

## 配置和运行
配置需要爬取的站点: `sites.json`
ps: 站点不需要 `.tumblr.com` 后缀
```
[
  {
    "site": "truenorthshow",
    "video": true,
    "photo": true
  },
  {
    "site": "photosbygerardo",
    "video": true,
    "photo": true
  }
]
```
配置代理 : `proxies.json`
```
{
  "http": "socks5://127.0.0.1:1080",
  "https": "socks5://127.0.0.1:1080"
}
```
然后保存文件,双击运行 `./tumblr.exe`(还没编译好,你可以自行编译).
mac 用户可以直接运行  `./tumblr`

96qbhy@gmail.com