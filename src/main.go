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
Version: 17.12.28
*/
package main

import (
	"os"
	"fmt"
	"path"
	"trickyunits/gini"
	"trickyunits/mkl"
	"trickyunits/qff"
	"trickyunits/qstr"
	"trickyunits/ansistring"
)

type ac struct{
	col int
	flag int
}

var cols = map[string] ac {}

func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - main.go","17.12.28")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - main.go","GNU General Public License 3")
cols["lblue"] = ac{ansistring.A_Blue,ansistring.A_Bright}
cols["yellow"] = ac{ansistring.A_Yellow,0}
cols["cyan"] = ac{ansistring.A_Cyan,0}
cols["red"] = ac{ansistring.A_Cyan,0}
cols["magenta"] = ac{ansistring.A_Magenta,0}
cols["bcyan"] = ac{ansistring.A_Cyan,ansistring.A_Blink}

}



func aprint(col string,s ...string){
	for _,s1:=range s{
		fmt.Print(ansistring.SCol(s1,cols[col].col,cols[col].flag))
	}
}

func aprintln(col string, s ...string){
	for _,s1:=range s{
		fmt.Print(ansistring.SCol(s1,cols[col].col,cols[col].flag))
	}
	fmt.Println()
}

func main(){
	version:=mkl.Newest()
	aprint  ("lblue",  "Ryanna ")
	aprintln("yellow", version )
	aprintln("cyan","(c) Jeroen P. Broks 2017-20",qstr.Left(version,2),"\n")
	if len(os.Args)<2 {
		aprint  ("red","usage: ")
		aprint  ("yellow","Ryanna ")
		aprint  ("cyan","<Project> ")
		aprintln("magenta","[buildmode]\n\n")
		aprint  ("yellow","Ryanna saves project files in GINI format if they get modified during the process only.\nRyanna can support multiple build modes if none are given Ryanna uses the mode top in line.\n\n")
		aprintln("yellow","This version of Ryanna was built on the next source files:")
		aprintln("cyan",mkl.ListAll())
		return
	}
	project = os.Args[1]
	if path.Ext(project)=="" { project+=".rpf" } // rpf = Ryanna Project File
	if !qff.Exists(project) {
		if yes("","The project "+project+" does not yet exist!\nShall I create it for you") {
			prjgini = gini.TGINI{}
			aprint("yellow","Creating: ")
			aprintln("cyan",project)
			prjgini.SaveSource(project)
		} else {
			return
		}
	} else {
		aprint("yellow","Reading: ")
		aprintln("cyan",project)
		prjgini = gini.ReadFromFile(project)
	}
}
