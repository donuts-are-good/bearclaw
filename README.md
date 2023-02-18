![bearclaw](https://user-images.githubusercontent.com/96031819/218302524-121cd81a-b552-45e5-b46e-5689bbf08390.png)
# bearclaw
a tiny static site generator, written in Go

![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

## what's a static site generator
a static site is a site with no fancy clicky things, signups, comments, just plain html. writing a raw html page for your blog makes it load very fast, but there's a lot of typing involved when designing it. markdown is a simpler and faster way to write pages, and bearclaw converts those markdown files to html for you, with your style template. 

no node-modules, no react, no fancy stuff or cool emojis. it just works.

## how do we use it?
bearclaw can be run on-demand, or it can rebuild automatically when it sees changes. there are 3 folders: `markdown`, `output`, and `templates`.

- **markdown** - all your new posts go here
- **output** - bearclaw puts all of your html here
- **templates** - header.html and footer.html

that's it! point your webserver at `output` or handle it however is best for your case.

**tip:** you can run `bearclaw` and it will run once, or you can use `./bearclaw --watch` to watch the current folder for changes.

## issues

if you run in to issues, there's a short bug report form you can fill out, or you can contribute with a pull-request.

## screenshot

![image](https://user-images.githubusercontent.com/96031819/218305635-75bdf421-e412-4b90-9f4a-26947219bf51.png)

## greetz

the Dozens, code-cartel, offtopic-gophers, the garrison, and the monster beverage company.

## license

this code uses the MIT license, not that anybody cares. If you don't know, then don't sweat it.

made with ☕ by 🍩 😋 donuts-are-good


## donate

If you would like to be an official energy drink sponsor of this project, you can contribute however you like.

**Bitcoin**: `bc1qg72tguntckez8qy2xy4rqvksfn3qwt2an8df2n`

**Monero**: `42eCCGcwz5veoys3Hx4kEDQB2BXBWimo9fk3djZWnQHSSfnyY2uSf5iL9BBJR5EnM7PeHRMFJD5BD6TRYqaTpGp2QnsQNgC` 

😆👏 Thanks
