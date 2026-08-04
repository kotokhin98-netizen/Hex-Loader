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

// checkArduinoCLI проверяет, доступен ли arduino-cli в PATH
func checkArduinoCLI() error {
	_, err := exec.LookPath("arduino-cli")
	return err
}

func main() {
	a := app.NewWithID("com.example.hexloader")
	w := a.NewWindow("Загрузчик HEX в Arduino")
	w.Resize(fyne.NewSize(550, 280))

	var hexPath string
	var portPath string
	var fqbn string

	hexLabel := widget.NewLabel("Файл не выбран")
	portLabel := widget.NewLabel("Порт не выбран")
	statusLabel := widget.NewLabel("Готов к работе")

	// Кнопки
	btnSelectHex := widget.NewButton("Выбрать HEX", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()
			hexPath = reader.URI().Path()
			hexLabel.SetText(hexPath)
			statusLabel.SetText("HEX файл выбран")
		}, w)
	})

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

		if len(boards) == 1 {
			selectBoard(boards[0], &portPath, &fqbn, portLabel, statusLabel, w)
			return
		}

		items := make([]string, len(boards))
		for i, b := range boards {
			label := b.Port
			if b.FQBN != "" {
				label += " (" + b.FQBN + ")"
			} else {
				label += " (тип не определён)"
			}
			items[i] = label
		}
		selected := widget.NewSelect(items, func(s string) {
			port := strings.Split(s, " ")[0]
			for _, b := range boards {
				if b.Port == port {
					selectBoard(b, &portPath, &fqbn, portLabel, statusLabel, w)
					break
				}
			}
		})
		dialog.ShowCustom("Выберите плату", "OK", selected, w)
	})

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
			showFQBNInputDialog(w, &fqbn, statusLabel)
			if fqbn == "" {
				return
			}
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

	// Проверяем наличие arduino-cli ДО создания содержимого
	// Если нет — отключаем кнопки и показываем статус
	if err := checkArduinoCLI(); err != nil {
		btnSelectHex.Disable()
		btnSelectPort.Disable()
		btnUpload.Disable()
		statusLabel.SetText("❌ arduino-cli не найден")
	}

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

func selectBoard(board BoardInfo, portPath *string, fqbn *string, portLabel *widget.Label, statusLabel *widget.Label, w fyne.Window) {
	*portPath = board.Port
	portLabel.SetText(board.Port)
	if board.FQBN != "" {
		*fqbn = board.FQBN
		portLabel.SetText(board.Port + " (" + board.FQBN + ")")
		statusLabel.SetText("Порт выбран: " + board.Port + " (" + board.FQBN + ")")
	} else {
		statusLabel.SetText("Тип платы не определён. Выберите FQBN.")
		showFQBNInputDialog(w, fqbn, statusLabel)
		if *fqbn != "" {
			portLabel.SetText(board.Port + " (" + *fqbn + ")")
			statusLabel.SetText("Порт выбран: " + board.Port + " (" + *fqbn + ")")
		} else {
			portLabel.SetText(board.Port + " (тип не выбран)")
			statusLabel.SetText("FQBN не выбран")
		}
	}
}

func showFQBNInputDialog(w fyne.Window, fqbn *string, statusLabel *widget.Label) {
	commonFQBNs := []string{
		"arduino:avr:uno",
		"arduino:avr:nano",
		"arduino:avr:mega",
		"arduino:avr:leonardo",
		"esp32:esp32:esp32",
		"esp8266:esp8266:generic",
	}
	selectWidget := widget.NewSelect(commonFQBNs, func(s string) {
		*fqbn = s
		statusLabel.SetText("FQBN выбран: " + s)
	})
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Или введите свой FQBN вручную")
	entry.OnSubmitted = func(s string) {
		if s != "" {
			*fqbn = s
			statusLabel.SetText("FQBN введён: " + s)
		}
	}
	content := container.NewVBox(
		widget.NewLabel("Выберите или введите FQBN для вашей платы:"),
		selectWidget,
		widget.NewLabel("или"),
		entry,
	)
	dialog.ShowCustom("Выбор FQBN", "OK", content, w)
}

// getAvailableBoards возвращает список плат через текстовый парсинг
func getAvailableBoards() ([]BoardInfo, error) {
	cmd := exec.Command("arduino-cli", "board", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("ошибка запуска arduino-cli: %v", err)
	}

	lines := strings.Split(out.String(), "\n")
	var boards []BoardInfo
	for _, line := range lines {
		if !strings.Contains(line, "/dev/") && !strings.Contains(line, "COM") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		port := fields[0]
		if !strings.HasPrefix(port, "/dev/") && !strings.HasPrefix(port, "COM") {
			continue
		}
		// Ищем FQBN
		var fqbn string
		for _, field := range fields {
			if strings.Count(field, ":") >= 2 {
				fqbn = field
				break
			}
		}
		boards = append(boards, BoardInfo{Port: port, FQBN: fqbn})
	}
	return boards, nil
}

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
