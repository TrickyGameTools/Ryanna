--[[
  use.lua
  Ryanna - Script
  version: 18.06.06
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
mkl.version("Ryanna - Builder for jcr based love projects - use.lua","18.06.06")
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
	local debug --= true
	local wimp = string.upper(imp)
	local ret
	local inits={}
	local function dc(txt) if debug then CSay("DEBUG> "..txt) end end
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
				  if k2=='init' then inits[#inits+1]=v2 dc("Init caught in: "..k) else
				     dc("In "..k.." there is a "..type(v2).." named "..k2)
					   if ret[k2] then print("WARNING! Duplicate identifier '"..k2.."' found!") end
					   ret[k2] = v2
					end   
				end
			end
		else
			ret[k] = v
		end
	end
	--for k,_ in spairs(ret) do print("= "..k) end -- debug line
	for f in each(inits) do f(ret) dc("Initcall") end
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

destroylib=libdestroy
