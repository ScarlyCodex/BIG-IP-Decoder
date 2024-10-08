# BIG-IP-Decoder
#### Output
![Screenshot from 2024-10-08 02-30-03](https://github.com/user-attachments/assets/bbc9b98c-0c85-459b-8e1b-9f683696daf3)
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
