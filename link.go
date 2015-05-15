package main

import (
	"github.com/codeskyblue/go-sh"
)

/**
 *
 */
type dirSync struct {
	from, to, user, host string
}

/**
 * [func description]
 * @param  {[type]} link *dirSync)     configure(from, to string [description]
 * @return {[type]}      [description]
 */
func (link *dirSync) configure(from, to string) {

	link.from = from
	link.to = to

}

func (link *dirSync) List() {

}

func (link *dirSync) Remove() {

}

func (link *dirSync) Add() {

	msg, err := sh.Command("lsyncd", " -nodaemon", "-rsyncssh", link.from, link.user+"@"+link.host, link.to).Output()
	// Base rsync call for transferring files
	// rsync -chavzP --delete --stats ./ root@os.dev.digitalpacific.com.au:/tmp/test

	//
	// watchout(){ notify_loop /Users/work/code/fig-esm/; rsync -chavzP --stats --delete /Users/work/code/fig-esm/ root@os.dev.digitalpacific.com.au:/tmp/test && watchout; }; watchout > /dev/null &

	lsyncd[OPTIONS] - rsyncssh[SOURCE][HOST][TARGETDIR]
	// Convert byte array to string
	output = string(msg[:])

}
