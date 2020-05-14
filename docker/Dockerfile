FROM golang:windowsservercore-1809
RUN mkdir c:\\app ; cd c:\\app
RUN go get github.com/bregydoc/gtranslate ; \
    go get github.com/getlantern/systray ; \
    go get github.com/go-toast/toast ; \
    go get github.com/akavel/rsrc
COPY entrypoint.ps1 c:\\app\\entrypoint.ps1
COPY goTranslate.manifest c:\\app\\goTranslate.manifest
WORKDIR c:\\app
ENTRYPOINT ["powershell", "c:\\app\\entrypoint.ps1"]
