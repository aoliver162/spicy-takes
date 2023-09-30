
htmx: | vendored
	curl -o web/vendored/htmx.min.js -L https://unpkg.com/htmx.org/dist/htmx.min.js

vendored:
	mkdir -p web/$@
