package logger

import (
	"encoding/json"
	"fmt"
	"os"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LogInfo(ctx *gin.Context, info interface{}) {
	// value := ctx.Value(constant.Context_Logger_Object)
	// if value == nil {
	// 	return
	// }

	// context.Set(ctx.Request, "", "")

	// ctxLogger, ok := value.(ContextLogger)
	// if ok {
	// 	ctxLogger.LogInfo(info)
	// 	ctx = context.WithValue(ctx, constant.Context_Logger_Object, ctxLogger)
	// }

	logger, exist := ctx.Get(constant.Context_Logger_Object)
	if exist && logger != nil {
		ctxLogger := logger.(ContextLogger)
		ctxLogger.LogInfo(info)
		ctx.Set(constant.Context_Logger_Object, ctxLogger)
	}

	contextName := ctx.GetString(constant.Context_ID)
	body, _ := json.Marshal(info)
	fmt.Println(contextName + " : " + string(body))
}

func PrintLog(ctx *gin.Context) {
	ctx.Set("", "")
	value, ok := ctx.Get(constant.Context_Logger_Object)
	if ok {
		logger := value.(ContextLogger)

		go logger.PrintLog()
		return
	}
}

type ContextLogger struct {
	ContextID string
	Ctx       *gin.Context
	LogStack  []interface{}
}

func CreateContextLogger(ctx *gin.Context) ContextLogger {
	return ContextLogger{
		ContextID: uuid.NewString(),
		Ctx:       ctx}
}

func (logger *ContextLogger) LogInfo(info interface{}) {
	logger.LogStack = append(logger.LogStack, info)
	// contextName := logger.Ctx.GetString(constant.Context_ID)
	// body, _ := json.Marshal(info)
	// fmt.Println(contextName + " : " + string(body))
}

func (logger *ContextLogger) PrintLog() {
	if len(logger.LogStack) < 1 {
		return
	}

	fmt.Fprintf(os.Stdout, "Red: \033[0;31m %s None: \033[0m %s", "red string", "colorless string")
	fmt.Println("Log stack : " + logger.ContextID + " - " + logger.Ctx.Request.RequestURI)
	for _, value := range logger.LogStack {
		fmt.Println(value)
	}
}
