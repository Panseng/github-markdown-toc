package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"regexp"

	"github-markdown-toc/ghtoc"
)

var (
	pathsDesc  = "Local path or URL of the document to grab TOC. Read MD from stdin if not entered."
	paths      = kingpin.Arg("path", pathsDesc).Strings()
	serial     = kingpin.Flag("serial", "Grab TOCs in the serial mode").Bool()
	hideHeader = kingpin.Flag("hide-header", "Hide TOC header").Bool()
	hideFooter = kingpin.Flag("hide-footer", "Hide TOC footer").Default("false").Bool() // 屏蔽打印脚注
	startDepth = kingpin.Flag("start-depth", "Start including from this level. Defaults to 0 (include all levels)").Default("0").Int()
	depth      = kingpin.Flag("depth", "How many levels of headings to include. Defaults to 0 (all)").Default("0").Int()
	noEscape   = kingpin.Flag("no-escape", "Do not escape chars in sections").Bool()
	token      = kingpin.Flag("token", "GitHub personal token").String()
	indent     = kingpin.Flag("indent", "Indent space of generated list").Default("2").Int()
	debug      = kingpin.Flag("debug", "Show debug info").Bool()
	keepLink  = kingpin.Flag("keep-link", "TOC keep title's link. Defaults to not keep").Default("false").Bool()
)

// check if there was an error (and panic if it was)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Entry point
func main() {
	kingpin.Version(ghtoc.Version)
	kingpin.Parse()

	if *token == "" {
		*token = os.Getenv("GH_TOC_TOKEN")
	}

	// 在win10环境，*paths会包含参数，导致错误，需要排除参数
	tem := []string{}
	for _,p := range *paths{
		if found, _ := regexp.MatchString("(^[a-zA-Z]:)|(^http)", p); found{ // 只处理合法路径、url
			tem = append(tem, p)
		}
	}
	pathsCount := len(tem)

	// read file paths | urls from args
	absPathsInToc := pathsCount > 1
	ch := make(chan *ghtoc.GHToc, pathsCount)

	for _, p := range tem {
		ghdoc := ghtoc.NewGHDoc(p, absPathsInToc, *startDepth, *depth, !*noEscape, *token, *indent, *debug, *keepLink)
		if *serial {
			ch <- ghdoc.GetToc()
		} else {
			go func(path string) { ch <- ghdoc.GetToc() }(p)
		}
	}

	if !*hideHeader && pathsCount == 1 {
		fmt.Println()
		fmt.Println("Table of Contents")
		fmt.Println("=================")
		fmt.Println()
	}

	for i := 1; i <= pathsCount; i++ {
		toc := <-ch
		// #14, check if there's really TOC?
		if toc != nil {
			check(toc.Print(os.Stdout))
		}
	}

	// read md from stdin
	if pathsCount == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		check(err)

		file, err := ioutil.TempFile(os.TempDir(), "ghtoc")
		check(err)
		defer os.Remove(file.Name())

		check(ioutil.WriteFile(file.Name(), bytes, 0644))
		check(ghtoc.NewGHDoc(file.Name(), false, *startDepth, *depth, !*noEscape, *token, *indent, *debug, *keepLink).
			GetToc().
			Print(os.Stdout))
	}

}