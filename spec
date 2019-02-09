# Test specification for csv2json application
# csv2json(src STRING, noHeader BOOL, dst STRING)

Environments:

	Source file exists:
		True.							[property fileExists]
		False.							[Error][single]

	
Parameters:

# no-header flag
	No-header flag specified:
		True.							[if fileExists][property noHeader][single]
		False.							[if fileExists]

# Source file
	
	Source file type:
		# default stdin
		Stdin.
		Diskfile.						[if fileExists]
	Number of records in file:
		0.								[property noRecord][single]
		>0.								[property hasRecords]
	Size:
		Empty.							[if noHeader && noRecord][single]			
		Not Empty.						[if !noHeader || !noRecord]
	All records have same number of fields:
		True.							[if !noRecord] [property validFieldContent]
		False.							[if !noRecord][single]
	Number of fields per record:
		1.								[if validFieldContent]
		>1.								[if validFieldContent]
	Existence of enclosed fields in records:
		True.							[if hasRecords] [property hasEnclosedFields]
		False.							[if hasRecords]
	Existence of field with DQUOTE or CRLF or COMMA:
		True.							[if hasRecords && hasEnclosedFields]
		False.							[if hasRecords]



# Destination file
	Destination file type:
		#default = stdout
		Stdout.
		Diskfile.
