dev_folder = $(abspath $(CURDIR)/..)

repo_folder = $(dev_folder)/danmux.github.com
pub_folder = $(dev_folder)/public

run:
	hugo server --buildDrafts --watch --verbose=true

build:
	rm -rf $(pub_folder)
	hugo -d $(pub_folder)
	cp -R $(pub_folder)/* $(repo_folder)

deploy: build

	git --git-dir=$(repo_folder)/.git --work-tree=$(repo_folder) add --all .
	git --git-dir=$(repo_folder)/.git --work-tree=$(repo_folder) commit -am"releasing"
	git --git-dir=$(repo_folder)/.git --work-tree=$(repo_folder) push
