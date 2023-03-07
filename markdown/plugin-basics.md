# plugin basics

bearclaw supports plugins. plugins are html documents that may or may not contain their own CSS or JS. 


## yeah but what is it

think of it like a "template" you can stick anywhere. 


## how to use it

plugins should be added to the `plugins` folder as a `.zip`. bearclaw will automatically extract the zip and put it where it needs to go. 

plugins normally contain html files, so let's say you have a page-hit counter that goes up every time a new person shows up. let's call it "pagehits". your plugin would arrive as `pagehits.zip` which you'd drop in to the `bearclaw/plugins` folder. 

run bearclaw, and that zip file will turn itself into a folder with some html files inside, in this example lets say it is `today.html` and `total.html`. today.html shows you how many hits your page had today, and total.html has how many hits your page has in total.

To call a plugin in your markdown pages, use a plugin comment like this:

`<!-- plugin "./plugins/pagehits/total.html" -->`

in our pagehits example, this should auto-insert our total page hits plugin every time we run `bearclaw`. bearclaw will try and rebuild everything in `markdown` folder every time it is run.

# [<< Go back](README.md)