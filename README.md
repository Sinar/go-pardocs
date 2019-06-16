# go-pardocs
Tools to process published Parliament documents (PDFs only) into more accessible form. Spiritual successor of https://github.com/leowmjw/parliamentMY-QA-blast


Assumes: 
- OSX dev environment
- Go v1.12 and above (uses go mod)
- Release for Linux, Windows, OSX available as cross-compile

## Usage

### Planning
```bash
$ ./go-pardocs plan

```

### Splitting
```bash
$ ./go-pardocs split 

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