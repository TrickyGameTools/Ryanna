# Important note

At the moment I wrote this document, Ryanna has only been tested in Darwin (mac). Problems for other systems are not expected, but you need to keep it in mind.

# What you need

1. First of all you need the Go compiler, since Ryanna was written in Go
2. You need the contents of [this repository](https://github.com/Tricky1975/trickyunits_go) to be present in a folder named "trickyunits", and it should be located in an "src" folder Go has access to.
3. You need the [jcr6 cli tools](https://github.com/Tricky1975/jcr6cli) to be installed, Ryanna uses this as a dependency.
4. The utility "zip" must be installed. For Mac and most Linux distros this is present by default, Windows users will need to read my notes below.
5. The zip files of [LÖVE](http://love2d.org) for Mac and Windows (both 32bit and 64bit).
6. And of course this source code in a directory Go can access of course.

# Note to Windows users:

Since Windows is not accustomed to dependencies the way Unix based systems are you need to have a few things in order.
1. Best course of action is to have jcr6_add.exe from the JCR6 cli tools in the same folder as were Ryanna is located.
2. Zip is not installed by default, and also note that WinZIP or any similar tools are not going to cut the bill. Zip is a command line utility, having several programs, most notably zip.exe and unzip.exe. Ryanna needs them both. Best is to copy these two files to the same folder as were Ryanna.exe is located after you are done compiling. [this is one place where you can download these](http://stahlworks.com/dev/index.php?tool=zipunzip)


For the easiness of this doc I'll pick the folder "Ryanna" in your home directory as the folder to install it in.
This doc has been written for Love 11.1, but basically Ryanna should support any version of 0.10.2 or higher but then you'll need to adept what I typed below, for that version... of course.

Linux and Mac users should open the terminal and type the following then:
~~~sh
mkdir ~/Ryanna
go build -o ~/Ryanna Ryanna
cd ~/Ryanna
mkdir love
cd love
mkdir 11.1
cd 11.1
# Now download all the zip files from love2d.org for Windows and Mac. I said, zipfiles... NOT THE INSTALLERS! #
mv love-11.1-win32.zip WIN32.zip
mv love-11.1-win64.zip WIN64.zip
mv love-11.1-macos.zip MAC64.zip
# NOTE! Use a browser to download these. I've tried using curl to download these myself, but that will only result into empty zip files.
~~~

And now the way to go for Windows users:
Open the command line (this has been written for cmd, but I think SuperShell must also be able to handle this).
Make sure the prompt says C:\users\your username> before you begin!
~~~bat
md Ryanna
go build -o C:\users\your_username\Ryanna Ryanna
cd Ryanna
md love
cd love
md 11.1
cd 11.1
rem Now download all the zip files from love2d.org for Windows and Mac. I said, zipfiles... NOT THE INSTALLERS! 
ren love-11.1-win32.zip WIN32.zip
ren love-11.1-win64.zip WIN64.zip
ren love-11.1-macos.zip MAC64.zip
~~~
WARNING! Take good note of the difference between *ren* (REName) and *rem* (REMark).


Then you must also have the jcrx tool at the ready.
In the folder where Ryanna is located there must be a subfolder named "jcrx"
It needs:
"jcrx_darwin" -- The executable for exporting to Mac
"jcrx_windows" -- the executable for exporting to Windows (.exe may not be present. Ryanna puts that in by itself).
"jcrx_linux" -- The executable for exporting to Linux (any distro)

# IMPORTANT NOTE TO WINDOWS USERS WHEN IT COMES TO USING RYANNA TO EXPORT TO MAC AND LINUX.

When using Ryanna to export to Mac and Linux, the game will initially NOT work. That is NOT a bug in Ryanna, but an issue with Windows being a DOS based system and Mac and Linux being Unix based systems. Where for Windows the extensions .exe, .com and .bat suffice to recognise a file as being exectuable, Unix based systems originally never had extensions and the extensions used today, well that's just a habit copied from MS-DOS, as the extensions WERE handy in quickly seeing what kind of data a file has. ;)

In stead Unix uses the attribute system, which is completely different from the attribute system in Windows, and no, there's no way to get it compatible, no matter WHAT you try. 

In Unix the "x" attribute has to tell the system a file is executable.

On Mac these two files must have the "x" attribute:
- mygame.app/contents/macos/love
- mygame.app/contents/resources/jcrx

On Linux LOVE itself is mostly downloaded as a dependency and then you don't have to worry about that one, but the jcrx file coming with the .love file (and other resource files) must have the x attribute.


There is no way to make this happen in Windows, and unfortunately, there are also no archivers supporting attaching "x" to files (and why not is completely beyond me). 

If you don't want your users to have to deal with this problem there is only one way to go. Install Virtual Box and install Linux on it. Any distro will do, and transfer your Mac and Linux exports to the virtual harddrive of your virtual Linux machine.

Now let's just put the games in your home folder in the subfolder RyannaOutput (just for keeping this clean), and the Mac version in the Mac subfolder and the Linux version in the Linux subfolder, and let's assume you exported the game as "mygame".

Open the terminal and type the next commands:
~~~shell
# Mac
chmod +x ~/RyannaOut/Mac/mygame.app/Contents/MacOS/love
chmod +x ~/RyannaOut/Mac/mygame.app/Conents/Resources/jcrx

# Linux
chmod +x ~/RyannaOut/Linux/jcrx
~~~
You can now zip the two exports from this harddrive and copy them back to your Windows systems. As long as they remain zipped, nothing's wrong.


# AppImage

You may want to convert the final output into an appimage to make sure your Linux users don't need any dependencies at all.
I am still investigating how to make Ryanna properly deal with that. I regret to tell you that from Windows it will never be possible to export into AppImage, sorry. If it can be done from Mac, is something I can't tell yet, I need to properly sort that one out.
For Linux I already got a script to easily do it, and but I didn't yet have to time to study it thoroughly. Once I did I will make sure the Linux version of Ryanna will be able to write everything in an AppImage, as I do support the portability of apps on Linux, as the giant need of (often conflicting) dependencies is a true achilles heel to the Linux system, I tell ya!
