## Hex Loader

GUI-загрузчик hex-файлов на базе **Arduino CLI** созданный на **Go** с использованием **Fyne** и **Arduino CLI** для Linux

Что умеет **Hex Loader**?
- Выбирать HEX-файлы через стандартный диалог.
- Обнаруживать подключенные платы (даже если они не распознаны `arduino-cli`).
- Предлагать выбрать `FQBN` вручную для нераспознанных плат (из списка или вводом своего).
- Загружать прошивку на плату с использованием `arduino-cli`.
- Работать без интернета (используя локально установленные ядра и инструменты).
- Работать в **Alt Linux p11** (и других дистрибутивах с установленными зависимостями).

### Alt Linux 11

Установка системных зависимостей (делается один раз)
```shell
epmi --auto gcc-c++ golang rpm-build-golang libXcursor-devel libX11-devel libGL-devel libXrender-devel libXfixes-devel libXi-devel libXinerama-devel libXrandr-devel libwayland-egl-devel libXxf86vm-devel libxkbcommon-devel libwayland-cursor-devel arduino-cli
```

#### Созданиие Go-проекта на примере Hex Loader
```
mkdir hex-loader && cd hex-loader
micro main.go   # (напишите код)
go mod init hex-loader
go get fyne.io/fyne/v2@latest
go install fyne.io/tools/cmd/fyne@latest   # опционально
go mod tidy
go run main.go   # проверка
```

#### Сборка Go-проекта

C оптимизацией размера бинарника
```shell
go build -ldflags="-s -w" -o hex-loader.bin main.go
```
- `s` — удаляет отладочную информацию.
- `-w` — удаляет DWARF-таблицы (ещё меньше размер).
или
```shell
go build -ldflags="-s -w" -trimpath -o hex-loader.bin main.go
```
-trimpath — убирает пути к исходникам из бинарника.

Запуск
```shell
./hex-loader.bin
```

Проверка зависимостей
```shell
ldd hex-loader.bin
```

Проверка работы `arduino-cli`

Проверить подключённые платы:
```shell
arduino-cli board list
```


### Ubuntu/Mint (находится в разработке)

Установка системных зависимостей (один раз)
```shell
sudo apt update
sudo apt install build-essential golang libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libwayland-dev libwayland-egl-dev arduino-cli
```

Создание и сборка проекта как в  Alt Linux p11

### Arduino CLI не может загрузить индексы в РФ из-за санкций

Варианты решения проблемы:
1. Используйте системный VPN
2. Скачайте и распакуйте вручную файлы library_index.tar.bz2 и package_index.tar.bz2 по ссылкам
    - https://downloads.arduino.cc/packages/package_index.tar.bz2
    - https://downloads.arduino.cc/libraries/library_index.tar.bz2
    в папку ~/.arduino15
    и перенесите с другого компьютера содержимое /home/user/.arduino15/packages/ и назначте права:
```shell
# Восстановить владельца и права для всех файлов в ~/.arduino15/
chown -R $USER:$USER ~/.arduino15/
find ~/.arduino15/packages/ -type f -exec chmod 644 {} \;
find ~/.arduino15/packages/ -type d -exec chmod 755 {} \;
find ~/.arduino15/packages/ -name "avrdude" -exec chmod +x {} \;
find ~/.arduino15/packages/ -name "*.so*" -exec chmod +x {} \;
# И другие исполняемые файлы (например, dfu-programmer, bossac и т.д.)
find ~/.arduino15/packages/ -path "*/bin/*" -exec chmod +x {} \;
```