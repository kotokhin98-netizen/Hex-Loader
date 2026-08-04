package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// BoardInfo хранит информацию о подключенной плате
type BoardInfo struct {
	Port string
	FQBN string
}

func main() {
	// Создаём приложение и окно
	a := app.New()
	w := a.NewWindow("Загрузчик HEX в Arduino")
	w.Resize(fyne.NewSize(500, 250))

	// Переменные для хранения выбранных данных
	var hexPath string
	var portPath string
	var fqbn string

	// Виджеты для отображения выбранных значений
	hexLabel := widget.NewLabel("Файл не выбран")
	portLabel := widget.NewLabel("Порт не выбран")
	statusLabel := widget.NewLabel("Готов к работе")

	// Кнопка выбора HEX файла
	btnSelectHex := widget.NewButton("Выбрать HEX", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				return // пользователь отменил выбор
			}
			defer reader.Close()
			hexPath = reader.URI().Path()
			hexLabel.SetText(hexPath)
			statusLabel.SetText("HEX файл выбран")
		}, w)
	})

	// Кнопка выбора порта (автоматическое определение)
	btnSelectPort := widget.NewButton("Выбрать порт", func() {
		boards, err := getAvailableBoards()
		if err != nil {
			dialog.ShowError(fmt.Errorf("не удалось получить список плат: %v", err), w)
			return
		}
		if len(boards) == 0 {
			dialog.ShowInformation("Нет плат", "Подключенные Arduino платы не найдены.", w)
			return
		}
		// Если найдена только одна плата, выбираем её автоматически
		if len(boards) == 1 {
			portPath = boards[0].Port
			fqbn = boards[0].FQBN
			portLabel.SetText(portPath + " (" + fqbn + ")")
			statusLabel.SetText("Порт выбран: " + portPath)
			return
		}
		// Если плат несколько, показываем список для выбора
		items := make([]string, len(boards))
		for i, b := range boards {
			items[i] = b.Port + " (" + b.FQBN + ")"
		}
		selected := widget.NewSelect(items, func(s string) {
			// Находим выбранную плату
			for _, b := range boards {
				if b.Port+" ("+b.FQBN+")" == s {
					portPath = b.Port
					fqbn = b.FQBN
					portLabel.SetText(s)
					statusLabel.SetText("Порт выбран: " + portPath)
					break
				}
			}
		})
		dialog.ShowCustom("Выберите плату", "OK", selected, w)
	})

	// Кнопка загрузки
	btnUpload := widget.NewButton("Загрузить", func() {
		if hexPath == "" {
			dialog.ShowInformation("Ошибка", "Сначала выберите HEX файл.", w)
			return
		}
		if portPath == "" {
			dialog.ShowInformation("Ошибка", "Сначала выберите порт.", w)
			return
		}
		if fqbn == "" {
			dialog.ShowInformation("Ошибка", "Не удалось определить тип платы (FQBN).", w)
			return
		}
		statusLabel.SetText("Загрузка...")
		err := uploadHex(hexPath, portPath, fqbn)
		if err != nil {
			statusLabel.SetText("Ошибка загрузки")
			dialog.ShowError(fmt.Errorf("ошибка загрузки: %v", err), w)
			return
		}
		statusLabel.SetText("Загрузка успешно завершена!")
		dialog.ShowInformation("Успех", "Прошивка загружена на плату.", w)
	})

	// Собираем интерфейс
	content := container.NewVBox(
		widget.NewLabel("Загрузчик HEX файлов в Arduino"),
		widget.NewSeparator(),
		container.NewHBox(btnSelectHex, hexLabel),
		container.NewHBox(btnSelectPort, portLabel),
		widget.NewSeparator(),
		btnUpload,
		statusLabel,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

// getAvailableBoards возвращает список доступных плат с их портами и FQBN
func getAvailableBoards() ([]BoardInfo, error) {
	cmd := exec.Command("arduino-cli", "board", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var boards []BoardInfo
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		// Пример строки: /dev/ttyACM0  Serial Port   arduino:avr:uno
		port := fields[0]
		if !strings.HasPrefix(port, "/dev/") && !strings.HasPrefix(port, "COM") {
			continue
		}
		// Последнее поле — FQBN (иногда может быть два последних, но берём последнее)
		fqbn := fields[len(fields)-1]
		// Дополнительная проверка: FQBN обычно содержит двоеточия
		if strings.Count(fqbn, ":") < 2 {
			// Если не похоже на FQBN, пропускаем (может быть лишний текст)
			continue
		}
		boards = append(boards, BoardInfo{Port: port, FQBN: fqbn})
	}
	return boards, nil
}

// uploadHex выполняет загрузку HEX файла на плату через arduino-cli
func uploadHex(hexPath, portPath, fqbn string) error {
	cmd := exec.Command(
		"arduino-cli",
		"upload",
		"-p", portPath,
		"--fqbn", fqbn,
		"--input-file", hexPath,
	)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%v: %s", err, stderr.String())
	}
	return nil
}