# bearclaw philosophy

before i add a feature or change some code, i ask myself:


- is this proposed solution easy to understand?

above all, the code must be easy to understand or it will otherwise scare away new developers or someone who must fix it themselves. 

use descriptive naming, use full words, use less 'lingo', be less fancy.

- is this change for me, or for the user?

is the thing being added for me the developer, or does this directly address something for the user?

- for the amount of code I'm adding, is the output or value of this commensurate for what it gives?

less things running, less moving parts, less load. stdlib is preferred over a library. abstraction to libraries trades brevity for mystery, don't do it.

- is this stylistically perpendicular to what the rest of the code looks like?

does what i am creating match the rest of the code? if no, why? am i unintentionally slipping a solo into a choral ensemble?

- who will maintain this?

simple solutions are maintainable. the tastiest product will rot on the vine if it is unrecognizable.

- will this incur cognitive debt for the user? for contributors?

is my solution flexing on the user? does this require anything new to learn? what concepts does it assume the user knows?

the user is always assumed to not be a developer

# [<< Go back](README.md)