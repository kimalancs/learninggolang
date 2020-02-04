package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议,所有的WWW文件都必须遵守这个标准。
// 设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。
// 浏览器和服务端之间通信的规则

// 后端返回的都是字节数据
// 浏览器根据HTML、CSS、JS的规则解析渲染后显示

// HTML称为超文本标记语言，是一种标识性的语言，定义网页内容的含义和格式 HyperText Markup Language
// 它包括一系列标签．通过这些标签可以将网络上的文档格式统一，使分散的Internet资源连接为一个逻辑整体
// HTML文本是由HTML命令组成的描述性文本，HTML命令可以说明文字，图形、动画、声音、表格、链接等
// 超文本是一种组织信息的方式，它通过超级链接方法将文本中的文字、图表与其他信息媒体相关联
// 这些相互关联的信息媒体可能在同一文本中，也可能是其他文件，或是地理位置相距遥远的某台计算机上的文件
// 这种组织信息方式将分布在不同位置的信息资源用随机方式进行连接，为人们查找，检索信息提供方便。

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)

	// str := `<h1 style="color:red">Hello</h1>`
	// w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/posts/Go", f1) // 根据访问的路径找对应的函数执行
	http.ListenAndServe("0.0.0.0:9090", nil) //起服务，监听，等待访问
}
