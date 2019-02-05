
SHELL_FILES := $(wildcard ./script/*.sh)
SHELL_REPORT := ./report/lint-shell.txt

.PHONY: lint-shell
lint-shell: $(SHELL_REPORT)
$(SHELL_REPORT): $(SHELL_FILES)
	docker run --rm -v "${PWD}:/mnt" koalaman/shellcheck:v0.6.0 \
    		--exclude=SC1090,SC2044 \
    		$?
	touch $(SHELL_REPORT)
