SHELL = /bin/bash

# Download selection of higher-res fotos from Deutsche Fotothek.
.PHONY: download-images
download-images: data/fotothek.jsonl
	go run dfdl.go < $<

data/fotothek.jsonl: data/fotothek.xml
	go run fotothek.go < $< > $@

data/fotothek.xml:
	@echo Use metha-sync and metha-cat on DF endpoint.


# Create thumbnail images.
images:
	python make.py

clean:
	rm -rf cache/images
