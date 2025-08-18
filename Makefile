# 设置服务名称（每次可生成一个微服务模块）
APP = oae

# 配置
OUTPUT_DIR := ./app/$(APP)/cmd/api/desc
WORK_DIR := ./app/$(APP)/cmd/api
API_FILE := $(OUTPUT_DIR)/$(APP).api
TEMPLATE_DIR := ./common/template

# 初始化 && 脚手架代码
init:
	@echo "Generating Service from $(APP)"
	mkdir -p $(OUTPUT_DIR) && cd $(OUTPUT_DIR)
	if [ ! -f $(API_FILE) ]; then \
		goctl api -o $(API_FILE) --home=$(TEMPLATE_DIR); \
	else \
		echo "$(API_FILE) already exists, skipping generation"; \
	fi
	goctl api go -api $(API_FILE) -dir $(WORK_DIR) --style=goZero --home=$(TEMPLATE_DIR)
