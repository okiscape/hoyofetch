VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ" || echo "unknown")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

build:
	go build -ldflags "\
		-X okiscape/hoyofetch/utils.Version=$(VERSION) \
		-X okiscape/hoyofetch/utils.BuildTime=$(BUILD_TIME) \
		-X okiscape/hoyofetch/utils.GitCommit=$(GIT_COMMIT)" \
		-o hoyofetch

version:
	@echo "Version:    $(VERSION)"
	@echo "Build time: $(BUILD_TIME)"
	@echo "Commit:     $(GIT_COMMIT)"
