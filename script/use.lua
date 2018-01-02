--[[
  use.lua
  Ryanna - Script
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
-- Importer

--[[
mkl.version("Ryanna - Builder for jcr based love projects - use.lua","18.01.02")
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
	wimp = string.upper(imp)
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
	pret = {} -- pre return
	for ename,entry in spairs(jcr.Entries) do
		if prefixed(ename,wimp+"/") and suffixed(ename,".LUA") then
			name = right(entry.entry,#entry.entry-(#imp+1))
			name = left(entry.entry,#entry.entry-4)
			pret[name] = PreProcess(entry.entry)
		end
	end
	-- Count it all
	cnt = 0
	for k,v in pairs(pret) do cnt = cnt + 1   lk = k end
	assert(cnt>0,"Ryanna Expanded Library is empty: "+imp)
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
				if v.me then print("WARNING! 'me' field set in module part.") else v.me = ret end
				ret[k] = v
			else 
				for k2,v2 in pairs(v) do
					if ret[k2] then print("WARNING! Duplicate identifier '"+k2+"' found!") end
					ret[k2] = v2
				end
			end
		else
			ret[k] = v
		end
	end
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
