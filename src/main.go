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
Version: 18.06.07
*/
package main

import (
//	"os"
	"fmt"
	"path"
	"flag"
	"trickyunits/gini"
	"trickyunits/mkl"
	"trickyunits/qff"
	"trickyunits/qstr"
	"trickyunits/ansistring"
	"time"
)

type ac struct{
	col int
	flag int
}

var cols = map[string] ac {}
var dependencies = []string{}

func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - main.go","18.06.07")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - main.go","GNU General Public License 3")
cols["lblue"] = ac{ansistring.A_Blue,ansistring.A_Bright}
cols["yellow"] = ac{ansistring.A_Yellow,0}
cols["cyan"] = ac{ansistring.A_Cyan,0}
cols["red"] = ac{ansistring.A_Red,0}
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
	flagtest:=flag.Bool("t",false,"Test build")
	flagrun:=flag.Bool("r",false,"Run immediately after building (only works if the project is build on this platform as well)")
	flag.Parse()
	Args=flag.Args()
	if len(Args)<1 {
		aprint  ("red","usage: ")
		aprint  ("yellow","Ryanna ")
		aprint  ("magenta","[flags] ")
		aprint  ("cyan","<Project> ")
		aprintln("magenta","[buildmode]\n\n")
		flag.PrintDefaults()
		aprint  ("yellow","Ryanna saves project files in GINI format if they get modified during the process only.\nRyanna can support multiple build modes if none are given Ryanna uses the mode top in line.\n\n")
		aprintln("yellow","This version of Ryanna was built on the next source files:")
		aprintln("cyan",mkl.ListAll())
		return
	}
	r_starttime:=time.Now()
	project = Args[0]
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
	if *flagtest { aprint("yellow","Flag: "); aprintln("cyan","Test build") }
	if *flagrun  { aprint("yellow","Flag: "); aprintln("cyan","Run after build") }
	// Let's get ready to rumble!
	askmeta()
	asksys()  // located in meta.go
	askdirs()
	loveversion()
	gather(*flagtest)
	release(*flagtest)
	r_endtime:=time.Now()
	dur:=r_endtime.Sub(r_starttime)
	sdur:=ansistring.SCol("Ryanna took ",ansistring.A_Yellow,0)
	/*
	sand:=false	
	if dur.Hours()>0 { sdur+=ansistrig.SCol(fmt.Sprintf("%d ",dur.Hours()),ansistring.A_Cyan,0)+ansistring.SCol("hours ",ansistring.A_Yellow,0); sand=true }
	*/
	sdur+=ansistring.SCol(dur.String(),ansistring.A_Cyan,0)
	sdur+=ansistring.SCol(" to build the entire package",ansistring.A_Yellow,0)
	fmt.Printf("\n\t%s\n\n",sdur)
}
