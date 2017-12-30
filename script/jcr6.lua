--[[
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
Version: 17.12.30
]]
local ldir = love.filesystem.getSourceBaseDirectory()

-- $IF WNDOWS
local jcrx = ldir+"\\jcrx"
-- $FI
-- $IF MAC
local jcrx = ldir+"//jcrx"
-- $FI
-- $IF LINUX
error("Sorry! I'm still working this part out for Linux")
-- $FI

function JCR_Dir(jfile)
	bt = io.popen(jcrx.." dirout '"..jfile.."' lua")
	sl = bt:readlines()
	assert(sl[1]=="OK",sl[2])	
	bt:close()
	s = ""
	for i=2,#sl do s = s .. sl[i] .. "\n" end
	f=load(s,"JCR_DIR("..jfile..")")
	ret.entries = f()
	ret.JCR_B = JCR_B
	ret.from = jfile
	ret.kind= 'JCR'
	return ret
end

function LOVE_Dir(skipwork) -- if set to true it will skip the directories swap and savegame as they are part of the working directories (at least they should be).
	local list = love.filesystem.getDirectoryItems( "" )
	local entries = {}
	for i,f in ipairs(list) do
		if (not skipwork) or (left(lower(f),5)~="swap/" and left(lower(f),9)~="savegame/") then
			entries[#entries+1] = { entry = f, LOVE = f, mainfile = love.filesystem.getSource() }
		end
	end
	ret = { entries = entries, ret.from = love.filesystem.getSource(), ret.kind="LOVE" }
	return ret
end

function JCR_B(j,nameentry)
	local mj
	if not nameentry then
		entry = string.upper(j)
		mj = jcr
		assert ( mj , "JCR not set!" )
	else
		entry = string.upper(nameentry)
		if type(mj)=='table' then
			mj = j
		else 
			mj = JCR_Dir(j)
		end
	end
	e = string.upper(entry)
	edata = mj.entries[e]
	assert(edata,"Entry "+entry+" not found")
	if not edata then return end -- Make sure nothing bad happens in case of a pcall
	if edata.LOVE then
		return love.filesystem.read(edata.LOVE)
	end
	bt = io.popen(jcrx.." typeout '"..mj.from.."' '"..entry.."'")
	sl = bt:readlines()
	assert(sl[1]=="OK",sl[2])
	bt:close()
	s = ""
	for i=2,#sl do s = s .. sl[i] .. "\n" end
	return s
end


function BaseDir() -- Basically only called by Ryanna and loaded based on Ryanna's findings.
	ret = {}
	ret.entries = {}
	ret.from = love.filesystem.GetSource()
	ret.kind = "MIXED"
	
	k = {}
	k[1] = LOVE_Dir()
	if RYANNA_LOAD_JCR then k[2] = JCR_Dir(ret.from) end
	for i,d in ipairs(k) do
		for key,entry in pairs(d) do ret.entries[#ret.entries+1] = entry end
	end
	return ret
end
BaseDir()


--[[
mkl.version("Ryanna - Builder for jcr based love projects - jcr6.lua","17.12.30")
mkl.lic    ("Ryanna - Builder for jcr based love projects - jcr6.lua","GNU General Public License 3")
]]
