# go-pardocs
Tools to process published Parliament documents (PDFs only) into more accessible form. Spiritual successor of https://github.com/leowmjw/parliamentMY-QA-blast


Assumes: 
- OSX dev environment
- Go v1.12 and above (uses go mod)
- Release for Linux, Windows, OSX available as cross-compile

## IMPORTANT!

Issue for the below is tracked at: https://github.com/hhrutter/pdfcpu/issues/80

It does seem that the new Written Questions in [Parlimen Malaysia](http://www.parlimen.gov.my/files/jindex/pdf/Pertanyaan%20Jawapan%20Bukan%20Lisan%2022019.pdf) 
is not purely PDF 1.4 compliant.  You may need to do the following workaround as the pdfcpu library we use does a validation 
for the expressed PDF version.

Check validity of Source PDF first!
```bash
$ hexdump -c ~/Desktop/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf  | head
  0000000   %   P   D   F   -   1   .   4  \n   %   �   �   �   �  \n   2
  0000010       0       o   b   j  \n   <   <   /   F   i   l   t   e   r
  0000020   /   F   l   a   t   e   D   e   c   o   d   e   /   L   e   n
  0000030   g   t   h       2   4   2   >   >   s   t   r   e   a   m  \n  

# Original file fails validation
$ ~/go/bin/pdfcpu validate  ~/Desktop/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf
  validating(mode=relaxed) /Users/mleow/Desktop/Pertanyaan Jawapan Bukan Lisan 22019.pdf ...
  validation error: dict=pagesDict entry=Tabs: unsupported in version 1.4
  This file could be PDF/A compliant but pdfcpu only supports versions <= PDF V1.7
```

After Hexedit (change 1.4 -> 1.7):
```bash
$ hexdump -c ~/Downloads/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf| head
  0000000   %   P   D   F   -   1   .   7  \n   %   �   �   �   �  \n   2
  0000010       0       o   b   j  \n   <   <   /   F   i   l   t   e   r
  0000020   /   F   l   a   t   e   D   e   c   o   d   e   /   L   e   n
  0000030   g   t   h       2   4   2   >   >   s   t   r   e   a   m  \n

# Now passes relaxed validation
$ ~/go/bin/pdfcpu validate  ~/Downloads/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf
  validating(mode=relaxed) /Users/mleow/Downloads/Pertanyaan Jawapan Bukan Lisan 22019.pdf ...
  validation ok
```
Possible tools in OSX is HexFiend / 0xED

## HOWTO Split file to smaller pieces for analysis / development

Assumes: [pdfcpu](https://github.com/hhrutter/pdfcpu) has been downloaded 

EXAMPLE: Split to 15 pages chunk
```bash
$ ~/go/bin/pdfcpu split  \ 
        ~/Downloads/Pertanyaan\ Jawapan\ Bukan\ Lisan\ 22019.pdf \
        raw/BukanLisan/split 15
```