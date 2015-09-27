/*

*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const processFileName string = ".links"

// Setup the properties required to sync directories
type dirSync struct {
	from string
	to   string
}

/**
 * Configure the link properties
 * @param  {[type]} link *dirSync)     configure(from, to string [description]
 * @return {[type]}      [description]
 */
func (link *dirSync) Configure(from, to string) {
	link.from = from
	link.to = to
}

/**
 * List all current links
 * @param  {[type]} link *dirSync)     List( [description]
 * @return {[type]}      [description]
 */
func (link *dirSync) List() {

	// Try to read a current process file
	linkData, err := ioutil.ReadFile(processFileName)

	// Soft error as it may not exists
	if err != nil {
		color.Red("No links available")
		os.Exit(1)
	}

	links := strings.Split(string(linkData[:]), "\n")

	fmt.Println("\tStatus\t\tPID\tFrom\tTo")

	// Display links
	for index, value := range links {

		var status string
		currentLink := strings.Split(value, " ")

		// Check if process is active
		if active := checkProcessActive(currentLink[0]); active {

			status = color.GreenString("Active")

		} else {

			status = color.RedString("Inactive")

		}

		fmt.Printf("[%d]\t%s\t%s\t%s\t%s\n", index, status, currentLink[0], currentLink[1], currentLink[2])

	}

}

/**
 * Remove a link
 * @param  {[type]} link *dirSync)     Remove( [description]
 * @return {[type]}      [description]
 */
func (link *dirSync) Remove() {

	// First show current links
	link.List()

	// Prompt user to choose which link to remove
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the link number you want to remove: ")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Validate that the link reference is correct and get the pid
	linkIndex, _ := strconv.Atoi(input)
	currentLink, err := getLinkByIndex(linkIndex)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if active := checkProcessActive(currentLink["pid"]); active == true {

		// Remove a link
		_, err := sh.Command("kill", currentLink["pid"]).Output()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// TODO: Remove link from file

	color.Green(fmt.Sprintf("Successfully removed link from:%s to:%s", currentLink["from"], currentLink["to"]))

}

/**
 * Add a new link
 * @param  {[type]} link *dirSync)     Add( [description]
 * @return {[type]}      [description]
 */
func (link *dirSync) Add() {

	// TODO: If a mac user check for notify_loop and any dependencies

	// Create new link
	newLink, err := launchLink(link.from, link.to)

	if err != nil {
		color.Red(fmt.Sprintf("Link could not be created: '%s'", err))
		os.Exit(1)
	}

	// Convert byte array to string
	processId := string(newLink[:])

	// Add new process to link file
	err = saveLinkProcess(processId, link.from, link.to)

	if err != nil {
		color.Red("Link could not be saved for later use. It was created however.")
		os.Exit(1)
	}
}

/**
 * Enable a folder sync
 * @param  {[type]} link *dirSync)     Up( [description]
 * @return nil
 */
func (link *dirSync) Up() {

	// First show current links
	link.List()

	// Prompt user to choose which link to enable
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the link number you want to enable: ")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	linkIndex, _ := strconv.Atoi(input)
	linkData, err := getLinkByIndex(linkIndex)

	// Make sure the link isnt active
	if active := checkProcessActive(linkData["pid"]); active {
		color.Yellow("Link is already active")
		os.Exit(1)
	}

	// Launch the process
	newLink, err := launchLink(linkData["from"], linkData["to"])

	if err != nil {
		color.Red(fmt.Sprintf("Link could not be initialised: '%s'", err))
		os.Exit(1)
	}

	// Convert byte array to string
	processId := string(newLink[:])

	err = updateLinkProcess(processId, linkIndex)

	if err != nil {
		color.Red("Link could not be saved for later use. It was created however.")
		os.Exit(1)
	}

	color.Green(fmt.Sprintf("Successfully enabled link from:%s to:%s", linkData["from"], linkData["to"]))

}

/**
 * Disable a folder sync
 * @param  {[type]} link *dirSync)     Down( [description]
 * @return nil
 */
func (link *dirSync) Down() {

	// First show current links
	link.List()

	// Prompt user to choose which link to remove
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the link number you want to disable: ")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get link
	linkIndex, _ := strconv.Atoi(input)
	linkData, err := getLinkByIndex(linkIndex)

	// Make sure the link is active
	if active := checkProcessActive(linkData["pid"]); !active {
		color.Yellow("Link is already inactive")
		os.Exit(1)
	}

	// We need to kill the process but leave the data in the file
	_, err = sh.Command("kill", linkData["pid"]).Output()

	if err != nil {
		color.Red("Could not disable the link. Try again or otherwise logout/reboot to reset.")
		os.Exit(1)
	}

	color.Green(fmt.Sprintf("Successfully disabled link from:%s to:%s", linkData["from"], linkData["to"]))
	color.Green("You can re-enable the link by running kubefactory link up")
}

/**
 * Initialise/re-initialise a link
 * @param  {[type]} from [description]
 * @param  {[type]} to   string)       (msg, err [description]
 * @return {[type]}      [description]
 */
func launchLink(from, to string) (msg []byte, err error) {

	msg, err = sh.Command("bash", "watchout(){", "notify_loop", from+";", "rsync", "-chavzP", "--stats", "--delete", from, to, "&&", "watchout;", "watchout", ">", "/dev/null", "&").Command("awk", "{print $2}").Output()

	return

	// watchout(){ notify_loop /Users/work/code/fig-esm/; rsync -chavzP --stats --delete /Users/work/code/fig-esm/ root@os.dev.digitalpacific.com.au:/tmp/test && watchout; }; watchout > /dev/null &

	// Base rsync call for transferring files
	// rsync -chavzP --delete --stats ./ root@os.dev.digitalpacific.com.au:/tmp/test
	// msg, err := sh.Command("lsyncd", " -nodaemon", "-rsyncssh", link.from, link.user+"@"+link.host, link.to).Output()

}

/**
 * Updates the processId of a link and saves to file
 * @param  {string} pid
 * @param  {int} linkIndex
 * @return {error} err
 */
func updateLinkProcess(pid string, linkIndex int) (err error) {

	// Get the link file
	links, err := getLinks()

	// Find the index
	for index, _ := range links {

		if index != linkIndex {
			continue
		}

		// Update the pid with the one given
		links[linkIndex]["pid"] = pid

	}

	return nil

}

/**
 * Writes an array of links to file
 * @param  {[]map[string]string} links
 * @return {error} err
 */
func writeLinkFile(links []map[string]string) (err error) {

	var linksData string

	for _, link := range links {

		linksData += fmt.Sprintf("%s %s %s\n", link["pid"], link["from"], link["to"])

	}

	// Write process file
	err = ioutil.WriteFile(processFileName, []byte(linksData), 0600)

	if err != nil {
		return errors.New("Couldnt write to file " + processFileName)
	}

	return nil

}

/**
 * Check to see if a given processId is active
 * @param  {string} pid
 * @return {bool}
 */
func checkProcessActive(pid string) bool {

	msg, err := sh.Command("ps", "-p", pid).Command("grep", "bash").Command("awk", "$1").Output()

	if err != nil {
		return false
	}

	output := string(msg[:])

	if output != pid {
		return false
	}

	return true

}

/**
 * Get an array of all links available
 * @param  {[type]} ) (links        []map[string]string, err error [description]
 * @return {[type]}   [description]
 */
func getLinks() (links []map[string]string, err error) {

	// Try to read a current process file
	linkData, err := ioutil.ReadFile(processFileName)

	if err != nil {
		return nil, errors.New("No links available")
	}

	lines := strings.Split(string(linkData[:]), "\n")

	// Build links
	for _, value := range lines {

		currentLink := strings.Split(value, " ")

		fmt.Println(currentLink)

		links = append(links, map[string]string{"pid": currentLink[0], "from": currentLink[1], "to": currentLink[2]})

	}

	return links, nil

}

/**
 * Get a link by index
 * @param  {[type]} index int)          (pid int, err error [description]
 * @return {[type]}       [description]
 */
func getLinkByIndex(index int) (pid map[string]string, err error) {

	links, _ := getLinks()

	for key, link := range links {
		if key != index {
			continue
		}

		return link, nil
	}

	return nil, errors.New("Index could not be found")
}

/**
 * Save the process id from a link for later use
 * @param  {[type]} processId string)       (err error [description]
 * @return {[type]}           [description]
 * TODO: Update to use objects for saving data
 */
func saveLinkProcess(processId, from, to string) (err error) {

	var processIds string

	newLink := fmt.Sprintf("%s %s %s\n", processId, from, to)

	// Try to read a current process file
	dat, err := ioutil.ReadFile(processFileName)

	// Soft error as it may not exists
	if err != nil {
		processIds = processId
	} else {
		processIds = string(dat[:]) + newLink
	}

	// Rewrite process file
	err = ioutil.WriteFile(processFileName, []byte(processIds), 0644)

	if err != nil {
		return errors.New("Couldnt write to file " + processFileName)
	}

	return nil

}

/**
 * Handle an error ungracefully
 * @param  {error} e error
 * @return nil
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}
