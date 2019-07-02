# mod-ldp

Current dev environment is to build the Docker image from inside the vagrant VM (folio/testing). 

Install Go and mod-ldp:
```
vagrant ssh
wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
sudo tar -xvf go1.12.6.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=/vagrant/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
cd $GOPATH
go get github.com/folio-org/mod-ldp
```

Build go binary¹:
```
CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s"
```

Build Docker image²
```
sudo docker build -t mod-ldp --no-cache .
```

Run Docker image
```
sudo docker run -p 8001:8001 --rm mod-ldp
```

## Okapi

Use the `scripts` to declare, deploy, enable, and test the module.


---

¹ The Go compile flags are for the minimal Docker image where the binary will be run. Without them, I get the error `exec user process caused "no such file or directory"`

² `--no-cache` to ensure Docker doesn't use a cached Go binary
