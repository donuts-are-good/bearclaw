# bearclaw origins

*hello, i'm [@donuts-are-good](https://github.com/donuts-are-good), the author of bearclaw. hopefully this document answers some of the questions about how bearclaw got made. if you have any questions, just ask.*

bearclaw happened because of challenge 2023 (i do some type of year-long programming challenge like this every year), which in short is a self-imposed challenge to push 1 fully formed project every week.  it just made sense to want to live-blog about the projects and how they're going along the way, if anything just to look back at the end of the year and see what can be learned.

the blog never got made, but bearclaw did in hopes that some day a blog would happen, as of this time, Mar 6th 2023, donuts-are-good.com is still a 404, and the completed project count is approaching 20. it's a very lucky thing that there wasn't a stipulation that a project have complete documentation or branding package to be considered complete :)

## but why bearclaw?

normally html with something like bootstrap is quick enough to write by hand, and i know the bootstrap side well, so it made sense that whatever blog i come up with wouldn't be on some platform somewhere, i'd just set up my own site and host it on github pages or something. as I was writing the html for an article, the main page, and a list of posts, i was thinking of how easy it used to be to use server-side includes to add 'template' stuff to your website.

```
<!--#include virtual="/cgi-bin/counter.pl" --> 
```
*server side includes: this is how you'd insert a page hit counter, for example, in the late 17th century*

**imagine this scenario:** you have a header, a footer, and a menu and you have this code on every blog post. now imagine you want to add an item to your menu. unless you're running react or something that shares code, you're going to be modifying every one of those html files to update your menu.

**initially, bearclaw was going to be just that. `header.html`, `footer.html`, and `menu.html`.**

those changes took very little time to put together, so rather than actually start writing articles, i figured id make bearclaw this week's project. fast forward to today, it's still going, so here we are.

# [<< Go back](README.md)