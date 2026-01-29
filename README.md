# training-room

Go 语言练习项目。每个练习独立可运行，便于按主题学习和复现。

## 目录结构设计

```
training-room/
├── go.mod                 # 根模块，所有练习共用
├── README.md
├── Makefile               # 可选：列出/运行练习的快捷命令
│
├── core-semantics/        # 按主题/类别划分
│   └── make-new/          # 每个练习一个子目录
│       └── main.go        # 必须是 package main，可独立运行
│
├── concurrency/           # 示例：并发相关练习
│   └── channel-basics/
│       └── main.go
│
└── ...                    # 其他主题：stdlib、testing、errors 等
```

### 设计原则

1. **一个练习 = 一个目录 + 一个 `main.go`**  
   每个练习目录内只有 `package main` 和 `func main()`，用 `go run ./路径` 即可运行，无需切目录。

2. **按主题分目录**  
   如 `core-semantics`、`concurrency`、`stdlib`，方便以后扩展，也便于 `go run ./core-semantics/...` 按类别跑。

3. **根目录一个 `go.mod`**  
   所有练习都在同一 module 下，依赖统一、工具链一致；需要时可在某练习目录加 `go.work` 做本地多模块实验。

## 如何运行

### 运行单个练习（推荐）

在项目根目录执行：

```bash
# 运行 core-semantics/make-new
go run ./core-semantics/make-new

# 运行其他练习（替换为实际路径）
go run ./concurrency/channel-basics
```

无需 `cd` 到练习目录，路径即练习名，便于记忆和脚本化。

### 列出所有可运行练习

```bash
go run ./cmd/list-exercises
# 或
make list
```

### 运行所有练习（可选）

若提供了 `make run-all` 或脚本，可在根目录一键跑完所有练习（适合回归或演示）；默认不要求实现，按需添加。

## 添加新练习

1. 在合适主题下新建目录，例如 `concurrency/channel-basics/`。
2. 在该目录下创建 `main.go`，包名为 `package main`，并实现 `func main()`。
3. 在根目录执行：`go run ./concurrency/channel-basics`。

无需改 `go.mod`，新练习自动纳入同一模块。

## 当前练习列表

| 路径 | 说明 |
|------|------|
| `core-semantics/make-new` | `new()` 与 slice/map/channel 的语义与 nil 行为 |
