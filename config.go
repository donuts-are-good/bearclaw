package main

import _ "embed"

// embedding the templates
var (
	//go:embed templates/header.html
	headerHTML string
	//go:embed templates/footer.html
	footerHTML string
)

// paths
var (
	inFolder       = "./markdown"  // your markdown articles go in here
	outFolder      = "./output"    // your rendered html will end up here
	templateFolder = "./templates" // your header and footer go here
	pluginsFolder  = "./plugins"   // your plugins go here
)

// config
var (

	// author vars
	author_name  = "@donuts-are-good"
	author_bio   = "i like Go and jelly filled pastries :)"
	author_links = []string{
		"https://github.com/donuts-are-good/",
		"https://github.com/donuts-are-good/bearclaw",
	}

	// content vars
	site_name        = "bearclaw blog"
	site_description = "a blog about a tiny static site generator in Go!"
	site_link        = "https://" + "bearclaw.blog"
	site_license     = "MIT License " + author_name + " " + site_link
)
