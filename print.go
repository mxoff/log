package log

func (l *logType) Println(message string) {
	l.types = "INFO"
	l.color = colors["white"]
	l.message = message
	l.print()
}

func (l *logType) Errorln(message string) {
	l.types = "ERRO"
	l.color = colors["red"]
	l.message = message
	l.print()
}

func (l *logType) Warningln(message string) {
	l.types = "WARN"
	l.color = colors["yellow"]
	l.message = message
	l.print()
}

func (l *logType) Debugln(message string) {
	l.types = "DEBU"
	l.color = colors["gray"]
	l.message = message
	l.print()
}

func (l *logType) ErrorNoNil(err error) {
	if err != nil {
		l.types = "ERRO"
		l.color = colors["red"]
		l.message = err.Error()
		l.print()
	}
}
