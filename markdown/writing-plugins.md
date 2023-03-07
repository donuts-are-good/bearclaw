# writing plugins

plugins are just zip folders that contain html files. these html files can have inline CSS and JS.

- structure your files to be self-contained, meaning they should not rely on resources to be transported with them. 
- they should only consist of html documents with all necessary scripts and styles inlined, or remotely referenced. it's okay to add documentation but don't count on those files being distributed with the html.
- for predictable behavior, use names with no spaces or special characters
- each plugin should be encapsulated by a `<div>` for clear separations between plugin content and bearclaw content. 

structure your plugin zips like this:

[socialIconsPlugin.zip]
- youtube-subscribe.html
- github-fork.html
- twitter-follow.html

don't use an inner folder like this:

[socialIconsPlugin.zip]
- socialIconsPlugin/
    - youtube-subscribe.html
    - github-fork.html
    - twitter-follow.html

# [<< Go back](README.md)