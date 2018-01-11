package main

import "trickyunits/mkl"

// Licensed under the GNU

func init(){
	script["preprocess"] = `--[[
  preprocess.lua
  
  version: 18.01.11
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


preiftrueerror = false

local defs = {}
local plat = string.upper(love.system.getOS( ))
if plat == "OS X" then defs["$MAC"] = true defs["$DARWIN"]=true defs["$OSX"]=true end
if plat == "WINDOWS" then defs["$WIN"] = true defs["$WINDOWS"] = true defs["$WINDHOOS"] = true defs["$MICROSCHOFT"] = true end
if plat == "LINUX" then defs["$LINUX"] = true defs["$INSTABIEL"] = true end
if plat == "ANDROID" then defs["$ANDROID"] = true defs["$MOBILE"] = true end
if plat == "IOS" then defs["$IOS"] = true defs["$MOBILE"]=true end
if RYANNA_BUILDTYPE=='test' then defs["$TEST"]=true end

local prid = {
	["IF"] = function(sl,h,m,n,ld)
		assert(not h,"Duplicate $IF in line: "..n)
		assert(#sl>2,"$IF expects statements in line: "..n)
		local pline = "local chk = {}\n"
		-- global defs
		for k,v in pairs(defs) do pline = pline .. "chk['"..k.."'] = true\n" end -- This makes sure we got all locals in our little function.
		-- localdefs defs
		for k,v in pairs(ld) do pline = pline .. "chk['"..k.."'] = true\n" end -- This makes sure we got all locals in our little function.
		pline = pline .."\n\nreturn "
		for i=3,#sl do
			local w=string.upper(sl[i])
			if w=="OR" or w=="AND" then pline = pline .. string.lower(w) .. " "
			elseif prefixed(w,1)=="!" then pline = pline .. " (not chk['"+w+"']) "
			else   pline = pline .. "(chk['"..w.."']) " end
		end
		local ok,chkf = pcall(load(pline,"$IF"))
		if not ok then
			print("$IF went wrong in line: "..n)
			print("-- GENERATED CODE --")
			print(pline)
			print("-- END CODE --")
			print("error: "..chkf)
			error("Invalid $IF call in line: "..n)
		end
		local mute = not chkf
		return true,mute
	end,
	
	["ELSE"] = function(sl,h,m,n,ld)
		assert(h,"$ELSE without $IF in "..n)
		return true,not m
	end,
	
	["FI"] = function(sl,h,m,n,ld)
		assert(h,"$FI without $IF in "..n)
		return false,false
	end,
	
	["DEFINE"] = function(sl,h,m,n,ld)
		if h and m then return h,m end
		assert(#sl>2,"$DEFINE expects options in line: "..n)
		d = string.upper(sl[3])
		if prefixed(d,"#") then
			ld[d]=true
		else
			defs[d]=true
		end
		return h,m
	end,

	["UNDEF"] = function(sl,h,m,n,ld)
		if h and m then return h,m end
		assert(#sl>2,"$UNDEF expects options in line: "..n)
		d = string.upper(sl[3])
		if prefixed(d,"#") then
			ld[d]=nil
		else
			defs[d]=nil
		end
		return h,m
	end,
	
	["USE"] = function(sl,h,m,n,ld)
		if h and m then return h,m end
		assert(#sl>2,"$USE expects libraries in line: "..n)		
		local bfs = mysplit(sl[3],"/")
		local bf = bfs[#bfs]
		local exs = mysplit(bf,".")
		local id = exs[1]
		local x = string.upper(sl[4] or "")
		local pre = id .. " = "..id.." or "
		if x == "NOAS" then
			pre = ""
		end
		if x == "AS" then
			assert(sl[5],"AS without identifier")
			asid = sl[5]
			if prefixed(asid,"#") then asid = right(asid,#asid-1) pre = "local " else pre = " " end
			pre = pre .. asid .. " = " ..asid.. " or "
		end
		return h,m,pre .. " Use('"..sl[3].."') "
	end

}

function PreProcess(file)
  local debug = false
	local d = JCR_Lines(file)
	local haveif
	local muteif
	local ret = ""
	local localdefs = {}
	print("Compiling: "..file)
	if type(d)=='string' then
	   --print(d)
  end	   
	for lnum,line in ipairs(d) do
	  if debug then print ("Processing line: "..lnum.."> "..line) print (prefixed(trim(line),"-- $")) end
		local sline = mysplit(trim(line))
		if prefixed(trim(line),"-- $") then
		  if debug then print("compiler directivefound: "..line) end
			sline[2]=string.upper(sline[2])
			local cmd = string.upper(sline[2])
			local cmd = right(cmd,#cmd-1)
			assert(prid[cmd],"UNKNOWN PRE-PROCESSOR DIRECTIVE in line "..lnum.." ("..cmd..")")
			haveif,muteif,rl = prid[cmd](sline,haveif,muteif,lnum,localdefs)
			ret = ret .. (rl or line) .. "\n"
		elseif haveif and muteif then
			ret = ret .. "-- muted: "..line.."\n" -- This takes extra ram, but this way the line numbers in traceback routines remain.
		else
			ret = ret .. line .."\n"
		end
	end
	if debug then print(ret) end
	local f = load(ret,file)
	return f()
end
`

	script["jcr6"] = `--[[
  jcr6.lua
  Ryanna - Script
  version: 18.01.12
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
local ldir = love.filesystem.getSourceBaseDirectory()

local jcrx
if RYANNA_LOAD_JCR then
	if platform == "Windows" then
		jcrx = ldir.."\\jcrx"
	elseif platform == "OS X" then
		jcrx = ldir.."//jcrx"
	elseif platform == "Linux" then
		error("Sorry! I'm still working this part out for Linux")
	else 
		error("Sorry! This game cannot work on "..platform..". At least not in the way it's currently build. And no, that will very likely never be possible either.\nBut fear not, as it may be possible in a different building type.\nNotify the developer and tell him/her of this error.")
	end
end
-- $FI

function JCR_Dir(jfile)
	local jcall = "'"..jcrx.."' dirout '"..jfile.."' lua"
	print ("debug> ",jcall)
	local bt = io.popen(jcall)
	-- sl = bt:readlines()
	local sl = {}
	for rl in bt:lines() do sl[#sl+1]=rl end
	bt:close()
	assert(sl[1]=="OK","JCR-Dirout failure "..jfile.."\n"..(sl[2] or sl[1] or "No error message provided"))
	local s = ""
	for i=2,#sl do s = s .. sl[i] .. "\n" end
	local f=load(s,"JCR_DIR("..jfile..")")
	local ret={}
	assert(f,'Error parsing directory data')
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
			entries[string.upper(f)] = { entry = f, LOVE = f, mainfile = love.filesystem.getSource() }
		end
	end
	local ret = { entries = entries, from = love.filesystem.getSource(), kind="LOVE" }
	return ret
end

function LOVE_FullDir(adir) -- recursive dir
  local dir = adir or ""
  local list = love.filesystem.getDirectoryItems( dir or "" )
  local entries = {}
  local slash = "/"; if dir=="" then slash="" end
  for i,f in ipairs(list) do
      entries[string.upper(dir..slash..f)] = { entry = dir..slash..f, LOVE = dir..slash..f, mainfile = love.filesystem.getSource() }
      if love.filesystem.isDir(dir..slash..f) then 
         local te = LOVE_FullDir(dir..slash..f)
         for k,d in pairs(te.entries) do entries[k]=d end
      end
  end
  local ret = { entries = entries, from = love.filesystem.getSource(), kind="LOVE" }
  return ret

end 

function JCR_B(j,nameentry,lines)
	local mj,entry
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
	local e = string.upper(entry)
	local edata = mj.entries[e]
	--print(serialize('jcr',mj)) -- debug line
	assert(edata,"Entry "..entry.." not found")
	if not edata then return end -- Make sure nothing bad happens in case of a pcall
	if edata.LOVE then
		local rets = love.filesystem.read(edata.LOVE)
    if not lines then return rets end
	  local rett = mysplit(rets,"\n")
	  return rett
	end
	local bt = io.popen("'"..jcrx.."' typeout '"..mj.from.."' '"..entry.."'")
	if lines then
	 -- sl = bt:readlines()
	 local sl = {}
	 local s
	 for rsl in bt:lines() do sl[#sl+1]=rsl end 
    bt:close()
	 assert(sl[1]=="OK",sl[2] or sl[1] or "Unknown error from jcrx")
	 -- if lines then
		  s = {}
		  for i=2,#sl do s[#s+1] = sl[i] end
	 --[[ else
		  s = ""
		  for i=2,#sl do s = s .. sl[i] .. "\n" end
	 --end ]]
	 return s
	else
	  local head=bt:read(3)
	  local data=bt:read('*all')
	  assert(head=="OK\n","JCR_B failed "..(data or "Unprintable error"))
	  bt:close()
	  --print(data)
    return data
	end	
end

function JCR_GetDir(p1,p2,p3)
   local mj,dir,trimpath = p1,p2,p3
   if not p3 then mj,dir,trimpath=jcr,p1,p2 end
   local cd = upper(dir)
   local ret = {}
   for k,v in pairs(mj.entries) do
       if prefixed(k,cd) then
          local ename = v.entry
          if trimpath then ename=right(ename,#ename-#dir) end
          ret[#ret+1] = ename
       end
   end
   return ret
end

function JCR_D(file)
     local data = JCR_B(file)
     local fdata = love.filesystem.newFileData(data,file)
     return fdata
end

function JCR_Lines(j,nameentry)
	return JCR_B(j,nameentry,true)
end

function JCR_Exists(j,nameentry)
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
	local e = string.upper(entry)
	local edata = mj.entries[e]
	return edata~=nil
end

function JCR_HasDir(j,namedir)
	local mj,dir
	if not namedir then
		dir= string.upper(j)
		mj = jcr
		assert ( mj , "JCR not set!" )
	else
		dir= string.upper(namedir)
		if type(mj)=='table' then
			mj = j
		else 
			mj = JCR_Dir(j)
		end
	end
	if not suffixed(dir,"/") then dir = dir .. "/" end
	for ent,_ in pairs(mj.entries) do
	  -- print(ent,"\t",dir) -- debug line. Must be on comment if not in use!
		if prefixed(ent,dir) then return true end
	end
	return false
end



function BaseDir() -- Basically only called by Ryanna and loaded based on Ryanna's findings.
	ret = {}
	ret.entries = {}
	ret.from = love.filesystem.getSource()
	ret.kind = "MIXED"	
	local k = {}
	k[1] = LOVE_FullDir()
	if RYANNA_LOAD_JCR then k[2] = JCR_Dir(ret.from) end
	for i,d in ipairs(k) do
		for key,res in pairs(d) do 
		  if key=="entries" then
		    for ekey,edata in pairs(res) do
		      ret.entries[ekey] = edata 
		      --print("Adding "..i..": "..ekey)
		    end -- for ekey,edata
		  end -- if key==entres
		end -- for key,res  
	end -- for i,d
	return ret
end
jcr = BaseDir()
-- for k,e in pairs(ret.entries) do print(k) end -- debug


--[[
mkl.version("Ryanna - Builder for jcr based love projects - jcr6.lua","18.01.12")
mkl.lic    ("Ryanna - Builder for jcr based love projects - jcr6.lua","ZLib License")
]]
`

	script["use"] = `--[[
  use.lua
  Ryanna - Script
  version: 18.01.12
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
-- Importer

--[[
mkl.version("Ryanna - Builder for jcr based love projects - use.lua","18.01.12")
mkl.lic    ("Ryanna - Builder for jcr based love projects - use.lua","ZLib License")
]]



-- Now in stead of using this function, you can also make a call to the 
-- preprocessor. Especially for external libraries this is recommended 
-- As that will cause cause Ryanna to import them into the .love file
-- automatically (that happens for all calls made to libs/
-- .lua or .rel are allowed (to make sure stuff is found correctly,
-- but is not required)
function Use(imp,noreturn)
	-- single file
	local wimp = string.upper(imp)
	local ret
	if right(wimp,4)==".LUA" then
		ret = PreProcess(imp)
		if noreturn then return nil else return ret end
	end
	if JCR_Exists(imp..".lua") then return Use(imp..".lua",noreturn) end
	-- Is this a Ryanna Expanded Library then?
	if right(wimp,4)~=".REL" then
		assert(JCR_HasDir(imp..".rel"),"Nothing appears to match the use request: "..imp)
		return Use(imp..".rel",noreturn)
	end
	-- Import all the data
	local pret = {} -- pre return
	local name
	--print (serialize('jcr',jcr))
	for ename,entry in spairs(jcr.entries) do
	  if JCR_Exists(wimp.."/RyannaBuild.gini") then
	     local l=JCR_Lines(wimp.."/RyannaBuild.gini")
	     local req=nil
	     local rqs="require="
	     for i,ln in ipairs(l) do
	         if prefixed(ln:lower(),rqs) then
	            req=wimp.."/"..right(ln,#ln-#rqs)
	         end   
	     end
	     assert(req,"Special unit has not file to call set!")
	     return Use(req)
	  end   
		if prefixed(ename,wimp.."/") and suffixed(ename,".LUA") then
			name = right(entry.entry,#entry.entry-(#imp+1))
			name = left(entry.entry,#entry.entry-4)
			local allow=true
			local dn = mysplit(lower(name),"__")
			local plat = string.upper(love.system.getOS( ))
			if #dn>1 then
			   for i=2,#dn do
			       if dn[i]=="ignore" then allow=false end
			       if dn[i]=="windows" then allow=allow and plat=="WINDOWS" end
             if dn[i]=="osx" or dn=="darwin" or dn=="macos" or dn=="mac " then allow=allow and plat=="OS X" end
             if dn[i]=="linux" then allow = allow and plat=="LINUX" end
             if dn[i]=="android" then allow = allow and plat=="ANDROID" end
             if dn[i]=="ios" then allow = allow and plat=='IOS' end
             if dn[i]=="mobile" then allow = allow and ( plat=='IOS' or plat=="ANDROID") end -- I know Windows can be mobile too, but as LOVE does not take that possibility into account, I'm sorry for Windows phone users. Blame either microsoft of the love crew for that, but not me!
			   end
			end
			if allow then 
			   pret[name] = PreProcess(entry.entry)
			else print("Skipping: "..entry.entry) -- debug   
			end   
		end
	end
	-- Count it all
	local cnt = 0
	local lk
	for k,v in pairs(pret) do cnt = cnt + 1   lk = k end
	assert(cnt>0,"Ryanna Expanded Library is empty: "..imp)
	if noreturn then return end
	if cnt==1 then
		return pret[lk]
	end
	-- process the returning result
	ret = {}
	ret.me = ret
	for k,v in pairs(pret) do
		if type(v)=="table" then
			if v.nomerge then
			  local ks  = mysplit(k,"/")
			  local key = ks[#ks]
				if v.me then print("WARNING! 'me' field set in module part.") else v.me = ret end
				ret[key] = v
				--print('Sub '..key..' added')
			else 
				for k2,v2 in pairs(v) do
					if ret[k2] then print("WARNING! Duplicate identifier '"..k2.."' found!") end
					ret[k2] = v2
				end
			end
		else
			ret[k] = v
		end
	end
	--for k,_ in spairs(ret) do print("= "..k) end -- debug line
	return ret
end


-- If you really need to destroy a module this is the wisist way to do so
-- before you set it to 'nil'. Otherwise stuff will not be picked up by 
-- Lua's garbage collector correctly causing "big boom" in the process.
-- In other words, memory leaks.
function libdestroy(lib)
	if type(lib)~='table' then return end -- Only needed for tables
	if not lib.me then return end
	lib.me = nil
	local fld = {}
	for key,v in pairs(lib)  do libdestroy(v); fld[#fld+1]=key end
	for i,key in ipairs(fld) do lib[key] = nil end
end
`

	script["main"] = `--[[
  main.lua
  
  version: 18.01.12
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
mkl.version("Ryanna - Builder for jcr based love projects - main.lua","18.01.12")
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

	/* Lua */ mkl.Version("Ryanna - Builder for jcr based love projects - jcr6.lua","18.01.12")

	/* Lua */ mkl.Lic    ("Ryanna - Builder for jcr based love projects - jcr6.lua","ZLib License")

	/* Lua */ mkl.Version("Ryanna - Builder for jcr based love projects - use.lua","18.01.12")

	/* Lua */ mkl.Lic    ("Ryanna - Builder for jcr based love projects - use.lua","ZLib License")

	/* Lua */ mkl.Version("Ryanna - Builder for jcr based love projects - main.lua","18.01.12")

	/* Lua */ mkl.Lic    ("Ryanna - Builder for jcr based love projects - main.lua","ZLib License")



}