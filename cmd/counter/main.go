package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/AntonioMartinezFernandez/ai-image-recognizer/pkg/llamacpp"
	"github.com/AntonioMartinezFernandez/ai-image-recognizer/pkg/utils"
)

func main() {
	var category string
	var folder string
	var llamaServerURL string
	var maxTokens int64

	flag.StringVar(&category, "category", "", "Category to classify images (e.g., fruit, vegetable, animal)")
	flag.StringVar(&folder, "folder", "assets/counter", "Folder containing images to process")
	flag.StringVar(&llamaServerURL, "llama-server-url", "http://localhost:8080", "URL of the Llama server")
	flag.Int64Var(&maxTokens, "max-tokens", 1000, "Maximum tokens for the Llama request")

	flag.Parse()

	if category == "" {
		fmt.Printf("No category provided. Please use the category flag to specify a category. e.g., make counter category=fruit\n")
		os.Exit(1)
	}

	jpgFiles, err := utils.JpgFileNames(folder)
	if err != nil {
		fmt.Printf("Error reading assets folder: %v\n", err)
		os.Exit(1)
	}

	llamaCppClient := llamacpp.NewClient(
		llamaServerURL,
		maxTokens,
	)

	wg := sync.WaitGroup{}

	for _, jpgFile := range jpgFiles {
		wg.Add(1)
		go processImage(folder, jpgFile, category, llamaCppClient, &wg)
	}

	wg.Wait()
}

func processImage(assetsFolder, jpgFile string, category string, llamaCppClient *llamacpp.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	filePath := filepath.Join(assetsFolder, jpgFile)
	base64Image, err := utils.EncodeJpgBase64(filePath)
	if err != nil {
		fmt.Printf("Error encoding image %s: %v\n", jpgFile, err)
		return
	}

	response, err := llamaCppClient.CountSubjects(category, *base64Image)
	if err != nil {
		fmt.Printf("Error recognizing image %s: %v\n", jpgFile, err)
		return
	}

	fmt.Printf("> Image: %s - AI Result: %s\n", jpgFile, *response)
}
