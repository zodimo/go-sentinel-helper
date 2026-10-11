.PHONY: help tag-patch tag-minor tag-major test bench bench-allocs bench-record bench-compare

# Benchmark settings. BENCHTIME/BENCHCOUNT can be overridden on the command
# line, e.g. `make bench BENCHCOUNT=5`.
BENCHTIME  ?= 1s
BENCHCOUNT ?= 3
BENCH      ?= .
BENCHDIR   ?= bench
BENCH_FILE ?= $(BENCHDIR)/baseline.txt

help:
	@echo "Available targets:"
	@echo "  test           - Run tests"
	@echo "  bench          - Run all benchmarks (count=$(BENCHCOUNT), benchtime=$(BENCHTIME))"
	@echo "  bench-allocs   - Run benchmarks with a single count, allocs only (quick look)"
	@echo "  bench-record   - Run benchmarks and save results to $(BENCH_FILE)"
	@echo "  bench-compare  - Compare current run against $(BENCH_FILE) via benchstat"
	@echo "  tag-patch      - Increment patch version (v0.0.X -> v0.0.X+1)"
	@echo "  tag-minor      - Increment minor version (v0.X.0 -> v0.X+1.0)"
	@echo "  tag-major      - Increment major version (vX.0.0 -> vX+1.0.0)"


# Get the latest tag (globally) or default to v0.0.0
LATEST_TAG := $(shell git tag -l | sort -V | tail -n 1)
CURRENT_TAG := $(if $(LATEST_TAG),$(LATEST_TAG),v0.0.0)

# Parse current version
MAJOR := $(shell echo $(CURRENT_TAG) | awk -F. '{print $$1}' | sed 's/v//')
MINOR := $(shell echo $(CURRENT_TAG) | awk -F. '{print $$2}')
PATCH := $(shell echo $(CURRENT_TAG) | awk -F. '{print $$3}')

test:
	go test ./...

# bench runs every benchmark across all packages with -benchmem so each result
# carries ns/op, B/op and allocs/op. BENCHCOUNT > 1 gives benchstat something to
# compute variance from.
bench:
	go test ./... -run '^$$' -bench '$(BENCH)' -benchmem -benchtime=$(BENCHTIME) -count=$(BENCHCOUNT)

# bench-allocs is the quick single-pass version for a fast "did I add an alloc?"
# check while iterating.
bench-allocs:
	go test ./... -run '^$$' -bench '$(BENCH)' -benchmem -benchtime=$(BENCHTIME)

# bench-record runs the suite and writes the raw output to $(BENCH_FILE), which
# is what bench-compare and benchstat consume. Committed as the baseline.
bench-record:
	@mkdir -p $(BENCHDIR)
	go test ./... -run '^$$' -bench '$(BENCH)' -benchmem -benchtime=$(BENCHTIME) -count=$(BENCHCOUNT) | tee $(BENCH_FILE)
	@echo "Wrote $(BENCH_FILE)"

# bench-compare diffs a fresh run against the recorded baseline using benchstat.
# Requires: go install golang.org/x/perf/cmd/benchstat@latest
bench-compare:
	@command -v benchstat >/dev/null || { echo "benchstat not found: go install golang.org/x/perf/cmd/benchstat@latest"; exit 1; }
	@mkdir -p $(BENCHDIR)
	go test ./... -run '^$$' -bench '$(BENCH)' -benchmem -benchtime=$(BENCHTIME) -count=$(BENCHCOUNT) > $(BENCHDIR)/current.txt
	benchstat $(BENCH_FILE) $(BENCHDIR)/current.txt


tag-patch:
	@echo "Current version: $(CURRENT_TAG)"
	@new_patch=$$(($(PATCH) + 1)); \
	NEW_TAG="v$(MAJOR).$(MINOR).$$new_patch"; \
	echo "New version: $$NEW_TAG"; \
	git tag $$NEW_TAG; \
	echo "Created tag $$NEW_TAG"

tag-minor:
	@echo "Current version: $(CURRENT_TAG)"
	@new_minor=$$(($(MINOR) + 1)); \
	NEW_TAG="v$(MAJOR).$$new_minor.0"; \
	echo "New version: $$NEW_TAG"; \
	git tag $$NEW_TAG; \
	echo "Created tag $$NEW_TAG"

tag-major:
	@echo "Current version: $(CURRENT_TAG)"
	@new_major=$$(($(MAJOR) + 1)); \
	NEW_TAG="v$$new_major.0.0"; \
	echo "New version: $$NEW_TAG"; \
	git tag $$NEW_TAG; \
	echo "Created tag $$NEW_TAG"
