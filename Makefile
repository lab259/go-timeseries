GOPATHV=$(CURDIR)/.gopath
GOPATHCMD=GOPATH=$(GOPATHV)

COVERDIR=$(CURDIR)/.cover
COVERAGEFILE=$(COVERDIR)/cover.out

.PHONY: deps deps-ci coverage coverage-ci test test-watch coverage coverage-html

test:
	@${GOPATHCMD} ginkgo --failFast ./...

test-watch:
	@${GOPATHCMD} ginkgo watch -cover -r ./...

coverage-ci:
	@mkdir -p $(COVERDIR)
	@ginkgo -r -covermode=count --cover --trace ./
	@echo "mode: count" > "${COVERAGEFILE}"
	@find . -type f -name *.coverprofile -exec grep -h -v "^mode:" {} >> "${COVERAGEFILE}" \; -exec rm -f {} \;

coverage:
	@mkdir -p $(COVERDIR)
	@${GOPATHCMD} ginkgo -r -covermode=count --cover --trace ./
	@echo "mode: count" > "${COVERAGEFILE}"
	@find . -type f -name *.coverprofile -exec grep -h -v "^mode:" {} >> "${COVERAGEFILE}" \; -exec rm -f {} \;
	@sed -i -e "s|_$(CURDIR)/|./|g" "${COVERAGEFILE}"

coverage-html:
	@$(GOPATHCMD) go tool cover -html="${COVERAGEFILE}" -o .cover/report.html

deps:
	@mkdir -p ${GOPATHV}
	@$(GOPATHCMD) go get -v -t ./...

deps-ci:
	@go get -v -t ./...

vet:
	@$(GOPATHCMD) go vet ./...

fmt:
	@$(GOPATHCMD) go vet ./...