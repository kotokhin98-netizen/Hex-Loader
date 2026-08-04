# Arduino GitHub Downloader

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Fyne](https://img.shields.io/badge/Fyne-UI-blueviolet)](https://fyne.io/)

## Hex Loader

**GUI-загрузчик hex-файлов** на базе **Arduino CLI**, созданный на **Go** с использованием **Fyne** для Linux.

### 🚀 Что умеет Hex Loader?

- Выбирать HEX-файлы через стандартный диалог.
- Обнаруживать подключенные платы (даже если они не распознаны `arduino-cli`).
- Предлагать выбрать `FQBN` вручную для нераспознанных плат (из списка или вводом своего).
- Загружать прошивку на плату с использованием `arduino-cli`.
- Работать без интернета (используя локально установленные ядра и инструменты).
- Работать в **Alt Linux p11** и других дистрибутивах с установленными зависимостями.
- Красивый интерфейс с обводкой кнопок.

![Img](/img/1.png)
![Img](/img/2.png)

---

### 📦 Alt Linux p11

**Установка системных зависимостей** (делается один раз):
```bash
epmi --auto gcc-c++ golang rpm-build-golang libXcursor-devel libX11-devel libGL-devel libXrender-devel libXfixes-devel libXi-devel libXinerama-devel libXrandr-devel libwayland-egl-devel libXxf86vm-devel libxkbcommon-devel libwayland-cursor-devel arduino-cli
```

#### Создание Go-проекта (на примере Hex Loader)

```bash
mkdir hex-loader && cd hex-loader
micro main.go   # вставьте код
go mod init hex-loader
go get fyne.io/fyne/v2@latest
go install fyne.io/tools/cmd/fyne@latest   # опционально
go mod tidy
go run main.go   # проверка
```

#### Сборка Go-проекта

С оптимизацией размера бинарника:
```bash
go build -ldflags="-s -w" -trimpath -o hex-loader main.go
```

- `-s` — удаляет отладочную информацию.
- `-w` — удаляет DWARF-таблицы.
- `-trimpath` — убирает пути к исходникам из бинарника.

**Запуск:**
```bash
./hex-loader
```

**Проверка зависимостей:**
```bash
ldd hex-loader
```

**Проверка работы `arduino-cli`** (подключённые платы):
```bash
arduino-cli board list
```

---

### 🐧 Ubuntu / Linux Mint

**Установка системных зависимостей** (один раз):
```bash
sudo apt update
sudo apt install build-essential golang libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libwayland-dev libwayland-egl-dev arduino-cli
```

**Создание и сборка проекта** — аналогично Alt Linux p11.

---

### 🌍 Arduino CLI не может загрузить индексы в РФ из-за санкций

**Варианты решения проблемы:**

1. **Используйте системный VPN** — самый надёжный способ.
2. **Скачайте и распакуйте вручную** файлы индексов:
   - `https://downloads.arduino.cc/packages/package_index.tar.bz2`
   - `https://downloads.arduino.cc/libraries/library_index.tar.bz2`

   в папку `~/.arduino15`.
3. **Перенесите папку `~/.arduino15/packages/`** с другого компьютера, где уже установлены ядра.

**После переноса восстановите права:**
```bash
# Восстановить владельца
chown -R $USER:$USER ~/.arduino15/

# Права на файлы и папки
find ~/.arduino15/packages/ -type f -exec chmod 644 {} \;
find ~/.arduino15/packages/ -type d -exec chmod 755 {} \;

# Исполняемые файлы
find ~/.arduino15/packages/ -name "avrdude" -exec chmod +x {} \;
find ~/.arduino15/packages/ -name "*.so*" -exec chmod +x {} \;
find ~/.arduino15/packages/ -path "*/bin/*" -exec chmod +x {} \;
```

---

## 📄 Лицензия

Этот проект распространяется под лицензией **MIT**. Подробнее см. в файле [LICENSE](LICENSE).

---

## 🙏 Благодарности

Программа разработана при участии **AI-ассистента** (DeepSeek AI) и доработана на основе обратной связи сообщества.

Если у вас есть идеи по улучшению, создавайте **Issue** или отправляйте **Pull Request**!

---

## 📝 Дополнительная информация

- Скрипт тестировался в **ALT Education 11.2** и **Ubuntu 22.04/24.04**, но должен работать в любом дистрибутиве Linux с графическим интерфейсом.
- Для корректной установки плат, перед запуском Hex Loader запустите `arduino-cli` с VPN, чтобы загрузить индексы.

---

### 🛠️ Планы по улучшению

- [ ] Добавить прогресс-бар при загрузке
- [ ] Сохранение последнего выбранного FQBN
- [ ] Автообновление списка портов
- [ ] Упаковка в AppImage/DEB-пакет
```