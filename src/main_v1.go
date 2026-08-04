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

func main() {
	// Создаём приложение и окно
	a := app.New()
	w := a.NewWindow("Загрузчик HEX в Arduino")
	w.Resize(fyne.NewSize(500, 250))

	// Переменные для хранения выбранных данных
	var hexPath string
	var portPath string

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
		ports, err := getAvailablePorts()
		if err != nil {
			dialog.ShowError(fmt.Errorf("не удалось получить список портов: %v", err), w)
			return
		}
		if len(ports) == 0 {
			dialog.ShowInformation("Нет портов", "Подключенные Arduino платы не найдены.", w)
			return
		}
		// Если найден только один порт, выбираем его автоматически
		if len(ports) == 1 {
			portPath = ports[0]
			portLabel.SetText(portPath)
			statusLabel.SetText("Порт выбран: " + portPath)
			return
		}
		// Если портов несколько, показываем список для выбора
		items := make([]string, len(ports))
		for i, p := range ports {
			items[i] = p
		}
		selected := widget.NewSelect(items, func(s string) {
			portPath = s
			portLabel.SetText(portPath)
			statusLabel.SetText("Порт выбран: " + portPath)
		})
		dialog.ShowCustom("Выберите порт", "OK", selected, w)
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
		statusLabel.SetText("Загрузка...")
		// Для простоты используем FQBN arduino:avr:uno (самая распространённая плата)
		// В более продвинутой версии можно добавить выбор FQBN
		fqbn := "arduino:avr:uno"
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

// getAvailablePorts возвращает список доступных последовательных портов
// с помощью команды arduino-cli board list
func getAvailablePorts() ([]string, error) {
	cmd := exec.Command("arduino-cli", "board", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var ports []string
	for _, line := range lines {
		// Ищем строки, содержащие путь к порту (начинаются с /dev/ или COM)
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		// Первое поле в выводе arduino-cli board list — это путь к порту
		potentialPort := fields[0]
		if strings.HasPrefix(potentialPort, "/dev/") || strings.HasPrefix(potentialPort, "COM") {
			ports = append(ports, potentialPort)
		}
	}
	return ports, nil
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
