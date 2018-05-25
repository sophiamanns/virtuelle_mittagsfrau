data/fotothek.jsonl: data/fotothek.xml
	go run fotothek.go < $< > $@

data/fotothek.xml:
	@echo Use metha-sync and metha-cat on DF endpoint.

images:
	python make.py

clean:
	rm -rf cache/images

