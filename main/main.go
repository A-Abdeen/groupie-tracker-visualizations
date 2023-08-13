package main

import (
	gt "gt/webapp/API"
	handlers "gt/webapp/handlers"
	server "gt/webapp/server"
)

func main() {
	(gt.FindArtist(4, "Relations"))
	(gt.FindArtist(4, "Locations"))
	(gt.FindArtist(4, "Dates"))
	(gt.Relations(4))
	(gt.Locations(4))
	(gt.Dates(4))
	// Initialize html & tmpl global array for all handlers
	if handlers.HtmlTmpl == nil {
		handlers.Init()
	}
	server.Server()
}
