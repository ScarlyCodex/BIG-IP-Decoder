# BIG-IP-Decoder
#### Demo
![BIGIPDECODER](https://github.com/user-attachments/assets/7787425f-961c-43ef-bfd1-90051df9a2cb)
# Installation
#### VIA go install
```
go install github.com/ScarlyCodex/BIG-IP-Decoder@latest
sudo mv ~/go/bin/BIG-IP-Decoder /usr/local/bin/big_ip_decoder
```
#### VIA git clone
```
git clone https://github.com/ScarlyCodex/BIG-IP-Decoder.git
cd BIG-IP-Decoder
go build -gccgoflags="-s -w" -o big_ip_decoder big_ip_decoder.go
sudo mv big_ip_decoder /usr/local/bin/big_ip_decoder
```
# Usage
`big_ip_decoder <url>`
