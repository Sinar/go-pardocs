# go-pardocs
Tools to process published Parliament documents (PDFs only) into more accessible form. Spiritual successor of https://github.com/leowmjw/parliamentMY-QA-blast


Assumes: 
- OSX dev environment
- Go v1.13 and above (uses go mod)
- Release for Linux, Windows, OSX available as cross-compile

## Usage

### Planning
```bash
$ ./go-pardocs plan -session <name> -type <L|BL> [-force] [-dir <workspace>] <sourcePDFPath>
Example:
	./go-pardocs plan -session par14sesi1 -type L ./raw/Lisan/JDR12032019.pdf
	./go-pardocs plan -session par13sesi3 -type L ./raw/Lisan/JWP DR 161018.pdf
	./go-pardocs plan -session par12sesi1 -type L ./raw/Lisan/20140327__DR_JawabLisan_clean.pdf

```

### Splitting
```bash
$ ./go-pardocs split -session par14sesi1 -type BL <file>

```
## Output
```bash
$ ls ./splitout
...                                 par14sesi1-soalan-BukanLisan-3.pdf
README.md                           par14sesi1-soalan-BukanLisan-4.pdf
par14sesi1-soalan-BukanLisan-1.pdf  par14sesi1-soalan-BukanLisan-5.pdf
par14sesi1-soalan-BukanLisan-2.pdf  par14sesi1-soalan-BukanLisan-6.pdf
...
```

### Splitting with optional date prefix
Example: Parlimen 14 Sesi 2 Mesyuarat 3; 04 Disember  2019 

Run the plan
```bash
$ ./go-pardocs plan -session 20191204-par14sesi2mesy3 -type L ./raw/Lisan/JDR04122019.pdf
```
Split with the date prefix in session parameter
```bash
$ ./go-pardocs split -session 20191204-par14sesi2mesy3 -type L ./raw/Lisan/JDR04122019.pdf
```
## Output
```bash
$ ls ./splitout
20191204-par14sesi2mesy3-soalan-Lisan-1.pdf
20191204-par14sesi2mesy3-soalan-Lisan-10.pdf
20191204-par14sesi2mesy3-soalan-Lisan-11.pdf
20191204-par14sesi2mesy3-soalan-Lisan-12.pdf
20191204-par14sesi2mesy3-soalan-Lisan-13.pdf
...
20191204-par14sesi2mesy3-soalan-Lisan-6.pdf
20191204-par14sesi2mesy3-soalan-Lisan-7.pdf
20191204-par14sesi2mesy3-soalan-Lisan-8.pdf
20191204-par14sesi2mesy3-soalan-Lisan-9.pdf

```

## IMPORTANT!

The API for split seems to be broken for certain malformed PDF [ most older ones from Parliamant ;P ]. 

Issue for the above is tracked at: https://github.com/hhrutter/pdfcpu/issues/87

The program will try to use the API but if it fails; the fall-back is to using the ```pdfcpu``` command directly.  
```bash
2019/06/16 21:01:09 Unexpected error split via API:  dict=pagesDict entry=Tabs: unsupported in version 1.4
This file could be PDF/A compliant but pdfcpu only supports versions <= PDF V1.7

2019/06/16 21:01:09 Falling back to split using pdfcpu CLI ..

```

pdfcpu is assumed to be in the default bin folder of Golang installation ```~/go/bin/pdfcpu``` AND  ```pdfcpu version 0.1.23 or above```

## HOWTO Split file to smaller pieces for analysis / development

Assumes: [pdfcpu](https://github.com/hhrutter/pdfcpu) has been downloaded 

EXAMPLE: Split to 15 pages chunk
```bash
$ ~/go/bin/pdfcpu split  \ 
        ~/Downloads/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf \
        raw/BukanLisan/split 15
```