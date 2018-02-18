--[[
  jcr6.lua
  Ryanna - Script
  version: 18.02.18
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
local winspace
if RYANNA_LOAD_JCR then
	if platform == "Windows" then	  	
		jcrx = ldir.."\\jcrx.exe"
		jcrx = jcrx:gsub("/","\\")
		winspace="DIE_VIEZE_VUILE_FUCK_WINDOWS_HEEFT_EEN_SPATIEBALK_NODIG"
	elseif platform == "OS X" then
		jcrx = ldir.."//jcrx"
	elseif platform == "Linux" then
		error("Sorry! I'm still working this part out for Linux")
	else 
		error("Sorry! This game cannot work on "..platform..". At least not in the way it's currently build. And no, that will very likely never be possible either.\nBut fear not, as it may be possible in a different building type.\nNotify the developer and tell him/her of this error.")
	end
end
-- $FI

function Dir2JCR(jfile)
   if not RYANNA_LOAD_JCR then return false,"This is no JCR project" end
   local jcall = "'"..jcrx.."' transform '"..jfile.."'"
   local bt = io.popen(jcall)
   -- sl = bt:readlines()
   local sl = {}
   for rl in bt:lines() do sl[#sl+1]=rl end
   bt:close()
   local s = ""
   for i=2,#sl do s = s .. sl[i] .. "\n" end
   return sl[1]=="OK",s
end   

function JCR_Dir(jfile)
  local jcall
  if platform == "Windows" then
     jcall = '"'..jcrx..'" dirout '..jfile:gsub(" ",winspace).." lua"
  else
	   jcall = " \""..jcrx.."\" dirout \""..jfile.."\" lua "
	end
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
		if type(j)=='table' then
			mj = j
		else 
			mj = JCR_Dir(j)
		end
	end
	local e = string.upper(entry)
	local edata = mj.entries[e]
  assert(edata,"Entry "..entry.." not found")
	if edata.aliasfor then 
	   local af=edata.aliasfor	   
	   edata=mj.entries[edata.aliasfor:upper()] 
	   assert(edata,"Entry '"..af.."' from alias "..entry.." not found")
	   entry=af 
	end
	--print(serialize('jcr',mj)) -- debug line
	
	if not edata then return end -- Make sure nothing bad happens in case of a pcall
	if edata.LOVE then
		local rets = love.filesystem.read(edata.LOVE)
    if not lines then return rets end
	  local rett = mysplit(rets,"\n")
	  return rett
	end
	local bt
	if platform=='Windows' then
	  bt = io.popen('"'..jcrx..'" typeout '..mj.from:gsub(" ",winspace).." "..entry:gsub(" ",winspace))
	else 
	  bt = io.popen("'"..jcrx.."' typeout '"..mj.from.."' '"..entry.."'")
	end
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

function JCR_TRUEEXTRACT(arc,source,target)
    local cmd="'"..jcrx.."' extract '"..arc.."' '"..source.."' '"..target.."'"
    local bt = io.popen(cmd)
    assert(bt,"Pipe open failed in extraction request")
    local sl={}
    for rsl in bt:lines() do sl[#sl+1]=rsl end 
    bt:close()
    assert(sl[1]=="OK",(sl[2] or sl[1] or "Unknown error from jcrx").." \n>"..cmd)    
end

function JCR_Extract(p1,p2,p3)
  local data
  if type(p1)=='table' and type(p2)=='string' and type(p3)=='string' then
   love.filesystem.write(JCR_B(p1,p2),p3)
   return
  elseif type(p1)=='string' and type(p2)=='string' and type(p3)=='nil' then
   love.filesystem.write(JCR_B(p1),p2)    
   return
  elseif type(p1)=='string' and type(p2)=='string' and type(p3)=='string' then
   if suffixed(p1,".jcr") then
      local d = ""
      if not prefixed(p3,"/") then d=love.filesystem.getSaveDirectory( ).."/" end
      JCR_TRUEEXTRACT(p1,p2,d..p3)
   else
      love.filesystem.write(JCR_B(p1,p2),p3)   
   end         
  else
     error("JCR_Extract("..sval(p1)..","..sval(p2)..","..sval(p3).."): Invalid parameters input")
  end        
end

function JCR_GetDir(p1,p2,p3)
   local mj,dir,trimpath = p1,p2,p3
   if p3==nil then mj,dir,trimpath=jcr,p1,p2 end
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
  local ret
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
	if love.filesystem.isFile('alias.data') then
	   local aliasstring = love.filesystem.read('alias.data')
	   local aliaslines = mysplit(aliasstring,"\n")
	   for i,aline in ipairs(aliaslines) do
	       local asplit = {""}
	       local ai = 1
	       local wait=0
	       for i=1,#aline do
	           if mid(aline,i,4)==" => " then
	              ai=2
	              asplit[2]=""
	              wait=3
	           elseif wait>0 then
	              wait=wait-1
	           else
	              asplit[ai]=asplit[ai]..mid(aline,i,1)
	           end   
	       end
	       local source = trim(asplit[1]):upper()
	       local puretarget = trim(asplit[2])
	       local target = puretarget:upper()
	       assert(source,"Syntax error in alias file line: "..i)
	       assert(target,"Syntax error in alias file line: "..i.."   '=>' expected")
	       assert(ret.entries[source],"Alias error -- Original ("..source..") doesn't exit ("..aline..") line: "..i)
	       if ret.entries[target] then print("WARNING! Alias will overwrite existing entry '"..target.."' (line: "..i..")") end
	       ret.entries[target]={ aliasfor=source}
	       --[[
	       for k,v in pairs(ret.entries[source]) do ret.entries[target][k] = v end
	       ret.entries[target].entry = puretarget
	       ]]
	       print("Artificially aliased "..source.." into "..target)
	       --print(serialize("jcr",ret)) -- debug
	   end
	end   
	return ret
end

-- for k,e in pairs(ret.entries) do print(k) end -- debug


--[[
mkl.version("Ryanna - Builder for jcr based love projects - jcr6.lua","18.02.18")
mkl.lic    ("Ryanna - Builder for jcr based love projects - jcr6.lua","ZLib License")
]]
