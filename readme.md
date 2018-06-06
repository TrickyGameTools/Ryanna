![image](https://user-images.githubusercontent.com/11202073/41039184-f074b4fe-6998-11e8-895d-fd8dd4e2e8aa.png)

# Ryanna

Ryanna is a quick building tool written in Go.
It's been set up to gather everything a project of mine written in LÖVE will need and pack it all together, attach the proper LÖVE version and other dependencies to it to create a ready-to-go release.
It's also able to create a test build which I can only use on my own computer, but which only takes seconds to build and which would normally need no rebuilds for every single change.

Ryanna also provides some functionality that LÖVE doesn't have:
- For starters it adds a pre-processor for Lua, which can empower Lua
- It's $USE directive can look for external libraries I wrote and automatically only include those in the package I need
- JCR6 support, which is for game resource packages far more sophistacated than zip (hence the need of the jcrx dependency which Ryanna automatically puts in).
- A few quick functions that I consider standard, but which were never standard to Lua
- Automatically creates packages for both mac and windows
- (Linux is planned, but due to Linux' needlessly complicated approach to things, this requires more study).


Ryanna has in the first place only been set up to make my own life easier when using LÖVE, so it may not be the tool suitable for anybody, but if you do find some use for it, go ahead use it.


The tool Ryanna was named after the main protagonist of Sixty-Three Fires of Lung who bears the same name, hence her picture also serves as the project's mascot ;)
