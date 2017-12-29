# Script

Here you'll find a collection of .lua files which need to be merged 
into the go files prior to compiling. In the repository this 
should normally already have been done for you.

I chose this setup because... It was the quickest way to include these
files into Go as strings, while I could still benefit from Syntax
Highlighters and parse error checkers.

My sincerest apologies for putting the linkup script in Python 2.7
I wanted to do this with Lua, but Lua has no default directory reading 
routines (that requires an external third party library in C, or so
it seems, and Lua didn't want to rely on external parties when they
didn't need to). Basically all Python will do is grab all the files 
together and dump them in one big Go source file in the src directory.
