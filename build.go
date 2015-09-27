package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Save required diretory locations for methods
type builder struct {
	destination    string
	templateSource string
	envSource      string
	replaceMarker  string
	ignoreDirs     []string
}

/**
 * Saves the required configuration options
 * @param  {[type]} build *builder)     Configure(destination, templateSource, envSource string [description]
 * @return {[type]}       [description]
 */
func (build *builder) Configure(destination, templateSource, envSource, replaceMarker string) {

	build.ignoreDirs = append(build.ignoreDirs, "partials")
	build.replaceMarker = replaceMarker

	// Ensure all required directory entries exist
	directories := []string{destination, templateSource, envSource}
	errCount := 0

	for _, directory := range directories {

		color.White(fmt.Sprintf("Checking if '%s' exists \r\n", directory))

		file, err := exists(directory)

		if file == false || err != nil {
			color.Yellow(fmt.Sprintf("The file '%s' could not be found \r\n", directory))
			errCount++
		} else {
			color.Green(fmt.Sprintf("Found '%s'", directory))
		}

		// Handle the destination directory for dynamic creation
		if file == false && directory == destination {

			color.White(fmt.Sprintf("Lets create '%s'", directory))
			err = os.MkdirAll(destination, 0755)

			if err != nil {
				color.Red(fmt.Sprintf("Could not create directory '%s'", destination))
				os.Exit(1)
			}

			color.Green(fmt.Sprintf("Successfully created '%s'", destination))
			errCount--
		}

	}

	// Cancel processing
	if errCount > 0 {
		color.Red(fmt.Sprint("We cannot continue without these directories"))
		os.Exit(1)
	}

	build.destination = destination
	build.templateSource = templateSource
	build.envSource = envSource

}

/**
 * Build the template files for the environment
 * @param  {[type]} build *builder)     Build( [description]
 * @return {[type]}       [description]
 */
func (build *builder) Build() {

	color.White("Checking for template files")

	// Set variables to point to folders
	templateFiles, _ := ioutil.ReadDir(build.templateSource)
	envFiles, _ := ioutil.ReadDir(build.envSource)

	build.BuildTemplatesFromFolder(templateFiles, build.templateSource, "")
	build.BuildTemplatesFromFolder(envFiles, build.envSource, "")

}

/**
 * [func description]
 * @param  {[type]} build *builder)     func BuildTemplatesFromFolder(templateFiles []os.FileInfo [description]
 * @return {[type]}       [description]
 */
func (build *builder) BuildTemplatesFromFolder(templateFiles []os.FileInfo, baseDir, subDir string) {

fileLoop:
	// Loop through template files for processing
	for _, file := range templateFiles {

		// If file is dir, recurse function
		if file.IsDir() {

			for _, ignore := range build.ignoreDirs {
				if file.Name() == ignore {
					continue fileLoop
				}
			}

			// Create the directory
			err = os.MkdirAll(build.destination+subDir+file.Name(), 0755)

			if err != nil {
				color.Red(fmt.Sprintf("Could not create directory '%s'", build.destination+subDir+file.Name()))
				os.Exit(1)
			}

			templateFiles, _ := ioutil.ReadDir(baseDir + subDir + file.Name())
			build.BuildTemplatesFromFolder(templateFiles, baseDir, "/"+file.Name())
			continue

		}

		// Skip if file is not a .yaml
		if filepath.Ext(file.Name()) != ".yaml" {

			color.Yellow(fmt.Sprintf("File %s is not a template", file.Name()))
			continue

		}

		// Copy template file to dist
		filename := baseDir + subDir + file.Name()
		distFilename := build.destination + file.Name()
		// err := CopyFile(filename, build.destination)
		msg, err := sh.Command("cp", "-rfv", filename, build.destination).Output()

		if err != nil {
			color.Yellow(fmt.Sprintf("Could not copy %s to dist folder, %s", filename, err))
			continue
		} else {
			color.Green(fmt.Sprintf("%s", msg))
		}

		// Translate the template file
		parseResult, err := build.ParseTemplate(distFilename)

		if err != nil {
			color.Yellow(fmt.Sprintf("%s", err.Error()))
		}

		color.White(fmt.Sprintf("%s", parseResult))

	}

}

/**
 * Get a template file and translate any content using environment files
 * @param {[type]} template       os.FileInfo [description]
 * @param {[type]} environmentDir string)     (contents     string, err error [description]
 */
func (build *builder) ParseTemplate(filename string) (output string, err error) {

	msg, err := sh.Command("python2.7", "yamlthingy.py", "--templatefile", filename, "--marker", build.replaceMarker, "--envdir", build.envSource+"partials/").Output()

	// Convert byte array to string
	output = string(msg[:])

	return output, err
}
