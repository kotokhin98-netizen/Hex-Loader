```
arduino-cli board list
Скачивание индекса: library_index.tar.bz2 Ответ сервера: 403 Forbidden
Ошибка при инициализации экземпляра: Ошибка скачивания индекса 'https://downloads.arduino.cc/libraries/library_index.tar.bz2': Ответ сервера: 403 Forbidden
Ошибка при инициализации экземпляра: Загрузка индексного файла: загрузка индексного файла json /home/user/.arduino15/package_index.json: open /home/user/.arduino15/package_index.json: no such file or directory
Ошибка при инициализации экземпляра: Ошибка загрузки аппаратной платформы: обнаружение builtin:serial-discovery не найдено
Ошибка при инициализации экземпляра: Ошибка загрузки аппаратной платформы: обнаружение builtin:mdns-discovery не найдено
Ошибка при инициализации экземпляра: Ошибка загрузки аппаратной платформы: обнаружение builtin:serial-discovery не найдено
Ошибка при инициализации экземпляра: Ошибка загрузки аппаратной платформы: обнаружение builtin:mdns-discovery не найдено
Ошибка при инициализации экземпляра: Загрузка индексного файла: чтение library_index.json: open /home/user/.arduino15/library_index.json: no such file or directory
Платы не найдены.
```

```
Скачивание пропущенных инструментов builtin:serial-discovery@1.5.2...
builtin:serial-discovery@1.5.2 Ответ сервера: 403 Forbidden
Ошибка при инициализации экземпляра: скачивание builtin:serial-discovery@1.5.2 инструмента: Ответ сервера: 403 Forbidden
Скачивание пропущенных инструментов builtin:mdns-discovery@1.1.0...
builtin:mdns-discovery@1.1.0 Ответ сервера: 403 Forbidden
Ошибка при инициализации экземпляра: скачивание builtin:mdns-discovery@1.1.0 инструмента: Ответ сервера: 403 Forbidden
Порт         Протокол Тип               Наименование платы FQBN Ядро
/dev/ttyUSB0 serial   Serial Port (USB) Неизвестный
```

```
go build -ldflags="-s -w" -trimpath -o hex-loader main_v2.go 
# command-line-arguments
./main_v2.go:148:3: portLabel.SetText(board.Port) (no value) used as value or type
```