PHONY: generate
generate:
		IF NOT EXIST pkg\\yahooScraper_v1 mkdir pkg\\yahooScraper_v1
		protoc --go_out=pkg/yahooScraper_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/yahooScraper_v1 --go-grpc_opt=paths=source_relative \
		api/yahooScraper_v1/yahooStocks.proto
		move pkg\\yahooScraper_v1\\api\\yahooScraper_v1\\* pkg\\yahooScraper_v1\\
		rmdir /s /q pkg\\yahooScraper_v1\\api



