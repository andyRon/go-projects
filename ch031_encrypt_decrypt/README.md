文件加密和解密
---

ch031_encrypt_decrypt

// TODO


### 配置encrypt.go文件的编译运行参数

在“Go tool arguments”文本框中输入“-ldflags "-w　-s" -o encrypt.exe”，表示去除可执行文件的符号表，并编译输出名称为“encrypt.exe”的可执行程序，
在“Program arguments”文本框中输入“10088　a.zip”，表示可执行程序的两个命令行参数为“10088”和“a.zip”。


### ref

[Go语言从入门到项目实战](https://book.douban.com/subject/36049170/)