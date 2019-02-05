
YAML_FILES := $(wildcard ./*.yml)
YAML_REPORT := ./report/lint-yaml.txt

.PHONY: lint-yaml
lint-yaml: $(GO_LINT)
$(YAML_REPORT): $(YAML_FILES)
	docker run --rm -v "${CLONE_DIR}:/workdir" giantswarm/yamllint $?
	touch $(YAML_REPORT)
