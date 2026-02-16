package main

import (
	"bufio"
	"fmt"
	"os"
)

// Leitura simples: lê todo o conteúdo do arquivo de uma vez
func lerArquivoSimples(nome string) {
	dados, err := os.ReadFile(nome)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo (simples):", err)
		return
	}
	fmt.Println("Conteúdo do arquivo (simples):")
	fmt.Println(string(dados))
}

// Forma mais simples de criar e escrever em um arquivo
func escreverArquivoSimples() {
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

	fmt.Println("Arquivo (simples) criado e escrito com sucesso!")
}

// Leitura avançada: lê o arquivo linha a linha usando buffer
func lerArquivoBuffer(nome string) {
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
}

// Forma mais performática usando buffer
func escreverArquivoBuffer() {
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

	fmt.Println("Arquivo (buffer) criado e escrito com sucesso!")
}

func main() {
	// escreverArquivoSimples()
	// escreverArquivoBuffer()

	// lerArquivoSimples("exemplo_simples.txt")
	// lerArquivoBuffer("exemplo_buffer.txt")

	lerArquivoBuffer("TA_PRECO_MEDICAMENTO.csv")
}
