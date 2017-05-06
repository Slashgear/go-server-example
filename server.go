package main

import (
    "net/http"
    "log"
    "os"
    "net"
    "math/rand"
    "html/template"
    "strconv"
)


type data struct {
  Title string
  R int
  G int
  B int
  Msg string
  Hostname string
  Ips []net.IP
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    data := new(data)
    title, present := os.LookupEnv("APP_TITLE")
    if present {
      data.Title = title
    } else {
      data.Title = "Title"
    }
    msg, present := os.LookupEnv("APP_MSG")
    if present {
      data.Msg = msg
    } else {
      data.Msg = "Hello world !"
    }
    red, present := os.LookupEnv("APP_R")
    if present {
      data.R, _ = strconv.Atoi(red)
    } else {
      data.R = rand.Intn(256)
    }
    green, present := os.LookupEnv("APP_G")
    if present {
      data.G, _ = strconv.Atoi(green)
    } else {
      data.G = rand.Intn(256)
    }
    blue, present := os.LookupEnv("APP_B")
    if present {
      data.B, _ = strconv.Atoi(blue)
    } else {
      data.B = rand.Intn(256)
    }
    hostname, _ := os.Hostname()
    data.Hostname = hostname
    ifaces, _ := net.Interfaces()
    for _, i := range ifaces {
      addrs, _ := i.Addrs()
      // handle err
      for _, addr := range addrs {
        var ip net.IP
        switch v := addr.(type) {
        case *net.IPNet:
          ip = v.IP
        case *net.IPAddr:
          ip = v.IP
        }
        data.Ips = append(data.Ips, ip)
      }
    }
    t, _ := template.ParseFiles(`index.html`)
    t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", sayhelloName)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
