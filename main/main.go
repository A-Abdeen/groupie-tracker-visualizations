package main

import (
	gt "gt/webapp/API"
	server "gt/webapp/server"
)

func main() {
	(gt.FindArtist(4, "Relations"))
	(gt.FindArtist(4, "Locations"))
	(gt.FindArtist(4, "Dates"))
	(gt.Relations(4))
	(gt.Locations(4))
	(gt.Dates(4))
	server.Server()
}
