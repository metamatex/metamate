SOURCE_MAKE=. ./.make/make.sh
SHELL := /bin/bash

build:
	@${SOURCE_MAKE} && build

build_metactl:
	@${SOURCE_MAKE} && build_metactl

build_metamate:
	@${SOURCE_MAKE} && build_metamate

chore:
	@${SOURCE_MAKE} && chore

release:
	@${SOURCE_MAKE} && release

generate:
	@${SOURCE_MAKE} && generate

deploy:
	@${SOURCE_MAKE} && deploy