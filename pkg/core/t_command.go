package core

// TCommand представляет собой байтовый код управляющего символа терминала.
type TCommand byte

const (
	// =========================================================================
	// ОСНОВНЫЕ УПРАВЛЯЮЩИЕ КОДЫ (СИГНАЛЫ И ЗАВЕРШЕНИЕ)
	// =========================================================================

	// TCommandETX — Прервать выполнение текущей команды (Ctrl+C / SIGINT)
	TCommandETX TCommand = 0x03

	// TCommandEOT — Конец ввода (EOF) / Выход из терминала (Ctrl+D)
	TCommandEOT TCommand = 0x04

	// TCommandSUB — Приостановить процесс и отправить в фон (Ctrl+Z / SIGTSTP)
	TCommandSUB TCommand = 0x1A

	// TCommandFS — Жестко убить процесс с дампом памяти (Ctrl+\ / SIGQUIT)
	TCommandFS TCommand = 0x1C

	// =========================================================================
	// ФОРМАТИРОВАНИЕ, СТИРАНИЕ И НАВИГАЦИЯ
	// =========================================================================

	// TCommandBS — Стереть один символ слева (Backspace / Ctrl+H)
	TCommandBS TCommand = 0x08

	// TCommandHT — Горизонтальная табуляция / Автодополнение (Tab / Ctrl+I)
	TCommandHT TCommand = 0x09

	// TCommandLF — Перевод строки / Выполнение команды (Enter / Line Feed / Ctrl+J)
	TCommandLF TCommand = 0x0A

	// TCommandFF — Очистить экран / Перевод страницы (Clear screen / Ctrl+L)
	TCommandFF TCommand = 0x0C

	// TCommandCR — Возврат каретки (Carriage Return / Ctrl+M)
	TCommandCR TCommand = 0x0D

	// TCommandESC — Начало управляющей ANSI-последовательности (Escape / Ctrl+[)
	TCommandESC TCommand = 0x1B

	// =========================================================================
	// УПРАВЛЕНИЕ ПОТОКОМ ВВОДА-ВЫВОДА (FLOW CONTROL)
	// =========================================================================

	// TCommandDC1 — Разморозить вывод терминала после Ctrl+S (XON / Ctrl+Q)
	TCommandDC1 TCommand = 0x13

	// TCommandDC3 — Заморозить вывод терминала на экран (XOFF / Ctrl+S)
	TCommandDC3 TCommand = 0x11

	// =========================================================================
	// ДОПОЛНИТЕЛЬНЫЕ СИСТЕМНЫЕ КОДЫ
	// =========================================================================

	// TCommandBEL — Звуковой сигнал терминала / Звонок (Bell / Ctrl+G)
	TCommandBEL TCommand = 0x07

	// TCommandNAK — Стереть всю текущую строку ввода (Ctrl+U)
	TCommandNAK TCommand = 0x15
)
