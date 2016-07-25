package main

import (
    "fmt"
    "time"
    "net/http"
    "net/url"
    "log"
    "os"
    "flag"
    "./conf"
)

//get方法请求apiAddress
func get(apiAddress string){
    client := &http.Client{}
    req, err := http.NewRequest("GET", apiAddress, nil)
    
    resp, err := client.Do(req)

    if err != nil {
        log.Println("====GET: ",apiAddress," 不可访问，请检查url是否正确")
        return
    }
    defer resp.Body.Close()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           
    if resp.StatusCode == 200 {
        log.Println("====GET: ",apiAddress," 可访问，且http response code = 200")
    } else {
        log.Println("====GET: ",apiAddress," 可访问，但http response code = ",resp.StatusCode)
    }
}

//post方法请求apiAddress
func post(apiAddress string){
    resp, err := http.PostForm(apiAddress, url.Values{"name": {"mamian"}})

    if err != nil {
        log.Println("====POST: ",apiAddress," 不可访问，请检查url是否正确")
        return
    }
    defer resp.Body.Close()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           
    if resp.StatusCode == 200 {
        log.Println("====POST: ",apiAddress," 可访问，且http response code = 200")
    } else {
        log.Println("====POST: ",apiAddress," 可访问，但http response code = ",resp.StatusCode)
    }
}


//=======================================================================================
//隔hour小时get方法请求一次apiAddress
func timer_get(interval int, timeunit string, apiAddress string) {
    if interval <= 0 {
        log.Println("========无法获取interval参数或参数错误")
        return
    }

    timer := time.NewTicker(time.Duration(interval) * time.Hour)
    switch timeunit{
        case "Day":
            timer = time.NewTicker(time.Duration(interval) * 24 * time.Hour)
        case "Hour":
            timer = time.NewTicker(time.Duration(interval) * time.Hour)
        case "Minute":
            timer = time.NewTicker(time.Duration(interval) * time.Minute)
        case "Second":
            timer = time.NewTicker(time.Duration(interval) * time.Second)
    }
    for {
        select {
        case <-timer.C:
            get(apiAddress)
        }
    }
}


func timer_post(interval int, timeunit string, apiAddress string) {
    if interval <= 0 {
        log.Println("========无法获取interval参数或参数错误")
        return
    }
    timer := time.NewTicker(time.Duration(interval) * time.Hour)
    switch timeunit{
        case "Day":
            timer = time.NewTicker(time.Duration(interval) * 24 * time.Hour)
        case "Hour":
            timer = time.NewTicker(time.Duration(interval) * time.Hour)
        case "Minute":
            timer = time.NewTicker(time.Duration(interval) * time.Minute)
        case "Second":
            timer = time.NewTicker(time.Duration(interval) * time.Second)
    }
    for {
        select {
            case <-timer.C:
                post(apiAddress)
        }
    }
}

func timer(methed string, interval int, timeunit string, apiAddress string){
    log.Println("========schedule start!")
    switch methed {
        case "get" :
            timer_get(interval, timeunit, apiAddress)
        case "post" :
            timer_post(interval, timeunit, apiAddress)
    }
}
//=======================================================================================
func main() {
    //接口地址的根目录
    rootUrl := flag.String("rootUrl", "", "rootUrl+apiPath=apiAddress")
    flag.Parse()

    config := conf.Load("conf/conf.json")

    //如果LogPath不为""，输出日志到LogPath，否则日志输出到控制台
    if len(config.LogPath) > 0 {
        logfile,err := os.OpenFile(config.LogPath,os.O_CREATE|os.O_APPEND|os.O_RDWR,0);
        if err!=nil {
            fmt.Printf("%s\r\n",err.Error());
            os.Exit(-1);
        }
        defer logfile.Close();
        log.SetOutput(logfile)
    }

    log.Println("========load conf file success!")

    for index, eachtask := range config.Tasks {//eachtask.ImmediateExe)eachtask.Method)eachtask.Params)
        if index == len(config.Tasks)-1 {
            timer(eachtask.Method, eachtask.Interval, eachtask.TimeUnit, *rootUrl+eachtask.Url)
        } else {
            go timer(eachtask.Method, eachtask.Interval, eachtask.TimeUnit, *rootUrl+eachtask.Url)
        }
    }
}