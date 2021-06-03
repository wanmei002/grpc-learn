package main

import (
    "encoding/json"
    "fmt"
)

func main(){
    str := "{\"LoadBalancingPolicy\": \"round_robin\"}"
    
    var m map[string]string
    err := json.Unmarshal([]byte(str), &m)
    if err != nil {
        fmt.Println("json err:", err)
        return
    }
    
    fmt.Printf("%+v", m)
}
