SRCPATH		:= $(shell pwd)
PKG_DIR		= $(SRCPATH)/releases

package: 
	rm -rf $(PKG_DIR)
	mkdir -p $(PKG_DIR)
	python misc/release.py --outdir $(PKG_DIR)

# doesn't verify that tag and .version match
fakepackage: 
	rm -rf $(PKG_DIR)
	mkdir -p $(PKG_DIR)
	python misc/release.py --outdir $(PKG_DIR) --fake-release


.PHONY: package fakepackage