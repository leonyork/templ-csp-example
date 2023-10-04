
init:
	go get
	go install github.com/a-h/templ/cmd/templ@v0.2.364
.PHONY: init

run: init
	templ generate
	go run main.go
.PHONY: run

test: run-lighthouse run-csp-evaluator
.PHONY: test

# Runs lighthouse to make the sure that page loads without errors
run-lighthouse:
	@docker build -t lighthouse test/lighthouse
	@docker run lighthouse || (echo "❌ Lighthouse spotted errors logged to the console or inspector issues" && exit 1)
.PHONY: run-lighthouse

# Curls the app and checks for a CSP header and "unsafe" in the CSP header
run-csp-evaluator: 
	@curl -is localhost:3000 | grep -q "Content-Security-Policy" || (echo "❌ Content-Security-Policy header not found" && exit 1)
	@curl -is localhost:3000 | grep "Content-Security-Policy" | grep -qv "unsafe" || (echo "❌ found 'unsafe' in Content-Security-Policy header" && exit 1)
	@echo ✅ CSP contains no unsafe policies
.PHONY: run-csp-evaluator