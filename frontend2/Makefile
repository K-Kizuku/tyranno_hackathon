.PHONY: init run build preview serve

init:
	bun install

run: init
	bun --bun astro dev

build:
	bun --bun astro build

preview:
	bun --bun astro preview

serve: build preview