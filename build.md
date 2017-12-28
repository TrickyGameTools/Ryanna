# Building


## Preparations to build

Ryanna requires the [tricky units for go](https://github.com/Tricky1975/trickyunits_go) reposity do be present.
If /mygoprojects is your go folder then it needs to be in /mygoprojects/src/trickyunits

## Building the tool itself

If you have the Ryanna folder in your src folder of Go you can best create a folder named bin and then type

~~~sh
	go build -o bin/Ryanna Ryanna/src
~~~

The file build.sh contains this command (if you are on Windows, copy its contents to build.bat if you wish).


## Further requirements

This will, of course, only build the Ryanna tool itself. If no errors occur you got yourself your tool.
It also requires [jcrx](https://github.com/Tricky1975/jcrx) to be present for the platforms you want to export for.
In the directory where the Ryanna executable is located you should create a directory named jcrx and in there create the folders "Windows", "Mac" and "Linux" respectively and put the JCRX tool in there.
(Please note, as Ryanna is not a real compiler, it can cross-built to any platform from any platform. Yeah, even windows->mac is possible. Ryanna can do it).

Ryanna does also require the tools zip and unzip to be installed in any directory it can reach through the path and the same goes for the [jcr6cli tools](https://github.com/Tricky1975/jcr6cli). (Please note only the variant written in Go will work here. The variant written in Blitzmax is deprecated and has some different approaches to argumentation rendering it incompatible with Ryanna).

 
## Lastly you'll need Love itself.

As long as I don't know any different Ryanna should work with any Love version.
In the folder were Ryanna is located you should create a directory named "Love"
In there you should create a folder containing the version number of Love you want to use. (This is important to enable Ryanna to work with older versions of Love if projects require this).
In there you can place the zips for Love you can download from the [love2d.org](https://love2d.org) and place them in the folder of the correct version number.
- The zip with the 32bit version for Windows should be named Windows_32.zip
- The zip with the 64bit version for Windows should be named Windows_64.zip
- The zip with the 64bit version for Mac should be named Mac_64.zip
- For Linux all Ryanna can do for now is create a .love file with its needed jcr file files (or merge them if the project file requires it to). As Linux is really distro dependent and me not having a fully working Linux machine at my disposal to truly test things), this is all I can do for now. 
As Ryanna has been set up specifically for projects that are too large for mobile devices to handle (and since I don't think jcrx could really compile to mobile devices either), Ryanna cannot export to Android or iOS and also not to HTML5, nor am I interested in outputting to these systems. My original LoveBuilder might be able to fill the drill someday (no promises though).




That should set things right here.
