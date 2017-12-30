/*
	Ryanna
	
	
	
	
	(c) Jeroen P. Broks, 2017, All rights reserved
	
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
Version: 17.12.30
*/
package main


import (
	"os"
	"strings"
	"trickyunits/mkl"
	"trickyunits/qff"
	"trickyunits/dirry"
	"trickyunits/shell"
	"trickyunits/tree"

)


func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - gather.go","17.12.30")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - gather.go","GNU General Public License 3")
}


func md(d string){
	// Create new swap and make sure all the swap is clean!
	err:=os.RemoveAll(dirry.Dirry(d))
	if err!=nil { crash ( err.Error() ) }
	err=os.MkdirAll(dirry.Dirry(d),0777)
	if err!=nil { crash ( err.Error() ) }
}

func zip(dir,zipf string){
	od:=qff.PWD()
	err:=os.Chdir(dirry.Dirry(dir))
	if err!=nil { crash(err.Error()) }
	shell.Shell("zip -r -9 '"+dirry.Dirry(zipf)+"' *")
	os.Chdir(od)
}

func gather(test bool){
	aprintln("cyan","Organising swap")
	swap:="$AppSupport$/$LinuxDot$Phantasar Productions/Ryanna/swap/"
	swapbase:=swap+"BaseShit/"
	swapbuild:=swap+"Build/"
	zipf:=swapbuild+"zipped.zip"
	jcrf:=swapbuild+"packed.jcr"
	love:=swapbuild+"love.love"
	MainScript:=ask("MainScript","Main Script:","Script/Ryanna_Main.lua")
	sig:=""
	if prjgini.C("Package")=="JCR" { sig = ask("JCRSIG","JCR signature","" ) }
	calljcr:=test || prjgini.C("package")=="JCR"
	jif:=""
	if calljcr {
		jif += "SIGNATURE:"+sig+"\n"
		jif += "FATSTORAGE:BRUTE\n"
	}
	md(swapbase)
	md(swapbuild)
	for f,str := range script {
		bstr:=str
		bstr = strings.Replace(bstr,"$RyannaMainScript$",MainScript,-10)
		if prjgini.C("Package")=="JCR" { bstr=strings.Replace(bstr,"\"$RyannaLoadJCR$\"","true",-11) } else { bstr=strings.Replace(bstr,"\"$RyannaLoadJCR$\"","false",-12) }
		bstr = strings.Replace(bstr,"$RyannaVersion$",mkl.Newest(),-14)
		err := qff.WriteStringToFile(dirry.Dirry(swapbase+f+".lua"),bstr)
		if err!=nil { crash(err.Error()) }
	}
	zip(swapbase,zipf)
	// All preps done, now to gather it all
	for _,d:=range dirstoprocess {
		aprint  ("yellow","Gathering: ")
		aprintln("cyan",d)
		if test {
			jif += "IMPORT:"+d+"\n"
		} else if prjgini.C("Package")=="JCR" {
			jtree:=tree.GetTree(d,false)
			for _,jtf:=range jtree{
				jif += "FILE:"+d+"/"+jtf+"\n"
				jif += "TARGET:"+jtf+"\n"
				jif += "AUTHOR:"+prjgini.C("SOURCE['"+d+"/"+jtf+"'].AUTHOR")+"\n"
				jif += "NOTES:"+prjgini.C("SOURCE['"+d+"/"+jtf+"'].LICENSE")+"\n"
				jif += "STORAGE:BRUTE\n"
			}
		} else {
			zip(d,zipf)
		}
	}
	if calljcr {
		aprintln("cyan","Creating JCR6 work file")
		jifile:=dirry.Dirry(swapbuild)+"JCR_Instruction_File.jif"
		err:=qff.WriteStringToFile(jifile,jif)
		if err!=nil { crash(err.Error()) }
		shell.Shell("jcr6_add -jif '"+jifile+"' '"+dirry.Dirry(jcrf)+"'")
		aprintln("cyan","Merging")
		if platform=="windows"{
			shell.Shell("copy/b \""+dirry.Dirry(jcrf)+"\"+\""+dirry.Dirry(zipf)+"\" \""+dirry.Dirry(love)+"\"")
		} else {
			shell.Shell("cat \""+dirry.Dirry(jcrf)+"\" \""+dirry.Dirry(zipf)+"\" > \""+dirry.Dirry(love)+"\"")
		}
	}

}
