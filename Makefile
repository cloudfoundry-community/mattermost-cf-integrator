.PHONY: all dist build package test clean

FINAL_TAG="4.0.2"

PORT="8065"
VCAP_SERVICES='{"cleardb":[{"credentials":{"uri":"mysql://titi:toto@my.db.com:3306/mydbname?reconnect=true"},"label":"cleardb","name":"dbmattermost","plan":"spark","tags":["mysql"]}]}'
VCAP_APPLICATION='{"application_id":"05b41155-9773-48be-8c6c","application_name":"mattermost","application_uris":["mattermost-ah.test.io"],"application_version":"44478127-f1d3-4d37-995b","limits":{"disk":1024,"fds":16384,"mem":1024},"name":"mattermost-ah","space_id":"96ecb509-9063-41b7-ac36-147e5f145549","space_name":"development","uris":["mattermost.app.io"],"users":null,"version":"44478127-f1d3-4d37-995b" }'

GOPATH ?= $(GOPATH:)
GOFLAGS ?= $(GOFLAGS:)

GO=go

DIST_ROOT=dist

all: dist

dist: | build test package

build:
	@echo Building Mattermost integrator

	rm -Rf $(DIST_ROOT)
	$(GO) clean $(GOFLAGS) -i ./...
	$(GO) build $(GOFLAGS)
	$(GO) install $(GOFLAGS)

package:
	@ echo Packaging Mattermost integrator

	mkdir -p $(DIST_ROOT)
	
	curl https://releases.mattermost.com/$(FINAL_TAG)/mattermost-enterprise-$(FINAL_TAG)-linux-amd64.tar.gz -o $(DIST_ROOT)/mattermost-enterprise.tar.gz
	curl https://releases.mattermost.com/$(FINAL_TAG)/mattermost-team-$(FINAL_TAG)-linux-amd64.tar.gz -o $(DIST_ROOT)/mattermost-team.tar.gz

	tar -xvzf $(DIST_ROOT)/mattermost-enterprise.tar.gz -C $(DIST_ROOT)
	cp $(GOPATH)/bin/mattermost-cf-integrator $(DIST_ROOT)/mattermost
	echo "web: ./mattermost-cf-integrator" > $(DIST_ROOT)/mattermost/Procfile
	pushd $(DIST_ROOT) && zip -r mattermost-cf-enterprise mattermost/* && popd
	rm -Rf $(DIST_ROOT)/mattermost

	tar -xvzf $(DIST_ROOT)/mattermost-team.tar.gz -C $(DIST_ROOT)
	cp $(GOPATH)/bin/mattermost-cf-integrator $(DIST_ROOT)/mattermost
	echo "web: ./mattermost-cf-integrator" > $(DIST_ROOT)/mattermost/Procfile
	pushd $(DIST_ROOT) && zip -r mattermost-cf-team mattermost/* && popd
	rm -Rf $(DIST_ROOT)/mattermost

test:
	PORT=$(PORT) VCAP_SERVICES=$(VCAP_SERVICES) VCAP_APPLICATION=$(VCAP_APPLICATION) $(GO) test $(GOFLAGS) -v ./mci

clean:
	@echo Cleaning
	rm -Rf $(DIST_ROOT)
	go clean $(GOFLAGS) -i ./...

	rm -f .prebuild