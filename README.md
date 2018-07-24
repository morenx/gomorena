## Developer Setup

1. Download golang (1.10.3)

2. Download dep for golang

3. Download direnv

4. Download mysql client for mariadb

5. Download docker

6. Load `direnv` in your profile or shell rc file

```sh
eval "$(direnv hook zsh)"
```

7. Setup the gopath

```sh
mkdir -p ~/.gopaths/gomorena/src/github.com/morena/
cd ~/.gopaths/gomorena/src/github.com/morena/
git clone git@github.com:morena/gomorena.git
cd gomorena
cp envrc.sample .envrc
direnv allow
```

8. Install all dependencies with dep:

```sh
make deps
```

9. Install go linter tool of your choice. For `gometalinter`:

```sh
go get -u github.com/alecthomas/gometalinter
gometalinter --install
```

10. Use autotools make to start a local mysql instance and run the application:

```sh
make db
make run
```
