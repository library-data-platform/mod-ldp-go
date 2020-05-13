# mod-ldp

## Run stand-alone

Install:
```
go get github.com/folio-org/mod-ldp
```

Run:
```
export LDP_PW="ENTER_PASSWORD_HERE"
go build
./mod-ldp
```

Check http://localhost:8001/ldp/db/status to check database connection.

### Run with Okapi in VM

Current dev environment is to build the Docker image from inside the vagrant VM (folio/testing). 

First, ssh into the VM:
```
vagrant ssh
```

Install Go:

```
wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
sudo tar -xvf go1.12.6.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=/vagrant/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
cd $GOPATH
```

Install mod-ldp:
```
go get github.com/folio-org/mod-ldp
```

Build go binary¹
```
CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s"
```

Build Docker image²
```
sudo docker build -t mod-ldp --no-cache .
```

Run Docker image³
```
sudo docker run -p 8001:8001 --rm mod-ldp
```

## Okapi

Use the `scripts` to declare, deploy, enable, and test the module.
```
cd scripts
./okapi-declare-mod.sh
./okapi-deploy-mod.sh
./okapi-enable-mod-for-tenant.sh
./okapi-test-mod.sh

# Cleanup:
./okapi-delete-mod.sh
```

---

¹ The Go compile flags are for the minimal Docker image where the binary will be run. Without them, I get the error `exec user process caused "no such file or directory"`

² `--no-cache` to ensure Docker doesn't use a cached Go binary

³ For development of the Go server, it's probably easiest to just run the binary. For testing the Docker image, run the Docker image, and for testing integration with other FOLIO modules, deploy to Okapi.
