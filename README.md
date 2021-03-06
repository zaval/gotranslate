# goTranslate

Translate selected text.

* Support english and russian languages
* Automatic detect language by cyrillic characters

## Build

### Linux

* Building apps requires the `gtk3`, `libappindicator3` and `libwebkit2gtk-4.0-dev` development headers to be installed. For Debian or Ubuntu, you can may install these using:

`sudo apt-get install libgtk-3-dev libappindicator3-dev libwebkit2gtk-4.0-dev`

* Running app requires `xsel` and `notify-send` programs to be installed. For Debian or Ubuntu you can install these using:

`sudo apt-get install xsel libnotify-bin`

### Windows

To avoid opening a console at application startup, use these compile flags:

`go build -ldflags -H=windowsgui`

You'll also need to include a Manifest for your windows app. Create a file `goTranslate.manifest` like this:

```<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
   <assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
       <assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="App Name Here" type="win32"/>
       <dependency>
           <dependentAssembly>
               <assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
           </dependentAssembly>
       </dependency>
       <application xmlns="urn:schemas-microsoft-com:asm.v3">
           <windowsSettings>
               <dpiAwareness xmlns="http://schemas.microsoft.com/SMI/2016/WindowsSettings">PerMonitorV2, PerMonitor</dpiAwareness>
               <dpiAware xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">True</dpiAware>
           </windowsSettings>
       </application>
   </assembly>
```

Then either compile the manifest using the [rsrc](https://github.com/akavel/rsrc) tool, like this:

```
go get github.com/akavel/rsrc
rsrc -manifest goTranslate.manifest -o rsrc.syso
```

## Cross-compile for windws host using docker

* Create and switch to windows docker context using [windows-docker-machine](https://github.com/StefanScherer/windows-docker-machine)
* Build image
```
cd docker
docker build -t gotranslate-build .
mkdir build
docker run --rm -v C:$(pwd)/build:C:/build -it gotranslate-build:latest
``` 
* Compiled windows binary will appear in the `build` directory
* PROFIT

## Thanks

* [systray](https://github.com/getlantern/systray) a cross-platform Go library to place an icon and menu in the notification area.
* [gtranslate](https://github.com/bregydoc/gtranslate) Google Translate API for unlimited and free translations
* [go-toast](https://github.com/go-toast/toast) A go package for Windows 10 toast notifications.
