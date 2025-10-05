# Site Disponibility Monitorator

Primeira aplicação que estou criando em Go, com finalidade de estudar a linguagem.

Essa é uma aplicação de console que irá verificar a disponibilidade de alguns sites.

Para isso irei utilizar os seguintes recursos da linguagem:

- Declaração de variáveis e funções
- Controle de fluxo
- Requisições Http
- Tratamento de erros

## Start Application

Basta usar o comando `go run main.go` no seu terminal.

## Cheatsheet GoLang

### Visão geral

Criador: Google (Robert Griesemer, Rob Pike, Ken Thompson).

Paradigma: Compilado, tipagem estática.

### Instalação & comandos básicos

Instalar: baixar binário do site oficial ou usar gerenciador do SO.

Comandos principais:

`go version` — versão.

`go env` — variáveis de ambiente.

`go mod init <module>` — inicializa módulo.

`go build` — compila pacote.

`go run main.go` — executa diretamente.

`go test ./...` — roda testes.

`go fmt` — formata código.

`go vet` — checa problemas comuns.

### Estrutura de um arquivo `.go`

```
package main


import (
    "fmt"
)


func main() {
    fmt.Println("Hello, world")
}
```

`package main` define executável; outros pacotes são bibliotecas.

`import` agrupa dependências.

### Tipos básicos

Primitivos: `bool`, `string`.

Números: `int`, `float32` e `float64`.

### Declaração de variáveis

```
var x int = 10
var y = 20        // inferred type
z := 30           // short declaration, only inside functions
```

Zero values: `0`, `false` e `""`.

### Constantes

```
const Pi = 3.1415
const (
  A = iota // 0
  B        // 1
)
```

`const` é a palavra reservada que define uma constante.

O bloco de constantes acima é a criação de um enumerado. `iota` define o valor de A para 0, B segue a sequência, tendo valor 1, assim subsequêntemente.

### Controle de fluxo

`if`, `else` — sem parênteses em condições.

```
if condition {}
if else condition {}
else {}
```

`switch` — cases não precisam de `break` (break implícito). Pode usar `fallthrough`.

```
switch valor {
    case 1:
        fmt.Println("Do something...")
    case 2:
        fmt.Println("Do something 2...")
    default:
        fmt.Println("Default")
}
```

`for` — único laço; pode funcionar como `while`:

```
for i := 0; i < 10; i++ { }
for ; condition; { }
for range sliceOrMap { }
```

### Funções

```
func soma(a int, b int) int {
  return a + b
}


func swap(a, b string) (string, string) {
  return b, a
}
```

Suporta múltiplos retornos.
