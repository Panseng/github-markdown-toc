# github-markdown-toc
github markdown toc in go, readme toc in go, markdown &amp; readme 目录生成器（in go）

来源：[github-markdown-toc.go](https://github.com/ekalinin/github-markdown-toc.go)

输出win10本地readme目录命令 \
`go run main.go E:\code\golang\leetcode_go\notes\base.md`

输出的目录，默认去除链接，如果需要链接，则需要通过`--keep-link=true`来配置 如：\
`go run main.go --keep-link=true E:\code\golang\leetcode_go\notes\base.md`

