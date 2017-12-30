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


import(
	"strings"
	"trickyunits/mkl"
	"trickyunits/qff"
)


func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - askfiles.go","17.12.30")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - askfiles.go","GNU General Public License 3")
}

func scansrcdir(d string){
	dirs,err:=qff.GetDir(d,2,false)
	if err!=nil {
		aprint("red","ERROR! ")
		aprintln("yellow",err.Error())
		return
	}
	aprint  ("yellow","Scanning source: ")
	aprintln("cyan"  ,d)
	for _,sd:=range(dirs) {
		t:=pask("SOURCE['"+d+"/"+sd+"'].TYPE",sd+"    ","NONE")
		    ask("SOURCE['"+d+"/"+sd+"'].AUTHOR","Author of this dir: ",sd)
		    ask("SOURCE['"+d+"/"+sd+"'].LICENSE","License:","Unknown license")
		mayadd:=t=="ALL"
		mode:=prjgini.ListIndex("BuildModes",0)
		if len(Args)>1 { mode=Args[1] }
		ts:=strings.Split(t," ")
		for _,st:=range ts{
			mayadd = mayadd || st==mode
		}
		mayadd = mayadd && t!="NONE"
		if mayadd {
			dirstoprocess = append(dirstoprocess,d+"/"+sd)
		}
	}
}


func askdirs(){
	aprintln("yellow","\n\nTime to see which dirs we must add under what types and authors and licences")
	c:=0
	aprint  ("yellow","The next types are known. ALL will always include this dir to all directories and NONE will always ignore it: ")
	for i,t:=range prjgini.List("buildmodes"){
		c++
		if i!=0 {aprint("cyan",", ") }
		aprint("cyan",t)
	}
	aprint("yellow","\n\n")
	for _,d:=range prjgini.List("sources."+platform){
		scansrcdir(d)
	}
}
