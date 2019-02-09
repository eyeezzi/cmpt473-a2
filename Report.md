# CSV Specification 

MIME: txt/csv
RFC4180 https://tools.ietf.org/html/rfc4180

## BNF Grammar

```
file = [header CRLF] record *(CRLF record) [CRLF]

header = name *(COMMA name)

record = field *(COMMA field)

name = field

field = (escaped / non-escaped)

escaped = DQUOTE *(TEXTDATA / COMMA / CR / LF / 2DQUOTE) DQUOTE

non-escaped = *TEXTDATA

COMMA = %x2C

CR = %x0D ;as per section 6.1 of RFC 2234 [2]

DQUOTE =  %x22 ;as per section 6.1 of RFC 2234 [2]

LF = %x0A ;as per section 6.1 of RFC 2234 [2]

CRLF = CR LF ;as per section 6.1 of RFC 2234 [2]

TEXTDATA =  %x20-21 / %x23-2B / %x2D-7E
```

### Questions

1. How many tests did you generate? 8.
2. How many of these tests were successful/passing? 3 
3. How many tests would have been generated if you didn't use pairwise testing? 32.
	> Using the ACTS flag -Dcombine=all and optionally -Dchandler=no
4. What tradeoffs did you make as a result of pairwise testing?

One benefit of pairwise test is the reduction in the number of tests generated. The downside is that interactions of t >= 3 parameters are ignored. For example the generated pairwise tests does not cover the case below.

```
Src_File_Exists,No_Header,Src,Dest,File_With_Header,Number_Of_Records,Field_Type_In_Record,Same_Field_Count_Per_Record
TRUE,FALSE,DISKFILE,DISKFILE,FALSE,GTZERO,DONTCARE,FALSE
```

This implies that bugs that might be caused by a combination of 3 or more factors go untested. In the above example, we would not know what happens with a csv file containing mixed records.