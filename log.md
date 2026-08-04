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

```
2026/08/04 20:06:26 *** This application has not been migrated to the fyne.Do threading model ***
2026/08/04 20:06:26 *** The next major Fyne release will remove this safety! ***
2026/08/04 20:06:26 *** Read more at https://docs.fyne.io/started/goroutines ***
2026/08/04 20:06:48 *** Error in Fyne call thread, this should have been called in fyne.Do[AndWait] ***
2026/08/04 20:06:48 *** This error needs to be addressed as it could cause concurrent errors soon ***
2026/08/04 20:06:48 *** The next major Fyne release will disable this safety by default! ***
2026/08/04 20:06:48   From: ./main_v6.go:213
SIGABRT: abort
PC=0x495041 m=0 sigcode=0

goroutine 0 gp=0x1b75b00 m=0 mp=0x1b76ce0 [idle]:
runtime.futex(0x1b76e38, 0x80, 0x0, 0x0, 0x0, 0x0)
        runtime/sys_linux_amd64.s:569 +0x21 fp=0x7fff23bb7978 sp=0x7fff23bb7970 pc=0x495041
runtime.futexsleep(0x7fff23bb79f0?, 0x1b75b00?, 0x7fff23bb79f0?)
        runtime/os_linux.go:73 +0x30 fp=0x7fff23bb79c8 sp=0x7fff23bb7978 pc=0x4505b0
runtime.notesleep(0x1b76e38)
        runtime/lock_futex.go:47 +0x87 fp=0x7fff23bb7a00 sp=0x7fff23bb79c8 pc=0x426407
runtime.mPark(...)
        runtime/proc.go:1967
runtime.stoplockedm()
        runtime/proc.go:3263 +0x73 fp=0x7fff23bb7a58 sp=0x7fff23bb7a00 pc=0x45c273
runtime.schedule()
        runtime/proc.go:4143 +0x3a fp=0x7fff23bb7a98 sp=0x7fff23bb7a58 pc=0x45e7fa
runtime.park_m(0x2482dadec1e0)
        runtime/proc.go:4304 +0x285 fp=0x7fff23bb7af8 sp=0x7fff23bb7a98 pc=0x45ed05
runtime.mcall()
        runtime/asm_amd64.s:496 +0x55 fp=0x7fff23bb7b10 sp=0x7fff23bb7af8 pc=0x4918d5

goroutine 1 gp=0x2482dadec1e0 m=nil [chan send, 7 minutes, locked to thread]:
runtime.gopark(0x13d5fe0?, 0x2482dca75d40?, 0xc8?, 0x75?, 0x51487b?)
        runtime/proc.go:462 +0xce fp=0x2482dce8f5a0 sp=0x2482dce8f580 pc=0x48bd8e
runtime.chansend(0x2482dbb922a0, 0xd76a10, 0x1, 0xae967b?)
        runtime/chan.go:283 +0x3fc fp=0x2482dce8f610 sp=0x2482dce8f5a0 pc=0x41f07c
runtime.chansend1(0x2482dca75d40?, 0x0?)
        runtime/chan.go:161 +0x17 fp=0x2482dce8f640 sp=0x2482dce8f610 pc=0x41ec77
main.uploadWithProgress({0x2482dc4448a0?, 0xd4b4a8?}, {0x2482dcafb5ef?, 0x6fe3cb?}, {0xd2c365?, 0xa46eaf?}, 0x2482daf2bcc0, 0x2482db04c690)
        ./main_v6.go:221 +0x265 fp=0x2482dce8f740 sp=0x2482dce8f640 pc=0xae96a5
main.main.func3()
        ./main_v6.go:328 +0xfb fp=0x2482dce8f7e8 sp=0x2482dce8f740 pc=0xaea7db
fyne.io/fyne/v2/widget.(*Button).Tapped(0x2482db2aa280, 0xd00400?)
        fyne.io/fyne/v2@v2.8.0/widget/button.go:214 +0x9f fp=0x2482dce8f830 sp=0x2482dce8f7e8 pc=0x9b9e5f
fyne.io/fyne/v2/internal/driver/glfw.(*window).mouseClickedHandleTapDoubleTap(0x2482db1f4000, {0xd8df60, 0x2482db2aa280}, 0x2482dce50040)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/window.go:608 +0x182 fp=0x2482dce8f860 sp=0x2482dce8f830 pc=0xaa1ac2
fyne.io/fyne/v2/internal/driver/glfw.(*window).processMouseClicked(0x2482db1f4000, 0x1, 0x0, 0x0)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/window.go:569 +0x895 fp=0x2482dce8f928 sp=0x2482dce8f860 pc=0xaa1115
fyne.io/fyne/v2/internal/driver/glfw.(*window).mouseClicked(0x2482dafa14d0?, 0x2482db083280?, 0xa53600?, 0x1b72d20?, 0x2482dae37968?)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/window_desktop.go:452 +0xaf fp=0x2482dce8f958 sp=0x2482dce8f928 pc=0xaa53ef
fyne.io/fyne/v2/internal/driver/glfw.(*window).mouseClicked-fm(0x2482db1edd80?, 0x420600?, 0x0?, 0x0?)
        <autogenerated>:1 +0x26 fp=0x2482dce8f990 sp=0x2482dce8f958 pc=0xaa9e66
github.com/go-gl/glfw/v3.4/glfw.goMouseButtonCB(0x2482dae379d8?, 0x0, 0x0, 0x0)
        github.com/go-gl/glfw/v3.4/glfw@v0.1.0-pre.1.0.20260707082822-2a407d02d01a/input.go:341 +0x4e fp=0x2482dce8f9c0 sp=0x2482dce8f990 pc=0xa4fe8e
_cgoexp_aab3bd21ad33_goMouseButtonCB(0x248200000008?)
        _cgo_gotypes.go:2874 +0x25 fp=0x2482dce8f9e8 sp=0x2482dce8f9c0 pc=0xa56445
runtime.cgocallbackg1(0xa56420, 0x7fff23bb7b80, 0x0)
        runtime/cgocall.go:466 +0x2e5 fp=0x2482dce8faa0 sp=0x2482dce8f9e8 pc=0x41db25
runtime.cgocallbackg(0xa56420, 0x7fff23bb7b80, 0x0)
        runtime/cgocall.go:362 +0x130 fp=0x2482dce8fb08 sp=0x2482dce8faa0 pc=0x41d750
runtime.cgocallbackg(0xa56420, 0x7fff23bb7b80, 0x0)
        <autogenerated>:1 +0x29 fp=0x2482dce8fb30 sp=0x2482dce8fb08 pc=0x4958c9
runtime.cgocallback(0x2482dce8fb90, 0x488d95, 0xb71690)
        runtime/asm_amd64.s:1160 +0xcc fp=0x2482dce8fb58 sp=0x2482dce8fb30 pc=0x4932ec
runtime.systemstack_switch()
        runtime/asm_amd64.s:516 +0x8 fp=0x2482dce8fb68 sp=0x2482dce8fb58 pc=0x4918e8
runtime.cgocall(0xb71690, 0x2482dce8fbc8)
        runtime/cgocall.go:185 +0x75 fp=0x2482dce8fba0 sp=0x2482dce8fb68 pc=0x488d95
github.com/go-gl/glfw/v3.4/glfw._Cfunc_glfwPollEvents()
        _cgo_gotypes.go:1840 +0x3a fp=0x2482dce8fbc8 sp=0x2482dce8fba0 pc=0xa4c37a
github.com/go-gl/glfw/v3.4/glfw.PollEvents()
        github.com/go-gl/glfw/v3.4/glfw@v0.1.0-pre.1.0.20260707082822-2a407d02d01a/window.go:1010 +0x13 fp=0x2482dce8fbf0 sp=0x2482dce8fbc8 pc=0xa56333
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).pollEvents(...)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/loop_desktop.go:40
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).runGL(0x2482db983d10?)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/loop.go:159 +0x1f2 fp=0x2482dce8fcd8 sp=0x2482dce8fbf0 pc=0xa9b792
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).Run(0x2482db1dc4d0)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/driver.go:195 +0x72 fp=0x2482dce8fd20 sp=0x2482dce8fcd8 pc=0xa998f2
fyne.io/fyne/v2/app.(*fyneApp).Run(0x2482db1b6600)
        fyne.io/fyne/v2@v2.8.0/app/app.go:87 +0x232 fp=0x2482dce8fdb8 sp=0x2482dce8fd20 pc=0xad5dd2
fyne.io/fyne/v2/internal/driver/glfw.(*window).ShowAndRun(0x2482db1f4000)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/window.go:226 +0x64 fp=0x2482dce8fdd8 sp=0x2482dce8fdb8 pc=0xa9f0c4
main.main()
        ./main_v6.go:370 +0xce2 fp=0x2482dce8ff48 sp=0x2482dce8fdd8 pc=0xaea642
runtime.main()
        runtime/proc.go:290 +0x2d5 fp=0x2482dce8ffe0 sp=0x2482dce8ff48 pc=0x456d55
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dce8ffe8 sp=0x2482dce8ffe0 pc=0x493521

goroutine 2 gp=0x2482dadecd20 m=nil [force gc (idle), 2 minutes]:
runtime.gopark(0x6c54c7b74c9?, 0x6831681360c?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482dae52fa8 sp=0x2482dae52f88 pc=0x48bd8e
runtime.goparkunlock(...)
        runtime/proc.go:468
runtime.forcegchelper()
        runtime/proc.go:375 +0xb3 fp=0x2482dae52fe0 sp=0x2482dae52fa8 pc=0x457073
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae52fe8 sp=0x2482dae52fe0 pc=0x493521
created by runtime.init.7 in goroutine 1
        runtime/proc.go:363 +0x1a

goroutine 3 gp=0x2482daded2c0 m=nil [GC sweep wait]:
runtime.gopark(0x1?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482dae53788 sp=0x2482dae53768 pc=0x48bd8e
runtime.goparkunlock(...)
        runtime/proc.go:468
runtime.bgsweep(0x2482dae7a000)
        runtime/mgcsweep.go:324 +0x151 fp=0x2482dae537c8 sp=0x2482dae53788 pc=0x440cf1
runtime.gcenable.gowrap1()
        runtime/mgc.go:214 +0x17 fp=0x2482dae537e0 sp=0x2482dae537c8 pc=0x432057
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae537e8 sp=0x2482dae537e0 pc=0x493521
created by runtime.gcenable in goroutine 1
        runtime/mgc.go:214 +0x66

goroutine 4 gp=0x2482daded4a0 m=nil [GC scavenge wait, 2 minutes]:
runtime.gopark(0x3b9fec2d?, 0x3b9aca00?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482dae53f78 sp=0x2482dae53f58 pc=0x48bd8e
runtime.goparkunlock(...)
        runtime/proc.go:468
runtime.(*scavengerState).park(0x1b752c0)
        runtime/mgcscavenge.go:425 +0x49 fp=0x2482dae53fa8 sp=0x2482dae53f78 pc=0x43e769
runtime.bgscavenge(0x2482dae7a000)
        runtime/mgcscavenge.go:658 +0x59 fp=0x2482dae53fc8 sp=0x2482dae53fa8 pc=0x43ecf9
runtime.gcenable.gowrap2()
        runtime/mgc.go:215 +0x17 fp=0x2482dae53fe0 sp=0x2482dae53fc8 pc=0x432017
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae53fe8 sp=0x2482dae53fe0 pc=0x493521
created by runtime.gcenable in goroutine 1
        runtime/mgc.go:215 +0xa5

goroutine 18 gp=0x2482daf043c0 m=nil [GOMAXPROCS updater (idle), 7 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482dae4e788 sp=0x2482dae4e768 pc=0x48bd8e
runtime.goparkunlock(...)
        runtime/proc.go:468
runtime.updateMaxProcsGoroutine()
        runtime/proc.go:7095 +0xe7 fp=0x2482dae4e7e0 sp=0x2482dae4e788 pc=0x4655e7
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae4e7e8 sp=0x2482dae4e7e0 pc=0x493521
created by runtime.defaultGOMAXPROCSUpdateEnable in goroutine 1
        runtime/proc.go:7083 +0x37

goroutine 19 gp=0x2482daf04960 m=nil [finalizer wait, 7 minutes]:
runtime.gopark(0x466635?, 0x1c8?, 0x80?, 0xab?, 0x2482dae52601?)
        runtime/proc.go:462 +0xce fp=0x2482dae52620 sp=0x2482dae52600 pc=0x48bd8e
runtime.runFinalizers()
        runtime/mfinal.go:210 +0x107 fp=0x2482dae527e0 sp=0x2482dae52620 pc=0x430fc7
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae527e8 sp=0x2482dae527e0 pc=0x493521
created by runtime.createfing in goroutine 1
        runtime/mfinal.go:172 +0x3d

goroutine 20 gp=0x2482daf04b40 m=nil [cleanup wait, 7 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482dae4ef68 sp=0x2482dae4ef48 pc=0x48bd8e
runtime.goparkunlock(...)
        runtime/proc.go:468
runtime.(*cleanupQueue).dequeue(0x1b75780)
        runtime/mcleanup.go:522 +0xd4 fp=0x2482dae4efa0 sp=0x2482dae4ef68 pc=0x42dd14
runtime.runCleanups()
        runtime/mcleanup.go:718 +0x45 fp=0x2482dae4efe0 sp=0x2482dae4efa0 pc=0x42e385
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae4efe8 sp=0x2482dae4efe0 pc=0x493521
created by runtime.(*cleanupQueue).createGs in goroutine 1
        runtime/mcleanup.go:672 +0xa5

goroutine 21 gp=0x2482daf04d20 m=nil [select, 7 minutes]:
runtime.gopark(0x2482dae4f760?, 0x2?, 0x0?, 0x0?, 0x2482dae4f6e4?)
        runtime/proc.go:462 +0xce fp=0x2482dae4f560 sp=0x2482dae4f540 pc=0x48bd8e
runtime.selectgo(0x2482dae4f760, 0x2482dae4f6e0, 0x0?, 0x0, 0x0?, 0x1)
        runtime/select.go:351 +0xaa5 fp=0x2482dae4f690 sp=0x2482dae4f560 pc=0x46a125
fyne.io/fyne/v2/internal/async.(*UnboundedChan[...]).processing(0xd92480)
        fyne.io/fyne/v2@v2.8.0/internal/async/chan.go:43 +0xd9 fp=0x2482dae4f7c0 sp=0x2482dae4f690 pc=0xaa7fd9
fyne.io/fyne/v2/internal/async.NewUnboundedChan[...].gowrap1()
        fyne.io/fyne/v2@v2.8.0/internal/async/chan.go:20 +0x1b fp=0x2482dae4f7e0 sp=0x2482dae4f7c0 pc=0xaa7edb
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae4f7e8 sp=0x2482dae4f7e0 pc=0x493521
created by fyne.io/fyne/v2/internal/async.NewUnboundedChan[...] in goroutine 1
        fyne.io/fyne/v2@v2.8.0/internal/async/chan.go:20 +0x10a

goroutine 22 gp=0x2482daf04f00 m=nil [IO wait, 7 minutes]:
runtime.gopark(0x469570?, 0x1?, 0x0?, 0x0?, 0xb?)
        runtime/proc.go:462 +0xce fp=0x2482dc9afcc8 sp=0x2482dc9afca8 pc=0x48bd8e
runtime.netpollblock(0x4e6698?, 0x41cd86?, 0x0?)
        runtime/netpoll.go:575 +0xf7 fp=0x2482dc9afd00 sp=0x2482dc9afcc8 pc=0x44f857
internal/poll.runtime_pollWait(0x7faf21dd3200, 0x72)
        runtime/netpoll.go:351 +0x85 fp=0x2482dc9afd20 sp=0x2482dc9afd00 pc=0x48af65
internal/poll.(*pollDesc).wait(0x2482db1d17a0?, 0x2482dc9afe53?, 0x1)
        internal/poll/fd_poll_runtime.go:84 +0x27 fp=0x2482dc9afd48 sp=0x2482dc9afd20 pc=0x506447
internal/poll.(*pollDesc).waitRead(...)
        internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Read(0x2482db1d17a0, {0x2482dc9afe53, 0x10000, 0x10000})
        internal/poll/fd_unix.go:165 +0x2ae fp=0x2482dc9afde0 sp=0x2482dc9afd48 pc=0x50766e
os.(*File).read(...)
        os/file_posix.go:30
os.(*File).Read(0x2482db1ee108, {0x2482dc9afe53?, 0x2482dc9afe53?, 0x1?})
        os/file.go:144 +0x4f fp=0x2482dc9afe20 sp=0x2482dc9afde0 pc=0x51306f
github.com/fsnotify/fsnotify.(*inotify).readEvents(0x2482db15e8c0)
        github.com/fsnotify/fsnotify@v1.9.0/backend_inotify.go:357 +0xcf fp=0x2482dc9bffc8 sp=0x2482dc9afe20 pc=0xad2d0f
github.com/fsnotify/fsnotify.newBackend.gowrap1()
        github.com/fsnotify/fsnotify@v1.9.0/backend_inotify.go:155 +0x17 fp=0x2482dc9bffe0 sp=0x2482dc9bffc8 pc=0xad18f7
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dc9bffe8 sp=0x2482dc9bffe0 pc=0x493521
created by github.com/fsnotify/fsnotify.newBackend in goroutine 1
        github.com/fsnotify/fsnotify@v1.9.0/backend_inotify.go:155 +0x1f6

goroutine 23 gp=0x2482daf050e0 m=nil [chan receive, 7 minutes]:
runtime.gopark(0x2482daf122a0?, 0x2482daf52a80?, 0x0?, 0x0?, 0x2482dae506e8?)
        runtime/proc.go:462 +0xce fp=0x2482dae50690 sp=0x2482dae50670 pc=0x48bd8e
runtime.chanrecv(0x2482daf122a0, 0x2482dae507a0, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482dae50708 sp=0x2482dae50690 pc=0x42006e
runtime.chanrecv2(0x2482db1d8580?, 0x0?)
        runtime/chan.go:514 +0x12 fp=0x2482dae50730 sp=0x2482dae50708 pc=0x41fbb2
fyne.io/fyne/v2/app.watchFile.func1()
        fyne.io/fyne/v2@v2.8.0/app/settings_desktop.go:43 +0x77 fp=0x2482dae507e0 sp=0x2482dae50730 pc=0xada8b7
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae507e8 sp=0x2482dae507e0 pc=0x493521
created by fyne.io/fyne/v2/app.watchFile in goroutine 1
        fyne.io/fyne/v2@v2.8.0/app/settings_desktop.go:42 +0xe9

goroutine 24 gp=0x2482daf052c0 m=nil [select, 7 minutes]:
runtime.gopark(0x2482dae50f78?, 0x2?, 0x0?, 0x0?, 0x2482dae50efc?)
        runtime/proc.go:462 +0xce fp=0x2482db292d78 sp=0x2482db292d58 pc=0x48bd8e
runtime.selectgo(0x2482db292f78, 0x2482dae50ef8, 0x0?, 0x0, 0x0?, 0x1)
        runtime/select.go:351 +0xaa5 fp=0x2482db292ea8 sp=0x2482db292d78 pc=0x46a125
fyne.io/fyne/v2/internal/async.(*UnboundedChan[...]).processing(0xd923e0)
        fyne.io/fyne/v2@v2.8.0/internal/async/chan.go:43 +0xdc fp=0x2482db292fc0 sp=0x2482db292ea8 pc=0x706e7c
fyne.io/fyne/v2/internal/async.NewUnboundedChan[...].gowrap1()
        fyne.io/fyne/v2@v2.8.0/internal/async/chan.go:20 +0x1b fp=0x2482db292fe0 sp=0x2482db292fc0 pc=0x70775b
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db292fe8 sp=0x2482db292fe0 pc=0x493521
created by fyne.io/fyne/v2/internal/async.NewUnboundedChan[...] in goroutine 1
        fyne.io/fyne/v2@v2.8.0/internal/async/chan.go:20 +0x10a

goroutine 5 gp=0x2482db29c3c0 m=nil [GC worker (idle), 7 minutes]:
runtime.gopark(0x661bf71ff79?, 0x0?, 0xb8?, 0xf?, 0x53f04c?)
        runtime/proc.go:462 +0xce fp=0x2482dae50f40 sp=0x2482dae50f20 pc=0x48bd8e
runtime.gcBgMarkWorker(0x2482daf129a0)
        runtime/mgc.go:1791 +0xeb fp=0x2482dae50fc8 sp=0x2482dae50f40 pc=0x434b2b
runtime.gcBgMarkStartWorkers.gowrap1()
        runtime/mgc.go:1695 +0x17 fp=0x2482dae50fe0 sp=0x2482dae50fc8 pc=0x434a17
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae50fe8 sp=0x2482dae50fe0 pc=0x493521
created by runtime.gcBgMarkStartWorkers in goroutine 1
        runtime/mgc.go:1695 +0x105

goroutine 6 gp=0x2482db29c5a0 m=nil [chan receive, 7 minutes]:
runtime.gopark(0x0?, 0xc146c0?, 0x0?, 0x0?, 0x2482dbaaffa8?)
        runtime/proc.go:462 +0xce fp=0x2482db290f18 sp=0x2482db290ef8 pc=0x48bd8e
runtime.chanrecv(0x2482db2fc380, 0x0, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482db290f90 sp=0x2482db290f18 pc=0x42006e
runtime.chanrecv1(0x2482db122b90?, 0x0?)
        runtime/chan.go:509 +0x12 fp=0x2482db290fb8 sp=0x2482db290f90 pc=0x41fb92
github.com/godbus/dbus/v5.newConn.func1()
        github.com/godbus/dbus/v5@v5.2.2/conn.go:303 +0x2c fp=0x2482db290fe0 sp=0x2482db290fb8 pc=0xa6a66c
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db290fe8 sp=0x2482db290fe0 pc=0x493521
created by github.com/godbus/dbus/v5.newConn in goroutine 33
        github.com/godbus/dbus/v5@v5.2.2/conn.go:302 +0x4da

goroutine 66 gp=0x2482daf054a0 m=5 mp=0x2482daf00008 [syscall, 7 minutes]:
runtime.notetsleepg(0x1bbb0a0, 0xffffffffffffffff)
        runtime/lock_futex.go:123 +0x29 fp=0x2482dbab3fa0 sp=0x2482dbab3f78 pc=0x4266e9
os/signal.signal_recv()
        runtime/sigqueue.go:152 +0x98 fp=0x2482dbab3fc0 sp=0x2482dbab3fa0 pc=0x48e218
os/signal.loop()
        os/signal/signal_unix.go:23 +0x13 fp=0x2482dbab3fe0 sp=0x2482dbab3fc0 pc=0xa5bb93
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dbab3fe8 sp=0x2482dbab3fe0 pc=0x493521
created by os/signal.Notify.func1.1 in goroutine 50
        os/signal/signal.go:152 +0x1f

goroutine 27 gp=0x2482daf05680 m=nil [GC worker (idle), 2 minutes]:
runtime.gopark(0x6c54d2f7187?, 0x1?, 0x16?, 0x84?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482db2dc740 sp=0x2482db2dc720 pc=0x48bd8e
runtime.gcBgMarkWorker(0x2482daf129a0)
        runtime/mgc.go:1791 +0xeb fp=0x2482db2dc7c8 sp=0x2482db2dc740 pc=0x434b2b
runtime.gcBgMarkStartWorkers.gowrap1()
        runtime/mgc.go:1695 +0x17 fp=0x2482db2dc7e0 sp=0x2482db2dc7c8 pc=0x434a17
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2dc7e8 sp=0x2482db2dc7e0 pc=0x493521
created by runtime.gcBgMarkStartWorkers in goroutine 1
        runtime/mgc.go:1695 +0x105

goroutine 28 gp=0x2482daf05860 m=nil [GC worker (idle), 7 minutes]:
runtime.gopark(0x66110e24caf?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482db2dcf40 sp=0x2482db2dcf20 pc=0x48bd8e
runtime.gcBgMarkWorker(0x2482daf129a0)
        runtime/mgc.go:1791 +0xeb fp=0x2482db2dcfc8 sp=0x2482db2dcf40 pc=0x434b2b
runtime.gcBgMarkStartWorkers.gowrap1()
        runtime/mgc.go:1695 +0x17 fp=0x2482db2dcfe0 sp=0x2482db2dcfc8 pc=0x434a17
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2dcfe8 sp=0x2482db2dcfe0 pc=0x493521
created by runtime.gcBgMarkStartWorkers in goroutine 1
        runtime/mgc.go:1695 +0x105

goroutine 29 gp=0x2482daf05a40 m=nil [GC worker (idle), 7 minutes]:
runtime.gopark(0x661bf71df59?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482db2dd740 sp=0x2482db2dd720 pc=0x48bd8e
runtime.gcBgMarkWorker(0x2482daf129a0)
        runtime/mgc.go:1791 +0xeb fp=0x2482db2dd7c8 sp=0x2482db2dd740 pc=0x434b2b
runtime.gcBgMarkStartWorkers.gowrap1()
        runtime/mgc.go:1695 +0x17 fp=0x2482db2dd7e0 sp=0x2482db2dd7c8 pc=0x434a17
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2dd7e8 sp=0x2482db2dd7e0 pc=0x493521
created by runtime.gcBgMarkStartWorkers in goroutine 1
        runtime/mgc.go:1695 +0x105

goroutine 30 gp=0x2482dadedc20 m=nil [chan receive, 7 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482db2e2f08 sp=0x2482db2e2ee8 pc=0x48bd8e
runtime.chanrecv(0x2482db0070a0, 0x2482db2e2fb8, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482db2e2f80 sp=0x2482db2e2f08 pc=0x42006e
runtime.chanrecv2(0x0?, 0x0?)
        runtime/chan.go:514 +0x12 fp=0x2482db2e2fa8 sp=0x2482db2e2f80 pc=0x41fbb2
fyne.io/fyne/v2/internal/app.(*Lifecycle).RunEventQueue(...)
        fyne.io/fyne/v2@v2.8.0/internal/app/lifecycle.go:110
fyne.io/fyne/v2/app.(*fyneApp).Run.gowrap1()
        fyne.io/fyne/v2@v2.8.0/app/app.go:76 +0x5f fp=0x2482db2e2fe0 sp=0x2482db2e2fa8 pc=0xad5f7f
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2e2fe8 sp=0x2482db2e2fe0 pc=0x493521
created by fyne.io/fyne/v2/app.(*fyneApp).Run in goroutine 1
        fyne.io/fyne/v2@v2.8.0/app/app.go:76 +0xcf

goroutine 31 gp=0x2482db29cb40 m=nil [IO wait, 7 minutes]:
runtime.gopark(0x28ef310800?, 0x0?, 0x0?, 0x0?, 0xb?)
        runtime/proc.go:462 +0xce fp=0x2482dba8fcc8 sp=0x2482dba8fca8 pc=0x48bd8e
runtime.netpollblock(0x4e6698?, 0x41cd86?, 0x0?)
        runtime/netpoll.go:575 +0xf7 fp=0x2482dba8fd00 sp=0x2482dba8fcc8 pc=0x44f857
internal/poll.runtime_pollWait(0x7faf21dd3000, 0x72)
        runtime/netpoll.go:351 +0x85 fp=0x2482dba8fd20 sp=0x2482dba8fd00 pc=0x48af65
internal/poll.(*pollDesc).wait(0x2482dae88d80?, 0x2482dba8fe53?, 0x1)
        internal/poll/fd_poll_runtime.go:84 +0x27 fp=0x2482dba8fd48 sp=0x2482dba8fd20 pc=0x506447
internal/poll.(*pollDesc).waitRead(...)
        internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Read(0x2482dae88d80, {0x2482dba8fe53, 0x10000, 0x10000})
        internal/poll/fd_unix.go:165 +0x2ae fp=0x2482dba8fde0 sp=0x2482dba8fd48 pc=0x50766e
os.(*File).read(...)
        os/file_posix.go:30
os.(*File).Read(0x2482dadea108, {0x2482dba8fe53?, 0x0?, 0x2000000000000000?})
        os/file.go:144 +0x4f fp=0x2482dba8fe20 sp=0x2482dba8fde0 pc=0x51306f
github.com/fsnotify/fsnotify.(*inotify).readEvents(0x2482db368b40)
        github.com/fsnotify/fsnotify@v1.9.0/backend_inotify.go:357 +0xcf fp=0x2482dba9ffc8 sp=0x2482dba8fe20 pc=0xad2d0f
github.com/fsnotify/fsnotify.newBackend.gowrap1()
        github.com/fsnotify/fsnotify@v1.9.0/backend_inotify.go:155 +0x17 fp=0x2482dba9ffe0 sp=0x2482dba9ffc8 pc=0xad18f7
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dba9ffe8 sp=0x2482dba9ffe0 pc=0x493521
created by github.com/fsnotify/fsnotify.newBackend in goroutine 1
        github.com/fsnotify/fsnotify@v1.9.0/backend_inotify.go:155 +0x1f6

goroutine 32 gp=0x2482db29cd20 m=nil [chan receive, 7 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482db2e3e90 sp=0x2482db2e3e70 pc=0x48bd8e
runtime.chanrecv(0x2482db9bd0a0, 0x2482db2e3fa0, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482db2e3f08 sp=0x2482db2e3e90 pc=0x42006e
runtime.chanrecv2(0x0?, 0x0?)
        runtime/chan.go:514 +0x12 fp=0x2482db2e3f30 sp=0x2482db2e3f08 pc=0x41fbb2
fyne.io/fyne/v2/app.watchFile.func1()
        fyne.io/fyne/v2@v2.8.0/app/settings_desktop.go:43 +0x77 fp=0x2482db2e3fe0 sp=0x2482db2e3f30 pc=0xada8b7
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2e3fe8 sp=0x2482db2e3fe0 pc=0x493521
created by fyne.io/fyne/v2/app.watchFile in goroutine 1
        fyne.io/fyne/v2@v2.8.0/app/settings_desktop.go:42 +0xe9

goroutine 33 gp=0x2482db29cf00 m=nil [chan receive, 5 minutes]:
runtime.gopark(0x2482db221d50?, 0x3?, 0xb8?, 0x6f?, 0x2482db2b2cc0?)
        runtime/proc.go:462 +0xce fp=0x2482dae66e70 sp=0x2482dae66e50 pc=0x48bd8e
runtime.chanrecv(0x2482dbb92000, 0x2482dae66f78, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482dae66ee8 sp=0x2482dae66e70 pc=0x42006e
runtime.chanrecv2(0x2482dbc00100?, 0x17?)
        runtime/chan.go:514 +0x12 fp=0x2482dae66f10 sp=0x2482dae66ee8 pc=0x41fbb2
github.com/rymdport/portal/settings.OnSignalSettingChanged(0x2482dae66fc0)
        github.com/rymdport/portal@v0.4.2/settings/changed.go:23 +0x67 fp=0x2482dae66f90 sp=0x2482dae66f10 pc=0xaaaa47
fyne.io/fyne/v2/app.watchTheme.func1()
        fyne.io/fyne/v2@v2.8.0/app/app_xdg.go:144 +0xa5 fp=0x2482dae66fe0 sp=0x2482dae66f90 pc=0xad7325
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae66fe8 sp=0x2482dae66fe0 pc=0x493521
created by fyne.io/fyne/v2/app.watchTheme in goroutine 1
        fyne.io/fyne/v2@v2.8.0/app/app_xdg.go:136 +0x4f

goroutine 50 gp=0x2482db29d0e0 m=nil [chan receive, 7 minutes]:
runtime.gopark(0xa5b81a?, 0xc520a0?, 0x1?, 0x8a?, 0x2482db2de760?)
        runtime/proc.go:462 +0xce fp=0x2482db2de6d0 sp=0x2482db2de6b0 pc=0x48bd8e
runtime.chanrecv(0x2482dae7eee0, 0x0, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482db2de748 sp=0x2482db2de6d0 pc=0x42006e
runtime.chanrecv1(0x2482dae7eee0?, 0x2482db2de798?)
        runtime/chan.go:509 +0x12 fp=0x2482db2de770 sp=0x2482db2de748 pc=0x41fb92
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).catchTerm(0x2482db1dc4d0)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/driver_desktop.go:248 +0x74 fp=0x2482db2de7c8 sp=0x2482db2de770 pc=0xa9ad34
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).Run.gowrap1()
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/driver.go:194 +0x17 fp=0x2482db2de7e0 sp=0x2482db2de7c8 pc=0xa99a57
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2de7e8 sp=0x2482db2de7e0 pc=0x493521
created by fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).Run in goroutine 1
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/driver.go:194 +0x68

goroutine 51 gp=0x2482db29d2c0 m=nil [select, 7 minutes, locked to thread]:
runtime.gopark(0x2482db2defa8?, 0x2?, 0x1?, 0xd2?, 0x2482db2def94?)
        runtime/proc.go:462 +0xce fp=0x2482db2dee28 sp=0x2482db2dee08 pc=0x48bd8e
runtime.selectgo(0x2482db2defa8, 0x2482db2def90, 0x0?, 0x0, 0x0?, 0x1)
        runtime/select.go:351 +0xaa5 fp=0x2482db2def58 sp=0x2482db2dee28 pc=0x46a125
runtime.ensureSigM.func1()
        runtime/signal_unix.go:1091 +0x188 fp=0x2482db2defe0 sp=0x2482db2def58 pc=0x486668
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482db2defe8 sp=0x2482db2defe0 pc=0x493521
created by runtime.ensureSigM in goroutine 50
        runtime/signal_unix.go:1074 +0xc5

goroutine 7 gp=0x2482dbb00000 m=nil [IO wait, 5 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        runtime/proc.go:462 +0xce fp=0x2482dbc07a70 sp=0x2482dbc07a50 pc=0x48bd8e
runtime.netpollblock(0x0?, 0x41cd86?, 0x0?)
        runtime/netpoll.go:575 +0xf7 fp=0x2482dbc07aa8 sp=0x2482dbc07a70 pc=0x44f857
internal/poll.runtime_pollWait(0x7faf21dd2a00, 0x72)
        runtime/netpoll.go:351 +0x85 fp=0x2482dbc07ac8 sp=0x2482dbc07aa8 pc=0x48af65
internal/poll.(*pollDesc).wait(0x2482db082d80?, 0x2482dbb80000?, 0x0)
        internal/poll/fd_poll_runtime.go:84 +0x27 fp=0x2482dbc07af0 sp=0x2482dbc07ac8 pc=0x506447
internal/poll.(*pollDesc).waitRead(...)
        internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).ReadMsg(0x2482db082d80, {0x2482dbb80000, 0x10, 0x10}, {0x2482dbb82028, 0x1000, 0x1000}, 0x40000000)
        internal/poll/fd_unix.go:295 +0x37d fp=0x2482dbc07bd8 sp=0x2482dbc07af0 pc=0x50873d
net.(*netFD).readMsg(0x2482db082d80, {0x2482dbb80000?, 0x2482dbc07cb8?, 0x41f41f?}, {0x2482dbb82028?, 0x40?, 0x2482dbc07e50?}, 0x2482db29cf00?)
        net/fd_posix.go:91 +0x31 fp=0x2482dbc07c60 sp=0x2482dbc07bd8 pc=0x5b98d1
net.(*UnixConn).readMsg(0x2482db29a160, {0x2482dbb80000?, 0x0?, 0x2?}, {0x2482dbb82028?, 0x42c23e?, 0x2482dbc07d38?})
        net/unixsock_posix.go:115 +0x3d fp=0x2482dbc07cf0 sp=0x2482dbc07c60 pc=0x5d2a9d
net.(*UnixConn).ReadMsgUnix(0x2482db29a160, {0x2482dbb80000?, 0x2482dbb00000?, 0x2482daf01008?}, {0x2482dbb82028?, 0x2482dbb92000?, 0x2482dbb92070?})
        net/unixsock.go:143 +0x36 fp=0x2482dbc07d68 sp=0x2482dbc07cf0 pc=0x5d15d6
github.com/godbus/dbus/v5.(*oobReader).Read(0x2482dbb82008, {0x2482dbb80000?, 0xa771a7?, 0x2482dbc07e30?})
        github.com/godbus/dbus/v5@v5.2.2/transport_unix.go:40 +0x3c fp=0x2482dbc07de0 sp=0x2482dbc07d68 pc=0xa8233c
io.ReadAtLeast({0xd7e2e0, 0x2482dbb82008}, {0x2482dbb80000, 0x10, 0x10}, 0x10)
        io/io.go:335 +0x8e fp=0x2482dbc07e28 sp=0x2482dbc07de0 pc=0x4a0fae
io.ReadFull(...)
        io/io.go:354
github.com/godbus/dbus/v5.(*unixTransport).ReadMessage(0x2482db99cf90)
        github.com/godbus/dbus/v5@v5.2.2/transport_unix.go:123 +0x225 fp=0x2482dbc07f58 sp=0x2482dbc07e28 pc=0xa82a45
github.com/godbus/dbus/v5.(*Conn).inWorker(0x2482db04c2d0)
        github.com/godbus/dbus/v5@v5.2.2/conn.go:390 +0x37 fp=0x2482dbc07fc8 sp=0x2482dbc07f58 pc=0xa6ab97
github.com/godbus/dbus/v5.(*Conn).Auth.gowrap1()
        github.com/godbus/dbus/v5@v5.2.2/auth.go:118 +0x17 fp=0x2482dbc07fe0 sp=0x2482dbc07fc8 pc=0xa68897
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dbc07fe8 sp=0x2482dbc07fe0 pc=0x493521
created by github.com/godbus/dbus/v5.(*Conn).Auth in goroutine 33
        github.com/godbus/dbus/v5@v5.2.2/auth.go:118 +0x7c7

goroutine 61 gp=0x2482dbaac3c0 m=nil [chan receive, 7 minutes]:
runtime.gopark(0x2482dbab1e28?, 0x2482dae7e070?, 0x78?, 0xbe?, 0x2482dbab1c98?)
        runtime/proc.go:462 +0xce fp=0x2482dae62c48 sp=0x2482dae62c28 pc=0x48bd8e
runtime.chanrecv(0x2482dbab68c0, 0x0, 0x1)
        runtime/chan.go:667 +0x4ae fp=0x2482dae62cc0 sp=0x2482dae62c48 pc=0x42006e
runtime.chanrecv1(0x13e9d60?, 0xd81ae0?)
        runtime/chan.go:509 +0x12 fp=0x2482dae62ce8 sp=0x2482dae62cc0 pc=0x41fb92
fyne.io/fyne/v2/internal/driver/glfw.runOnMainWithWait(0x2482dcbccff0, 0x1)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/loop.go:53 +0xfc fp=0x2482dae62d50 sp=0x2482dae62ce8 pc=0xa9b25c
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).DoFromGoroutine.func1()
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/driver.go:103 +0x1c fp=0x2482dae62d70 sp=0x2482dae62d50 pc=0xa994bc
fyne.io/fyne/v2/internal/async.EnsureNotMain(0x2482dcbcd010)
        fyne.io/fyne/v2@v2.8.0/internal/async/goroutine.go:29 +0x42 fp=0x2482dae62df0 sp=0x2482dae62d70 pc=0x6fc8e2
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).DoFromGoroutine(0xae9860?, 0x2482dcbccff0, 0x60?)
        fyne.io/fyne/v2@v2.8.0/internal/driver/glfw/driver.go:102 +0x56 fp=0x2482dae62e10 sp=0x2482dae62df0 pc=0xa99456
fyne.io/fyne/v2.DoAndWait(0x2482dcbccff0)
        fyne.io/fyne/v2@v2.8.0/thread.go:9 +0x3a fp=0x2482dae62e38 sp=0x2482dae62e10 pc=0x6e5e3a
fyne.io/fyne/v2/internal/async.EnsureMain(0x2482dcbccff0)
        fyne.io/fyne/v2@v2.8.0/internal/async/goroutine.go:57 +0x16c fp=0x2482dae62eb8 sp=0x2482dae62e38 pc=0x6fcccc
fyne.io/fyne/v2/internal/driver/common.(*Canvas).Refresh(0x2482db1949c0, {0xd8e080?, 0x2482db2ac018?})
        fyne.io/fyne/v2@v2.8.0/internal/driver/common/canvas.go:282 +0x66 fp=0x2482dae62ee0 sp=0x2482dae62eb8 pc=0xa47806
fyne.io/fyne/v2/canvas.Refresh({0xd8e080, 0x2482db2ac018})
        fyne.io/fyne/v2@v2.8.0/canvas/canvas.go:27 +0x89 fp=0x2482dae62f18 sp=0x2482dae62ee0 pc=0x7a9de9
fyne.io/fyne/v2/canvas.(*Rectangle).Refresh(...)
        fyne.io/fyne/v2@v2.8.0/canvas/rectangle.go:76
fyne.io/fyne/v2/widget.(*progressRenderer).Refresh(0x2482db2ac000)
        fyne.io/fyne/v2@v2.8.0/widget/progressbar.go:92 +0x58 fp=0x2482dae62f58 sp=0x2482dae62f18 pc=0x9dfeb8
fyne.io/fyne/v2/widget.(*BaseWidget).Refresh(0x2482dae7e1c0?)
        fyne.io/fyne/v2@v2.8.0/widget/widget.go:127 +0x52 fp=0x2482dae62f78 sp=0x2482dae62f58 pc=0x9f3132
fyne.io/fyne/v2/widget.(*ProgressBar).SetValue(...)
        fyne.io/fyne/v2@v2.8.0/widget/progressbar.go:126
main.uploadWithProgress.func1()
        ./main_v6.go:213 +0x5c fp=0x2482dae62fe0 sp=0x2482dae62f78 pc=0xae98bc
runtime.goexit({})
        runtime/asm_amd64.s:1771 +0x1 fp=0x2482dae62fe8 sp=0x2482dae62fe0 pc=0x493521
created by main.uploadWithProgress in goroutine 1
        ./main_v6.go:207 +0x23b

rax    0xca
rbx    0x0
rcx    0x495043
rdx    0x0
rdi    0x1b76e38
rsi    0x80
rbp    0x7fff23bb79b8
rsp    0x7fff23bb7970
r8     0x0
r9     0x0
r10    0x0
r11    0x286
r12    0x7fff23bb7998
r13    0x2482dce56030
r14    0x1b75b00
r15    0xffffffffffffffff
rip    0x495041
rflags 0x286
cs     0x33
fs     0x0
gs     0x0
```