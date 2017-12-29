from glob import glob

output = "package main\n\nimport \"trickyunits/mkl\"\n\n// Licensed under the GNU\n\nfunc init(){\n"

files = glob("*")

for myfile in files:
	if myfile[-4:]=='.lua':
		print "Converting: ",myfile
		vname=myfile[:-4]
		output += "\tscript[\""+vname+"\"] = `"
		mkl=""
		with open (myfile, "r") as bt:
			datalines=bt.readlines()
		for data in datalines:
			output += data
			if data[:11]=="mkl.version":
				mkl += data.replace("mkl.version","\t/* Lua */ mkl.Version")
				mkl += "\n"
			if data[:7]=="mkl.lic":
				mkl += data.replace("mkl.lic","\t/* Lua */ mkl.Lic")
				mkl += "\n"
		output += "`\n\n"
output += mkl
output += "\n\n}"

outfile = open("../src/luascript.go","w")
outfile.write(output)
outfile.close()

