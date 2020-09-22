package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	var s SitemapIndex
	//路径访问往往从根目录开始而不是当前目录
	contents, _ := ioutil.ReadFile("assets/sitemap.xml")
	xml.Unmarshal(contents, &s)
	fmt.Println(s.Locations)
	fmt.Printf("I %s trying \n", "am")
	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}

}
