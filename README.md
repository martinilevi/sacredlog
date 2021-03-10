# sacredlog

Converts json logs in format below to plain CSV-like format. (corr is omitted)

```
{
	"time":""
	"corr":""
	"lvl":""
	"shortfile":""
	"msg":""
}
```

Try the example

```
cat samplefile.txt | go run sacredlog.go
```
