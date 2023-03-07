# code styleguide

*above all, realize others will do things differently and that's fine.*

## the rule of thirds...
...isn't a rule really.

bearclaw code normally takes shape as something representing what's called 'the rule of thirds'. it's not a rule, and it doesn't always apply, but generally when it makes sense, i'll try to do this:

- 1 line of comment
- 1 line of code
- 1 line of whitespace

i do this to help me, because i find it makes code more readable. 

not doing this doesn't mean your code isn't good, many people have many preferences about whitespace. conversely, writing your code this way doesn't magically make it good either. good code doesn't try to be fancy, it just works.

the rule of thirds is important to me because above all, code should be accessible and easily understood by those who don't necessarily have business looking at it.

## good comments are descriptive

*nobody is expected to be mark twain in the comments on the first pass. do your best to eventually add them or ask someone to help you add them.*

a good comment explains like a story what is happening. we dont use crazy words and fancy lingo, we just say in plain english, using normal words, what is happening. the author of bearclaw is motivated, not educated. nobody is born with a CS degree.

if you want a suggestion, try to eventually describe your functions as 'name will do ... with ... and return ...' when you can. these function descriptions help lesser beings like myself see what things do in our code editors.

```
// thisFunctionHere will take a thatThing and an alsoThis to create a coolThing and returns coolThing and an error if the coolThing isn't cool.

func thisFunctionHere(mything thatThing, mythis alsoThis) (coolThing, error) {}
```

## some tidbits
these didn't need their own category, but the coffee told me they'd be good to mention.

- if it doesnt need to involve a pointer, don't make it a pointer.

- whatever it is probably doesn't need to be an interface.

- nothing we are doing here should require a library

- it probably doesn't need generics.

- json/yaml/xml is for developers

- handle errors as if the fuzzer *is* the user

## about descriptive naming

use or add descriptive names when you can:

```
// more this
for postIndex, thisPost := range posts {
    // this is better
}

// less this
for i, s := range pp {
    // this is less readable
}
```

note: this one's tough, and I'd be guilty of it too, but when you can, and when it makes sense, and particularly when there's some repeated or nested ranges and if trees, please use descriptive naming to help out the beginners and the curious.

## on toolchain, build steps, 'helpers'

you don't need more toolchain. if you are adding dependencies, you're adding liabilities. bearclaw is too simple for that. always seek to eliminate reliance on things outside the stdlib.  currently bearclaw relies on blackfriday, a markdown interpreter, merely because clawmark (an internal markdown interpreter) isn't done yet. [1]

[1] - markdown is hard! how do you do multi-line bold text? how do you do negative index number ordered lists? etc.

## dependencies

bearclaw also uses `fsnotify` for the `--watch` feature to know when a file has been modified, but as we've seen it's far from foolproof, **and it seems like** it may have reduced the number of platforms we can support, though that hasn't been confirmed. dependencies should always be eliminated at the first chance.

help others whenever possible. every suggestion and addition will bring up questions, and that's cool. most people aren't mind readers and most things you do in a group setting need to be 'sold' somewhat. it's cool.

# [<< Go back](https://github.com/donuts-are-good/bearclaw/blob/master/markdown/README.md)