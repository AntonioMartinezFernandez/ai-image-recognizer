.PHONY: llm
llm:
	llama-server -hf ggml-org/SmolVLM-500M-Instruct-GGUF

.PHONY: build-recognizer
build-recognizer:
	go build -o ./bin/recognizer ./cmd/recognizer

.PHONY: build-counter
build-counter:
	go build -o ./bin/counter ./cmd/counter

.PHONY: recognizer
recognizer: build-recognizer
	./bin/recognizer -category=$(category)

.PHONY: counter
counter: build-counter
	./bin/counter -category=$(category)
