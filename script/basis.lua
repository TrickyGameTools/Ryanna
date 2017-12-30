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
-- basis script

--[[
mkl.version("Ryanna - Builder for jcr based love projects - basis.lua","17.12.30")
mkl.lic    ("Ryanna - Builder for jcr based love projects - basis.lua","GNU General Public License 3")
]]


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
