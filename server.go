package main

import (
    "fmt"
    "net/http"
    "strings"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Get the Alert!")

    fmt.Println("Raw Request:")
    fmt.Println(request)
    fmt.Println("")

    fmt.Println("Request解析")

    //HTTP方法
    fmt.Println("method: ", request.Method)
    // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
    fmt.Println("RequestURI: ", request.RequestURI)
    //URL类型,下方分别列出URL的各成员
    fmt.Println("URL_scheme: ", request.URL.Scheme)
    fmt.Println("URL_opaque: ", request.URL.Opaque)
    fmt.Println("URL_user: ", request.URL.User.String())
    fmt.Println("URL_host: ", request.URL.Host)
    fmt.Println("URL_path: ", request.URL.Path)
    fmt.Println("URL_RawQuery: ", request.URL.RawQuery)
    fmt.Println("URL_Fragment: ", request.URL.Fragment)

    //协议版本
    fmt.Println("proto: ", request.Proto)
    fmt.Println("protomajor: ", request.ProtoMajor)
    fmt.Println("protominor: ", request.ProtoMinor)

    //HTTP请求的头域
    for k, v := range request.Header {
        for _, vv := range v {
            fmt.Println("header key: " + k + "  value:" + vv)
        }
    }

    //判断是否multipart方式
    is_multipart := false
    for _, v := range request.Header["Content-Type"] {
        if strings.Index(v, "multipart/form-data") != -1 {
            is_multipart = true
        }
    }
	
    //解析body
    if is_multipart == true {
        request.ParseMultipartForm(128)
            fmt.Println("解析方式: ParseMultipartForm")
    } else {
        request.ParseForm()
        fmt.Println("解析方式: ParseForm")
    }
    //body内容长度
    fmt.Println("ContentLength: ", request.ContentLength)
	
    //是否在回复请求后关闭连接
    fmt.Println("Close: ", request.Close)
    //HOST
    fmt.Println("host: ", request.Host)
    //form
    fmt.Println("Form: ", request.Form)
    //postform
    fmt.Println("PostForm: ", request.PostForm)
    //MultipartForm
    fmt.Println("MultipartForm: ", request.MultipartForm)
    //解析MultipartForm内的文件
    files := request.MultipartForm.File
    for k, v := range files {
        fmt.Println(k)
        for _, vv := range v {
            fmt.Println(k + ":" + vv.Filename)
        }
    }

    //该请求的来源地址
    fmt.Println("RemoteAddr: ", request.RemoteAddr)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9999", nil)
}
