## Hex Loader

GUI-загрузчик hex-файлов на базе **Arduino CLI** созданный на **Go**

### Alt Linux 11

Установка системных зависимостей (делается один раз)
```shell
epmi --auto gcc-c++ golang rpm-build-golang libXcursor-devel libX11-devel libGL-devel libXrender-devel libXfixes-devel libXi-devel libXinerama-devel libXrandr-devel libwayland-egl-devel libXxf86vm-devel libxkbcommon-devel libwayland-cursor-devel
```

#### Созданиие Go-проекта

mkdir hex-loader && cd hex-loader
micro main.go   # (напишите код)
go mod init hex-loader
go get fyne.io/fyne/v2@latest
go install fyne.io/tools/cmd/fyne@latest   # опционально
go mod tidy
go run main.go   # проверка

#### Сборка Go-проекта

C оптимизацией размера бинарника
```shell
go build -ldflags="-s -w" -o hex-loader main.go
```
- s — удаляет отладочную информацию.
- -w — удаляет DWARF-таблицы (ещё меньше размер).
или
```shell
go build -ldflags="-s -w" -trimpath -o hex-loader main.go
```
-trimpath — убирает пути к исходникам из бинарника.

Проверка зависимостей
```shell
ldd hex-loader
```


### Ubuntu/Mint (находится в разработке)

Установка системных зависимостей (один раз)
```shell
sudo apt update
sudo apt install build-essential golang libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libwayland-dev libwayland-egl-dev
```

Создание и сборка проекта теже, что и для Alt Linux p11
