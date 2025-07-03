读取U盘数据并处理文件信息
----
ch033_usb

go处理外部存储设备（如U盘）的数据和普通目录一样，只需要找到对应挂载位置。
通常情况下，Mac OS 外部存储设备会自动挂载到 /Volumes 目录下；在Linux系统中，U盘通常挂载在/media或/mnt目录下；在Windows系统中，U盘则可能出现在D:、E:等盘符中。


### ref

https://www.oryoy.com/news/golang-shi-zhan-ru-he-gao-xiao-du-qu-u-pan-shu-ju-bing-chu-li-wen-jian-xin-xi.html