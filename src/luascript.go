package main

import "trickyunits/mkl"

// Licensed under the GNU

func init(){
	script["main"] = `--[[
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
-- basis script

--[[
mkl.version("Ryanna - Builder for jcr based love projects - basis.lua","17.12.30")
mkl.lic    ("Ryanna - Builder for jcr based love projects - basis.lua","GNU General Public License 3")
]]


RYANNA_MAIN_SCRIPT = "$RyannaMainScript$"
RYANNA_LOAD_JCR    = "$RyannaLoadJCR$"     -- quotes will be removed. I've set it up as a string to deceive parse error checking IDEs, as they would otherwise go crazy.

Ryanna = {
	RyannaVersion = "$RyannaVersion$",
	LuaVersion = _VERSION,
	LoveVersion = string.format("%d.%d.%d - %s",love.getVersion() ) -- This line is dirty code straight from the toilet, but I don't care :P
	
}


-- include use.Lua and jcr6.lua which now have not made their official entrace, so I gotta call them manually
function load_primary_dependencies()
	for _,dep in {"jcr6.lua","use.lua"}
		chunk, errormsg = love.filesystem.load( name )
		assert(chunk,errormsg)
		chunk()
	end
end load_primary_dependencies()

-- Table features -- 
function each(a) -- BLD: Can be used if you only need the values in a nummeric indexed tabled. (as ipairs will always return the indexes as well, regardeless if you need them or not)
local i=0
if type(a)~="table" then
   print("Each received a "..type(a).."!",255,0,0)
   return nil
   end
return function()
    i=i+1
    if a[i] then return a[i] end
    end
end

function ieach(a) -- BLD: Same as each, but now in reversed order
local i=#a+1
if type(a)~="table" then
   print("IEach received a "..type(a).."!",255,0,0)
   return nil
   end
return function()
    i=i-1
    if a[i] then return a[i] end
    end
end

--[[

    This function is written by Michal Kottman.
    http://stackoverflow.com/questions/15706270/sort-a-table-in-lua

]]
function spairs(t, order)
    -- collect the keys
    local keys = {}
    for k in pairs(t) do keys[#keys+1] = k end

    -- if order function given, sort by it by passing the table and keys a, b,
    -- otherwise just sort the keys 
    if order then
        table.sort(keys, function(a,b) return order(t, a, b) end)
    else
        table.sort(keys)
    end

    -- return the iterator function
    local i = 0
    return function()
        i = i + 1
        if keys[i] then
            return keys[i], t[keys[i]]
        end
    end
end

-- misc

function valstr(a)
return ({
   ['nil'] = function(a) return 'nil' end,
   ['number'] = function(a) return ''..a end,
   ['string'] = function(a) return a end,
   ['boolean'] = function(a) return ({[true]='true', [false]='false'})[a] end,
   ['function'] = function(a) return("valstr does not support functions") end,
   ['table'] = function(a) return('Valstr does not support tables') end})[type(a)](a)
end

strval = valstr

--[[
  
  This function was written by Wookai
  http://stackoverflow.com/questions/2282444/how-to-check-if-a-table-contains-an-element-in-lua
  
]]
function tablecontains(table, element)
  for _, value in pairs(table) do
    if value == element then
      return true
    end
  end
  return false
end





function isorcontains(v,e)
if type(v)=="table" then return tablecontains(v,e) end
return v==e
end     

--[[ The name of the person who came up with this is unknown,
      however he placed this script here:
      
      http://stackoverflow.com/questions/1426954/split-string-in-lua
      
]]      

function mysplit(inputstr, sep)
        if sep == nil then
                sep = "%s"
        end
        local t={} ; i=1
        for str in string.gmatch(inputstr, "([^"..sep.."]+)") do
                t[i] = str
                i = i + 1
        end
        return t
end

pper = string.upper
lower = string.lower
chr = string.char
printf = string.format
replace = string.gsub
rep = string.rep
substr = string.sub


function left(s,l)
return substr(s,1,l)
end

function right(s,l)
local ln = l or 1
local st = s or "nostring"
-- return substr(st,string.len(st)-ln,string.len(st))
return substr(st,-ln,-1)
end 

function mid(s,o,l)
local ln=l or 1
local of=o or 1
local st=s or ""
return substr(st,of,(of+ln)-1)
end 


function trim(s)
  return (s:gsub("^%s*(.-)%s*$", "%1"))
end
-- from PiL2 20.4

function findstuff(haystack,needle) -- BLD: Returns the position on which a substring (needle) is found inside a string or (array)table (haystrack). If nothing if found it will return nil.<p>Needle must be a string if haystack is a string, if haystack is a table, needle can be any type.
local ret = nil
local i
for i=1,len(haystack) do
    if type(haystack)=='table'  and needle==haystack[i] then ret = ret or i end
    if type(haystack)=='string' and needle==mid(haystack,i,len(needle)) then ret = ret or i end
    -- rint("finding needle: "..needle) if ret then print("found at: "..ret) end print("= Checking: "..i.. " >> "..mid(haystack,i,len(needle)))
    end
return ret    
end




function safestring(s)
local allowed = "qwertyuiopasdfghjklzxcvbnmmQWERTYUIOPASDFGHJKLZXCVBNM 12345678890-_=+!@#$%^&*():;/<>[]{}.,"
local i
local safe = true
local alt = ""
for i=1,len(s) do
    safe = safe and (findstuff(allowed,mid(s,i,1))~=nil)
    alt = alt .."\\"..string.byte(mid(s,i,1),1)
    end
-- print("DEBUG: Testing string"); if safe then print("The string "..s.." was safe") else print("The string "..s.." was not safe and was reformed to: "..alt) end    
local ret = { [true] = s, [false]=alt }
-- print("returning "..ret[safe])
return ret[safe]     
end 





-- Serializing
function TRUE_SERIALIZE(vname,vvalue,tabs,noenter)
local ret = ""
local work = {
                ["nil"]        = function() return "nil" end,
                ["number"]     = function() return vvalue end,
                ["function"]   = function() return "'!ERROR! -- I cannot handle functions!'" end,
                ["string"]     = function() return "\""..safestring(vvalue).."\"" end,
                ["boolean"]    = function() return ({[true]="true",[false]="false"})[vvalue] end,
                ["table"]      = function()
                                 local titype
                                 local tindex = {
                                                   ["number"]     = function(v) return v end,
                                                   ["boolean"]    = function(v) return ({[true]="true",[false]="false"})[v] end,
                                                   ["string"]     = function(v) return "\""..safestring(v).."\"" end
                                 }
                                 local wrongindex = function() print("!ERROR! Type "..titype.." can not be used as a table in serializing") return "!WRONGINDEX" end
                                 local ret = "{"
                                 local k,v
                                 local result
                                 local notfirst
                                 for k,v in pairs(vvalue) do
                                     if notfirst then ret = ret .. ",\n" else notfirst=true ret = ret .."\n" end
                                     titype = type(k)
                                     result = (tindex[titype] or wrongindex)(k)
                                     -- print(titype.."/"..k)
                                     ret = ret .. TRUE_SERIALIZE("["..result.."]",v,(tabs or 0)+1,true)
                                     end
                                 if notfirst then    
                                   ret = ret .."\n"    
                                   for i=1,(tabs or 0) do ret = ret .."\t" end   
                                   for i=1,len(vname.." = ") do ret = ret .. " " end
                                   end 
                                 ret = ret .. "}"  
                                 return ret  
                                 end 
                                   
             }
local letsgo = work[type(vvalue)] or function() print("!ERROR! Unknown type. Cannot serialize","Variable,"..vname..";Type Value,"..type(vvalue)) return "whatever" end
local i
for i=1,(tabs or 0) do ret = ret .."\t" end
if vname then 
   ret = ret .. vname .." = "..letsgo()
   else
   ret = letsgo()
   end 
if not noenter then ret = ret .."\n" end
return ret
end


function serialize(vname,variableitself)
local ret = ""
local v = variableitself 
if vname then 
   v = v or _G[vname] 
   if type(vname)~='string' then print("First variable must be the name to return in the serializer string") end
   end
ret = TRUE_SERIALIZE(vname,v)
-- JBCSYSTEM.Returner(ret)
return ret
end

function sval(a)
return 
  (({
     ['string']=function() return a end,
     ['number']=function() return a end,
     ['boolean']=function() if a then return 'true' else return 'false' end     end
  })[type(a)] or function() return "<< "..type(a).." >>" end)()
end  


-- All done, let's now load the main script and start it all up.
assert(RYANNA_MAIN_SCRIPT and RYANNA_MAIN_SCRIPT~="","There has no script been assigned as main script!")
Use(RYANNA_MAIN_SCRIPT)
`

	script["jcr6"] = `--[[
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
`

	/* Lua */ mkl.Version("Ryanna - Builder for jcr based love projects - basis.lua","17.12.30")

	/* Lua */ mkl.Lic    ("Ryanna - Builder for jcr based love projects - basis.lua","GNU General Public License 3")

	/* Lua */ mkl.Version("Ryanna - Builder for jcr based love projects - jcr6.lua","17.12.30")

	/* Lua */ mkl.Lic    ("Ryanna - Builder for jcr based love projects - jcr6.lua","GNU General Public License 3")



}