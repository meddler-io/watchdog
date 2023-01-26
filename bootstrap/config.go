package bootstrap

// DependencySchema
type DependencySchema struct {
	Identifier string `json:"id"`
	Alias      string `json:"alias"`
	Type       string `json:"asset"`
}

// MessageDataSpec
type MessageDataSpec struct {
	Dependencies []DependencySchema `json:"dependencies"`
}

// type MessageConfig struct {
// 	System   SystemConstants   `json:"system"`
// 	Process  ProcessConstants  `json:"process"`
// 	Reserved ReservedConstants `json:"reserved"`
// }

// MessageSpec...
type MessageSpec struct {
	MessageDataSpec
	Identifier string `json:"id"` // For changing status, ingesting data, persisting FS on Storage (Bucket Name)
	// SystemEnviron map[string]string `json:"system_environ"` // SystemEnviron. Variables to inject to Watchdog & override for a particular task
	Environ             map[string]string `json:"environ"`                                 // Environ. Variables to inject / override inside the actial process
	Entrypoint          []string          `json:"entrypoint"`                              // Override entrypoint
	Cmd                 []string          `json:"cmd" validate:"required,cmd"`             // Override CMD
	Args                []string          `json:"args" validate:"required,args"`           // Override ARGS
	SubstituteVariables bool              `json:"substitute_var" validate:"required,args"` // Parse ARGS for variables
	Variables           map[string]string `json:"variables"`                               // TODO: Variables: Replace placeholders with actual value

	Config Constants `json:"config"` // TODO: Config:

	SuccessEndpoint string `json:"success_endpoint" validate:"required,success_endpoint"` // success_endpoint
	FailureEndpoint string `json:"failure_endpoint" validate:"required,failure_endpoint"` // success_endpoint

}

type TaskResultBase struct {
	Status          string `json:"exec_status" validate:"required"`
	Message         string `json:"message" validate:"required"`
	WatchdogVersion string `json:"watchdog_version" validate:"required"`
	Identifier      string `json:"identifier" validate:"required"`
}

type TaskResult struct {
	TaskResultBase
	Response string `json:"response" ` // success_endpoint
}
