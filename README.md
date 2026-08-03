## Hex Loader

GUI-загрузчик hex-файлов на базе **Arduino CLI** созданный на **Go**

### Alt Linux 11

```shell
epmi --auto golang rpm-build-golang libXcursor-devel libX11-devel libGL-devel libXrender-devel libXfixes-devel libXi-devel libXinerama-devel libXrandr-devel libwayland-egl-devel libXxf86vm libwayland-egl-devel
```

#### Настройка Go-проекта

```shell
mkdir hex-loader
```
```shell
cd hex-loader/
```
```shell
go mod init hex-loader
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

#### Сборка Go-проекта

C оптимизацией размера бинарника
```shell
go build -ldflags="-s -w" -o xloader main.go
```
- s — удаляет отладочную информацию.
- -w — удаляет DWARF-таблицы (ещё меньше размер).


### Ubuntu/Mint (находится в разработке)
