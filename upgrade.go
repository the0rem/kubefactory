package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"os"
)

func Upgrade(currentRC, filename, upgradeInterval string) {

	// Run the upgrade
	msg, err := sh.Command("kubectl", "rolling-update", currentRC, "--filename="+filename, "--update-period="+upgradeInterval).Output()

	if err != nil {
		color.Red(fmt.Sprintf("Couldn't perform the upgrade: %s", err))
		os.Exit(1)
	}

	color.Green(string(msg[:]))

}
