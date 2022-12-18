gen: 
	cd proto; buf mod update && buf generate

lint_proto:
	cd proto; buf lint