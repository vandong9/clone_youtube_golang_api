package log

import (
	"encoding/json"
	"fmt"

	"com.vandong9.clone_youtube_golang_api/common/constant"
	"github.com/gin-gonic/gin"
)

func LogInfo(ctx *gin.Context, info interface{}) {
	contextName := ctx.GetString(constant.Context_ID)
	body, _ := json.Marshal(info)
	fmt.Println(contextName + " : " + string(body))
}
