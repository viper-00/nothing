.PHONY: clean build-all build-collector build-agent build-alertprocessor build-client pack-all pack-collector pack-agent pack-alertprocessor pack-client

clean:
	rm -r agent/agent_linux_x86_64
	rm -r alertprocessor/alertprocessor_linux_x86_64
	rm -r client/client_linux_x86_64
	rm -r collector/collector_linux_x86_64
	rm -r agent/agent
	rm -r alertprocessor/alerts
	rm -r client/client
	rm -r collector/collector
	rm -rf release
	rm -rf collector/release/
	rm -rf agent/release/
	rm -rf alertprocessor/release/
	rm -rf client/release/

build-all: build-collector build-agent build-alertprocessor build-client

build-collector:
	cd collector && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o collector_linux_x86_64

build-agent:
	cd agent && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o agent_linux_x86_64

build-alertprocessor:
	cd alertprocessor && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o alertprocessor_linux_x86_64

build-client:
	cd client && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o client_linux_x86_64

pack-all: pack-collector pack-agent pack-alertprocessor pack-client

pack-collector: build-collector
	mkdir -p release/collector_linux_x86_64
	cp collector/collector_linux_x86_64 release/collector_linux_x86_64
	cp collector/init.sql release/collector_linux_x86_64
	cp collector/config_example.json release/collector_linux_x86_64
	cp collector/.env-example release/collector_linux_x86_64
	cp collector/alerts.json release/collector_linux_x86_64
	cd release/ && tar -cvf collector_linux_x86_64.tar.gz collector_linux_x86_64
	rm -rf release/collector_linux_x86_64

pack-agent: build-agent
 	mkdir -p release/agent_linux_x86_64
	cp agent/agent_linux_x86_64 release/agent_linux_x86_64
	cp agent/config-example.json release/agent_linux_x86_64
	cd release/ && tar -cvf agent_linux_x86_64.tar.gz agent_linux_x86_64
	rm -rf release/agent_linux_x86_64

pack-alertprocessor: build-alertprocessor
	mkdir -p release/alertprocessor_linux_x86_64
	cp alertprocessor/alertprocessor_linux_x86_64 release/alertprocessor_linux_x86_64
	cp alertprocessor/config.json release/alertprocessor_linux_x86_64
	cp alertprocessor/.env-example release/alertprocessor_linux_x86_64
	cd release/ && tar -cvf alertprocessor_linux_x86_64.tar.gz alertprocessor_linux_x86_64
	rm -rf release/alertprocessor_linux_x86_64

pack-client: build-client
	mkdir -p release/client_linux_x86_64
	cp client/client_linux_x86_64 release/client_linux_x86_64
	cp client/config-example.json release/client_linux_x86_64
	cp -r client/frontend release/client_linux_x86_64
	cd release/ && tar -cvf client_linux_x86_64.tar.gz client_linux_x86_64
	rm -rf release/client_linux_x86_64