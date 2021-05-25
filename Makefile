test:
	go test .

.PHONY: release
## release: release nke (tagging and exec goreleaser)
release:
	git semv patch --bump
	goreleaser --rm-dist

.PHONY: releasedeps
releasedeps: git-semv goreleaser

.PHONY: git-semv
git-semv:
	brew tap linyows/git-semv
	brew install git-semv

.PHONY: goreleaser
goreleaser:
	brew install goreleaser/tap/goreleaser
	brew install goreleaser
