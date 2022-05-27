package log

func Logger() *logType {
	return &logType{}
	//return &logType{fields:  make(map[string]interface{})}
}
