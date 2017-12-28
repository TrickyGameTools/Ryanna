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


import (
	"trickyunits/mkl"
	"trickyunits/qstr"
)


func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - ask.go","17.12.29")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - ask.go","GNU General Public License 3")
}


func ask(tag,question,defaultanswer string) string {
	antwoord:=""
	for antwoord==""{
		aprint("yellow",question," ")
		if prjgini.C(tag)=="" {
			aprint("cyan","["+defaultanswer+"] ")
			antwoord=qstr.RawInput("")
			if antwoord=="" { antwoord=defaultanswer }
			if antwoord!="" { 
				prjgini.D(tag,antwoord)
				prjgini.SaveSource(project) 
			}
		} else {
			antwoord=prjgini.C(tag)
			aprintln("cyan",antwoord)
		}
	}
	return antwoord
}


func pask(tag,question,defaultanswer string) string{
	return ask(tag+"."+platform,question,defaultanswer)
}
