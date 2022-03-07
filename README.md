# github-markdown-toc
github markdown toc in go, readme toc in go, markdown &amp; readme 目录生成器（in go）

项目链接：[Github markdown toc](https://github.com/Panseng/github-markdown-toc)

## 简介
项目来源于：[ekalinin/github-markdown-toc.go](https://github.com/ekalinin/github-markdown-toc.go)

修改内容：
- [x] 增加了对输入windows系统绝对地址的支持
- [x] 增加了默认去除目录跳转链接，使得目录导航功能更加清晰：导航到当前文档的标题处
    - 通过 `--keep-link=true`配置，可以保留目录跳转链接

Todo：
- [ ] 输入markdown链接生成目录

## 使用
方式1：直接通过`run`命令运行源文件，如：
```go run main.go E:\code\golang\leetcode_go\notes\base.md```

方式2：
1. `build`编译为平台可运行文件；
2. 将可运行文件放到`$GOPATH\bin`文件夹；
3. 须将`$GOPATH\bin`文件夹添加到系统环境变量的`Path`中；
4. 通过命令运行`github-markdown-toc E:\code\golang\leetcode_go\notes\base.md`
