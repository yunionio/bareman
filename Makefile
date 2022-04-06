ROOT_DIR = $(CURDIR)
BUILD_DIR = $(ROOT_DIR)/_output
BIN_DIR = $(BUILD_DIR)/bin

prepare_dir:
	@mkdir -p $(BUILD_DIR)
	@mkdir -p $(BIN_DIR)

cmd/%: prepare_dir
	go build -o $(BIN_DIR)/$(shell basename $@) $(CURDIR)/$@

.PHONY: prepare_dir
