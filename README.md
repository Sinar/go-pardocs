# go-pardocs
Tools to process published Parliament documents (PDFs only) into more accessible form. Spiritual successor of https://github.com/leowmjw/parliamentMY-QA-blast


Assumes: 
- OSX dev environment
- Go v1.12 and above (uses go mod)
- Release for Linux, Windows, OSX available as cross-compile

## IMPORTANT!

Issue for the below is tracked at: https://github.com/hhrutter/pdfcpu/issues/87

The API for split seems to be broken for certain malformed PDF [ most older ones from Parliamant ;P ]. 

The program will try to use the API but if it fails; the fall-back is to using the ```pdfcpu``` command directly.  It assumes to be in the default bin folder of Golang installation ```~/go/bin/pdfcpu``` AND  ```pdfcpu version 0.1.23 or above```

## HOWTO Split file to smaller pieces for analysis / development

Assumes: [pdfcpu](https://github.com/hhrutter/pdfcpu) has been downloaded 

EXAMPLE: Split to 15 pages chunk
```bash
$ ~/go/bin/pdfcpu split  \ 
        ~/Downloads/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf \
        raw/BukanLisan/split 15
```