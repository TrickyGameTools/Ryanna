package main

import "trickyunits/mkl"

// Licensed under the GNU

func init(){
	script["basis"] = `--[[
	Ryanna - Script
	
	
	
	
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
]]
-- basis script

--[[
mkl.version("Ryanna - Builder for jcr based love projects - basis.lua","17.12.29")
mkl.lic    ("Ryanna - Builder for jcr based love projects - basis.lua","GNU General Public License 3")
]]


Ryanna = {
	RyannaVersion = "$RyannaVersion$",
	LuaVersion = _VERSION,
	LoveVersion = string.format("%d.%d.%d - %s",love.getVersion() ) -- This line is dirty code straight from the toilet, but I don't care :P
	
}
`

	/* Lua */ mkl.Version("Ryanna - Builder for jcr based love projects - basis.lua","17.12.29")

	/* Lua */ mkl.Lic    ("Ryanna - Builder for jcr based love projects - basis.lua","GNU General Public License 3")



}