package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
)

type Logger struct {
	Log           map[string]interface{}
	headerInfoMsg map[string]interface{}
}

func NewLogger() *Logger {
	return &Logger{
		Log: map[string]interface{}{},
	}
}

func (l *Logger) AddHeaderInfo(msg map[string]interface{}) {
	if l.headerInfoMsg == nil {
		l.headerInfoMsg = make(map[string]interface{}, len(msg))
	}
	for k, v := range msg {
		switch k {
		case "level", "time", "cursor", "function", "message":
			continue
		}
		l.headerInfoMsg[k] = v
	}
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Fatal(msg interface{}, format ...interface{}) {
	l.log(msg, "FATAL", format)
	panic("Fatal error")
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Error(msg interface{}, format ...interface{}) {
	l.log(msg, "ERROR", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Warn(msg interface{}, format ...interface{}) {
	l.log(msg, "WARN", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Info(msg interface{}, format ...interface{}) {
	l.log(msg, "INFO", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Debug(msg interface{}, format ...interface{}) {
	l.log(msg, "DEBUG", format)
}

func (l *Logger) JsonParseOut(input interface{}) {
	makeMsg := make(map[string]interface{})
	setFields(input, &makeMsg)

	makeMsg["TimeStamp"] = time.Now().Format("2006/01/02 15:04:05")

	result, err := json.Marshal(makeMsg)
	if err != nil {
		panic(err)
	}
	res := string(result)
	fmt.Println(res)
}

func (l *Logger) log(msg interface{}, logLevel string, variableStr []interface{}) {
	output := map[string]interface{}{
		"level":    logLevel,
		"time":     time.Now().Format(time.RFC3339),
		"cursor":   createCursor(),
		"function": createFunctionName(),
	}
	defer fin(output)

	// printf系の処理
	typedMsg, ok := msg.(string)
	if ok {
		output["message"] = fmt.Sprintf(typedMsg, variableStr...)
		if len(l.headerInfoMsg) > 0 {
			output["information"] = l.headerInfoMsg
		}
		return
	}

	// errorの出力
	_, ok = msg.(error)
	if ok {
		output["message"] = fmt.Sprintf("%+v", msg)
		if len(l.headerInfoMsg) > 0 {
			output["information"] = l.headerInfoMsg
		}
		return
	}
	// jsonに変換できる場合の処理
	if len(l.headerInfoMsg) > 0 {
		for k, v := range l.headerInfoMsg {
			output[k] = v
		}
	}
	output["message"] = msg
}
func fin(msg map[string]interface{}) {
	switch msg["level"] {
	case "FATAL", "ERROR", "WARN":
		fmt.Fprintln(os.Stderr, jsonParse(msg))
	case "INFO", "DEBUG":
		fmt.Fprintln(os.Stdout, jsonParse(msg))
	}
}

func setFields(input interface{}, output *map[string]interface{}) {
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("input is not a struct")
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("json")
		if fieldName == "" {
			fieldName = typ.Field(i).Name
		}
		(*output)[fieldName] = field.Interface()
	}
}
