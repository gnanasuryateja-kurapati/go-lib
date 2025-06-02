package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/gnanasuryateja-kurapati/go-lib/constants"
// 	simplelogger "github.com/gnanasuryateja-kurapati/go-lib/logger/simpleLogger"
// 	"github.com/gnanasuryateja-kurapati/go-lib/utils"
// )

// func main() {

// 	lggr, err := simplelogger.NewSimpleLogger(
// 		simplelogger.SimpleLoggerParams{
// 			ServiceName: "",
// 			LogLevel:    utils.StringToStringPtr(constants.LOG_LEVEL_INFO),
// 			Env:         "dev",
// 		},
// 	)
// 	if err != nil {
// 		fmt.Println("error initializing logger: ", err)
// 		return
// 	} else {
// 		lggr.Debug(context.Background(), "debug log should be printed now")
// 		lggr.Info(context.Background(), "info log always prints")
// 		return
// 	}

// }
