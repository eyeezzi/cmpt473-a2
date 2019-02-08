# Test specification for csv2json application
# csv2json(src STRING, noHeader BOOL, dst STRING)

Environments:

	Source file exists:
		true.
		false.

	
Parameters:

# File
	Source file type:
		# default stdin
		Stdin.
		Diskfile.

	Number of records in file:
		0.
		1.
		>1.

	All records have same number of fields:
		True.
		False.
	
	Existence of enclosed fields in records:
		True.
		False.
	
	Existence of field with DQUOTE or CRLF or COMMA:
		True.
		False.

# no-header flag
	No-header flag specified:
		True.
		False.

# Destination file
	Destination file type:
		#default = stdout
		Stdout.
		Diskfile.
