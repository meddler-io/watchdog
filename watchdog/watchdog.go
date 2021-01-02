package watchdog

import (
	"log"
	"os"

	"github.com/meddler-xyz/watchdog/config"
	"github.com/meddler-xyz/watchdog/executor"
)

func Start(env map[string]string) {

	// environment := make(map[string]string)
	environment := []string{
		"exec_timeout=2",
		"fprocess=sleep",
	}

	// environment["exec_timeout"] = "1000"
	// environment["fprocess"] = "ls"

	watchdogConfig := config.New(env)
	if watchdogConfig.InjectCGIHeaders {

	}

	commandName, arguments := watchdogConfig.Process()
	functionInvoker := executor.ForkFunctionRunner{
		ExecTimeout: watchdogConfig.ExecTimeout,
	}

	// commandName = "echo"
	// arguments = []string{"10"}
	req := executor.FunctionRequest{
		Process:                 commandName,
		ProcessArgs:             arguments,
		InputReader:             os.Stdin,
		OutputWriter:            os.Stdout,
		Environment:             environment,
		TractID:                 env["TraceId"],
		CurrentWorkingDirectory: env["CWD"],
	}

	err := functionInvoker.Run(req)
	if err != nil {
		log.Println(err)
	}
}