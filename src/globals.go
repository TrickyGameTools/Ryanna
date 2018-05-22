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


import "trickyunits/mkl"
import "trickyunits/gini"
import "runtime"
import "os"
import "path"

func init(){
mkl.Version("Ryanna - Builder for jcr based love projects - globals.go","18.05.22")
mkl.Lic    ("Ryanna - Builder for jcr based love projects - globals.go","GNU General Public License 3")
me,_ = os.Executable()
mydir = path.Dir(me)
}


var platform = runtime.GOOS
var me,mydir string

var project string
var prjgini gini.TGINI

var dirstoprocess []string

var script = map[string] string {}


var Args []string

var runfile string


