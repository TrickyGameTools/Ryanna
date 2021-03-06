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
Version: 18.05.05
*/
package main


import (
	"os"
//	"path"
	"strings"
	"trickyunits/mkl"
	"trickyunits/qff"
	"trickyunits/qstr"
)


func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - meta.go","18.05.05")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - meta.go","GNU General Public License 3")
}


func askmeta(){
	aprint("yellow","\n\n")
	ask("title","Project title:",qstr.StripAll(project))
	ask("author","Author:","Mr. X")
}


func asksys() { // strictly speaking not meta, but handier to have it here.
	pask("Release","Release dir:","")
	pask("Test","Test dir:","")
	 ask("Exe","Executable name",prjgini.C("Title"))
	if !qff.Exists(prjgini.C("Release."+platform)) {
		if yes("","Release dir does not yet exist. Create it? ") {
			os.MkdirAll(prjgini.C("Release."+platform),0777)
		} else {
			os.Exit(0)
		}
	}
	if !qff.Exists(prjgini.C("Test."+platform)) {
		if yes("","Test dir does not yet exist. Create it? ") {
			os.MkdirAll(prjgini.C("Test."+platform),0777)
		} else {
			os.Exit(0)
		}
	}
	yes("Mac64","Do you want to create a mac build")
	yes("Win32","Do you want to create a windows 32bit build")
	yes("Win64","Do you want to create a windows 64bit build")
	yes("Linux","I cannot build for Linux yet, but do you want to create a file in a separate Linux folder for further packaging")
	//pask("Libs","Where are your external libraries for Ryanna stored?",path.Dir(qff.PWD())+"/RyannaLibs")
	initlist("Libraries."+platform,"Tell me which directories you wish to use for finding libraries.")
	initlist("buildmodes","Tell me. Which build modes do you have in mind?")
	initlist("sources."+platform,"Now give me the list of source dirs, please")
	for {
		a := strings.ToUpper(ask("PACKAGE","Package as JCR or as ZIP:","JCR"))
		if a=="ZIP" || a=="JCR" { break }
	}
	initlist("CONFIG","Do you have any special settings in mind for the conf.lua file?")
}
