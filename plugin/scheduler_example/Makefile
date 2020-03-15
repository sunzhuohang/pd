moveRegionPlugin.so: *.go
	GO111MODULE=on GO_ENABLED=1 go build  -buildmode=plugin -o moveRegionPlugin.so *.go

.PHONY : clean

clean:
	rm moveRegionPlugin.so
