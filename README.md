Este repositório contém os códigos desenvolvidos para o meu trabalho de conclusão de curso.

**Tema:** geração uniforme de *k-trees* para aprendizado de redes bayesianas  
**Supervisor:** [Prof. Denis Deratani Mauá](http://www.ime.usp.br/~ddm/)  
**Mais informações:** [site do projeto](https://www.linux.ime.usp.br/~tmadeira/mac0499/)

[![GoDoc](https://godoc.org/github.com/tmadeira/tcc?status.svg)](https://godoc.org/github.com/tmadeira/tcc)
[![Build Status](https://travis-ci.org/tmadeira/tcc.svg?branch=master)](https://travis-ci.org/tmadeira/tcc)
[![Coverage Status](https://coveralls.io/repos/github/tmadeira/tcc/badge.svg?branch=master)](https://coveralls.io/github/tmadeira/tcc?branch=master)

## Instalação e testes ##

As implementações foram realizadas em [Go](https://golang.org/) e é necessário instalá-lo. Recomenda-se utilizar o próprio `go` para clonar este repositório.

Para baixar o código e rodar os testes, use:

```sh
$ export ${GOPATH:=$HOME/go}
$ mkdir -p $GOPATH
$ cd $GOPATH
$ go get github.com/tmadeira/tcc/...
$ go test -v github.com/tmadeira/tcc/...
```

As três primeiras linhas criarão o diretório `$HOME/go` e o usarão como `$GOPATH` caso essa variável de ambiente não esteja configurada na sua estação de trabalho.

Caso você não esteja acostumado com esse *workflow* (que é o padrão da linguagem Go), recomenda-se a leitura de [How to Write Go Code](https://golang.org/doc/code.html).

## Exemplos ##

Após baixar o código no seu `$GOPATH`, você pode usar alguns utilitários disponíveis na pasta `examples/`. Eles permitem codificar/decodificar *k-trees* e gerar *k-trees* aleatórias.

### code-ktree ###

Para codificar uma *k-tree*, use:

```sh
$ go install github.com/tmadeira/tcc/examples/code-ktree
$ $GOPATH/bin/code-ktree
```

A entrada desse utilitário deve ser dada no formato:

```n
k
m
x_1 y_1
...
x_m y_m
```

Onde:

- `n` é o número de vértices
- `k` é o parâmetro *k* da *k-tree*
- `m` é o número de arestas
- `x_i y_i` corresponde à *i-ésima* aresta (0 <= `x_i`, `y_i` < n)

Um exemplo de entrada pode ser encontrado em `examples/code-ktree/fig1a.txt`.

A saída será um par *(Q, S)* no formato de entrada esperado pelo utilitário **decode-ktree**.

### decode-ktree ###

Para decodificar um código *(Q, S)* em uma *k-tree*, use:

```sh
$ go install github.com/tmadeira/tcc/examples/decode-ktree
$ $GOPATH/bin/decode-ktree
```

A entrada desse utilitário deve ser dada no formato:

```k
Q_1
...
Q_k
s
p_1 l_1
...
p_s l_s
```

Onde:

- `k` é o tamanho de *Q*
- `Q_i` corresponde ao *i-ésimo* valor em *Q*
- `s` é o tamanho do Generalized Dandelion Code, *|S|*
- `p_i l_i` corresponde ao *i-ésimo* valor em *S*

Um exemplo de entrada pode ser encontrado em `examples/decode-ktree/code1a.txt`.

A saída será uma *k-tree* no formato de entrada esperado pelo utilitário **code-ktree**.

### generate-ktree ###

Para gerar uma *k-tree* aleatória, use:

```sh
$ go install github.com/tmadeira/tcc/examples/generate-ktree
$ $GOPATH/bin/generate-ktree
```

A entrada desse utilitário deve ser dada no formato:

```
n k
```

A saída será uma *k-tree* com *n* vértices no formato de entrada esperado pelo utilitário **code-ktree**.

## Referências ##

**Caminiti et al.** citado nos comentários do código se refere o artigo **Bijective Linear Time Coding and Decoding for k-Trees** escrito por Saverito Caminiti, Emanuele G. Fusco e Rossella Petreschi (2008).
