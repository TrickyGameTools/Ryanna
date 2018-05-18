--[[
  main.lua
  
  version: 18.04.21
  Copyright (C) 2017, 2018 Jeroen P. Broks
  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.
  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:
  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
]]
-- basis script

--[[
mkl.version("Ryanna - Builder for jcr based love projects - main.lua","18.04.21")
mkl.lic    ("Ryanna - Builder for jcr based love projects - main.lua","ZLib License")
]]


RYANNA_MAIN_SCRIPT = "$RyannaMainScript$"
RYANNA_LOAD_JCR    = "$RyannaLoadJCR$"     -- quotes will be removed. I've set it up as a string to deceive parse error checking IDEs, as they would otherwise go crazy.
RYANNA_TITLE       = "$RyannaTitle$"; love.window.setTitle(RYANNA_TITLE)
RYANNA_BUILDTYPE   = "$RyannaBuildType"    -- Will contain 'normal' in normal builds and 'test' in test builds. Handy for extra debugging features in Ryanna.

platform = love.system.getOS( )

Ryanna = {
	RyannaVersion = "$RyannaVersion$",
	LuaVersion = _VERSION,
	LoveVersion = string.format("%d.%d.%d - %s",love.getVersion() ) -- This line is dirty code straight from the toilet, but I don't care :P
	
}


function OlderLove(maj,min,sub,selfinc) -- returns true if the love version is older than the set number
   local lvmaj,lvmin,lvsub,lvcn = love.getVersion()
   if lvmaj<maj then return true end
   if lvmin<min and lvmaj==maj then return true end
   if lvsub<sub and lvmin==min and lvmaj==maj then return true end
   return selfinc and lvsub==sub and lvmin==min and lvmaj==maj 
end

function NewerLove(maj,min,sub,selfinc) -- returns true if the love version is newer than the set number
   local lvmaj,lvmin,lvsub,lvcn = love.getVersion()
   if lvmaj>maj then return true end
   if lvmin>min and lvmaj==maj then return true end
   if lvsub>sub and lvmin==min and lvmaj==maj then return true end
   return selfinc and lvsub==sub and lvmin==min and lvmaj==maj 
end


function TablePack(tab)
   local max=0
   for i,_ in pairs(tab) do 
       if i>max then max=i end
   end
   if max==0 then return end
   local temptab = {}
   for i=1,max do
       if tab[i] then temptab[#temptab+1]=tab[i] tab[i]=nil end
   end
   for i,v in ipairs(temptab) do tab[i]=v end
end

love.filesystem.isDir = love.filesystem.isDirectory

-- include use.Lua and jcr6.lua which now have not made their official entrace, so I gotta call them manually
function load_primary_dependencies()
	for _,dep in ipairs( {"jcr6.lua","use.lua","preprocess.lua"} ) do
		chunk, errormsg = love.filesystem.load( dep )
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



function secu_each(b) -- This will copy all elements of a table to a local table before processing. This takes a bit more time (prior to the workout and after it when the extra table has to be released), but is safer to use. Very useful when the table you wanna process is being altered during the process (like removing elements), but you're not wanting to make that influence the iteration. So if you want speed, use each. When you want a safe approach, use this.
	local i=0
	local a={}
	for i,e in ipairs(b) do a[i]=e end
	if type(a)~="table" then
	print("Each received a "..type(a).."!",255,0,0)
		return nil
	end
	return function()
		i=i+1
		if a[i] then return a[i] end
	end
end

function copytable(original,recuse)
   local ret = {}
   for k,v in pairs(original) do
       if type(v)=="table" and recurse then ret[k]=copytable(v,true)
       else ret[k]=v 
       end
   end    
   return ret    
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
    "for i=x,y,z do" is normally translated to BASIC as "for i=x to y step z"
    This basically translates to "for i=x until y step z"
    Of course I know "y-1" is an option, but that is only safe when using integers. When using non-integers, this is not the safest route to go
]]
function urange(start,einde,stappen)
	local i=start
	return function()
		if not i<einde then return nil end
		ret = i
		i = i + stappen
		return ret
	end
end

--[[

    This function is written by Michal Kottman.
    http://stackoverflow.com/questions/15706270/sort-a-table-in-lua

]]
function spairs(t, order)
    -- collect the keys
    local keys = {}
    local t2 = {}
    for k,v in pairs(t) do keys[#keys+1] = k  t2[k]=v end

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
            return keys[i], t2[keys[i]]
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

function StripDir(f)
   local mf = mysplit(f,"/")
   return mf[#mf]
end

-- This routine will need some more work to be more accurate, but for now it'll do.
function StripExt(f)
    local mf=mysplit(f,".")
    if #mf==1 then return f end
    local ret=mf[1]
    for i=2,#mf-1 do ret = ret .. "."..mf[i] end
    return ret
end

function StripAll(f)
    return StripExt(StripDir(f))
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

upper = string.upper
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
for i=1,#haystack do
    if type(haystack)=='table'  and needle==haystack[i] then ret = ret or i end
    if type(haystack)=='string' and needle==mid(haystack,i,#needle) then ret = ret or i end
    -- rint("finding needle: "..needle) if ret then print("found at: "..ret) end print("= Checking: "..i.. " >> "..mid(haystack,i,len(needle)))
    end
return ret    
end

function prefixed(str,prefix)
	return left(str,#(prefix..""))==prefix
end

function suffixed(str,suffix)
	return right(str,#(suffix..""))==suffix
end



function safestring(s)
local allowed = "qwertyuiopasdfghjklzxcvbnmmQWERTYUIOPASDFGHJKLZXCVBNM 12345678890-_=+!@#$%^&*():;/<>[]{}.,"
local i
local safe = true
local alt = ""
for i=1,#s do
    safe = safe and (findstuff(allowed,mid(s,i,1))~=nil)
    alt = alt .."\\"..string.byte(mid(s,i,1),1)
    end
-- print("DEBUG: Testing string"); if safe then print("The string "..s.." was safe") else print("The string "..s.." was not safe and was reformed to: "..alt) end    
local ret = { [true] = s, [false]=alt }
-- print("returning "..ret[safe])
return ret[safe]     
end 

function round(a)
   if a>0 then
      return math.floor(a+.5)
   else 
      return math.ceil(a-.5)
   end
end       



-- Serializing
function TRUE_SERIALIZE(vname,vvalue,tabs,noenter)
local ret = ""
local len = function(s) return #s end
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


function cleartable(tab,recurse)
   -- Warning about "recurse". If your table contains a 'cyclic reference' an infinite loop may be caused eventually leading to a stack overflow.
   -- I did write some code trying to prevent this, but I cannot guarantee this.
   -- Also if table references are used in other tables, 'recurse' will clear those tables as well!, so keep in mind that using 'recurse' should ONLY be done when you 'KNOW' what you are doing!
   local keys = {}
   --local temptab
   -- gathering keys! Setting things to nil already while using pairs on the same table is bound to make Lua to malfunction
   for key,value in pairs(tab) do
       keys[#keys+1]=key
   end
   -- And now it's time for destruction!
   for key in each(keys) do
       local value=tab[key]
       tab[key]=nil 
       if type(value)=='table' and recurse then
          cleartable(value)
       end   
   end    
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


jcr = BaseDir()

-- All done, let's now load the main script and start it all up.
assert(RYANNA_MAIN_SCRIPT and RYANNA_MAIN_SCRIPT~="","There has no script been assigned as main script!")
Use(RYANNA_MAIN_SCRIPT)


