## Hex Loader

GUI-загрузчик hex-файлов для Arduino на Go

### Alt Linux 11

```shell
epmi --auto golang rpm-build-golang libXcursor-devel libX11-devel libGL-devel libXrender-devel libXfixes-devel libXi-devel libXinerama-devel libXrandr-devel libwayland-egl-devel libXxf86vm libwayland-egl-devel
```

#### Настройка Go

```shell
mkdir hexloader
```
```shell
cd hexloader/
```
```shell
go mod init hexloader
```
```shell
go get fyne.io/fyne/v2@latest
```
```shell
go install fyne.io/tools/cmd/fyne@latest
```
```shell
micro main.go
```
```shell
go mod tidy
```
```shell
go run main.go
```

### Ubuntu/Mint
