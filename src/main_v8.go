// uploadWithProgress выполняет загрузку с отображением статуса
func uploadWithProgress(hexPath, portPath, fqbn string, statusLabel *widget.Label) error {
	// Канал для результата
	result := make(chan error)

	// Запускаем загрузку в отдельной горутине
	go func() {
		// Сначала показываем статус через fyne.Do
		fyne.Do(func() {
			statusLabel.SetText("⏳ Подождите, идёт загрузка...")
			statusLabel.Refresh()
		})

		// Даём время на отрисовку (50 мс достаточно)
		time.Sleep(50 * time.Millisecond)

		// Выполняем команду
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

	// Ждём завершения
	err := <-result

	// Обновляем статус после завершения
	if err != nil {
		fyne.Do(func() {
			statusLabel.SetText("❌ Ошибка загрузки")
			statusLabel.Refresh()
		})
		return err
	}

	fyne.Do(func() {
		statusLabel.SetText("✅ Загрузка успешно завершена!")
		statusLabel.Refresh()
	})
	return nil
}