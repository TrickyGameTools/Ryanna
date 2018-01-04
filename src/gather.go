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
Version: 18.01.05
*/
package main


import (
	"os"
	"fmt"
	"path"
	"strings"
	"trickyunits/mkl"
	"trickyunits/qff"
	"trickyunits/qstr"
	"trickyunits/dirry"
	"trickyunits/shell"
	"trickyunits/tree"

)

var libdebug = false

func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - gather.go","18.01.05")
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

func ziplib(dir,lib,zipf string){
	od:=qff.PWD()
	libe:=lib
	if !qstr.Suffixed(libe,".rel") { libe+=".rel" }
	err:=os.Chdir(dirry.Dirry(dir))
	if err!=nil { crash(err.Error()) }
	shell.Shell("zip -r -9 '"+dirry.Dirry(zipf)+"' Libs/"+libe)
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
	libs:=[]string{}
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
		bstr = strings.Replace(bstr,"$RyannaTitle$",prjgini.C("Title"),-10)
		if test { bstr = strings.Replace(bstr,"$RyannaBuildType","test",-20) } else {bstr = strings.Replace(bstr,"$RyannaBuildType","normal",-20)}
		if test || prjgini.C("Package")=="JCR" { bstr=strings.Replace(bstr,"\"$RyannaLoadJCR$\"","true",-11) } else { bstr=strings.Replace(bstr,"\"$RyannaLoadJCR$\"","false",-12) }
		bstr = strings.Replace(bstr,"$RyannaVersion$",mkl.Newest(),-14)
		err := qff.WriteStringToFile(dirry.Dirry(swapbase+f+".lua"),bstr)
		if err!=nil { crash(err.Error()) }
	}
	zip(swapbase,zipf)
	// All preps done, now to gather it all
	for _,d:=range dirstoprocess {
		aprint  ("yellow","Gathering: ")
		aprintln("cyan",d)
		jtree:=tree.GetTree(d,false)
		for _,f:=range jtree { // looking for external lib references. The folder LIBS/ is reserved for this!
			if path.Ext(f)==".lua" {
				lines:=qff.GetLines(d+"/"+f)
				for _,line:=range lines {
					l:=strings.ToUpper(qstr.MyTrim(line))
					   //01234567
					pr:="-- $USE LIBS/"
					if qstr.Left(l,len(pr))==pr {
						if libdebug { aprintln("magenta","Requested: "+l) }
						found:=false
						nl:=qstr.MyTrim(l[7:])
						for _,lb:=range libs {
							found=found || lb==nl
						}
						if !found { 
							libs = append(libs,nl) 
							if libdebug {
								aprintln("magenta","Requested new library: "+l)
							}
						}
					}
				}
			}
		}
		if test {
			jif += "IMPORT:"+d+"\n"
		} else if prjgini.C("Package")=="JCR" {
			for _,jtf:=range jtree{
				jif += "FILE:"+d+"/"+jtf+"\n"
				jif += "TARGET:"+jtf+"\n"
				jif += "AUTHOR:"+prjgini.C("SOURCE['"+d+"'].AUTHOR")+"\n"
				jif += "NOTES:"+prjgini.C("SOURCE['"+d+"'].LICENSE")+"\n"
				if strings.ToLower(path.Ext(jtf))==".mp3" {
					jif += "STORAGE:Store\n"
				} else {
					jif += "STORAGE:BRUTE\n"
				}
			}
		} else {
			zip(d,zipf)
		}
	}
	// Map external libraries if there are any
	libtrees:=[]string{}
	for _,lt:=range prjgini.List("Libraries."+platform){
		if test { 
			jif += "IMPORT:"+lt+"\n"
		}
		ltd,err:=qff.GetDir(lt+"/Libs",2,false)
		if err!=nil { crash(err.Error()) }
		for _,actlib:=range ltd {
			libtrees=append(libtrees,lt+"/Libs/"+actlib)
		}
	}
	// Include libraries if there are any
	for i:=0;i<len(libs);i++ {
		ok:=false
		lib:=libs[i]
		clb:=strings.ToUpper(lib)
		if !qstr.Suffixed(clb,".REL") { clb += ".REL" }
		for _,pl:=range libtrees { 
			cpl:=strings.ToUpper(pl)
			if libdebug { 
				fmt.Print(i,">",libs[i],">",clb,"\n\t",pl,">",cpl,"\n\t\t",path.Base(cpl)," >> ",path.Base(cpl)==clb,"\n")
			}
			if "LIBS/"+path.Base(cpl)==clb{
				ok=true
				aprint("yellow","Importing library: ")
				aprintln("cyan",pl)
				if qstr.Suffixed(cpl,clb){
					// are there any requests for new libs?
					jtree:=tree.GetTree(pl,false)
					for _,f:=range jtree { // looking for external lib references. The folder LIBS/ is reserved for this!
						if path.Ext(f)==".lua" {
							lines:=qff.GetLines(pl+"/"+f)
							for _,line:=range lines {
								l:=strings.ToUpper(qstr.MyTrim(line))
								   //01234567
								pr:="-- $USE LIBS/"
								if qstr.Left(l,len(pr))==pr {
									found:=false
									nl:=qstr.MyTrim(l[7:])
									for _,lb:=range libs {
										found=found || lb==nl
									}
									if !found { libs = append(libs,nl) }
								}
							}
						}
					}
					if test {
						// Do nothing. Everything that has to be done, has been done.
					} else if prjgini.C("Package")=="JCR" {
						for _,jtf:=range jtree{
							jif += "FILE:"+pl+"/"+jtf+"\n"
							jif += "TARGET:Libs/"+path.Base(pl)+"/"+jtf+"\n"
							jif += "AUTHOR:"+prjgini.C("SOURCE['"+pl+"/"+jtf+"'].AUTHOR")+"\n"
							jif += "NOTES:"+prjgini.C("SOURCE['"+pl+"/"+jtf+"'].LICENSE")+"\n"
							if strings.ToLower(path.Ext(jtf))==".mp3" {
								jif += "STORAGE:Store\n"
							} else {
								jif += "STORAGE:BRUTE\n"
							}
						}
					} else {
						ziplib(path.Dir(pl)+"..",path.Base(pl),zipf)
					}
				}
			}
		}
		if !ok { crash("Library '"+libs[i]+"' could not be located") }
	}
	// Alias
	prjgini.CL("ALIASES",true)
	for _,al:=range prjgini.List("ALIASES"){
		als:=strings.Split(al," => ")
		aprint("yellow","Alias request: ")
		aprintln("cyan",al)
		if len(als)!=2 {
			nferror(" Invalid alias request: "+al)
		} else if test {
			nferror("Aliasing in testing not yet supported: "+al)
		} else if prjgini.C("Package")=="JCR" {
			jif += "ALIAS:"+qstr.MyTrim(als[0])+"\nAS:"+qstr.MyTrim(als[1])+"\n"
		} else {
			nferror("Aliasing in full zip export not yet supported: "+al)
		}
	}
	
	// jcr build
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
