# Go developer related info

## Updating enumnerations (Go constants)
First install the stringer command.
```
go install golang.org/x/tools/cmd/stringer@latest
```

Then you may stringify:
```
cd {package folder)  // eg. cd device
go generate
```

## Module initialisation (for developers only)
```
go mod init github.com/siuyin/personal-assistant
```

## Installing godoc
```
go install golang.org/x/tools/cmd/godoc@latest

godoc -http :3999
```
Point your browser to localhost:3999 then search for `siuyin`.


