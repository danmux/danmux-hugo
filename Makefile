dev_folder = $(abspath $(CURDIR)/../..)

hugo_folder = $(dev_folder)/danmux-hugo/source
repo_folder = $(dev_folder)/danmux.github.com
pub_folder = $(dev_folder)/danmux-hugo/public

run:
	hugo server --buildDrafts --watch --verbose=true

build:
	
	hugo
	cp -R $(pub_folder)/* $(repo_folder)

deploy: build

	git --git-dir=$(repo_folder)/.git --work-tree=$(repo_folder) add --all .
	git --git-dir=$(repo_folder)/.git --work-tree=$(repo_folder) commit -am"releasing"
	git --git-dir=$(repo_folder)/.git --work-tree=$(repo_folder) push
