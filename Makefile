run:
	@air

templ:
	@templ generate --watch --proxy="http://localhost:3000" --cmd="go run ."