package bootstrap

import (
	"log"
	"os"
	"path/filepath"
)

//
func setupFileSystem() {

}

// Bootstrap...
func Bootstrap() (err error) {

	inputDir := filepath.Join(*BASEPATH, *INPUTDIR)
	outputDir := filepath.Join(*BASEPATH, *OUTPUTDIR)
	resultsSchema := filepath.Join(*BASEPATH, *RESULTSSCHEMA)

	log.Println("inputDir", inputDir)
	log.Println("outputDir", outputDir)
	log.Println("schemaDir", resultsSchema)

	err = os.RemoveAll(inputDir)
	if err != nil {
		return
	}
	err = os.RemoveAll(outputDir)
	if err != nil {
		return
	}

	err = os.MkdirAll(inputDir, os.ModePerm)
	if err != nil {
		return
	}
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return
	}

	return

}
