# using site.config
when you start bearclaw for the first time, it will walk you through a few questions and do its best to create a config gile. the config file doesn't do much now, but the plan is to use the config file to auto insert these variable values where they make sense in your template some day.

*you can see this example in the `site.config.example` file.*


```
# hello!

# this is your site's config file
# think of it like a list of settings

# if you have questions or need help, come to 
# https://github.com/donuts-are-good/bearclaw

# if you haven't run bearclaw yet, when you run it for the first time,
# it will check for a file like this one called site.config and if that
# file doesn't exist, it will walk you through a few questions and then
# generate it for you!

# you can put anything here, I put my github username 
author_name: @donuts-are-good

# this is a small description about you, the blog author 
author_bio: bearclaw author

# this is a comma-separated list of your links/sites. it can be just one if you want.
author_links: https://github.com/donuts-are-good/,https://github.com/donuts-are-good/bearclaw

# this is the name of your site 
site_name: bearclaw blog

# this is a description of your site 
site_description: a blog about a tiny static site generator in Go!

# this is the url of the blog you're setting up 
site_link: https://bearclaw.blog

# this is the license of the content on your blog 
# it might not seem important, but it can get very serious if 
# you don't know much about it.

# bearclaw is MIT licensed, but most people will want to put 
# some type of copyright here unless they have an open-source blog 
site_license: MIT License
```

# [<< Go back](https://github.com/donuts-are-good/bearclaw/blob/master/markdown/README.md)