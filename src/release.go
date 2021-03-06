/*
	Ryanna
	
	
	
	
	(c) Jeroen P. Broks, 2017, 2018, All rights reserved
	
		This program is free software: you can redistribute it and/or modify
		it under the terms of the GNU General Public License as published by
		the Free Software Foundation, either version 3 of the License, or
		(at your option) any later version.
		
		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
		GNU General Public License for more details.
		You should have received a copy of the GNU General Public License
		along with this program.  If not, see <http://www.gnu.org/licenses/>.
		
	Exceptions to the standard GNU license are available with Jeroen's written permission given prior 
	to the project the exceptions are needed for.
Version: 18.05.22
*/
package main

import(
	"trickyunits/qff"
	"trickyunits/mkl"
	"trickyunits/qstr"
	"trickyunits/shell"
	"trickyunits/dirry"
	"path"
	"fmt"
	"os"
)



func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - release.go","18.05.22")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - release.go","GNU General Public License 3")
}

func copydependencies(p,target string){
	tgt:=target
	if qstr.Right(tgt,1)!="/" {tgt+="/"}
	for _,dep := range dependencies {
		tf := tgt + path.Base(dep)
		tsize:=-1; if qff.Exists(tf) {tsize= qff.FileSize(tf) }
		osize:=qff.FileSize(dep)
		if osize!=tsize {
			aprint  ("yellow",p+": Copying dependency: ")
			aprintln("cyan",  dep)
			err:=qff.CopyFile(dep,tf); if err!=nil { crash(err.Error()) }
		}
	}
}

func release_darwin(target,suf string){
	icon:=pask("MACICON","Mac: icon: ","")
	pwd:=qff.PWD()
	err:=os.Chdir(target)	
	if err!=nil { crash(err.Error() ) }
	err=os.MkdirAll("MAC",0777)
	if err!=nil { crash(err.Error() ) }
	os.Chdir("MAC")
	if err!=nil { crash(err.Error() ) }
	exe:=prjgini.C("Exe")
	aprint("yellow","Mac: Creating: ")
	if suf!="" { exe+="."+suf }
	aprintln("cyan",exe+".app")
	exe += ".app"
	if !qff.Exists(exe) {
		aprintln("cyan","Unzipping love for Mac")
		uzc:="unzip '"+mydir+"/love/"+prjgini.C("LOVEVERSION")+"/Mac64.zip' "
		//aprintln("magenta",uzc) // debug line
		shell.Shell(uzc)
		err = os.Rename("love.app",exe)
		if err!=nil { crash(err.Error() ) }
	}
	aprintln("yellow","Mac: Attaching resource")
	swap:=dirry.Dirry("$AppSupport$/$LinuxDot$Phantasar Productions/Ryanna/swap/")
	swapbuild:=swap+"Build/"
	orilove:=swapbuild+"love.love"
	err=qff.CopyFile(orilove,exe+"/Contents/Resources/"+prjgini.C("Exe")+".love"); if err!=nil { crash(err.Error()) }
	err=qff.CopyFile(icon,exe+"/Contents/Resources/love.icns"); if err!=nil { crash(err.Error()) }
	aprintln("yellow","Mac: Attaching icon")
	err=qff.CopyFile(icon,exe+"/Contents/Resources/GameIcon.icns"); if err!=nil { crash(err.Error()) }
	err=qff.CopyFile(icon,exe+"/Contents/Resources/OS X AppIcon.icns"); if err!=nil { crash(err.Error()) }
	copydependencies("Mac",exe+"/Contents/Resources")
	if suf!="" || prjgini.C("Package")=="JCR" {
		aprintln("yellow","Mac: Attaching jcrx")
		err=qff.CopyFile(mydir+"/jcrx/jcrx_darwin",exe+"/Contents/Resources/jcrx"); if err!=nil { crash(err.Error()) }
	}
	if platform=="darwin" { runfile = "open "+target+exe+".app" }
	os.Chdir(pwd)
}

func release_windows(target,suf string,bit int,test bool) {
	wd:=fmt.Sprintf("WIN%d",bit)
	pwd:=qff.PWD()
	err:=os.Chdir(target)	
	if err!=nil { crash(err.Error() ) }
	err=os.MkdirAll(wd,0777)
	if err!=nil { crash(err.Error() ) }
	os.Chdir(wd)
	if err!=nil { crash(err.Error() ) }
	defer os.Chdir(pwd)
	exe:=prjgini.C("Exe")
	if suf!="" { exe+="."+suf }
	exe +=".exe"
	aprintln("yellow",wd+": Creating distribution")
	swap:=dirry.Dirry("$AppSupport$/$LinuxDot$Phantasar Productions/Ryanna/swap/")
	if !qff.Exists("love.exe") {
		aprintln("cyan",fmt.Sprintf("Unzipping love for Windows %d bit",bit))
		uzc:="unzip -j '"+mydir+"/love/"+prjgini.C("LOVEVERSION")+"/"+wd+".zip' "
		//aprintln("magenta",uzc) // debug line
		shell.Shell(uzc)
		//err = CopyFile("love.exe",exe) //os.Rename("love.exe",exe)
		//if err!=nil { crash(err.Error() ) }
	}
	if test || prjgini.C("Package")=="JCR" {
		
		aprintln("yellow",wd+": Attaching resource")		
		swapbuild:=swap+"Build/"
		orilove:=swapbuild+"packed.jcr"
		err=qff.CopyFile(orilove,prjgini.C("Exe")+".jcr"); if err!=nil { crash(err.Error()) }
	}
	aprintln("yellow",wd+": Creating game's executable")
	err=qff.MergeFiles("love.exe",swap+"Build/zipped.zip",exe)
	if err!=nil { crash(err.Error() ) }
	copydependencies(wd,path.Dir(exe))
	if suf!="" || prjgini.C("Package")=="JCR" {
		aprintln("yellow",wd+": Attaching jcrx")
		err=qff.CopyFile(mydir+"/jcrx/jcrx_windows","jcrx.exe"); if err!=nil { crash(err.Error()) }
	}
	if platform=="window" && bit==32 { runfile = target+exe+".exe" }
}

func release_linux(target,suf string) {
	aprintln("red","Linux will be taken care of later!")
}

func release(test bool){
	if test {
		// Testing builds will only be created in the same platform kind as the platform on which the builder is running.
		// It's pretty pointless to build for a platform you can't test anyway, since test builds only work on the computer on which they were built.
		switch platform {
			case "darwin":	release_darwin(prjgini.C("Test."+platform),"test.build")
			case "windows":	release_windows(prjgini.C("Test."+platform),"test.build",32,test)
			case "linux":	release_linux(prjgini.C("Test."+platform),"test.build")
		}
		return
	}
	if prjgini.C("MAC64")=="YES" { release_darwin (prjgini.C("Release."+platform),"")    }
	if prjgini.C("WIN32")=="YES" { release_windows(prjgini.C("Release."+platform),"",32,test) }
	if prjgini.C("WIN64")=="YES" { release_windows(prjgini.C("Release."+platform),"",64,test) }
	if prjgini.C("LINUX")=="YES" { release_linux  (prjgini.C("Release."+platform),"")    }
	
}
