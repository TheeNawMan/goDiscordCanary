# goDiscordCanary
A simple canary written in go.

# Prupose
This was designed for bug bounties to prove code execution in an enviornment. It grabs key information as proof and then sends it via a discord webhook on port 443 out of the enviornemnt. Since 443 is generally avaliable through firewalls the only thing standing in the way is content filters. The code is written in a way that you should be able to repace the url with your own webhook(untested) and get the data.

# Usage
1. Compile this into your favorite platorm using go build commands
2. Using a downloader or your favorite lolbins or gtfo bins execute this on device.
3. Profit from the data being in your discord channel.

   
I have tested this on all three platforms using arm. More testing will come in the future.

# Features
[X] Grabs public IP
[X] Grabs hostname
[ ] Universal Screenshot utility
[ ] Internal IP
[ ] Make url be taken as a .env or arqument for CTF usage
[ ] IF NEEDED: Lets make sure AV doesn't trigger by adding some hops in the code.
