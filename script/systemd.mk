
SYSTEMD_FILES := $(wildcard ./sample/systemd/*/*.service)
SYSTEMD_REPORT := ./report/lint-systemd.txt

.PHONY: lint-systemd
lint-systemd: $(SYSTEMD_REPORT)
$(SYSTEMD_REPORT): $(SYSTEMD_FILES)
	systemd-analyze verify --user $?
	touch $(SYSTEMD_REPORT)
