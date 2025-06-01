# AI Image Recognizer

Simple AI tool to identify the main subjects appearing in a series of locally stored images using llama.cpp.

Have two different working modes:

- COUNTER: Count an specific type of subjects in the image
- RECOGNIZER: Recognize the specific subtype of subject in the image

# Getting Started

1. Install [Golang](https://go.dev/doc/install)
1. Install [llama.cpp](https://github.com/ggml-org/llama.cpp)
1. Put your JPEG images in the `assets/counter` and `assets/recognizer` folders
1. Execute `make llm` and leave the llama.cpp server running in the open terminal

```bash
# Execute COUNTER mode (replace MAIN_CATEGORY with the type of subjects in the images)

make counter category=MAIN_CATEGORY

# Example: make counter category=nuts
```

EXECUTE RECOGNIZER MODE:

```bash
# Execute RECOGNIZER mode (replace MAIN_CATEGORY with the type of subjects in the images)

make recognizer category=MAIN_CATEGORY

# Example: make recognizer category=fruit
```
