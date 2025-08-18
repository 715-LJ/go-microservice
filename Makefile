# 设置服务名称（针对单个微服务模块）
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

deploy:
	@echo "Generating Deploy Service"
	@if [ ! -f "./deploy.sh" ]; then \
		echo "Error: deploy.sh not found"; \
		exit 1; \
	fi
	@if [ ! -x "./deploy.sh" ]; then \
		echo "Error: deploy.sh is not executable"; \
		exit 1; \
	fi
	@./deploy.sh && echo "Deploy completed successfully" || (echo "Deploy failed"; exit 1)

update:
	@echo "Generating Update Service"
	@if [ ! -f ./update.sh ]; then \
		echo "Error: update.sh not found"; \
		exit 1; \
	fi
	@if [ ! -x ./update.sh ]; then \
		echo "Error: update.sh is not executable"; \
		exit 1; \
	fi
	@./update.sh || (echo "Error: Update script failed"; exit 1)

