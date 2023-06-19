#brew install coreutils

.PHONY: jsmd5sum
jsmd5sum:
	md5sum redoc.standalone.js > redoc.standalone.js.md5sum

.PHONY: jscheck
jscheck:
	wget -O redoc.standalone.js https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js
	md5sum -c redoc.standalone.js.md5sum