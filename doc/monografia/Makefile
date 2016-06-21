all: monografia clean

monografia:
	pdflatex monografia.tex
	bibtex monografia
	makeindex monografia.tex
	pdflatex monografia.tex
	pdflatex monografia.tex

clean:
	rm -rf monografia.{aux,bbl,blg,ilg,ind,log,toc}
