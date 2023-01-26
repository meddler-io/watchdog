package main

import (
	"github.com/meddler-io/watchdog/logger"
	"github.com/meddler-io/watchdog/producer"
)

func main() {

	// go sendData()
	// go sendData()
	// go sendData()
	sendData()
	sendData()
	// go sendData2()
	// go sendData3()

}

func produce() {
	sendData()
}

func sendData() {

	err := producer.Produce(

		"user",
		"bitnami",
		"172.24.42.51",
		"tasks_test",

		`
  { 

	"config": {

		"system": {
			"base_path": "/tmp/test",
			"input_dir":  "input_dirs",
			"output_dir":  "output_dirs",
			"exec_timeout": "10" ,
			"git_mode" : "false",
			"git_auth_mode" :"token",
			"git_remote" : "https://github.com/cyclops-org/hawkeye-api.git",
			"git_auth_username": "studiogangster",
			"git_auth_password": "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"


		} ,

		"process": {
			"test": "variable"

		}

	},

	"substitute_var": true,
	"variables": {
		"input_dir" : "",
		"output_dir": "$output_dir",
		"git_path": "$git_path"
	},


	"cmd": ["echo"],
	"args": [  "HelloWorld $output_dir $git_path"   ],

	"id": "outputbucket" ,

	"environ": 
	{

		"exec_timeout": "10" ,  
		"$INPUTDIR": "inpudir",
		"TraceId":"5fde15c7ed17c3374c56990e" ,
		"fprocess":"echo \\$fprocess $fprocess $BASEPATH "  

	} ,
		
	"dependencies": [





	]
  }`)

	logger.Println(err)

}
