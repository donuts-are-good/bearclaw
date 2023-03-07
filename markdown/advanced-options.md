# advanced options

one-size-fits-all is generally okay, but not everyone has the same tastes. these are some advanced options that made sense to add to help out different setups.

you can view these by running bearclaw like this:

`bearclaw --help` or `bearclaw -h`

```
╭── ~/Projects/bearclaw 
╰─$ ./bearclaw --help

Usage of ./bearclaw:
  -input string
    	the input folder for markdown files (default "./markdown")
  -output string
    	the output folder for html files (default "./output")
  -plugins string
    	the plugins folder for plugins (default "./plugins")
  -templates string
    	the templates folder for header and footer html files (default "./templates")
  -watch
    	watch the content directories for changes
```

## details

### `--input`

example: `./bearclaw --input path/to/your/markdown/files`
this is the folder where your markdown files go

### `--output`

example: `./bearclaw --output path/to/your/html/folder`
this is the folder where your html files end up

### `--plugins`

example: `./bearclaw --plugins path/to/your/bearclaw/plugins`
this is the folder where your plugins go

### `--templates`

example: `./bearclaw --templates path/to/your/header/and/footer/files`
this is the folder where your templates go if you wanted to add others for theme switching.

### `--watch`

example: `./bearclaw --watch`
watch your templates and markdown folders for changes. 

# [<< Go back](https://github.com/donuts-are-good/bearclaw/blob/master/markdown/README.md)