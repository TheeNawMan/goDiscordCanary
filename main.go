package main

import(
    "os"
    "log"
    "github.com/theenawman/goDiscordCanary/pkg/webhook"
)

func main() {
    var username = os.Hostname()
    var content = os.Getenv("IP")
    var url = "https://discord.com/api/webhooks/<<CHANGE THIS>>"
    var image_url = "https://rshell.nyc3.cdn.digitaloceanspaces.com/cosmo.png"

    image := webhook.Image{
        Url: &image_url,
    }

    embed := webhook.Embed{
        Image: &image,
    }

    message := webhook.Message{
        Username: &username,
        Content:  &content,
        Embeds:   &[]discordwebhook.Embed{embed},
    }

    err := webhook.SendMessage(url, message)
    if err != nil {
        log.Fatal(err)
    }
}