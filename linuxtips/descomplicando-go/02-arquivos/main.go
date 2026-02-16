package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Leitura simples: lê todo o conteúdo do arquivo de uma vez
func lerArquivoSimples(nome string) {
	start := time.Now()
	dados, err := os.ReadFile(nome)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo (simples):", err)
		return
	}
	duration := time.Since(start)
	fmt.Println("Conteúdo do arquivo (simples):")
	fmt.Println(string(dados))
	fmt.Printf("Tempo de leitura simples: %v\n", duration)
}

// Forma mais simples de criar e escrever em um arquivo
func escreverArquivoSimples() {
	start := time.Now()
	arquivo, err := os.Create("exemplo_simples.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo (simples):", err)
		return
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString("Olá, este é um exemplo de escrita simples em um arquivo em Go!")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo (simples):", err)
		return
	}

	duration := time.Since(start)
	fmt.Println("Arquivo (simples) criado e escrito com sucesso!")
	fmt.Printf("Tempo de escrita simples: %v\n", duration)
}

// Leitura avançada: lê o arquivo linha a linha usando buffer
func lerArquivoBuffer(nome string) {
	start := time.Now()
	arquivo, err := os.Open(nome)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo (buffer):", err)
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	fmt.Println("Conteúdo do arquivo (buffer):")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo (buffer):", err)
	}
	duration := time.Since(start)
	fmt.Printf("Tempo de leitura buffer: %v\n", duration)
}

// Forma mais performática usando buffer
func escreverArquivoBuffer() {
	start := time.Now()
	arquivo, err := os.Create("exemplo_buffer.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo (buffer):", err)
		return
	}
	defer arquivo.Close()

	writer := bufio.NewWriter(arquivo)
	_, err = writer.WriteString("Olá, este é um exemplo de escrita performática em um arquivo em Go!")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo (buffer):", err)
		return
	}
	writer.Flush()

	duration := time.Since(start)
	fmt.Println("Arquivo (buffer) criado e escrito com sucesso!")
	fmt.Printf("Tempo de escrita buffer: %v\n", duration)
}

func main() {
	// escreverArquivoSimples()
	// escreverArquivoBuffer()

	// lerArquivoSimples("exemplo_simples.txt")
	// lerArquivoBuffer("exemplo_buffer.txt")

	// lerArquivoSimples("TA_PRECO_MEDICAMENTO.csv")
	lerArquivoBuffer("TA_PRECO_MEDICAMENTO.csv")
}
