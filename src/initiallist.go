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
Version: 17.12.29
*/

package main

import(
	"trickyunits/mkl"
	"trickyunits/qstr"
	"trickyunits/ansistring"
)


func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - initiallist.go","17.12.29")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - initiallist.go","GNU General Public License 3")
}

func initlist(tag,question string) {
	if prjgini.ListExists(tag) { return }
	aprintln("yellow",question)
	aprintln("lblue","You can end by just pressing ENTER")
	for{
		myline:=qstr.RawInput(ansistring.SCol(">",ansistring.A_Red,0))
		if myline!="" {
			prjgini.Add(tag,myline)
		} else {
			prjgini.SaveSource(project)
			return
		}
	}
}
