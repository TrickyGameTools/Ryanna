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

