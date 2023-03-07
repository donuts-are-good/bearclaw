# writing plugins

plugins are just zip folders that contain html files. these html files can have inline CSS and JS.

- structure your files to be self-contained, meaning they should not rely on resources to be transported with them. 
- they should only consist of html documents with all necessary scripts and styles inlined, or remotely referenced. it's okay to add documentation but don't count on those files being distributed with the html.
- for predictable behavior, use names with no spaces or special characters
- each plugin should be encapsulated by a `<div>` for clear separations between plugin content and bearclaw content. 
- don't forget your plugin.json, it looks like this:

```
{
    "plugin_name": "blog-elements",
    "plugin_version": "1.0.0",
    "plugin_author": "@donuts-are-good",
    "plugin_description": "this plugin is a test of the bearclaw plugin syste. it contains download.html which will insert a bar that contains links to download bearclaw and bearclaw source code.",
    "plugin_license":"MIT License, @donuts-are-good",
    "plugin_link": "https://github.com/donuts-are-good/blog-elements" 
}
```

structure your plugin zips like this:

[socialIconsPlugin.zip]
- plugin.json
- youtube-subscribe.html
- github-fork.html
- twitter-follow.html

don't use an inner folder like this:

[socialIconsPlugin.zip]
- socialIconsPlugin/
    - plugin.json
    - youtube-subscribe.html
    - github-fork.html
    - twitter-follow.html

# [<< Go back](https://github.com/donuts-are-good/bearclaw/blob/master/markdown/README.md)