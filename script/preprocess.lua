--[[
  preprocess.lua
  
  version: 18.01.02
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

local prid = {
	["IF"] = function(sl,h,m,n,ld)
		assert(not h,"Duplicate $IF in line: "..n)
		assert(#sl>2,"$IF expects statements in line: "..n)
		local pline = "local chk = {}\n"
		-- global defs
		for k,v in pairs(defs) do pline = pline .. "chk['"..k.."'] = true\n" end -- This makes sure we got all locals in our little function.
		-- localdefs defs
		for k,v in pairs(ld) do pline = pline .. "chk['"..k.."'] = true\n" end -- This makes sure we got all locals in our little function.
		pline = "\n\nreturn "
		for i=3,#sl do
			w=string.upper(sl[i])
			if w=="OR" or w=="AND" then pline = pline .. string.lower(w) .. " "
			elseif prefixed(w,1)=="!" then pline = pline .. " (not chk['"+w+"']) "
			else   pline = pline .. "(chk['"+w+"') " end
		end
		ok,chkf = pcall(load(pline,"$IF"))
		if not ok then
			print("$IF went wrong in line: "..n)
			print("-- GENERATED CODE --")
			print(pline)
			print("-- END CODE --")
			print("error: "..chkf)
			error("Invalid $IF call in line: "..n)
		end
		local mute = not chkf()
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
	local d = JCR_Lines(file)
	local haveif
	local muteif
	local ret = ""
	local localdefs = {}
	for lnum,line in ipairs(d) do
		sline = mysplit(trim(line))
		if prefixed(trim(line),"-- $") then
			sline[2]=strings.upper(sline[2])
			cmd = strings.upper(sline[2])
			assert(prid[cmd],"UNKNOWN PRE-PROCESSOR DIRECTIVE in line "..lnum.." ("..cmd..")")
			haveif,muteif,rl = prid[cmd](sline,haveif,muteif,lnum,localdefs)
			ret = ret .. (rl or line) .. "\n"
		elseif haveif and muteif then
			ret = ret .. "-- muted: "..line.."\n" -- This takes extra ram, but this way the line numbers in traceback routines remain.
		else
			ret = ret .. line .."\n"
		end
	end
	return ret
end
