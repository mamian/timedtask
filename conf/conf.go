package conf

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
)

type Config struct {
    LogPath string
    Tasks []TimeTask
}

type TimeTask struct {
    Url string
    TimeUnit string
    Interval int
    ImmediateExe bool
    Method string
    Params map[string] []string//可优化为map[string] interface{}
}

func Load(confpath string) Config {
    fmt.Println("start loading config file:"+confpath)
    content, err := ioutil.ReadFile(confpath)
    if err!=nil{
        fmt.Print("load config Error: ",err)
        os.Exit(-1);
    }
    var conf Config
    err=json.Unmarshal(content, &conf)
    if err!=nil{
        fmt.Print("Unmarshal json config Error: ",err)
        os.Exit(-1);
    }
    fmt.Println("finish loading config file!")
    return conf
}