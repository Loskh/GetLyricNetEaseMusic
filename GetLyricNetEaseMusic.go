package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/widuu/gojson"
)

func GetLrc(id string) (result string) {
	u, _ := url.Parse("http://music.163.com/api/song/lyric")
	q := u.Query()
	q.Set("id", id)
	q.Set("lv", "-1")
	q.Set("kv", "-1")
	q.Set("tv", "-1")
	u.RawQuery = q.Encode()
	res, err1 := http.Get(u.String())
	if err1 != nil {
		log.Fatal(err1)
	}
	result0, err2 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err2 != nil {
		log.Fatal(err2)
	}
	result = string(result0)
	return
}
func GetName(id string) (name string) {
	u, _ := url.Parse("http://music.163.com/api/song/detail/")
	q := u.Query()
	q.Set("id", id)
	q.Set("ids", "["+id+"]")
	u.RawQuery = q.Encode()
	res, err1 := http.Get(u.String())
	if err1 != nil {
		log.Fatal(err1)
	}
	result0, err2 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err2 != nil {
		log.Fatal(err2)
	}
	result := string(result0)
	name = gojson.Json(result).Get("songs").Getkey("name", 1).Tostring()
	return
}
func SaveLrc(lrc, name string) {
	fileName := name + ".lrc"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(lrc)
}
func main() {
	var url string
	r, _ := regexp.Compile("[1-9][0-9]{3,}")
	fmt.Println("输入歌曲地址：")
	fmt.Scanln(&url)
	id := r.FindString(url)
	name := GetName(id)
	fmt.Println("音乐名为：" + name)
	result := GetLrc(id)
	//	fmt.Println(name)
	//	fmt.Println(result)
	lrc1 := gojson.Json(result).Get("lrc").Get("lyric").Tostring()
	SaveLrc(lrc1, name)
	lrc2 := gojson.Json(result).Get("tlyric").Get("lyric").Tostring()
	if len(lrc2) != 0 {
		SaveLrc(lrc2, name+"(CHS)")
	}
	fmt.Println("歌词下载成功")
}
