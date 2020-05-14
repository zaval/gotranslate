git clone https://github.com/zaval/gotranslate.git
mv goTranslate.manifest gotranslate\\goTranslate.manifest
cd gotranslate
rsrc -manifest goTranslate.manifest -o rsrc.syso
go build -ldflags -H=windowsgui
cp gotranslate.exe c:\\build\\gotranslate.exe
exit