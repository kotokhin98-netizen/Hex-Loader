package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// BoardInfo хранит информацию о подключенной плате
type BoardInfo struct {
	Port string
	FQBN string
}

// PortInfo для парсинга JSON вывода arduino-cli
type PortInfo struct {
	Port struct {
		Address string `json:"address"`
	} `json:"port"`
	MatchingBoards []struct {
		FQBN string `json:"fqbn"`
	} `json:"matching_boards"`
}

// CoreInfo для парсинга JSON вывода arduino-cli core list
type CoreInfo struct {
	ID        string `json:"id"`
	Installed string `json:"installed"`
}

// myTheme — кастомная тема с увеличенным размером
type myTheme struct {
	fyne.Theme
}

func (t myTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 16
	case theme.SizeNamePadding:
		return 14
	default:
		return t.Theme.Size(name)
	}
}

func (t myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xFF}
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0xE0, G: 0xE0, B: 0xE0, A: 0xFF}
	default:
		return t.Theme.Color(name, variant)
	}
}

// buttonWithBorder — кастомная кнопка с обводкой
type buttonWithBorder struct {
	widget.Button
}

func newButtonWithBorder(label string, onTap func()) *buttonWithBorder {
	btn := &buttonWithBorder{}
	btn.ExtendBaseWidget(btn)
	btn.Text = label
	btn.OnTapped = onTap
	return btn
}

func (b *buttonWithBorder) CreateRenderer() fyne.WidgetRenderer {
	return &buttonRenderer{btn: b}
}

// buttonRenderer — рендерер кнопки с обводкой
type buttonRenderer struct {
	btn       *buttonWithBorder
	label     *canvas.Text
	bg        *canvas.Rectangle
	border    *canvas.Rectangle
	objects   []fyne.CanvasObject
	container *fyne.Container
}

func (r *buttonRenderer) init() {
	if r.container != nil {
		return
	}
	r.label = canvas.NewText(r.btn.Text, theme.Color(theme.ColorNameForeground))
	r.label.Alignment = fyne.TextAlignCenter
	r.label.TextSize = theme.TextSize()

	r.bg = canvas.NewRectangle(theme.Color(theme.ColorNameButton))
	r.bg.CornerRadius = 8

	r.border = canvas.NewRectangle(color.Transparent)
	r.border.StrokeWidth = 2.5
	r.border.StrokeColor = color.NRGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xFF}
	r.border.CornerRadius = 8

	r.container = container.NewStack(r.bg, r.border, r.label)
	r.objects = []fyne.CanvasObject{r.container}
}

func (r *buttonRenderer) Layout(size fyne.Size) {
	r.init()
	r.container.Resize(size)
}

func (r *buttonRenderer) MinSize() fyne.Size {
	r.init()
	minSize := r.container.MinSize()
	return fyne.NewSize(minSize.Width+40, minSize.Height+16)
}

func (r *buttonRenderer) Refresh() {
	r.init()
	r.label.Text = r.btn.Text
	r.label.Color = theme.Color(theme.ColorNameForeground)
	r.label.TextSize = theme.TextSize()

	if r.btn.Disabled() {
		r.bg.FillColor = theme.Color(theme.ColorNameDisabled)
		r.border.StrokeColor = color.NRGBA{R: 0x66, G: 0x66, B: 0x66, A: 0x80}
	} else {
		r.bg.FillColor = theme.Color(theme.ColorNameButton)
		r.border.StrokeColor = color.NRGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xFF}
	}
	r.bg.CornerRadius = 8
	r.border.CornerRadius = 8

	r.label.Refresh()
	r.bg.Refresh()
	r.border.Refresh()
	canvas.Refresh(r.btn)
}

func (r *buttonRenderer) Objects() []fyne.CanvasObject {
	r.init()
	return r.objects
}

func (r *buttonRenderer) Destroy() {}

// checkArduinoCLI проверяет, доступен ли arduino-cli в PATH
func checkArduinoCLI() error {
	_, err := exec.LookPath("arduino-cli")
	return err
}

// checkUserInGroup проверяет, состоит ли текущий пользователь в указанной группе
func checkUserInGroup(groupName string) (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}

	// Используем id -Gn для более надежного получения списка групп
	cmd := exec.Command("id", "-Gn", currentUser.Username)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	groups := strings.Fields(string(output))
	for _, g := range groups {
		if g == groupName {
			return true, nil
		}
	}
	return false, nil
}

// showPermissionWarning показывает предупреждение о правах
func showPermissionWarning(a fyne.App, w fyne.Window) {
	var d dialog.Dialog // Переменная для ссылки на диалог
	
	label := widget.NewLabel(
		"⚠️ У вас нет прав для работы с последовательными портами.\n\n" +
			"Чтобы Hex Loader мог определять и прошивать платы,\n" +
			"добавьте пользователя в группу dialout:\n\n" +
			"  sudo usermod -a -G dialout $USER\n\n" +
			"После этого выйдите из системы и зайдите заново.\n\n" +
			"Вы всё равно можете использовать программу,\n" +
			"но порты могут не определяться.",
	)
	
	btnOk := widget.NewButton("Понятно", func() {
		if d != nil {
			d.Hide() // Закрываем диалог при клике
		}
	})
	
	content := container.NewVBox(label, btnOk)
	d = dialog.NewCustom("Внимание", "", content, w)
	d.Resize(fyne.NewSize(500, 300))
	d.Show()
}

// uploadWithProgress выполняет загрузку с отображением статуса
func uploadWithProgress(hexPath, portPath, fqbn string, statusLabel *widget.Label) error {
	result := make(chan error)

	go func() {
		// Показываем статус через fyne.Do (безопасно для UI)
		fyne.Do(func() {
			statusLabel.SetText("⏳ Подождите, идёт загрузка...")
		})

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
			result <- fmt.Errorf("ошибка загрузки: %v: %s", err, stderr.String())
			return
		}
		result <- nil
	}()

	err := <-result

	if err != nil {
		fyne.Do(func() {
			statusLabel.SetText("❌ Ошибка загрузки")
		})
		return err
	}

	fyne.Do(func() {
		statusLabel.SetText("✅ Загрузка успешно завершена!")
	})
	return nil
}

func main() {
	a := app.NewWithID("com.example.hexloader")
	a.Settings().SetTheme(&myTheme{theme.DefaultTheme()})

	w := a.NewWindow("Загрузчик HEX в Arduino")
	w.Resize(fyne.NewSize(700, 420))

	var hexPath string
	var portPath string
	var fqbn string

	hexLabel := widget.NewLabel("Файл не выбран")
	portLabel := widget.NewLabel("Порт не выбран")
	statusLabel := widget.NewLabel("Готов к работе")

	// Кнопка выбора HEX с фильтром
	btnSelectHex := newButtonWithBorder("📂 Выбрать HEX", func() {
		fileFilter := storage.NewExtensionFileFilter([]string{".hex"})
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				fyne.Do(func() {
					dialog.ShowError(err, w)
				})
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()
			hexPath = reader.URI().Path()
			fyne.Do(func() {
				hexLabel.SetText(hexPath)
				statusLabel.SetText("HEX файл выбран")
			})
		}, w)
		fileDialog.SetFilter(fileFilter)
		fileDialog.Show()
	})

	btnSelectPort := newButtonWithBorder("🔌 Выбрать порт", func() {
		boards, err := getAvailableBoards()
		if err != nil {
			fyne.Do(func() {
				dialog.ShowError(fmt.Errorf("не удалось получить список плат: %v", err), w)
			})
			return
		}
		if len(boards) == 0 {
			fyne.Do(func() {
				dialog.ShowInformation("Нет плат", "Подключенные Arduino платы не найдены.", w)
			})
			return
		}

		if len(boards) == 1 {
			selectBoard(boards[0], &portPath, &fqbn, portLabel, statusLabel, w)
			return
		}

		items := make([]string, len(boards))
		for i, b := range boards {
			label := b.Port
			if b.FQBN != "" && strings.Count(b.FQBN, ":") >= 2 {
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
		fyne.Do(func() {
			dialog.ShowCustom("Выберите плату", "OK", selected, w)
		})
	})

	btnUpload := newButtonWithBorder("⬆️ Загрузить", func() {
		if hexPath == "" {
			fyne.Do(func() {
				dialog.ShowInformation("Ошибка", "Сначала выберите HEX файл.", w)
			})
			return
		}
		if portPath == "" {
			fyne.Do(func() {
				dialog.ShowInformation("Ошибка", "Сначала выберите порт.", w)
			})
			return
		}
		if fqbn == "" {
			showFQBNInputDialog(w, &fqbn, statusLabel)
			if fqbn == "" {
				return
			}
		}

		// Фиксируем текущие значения для избежания гонок данных
		currentHex := hexPath
		currentPort := portPath
		currentFqbn := fqbn

		// Блокируем кнопку на время прошивки
		btnUpload.Disable()

		go func() {
			// Гарантированно разблокируем кнопку после завершения
			defer fyne.Do(func() { btnUpload.Enable() })

			// Проверяем, установлено ли ядро
			coreName := strings.Split(currentFqbn, ":")[0] + ":" + strings.Split(currentFqbn, ":")[1]
			if !isCoreInstalled(coreName) {
				fyne.Do(func() {
					dialog.ShowError(fmt.Errorf("ядро %s не установлено. Установите его через arduino-cli", coreName), w)
				})
				return
			}

			err := uploadWithProgress(currentHex, currentPort, currentFqbn, statusLabel)
			if err != nil {
				fyne.Do(func() {
					dialog.ShowError(fmt.Errorf("ошибка загрузки: %v", err), w)
				})
			}
		}()
	})

	// Проверка arduino-cli
	if err := checkArduinoCLI(); err != nil {
		btnSelectHex.Disable()
		btnSelectPort.Disable()
		btnUpload.Disable()
		statusLabel.SetText("❌ arduino-cli не найден")
	}

	// Проверка прав на группы (только для Linux)
	if runtime.GOOS == "linux" {
		inDialout, _ := checkUserInGroup("dialout")
		inUucp, _ := checkUserInGroup("uucp")
		if !inDialout && !inUucp {
			go showPermissionWarning(a, w)
			statusLabel.SetText("⚠️ Нет прав на доступ к портам (нужна группа dialout)")
		}
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
	fyne.Do(func() {
		portLabel.SetText(board.Port)
	})
	
	if board.FQBN != "" && strings.Count(board.FQBN, ":") >= 2 {
		*fqbn = board.FQBN
		fyne.Do(func() {
			portLabel.SetText(board.Port + " (" + board.FQBN + ")")
			statusLabel.SetText("Порт выбран: " + board.Port + " (" + board.FQBN + ")")
		})
	} else {
		fyne.Do(func() {
			statusLabel.SetText("Тип платы не определён. Выберите FQBN.")
		})
		showFQBNInputDialog(w, fqbn, statusLabel)
		if *fqbn != "" {
			fyne.Do(func() {
				portLabel.SetText(board.Port + " (" + *fqbn + ")")
				statusLabel.SetText("Порт выбран: " + board.Port + " (" + *fqbn + ")")
			})
		} else {
			fyne.Do(func() {
				portLabel.SetText(board.Port + " (тип не выбран)")
				statusLabel.SetText("FQBN не выбран")
			})
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
	
	var dialogObj *dialog.CustomDialog
	
	selectWidget := widget.NewSelect(commonFQBNs, func(s string) {
		*fqbn = s
		fyne.Do(func() {
			statusLabel.SetText("FQBN выбран: " + s)
		})
	})
	
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Или введите свой FQBN вручную")
	entry.OnSubmitted = func(s string) {
		if s != "" {
			*fqbn = s
			fyne.Do(func() {
				statusLabel.SetText("FQBN введён: " + s)
			})
			if dialogObj != nil {
				dialogObj.Hide() // Закрываем диалог по Enter
			}
		}
	}

	okButton := newButtonWithBorder("✅ OK", func() {
		if *fqbn == "" && len(commonFQBNs) > 0 {
			*fqbn = commonFQBNs[0]
			fyne.Do(func() {
				statusLabel.SetText("FQBN выбран: " + *fqbn)
			})
		}
		if dialogObj != nil {
			dialogObj.Hide()
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Выберите или введите FQBN для вашей платы:"),
		selectWidget,
		widget.NewLabel("или"),
		entry,
		container.NewCenter(okButton),
	)

	dialogObj = dialog.NewCustom("Выбор FQBN", "", content, w)
	dialogObj.Resize(fyne.NewSize(500, 350))
	dialogObj.Show()
}

// getAvailableBoards возвращает список плат через JSON парсинг
func getAvailableBoards() ([]BoardInfo, error) {
	cmd := exec.Command("arduino-cli", "board", "list", "--format", "json")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("ошибка запуска arduino-cli: %v", err)
	}

	var ports []PortInfo
	if err := json.Unmarshal(out, &ports); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	var boards []BoardInfo
	for _, p := range ports {
		fqbn := ""
		if len(p.MatchingBoards) > 0 {
			fqbn = p.MatchingBoards[0].FQBN
		}
		boards = append(boards, BoardInfo{
			Port: p.Port.Address,
			FQBN: fqbn,
		})
	}
	return boards, nil
}

// isCoreInstalled проверяет, установлено ли ядро через JSON
func isCoreInstalled(coreName string) bool {
	cmd := exec.Command("arduino-cli", "core", "list", "--format", "json")
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	var cores []CoreInfo
	if err := json.Unmarshal(out, &cores); err != nil {
		return false
	}

	for _, core := range cores {
		if core.ID == coreName && core.Installed != "" {
			return true
		}
	}
	return false
}