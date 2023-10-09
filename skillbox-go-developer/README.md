## Finish project usage  

### collect repo
```
(in home directory)
git clone https://github.com/paveltrpn/skillbox-go-developer
```
  
### start all services:
```
cd ~/skillbox-go-developer/assets/simulator/
go run main.go
cd ~/skillbox-go-developer/cmd/diploma/
go run server.go --port {number}
```
setting --port is optional
default server port is 8585
  
### then open data wb page and see results
```
cd ~
firefox code/skillbox-go-developer/web/index.html
```
default client port is 8585