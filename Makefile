# training-room 练习项目
# 运行单个练习：make run EXERCISE=./core-semantics/make-new
# 列出所有练习：make list
# 运行所有练习：make run-all

.PHONY: list run run-all

list:
	@go run ./cmd/list-exercises

run:
ifndef EXERCISE
	@echo "用法: make run EXERCISE=./core-semantics/make-new"
	@make list
else
	@go run $(EXERCISE)
endif

# 收集所有含 main.go 的目录（排除 cmd）并依次 go run
run-all:
	@for dir in $$(find . -name main.go -not -path "./cmd/*" | xargs -I{} dirname {}); do \
		echo ">>> $$dir"; go run ./$$dir || true; echo; \
	done
