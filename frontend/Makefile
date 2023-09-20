dev: deps
	bun run dev

build: deps test
	bunx --bun vite build

preview: build
	bun run preview

deps:
	bun install

test: lint prettier typecheck vitest

lint:
	bun run lint

prettier:
	bunx prettier --check src

prettier/write:
	bunx prettier --write src

vitest: deps
	bunx vitest run src

vitest/watch: deps
	bunx vitest src

typecheck:
	bunx tsc --noEmit

typecheck/watch:
	bunx tsc --noEmit --watch

gen:
	bunx buf generate buf.build/connectrpc/eliza --template ../buf.gen.ts.yaml  
