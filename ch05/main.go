package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main(){
    path := "C:\\Users\\v_zzzzhou\\Desktop\\test123\\aaaa_files"
    // aaaa_f
    path1 := "C:\\Users\\v_zzzzhou\\Desktop\\test123\\aaaa_f"
    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println("打开目录失败")
        return
    }
    f := "C:\\Users\\v_zzzzhou\\Desktop\\test123\\aaaa.html"
    hl, err := os.Open(f)
    if err != nil {
        fmt.Println("open", f, "failed")
        return
    }
    res, err := ioutil.ReadAll(hl)
    if err != nil {
        fmt.Println("read all failed; err", err)
        return
    }
    resA := string(res)
    for _, file := range files {
        // ./aaaa_files/0
        os.Rename(path + "\\" + file.Name(), path1+"\\"+file.Name()+".jpg")
        resA = strings.Replace(resA, "./aaaa_files/"+file.Name(), "./aaaa_f/"+file.Name()+".jpg", 1)
    }
    f1 := "C:\\Users\\v_zzzzhou\\Desktop\\test123\\bbbb.html"
    err = ioutil.WriteFile(f1, []byte(resA), os.ModePerm)
    
    if err != nil {
        fmt.Println("write file failed; err:", err)
    }
    
    fmt.Println("done")
}
