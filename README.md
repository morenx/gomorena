## Developer Setup

1. Download golang (1.10.3), dep for golang, autotools, direnv, mysql client
   and docker
2. Setup the gopath

```sh
mkdir -p ~/.gopaths/gomorena/src/github.com/morena/
cd ~/.gopaths/gomorena/src/github.com/morena/
git clone git@github.com:morena/gomorena.git
cd gomorena
cp envrc.sample .envrc
direnv allow
```

3. Install all dependencies with dep:

```sh
make dep
```

4. Use autotools make to start a local mysql instance and run the application:

```sh
make db
make run
```
