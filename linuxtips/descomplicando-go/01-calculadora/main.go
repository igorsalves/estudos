package main

import "fmt"

func lerNumero(mensagem string) float64 {
	var numero float64
	fmt.Print(mensagem)
	fmt.Scanln(&numero)
	return numero
}

func lerOperacao() string {
	var op string
	fmt.Print("Escolha a operação (+, -, *, /): ")
	fmt.Scanln(&op)
	return op
}

func calcular(num1, num2 float64, operacao string) (float64, error) {
	switch operacao {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Erro: divisão por zero não é permitida.")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Operação inválida.")
	}
}

func exibirResultado(resultado float64, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}
}

func main() {
	fmt.Println("Calculadora Simples em Go")
	num1 := lerNumero("Digite o primeiro número: ")
	num2 := lerNumero("Digite o segundo número: ")
	operacao := lerOperacao()
	resultado, err := calcular(num1, num2, operacao)
	exibirResultado(resultado, err)
}
