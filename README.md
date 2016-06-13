Este repositório contém os códigos desenvolvidos para o meu trabalho de conclusão de curso.

[![Build Status](https://travis-ci.org/tmadeira/tcc.svg?branch=master)](https://travis-ci.org/tmadeira/tcc)
[![Coverage Status](https://coveralls.io/repos/github/tmadeira/tcc/badge.svg?branch=master)](https://coveralls.io/github/tmadeira/tcc?branch=master)

## Informações ##

**Tema:** geração uniforme de *k-trees* para aprendizado de redes bayesianas  
**Supervisor:** [Prof. Denis Deratani Mauá](http://www.ime.usp.br/~ddm/)  
**Mais:** [site do projeto](https://www.linux.ime.usp.br/~tmadeira/mac0499/)

## Instalação e uso ##

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

## Referências ##

**Caminiti et. al** citado nos comentários do código se refere o artigo **Bijective Linear Time Coding and Decoding for k-Trees** escrito por Saverito Caminiti, Emanuele G. Fusco e Rossella Petreschi (2008).
