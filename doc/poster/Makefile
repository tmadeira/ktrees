all: poster clean

poster:
	pdflatex poster.tex
	bibtex poster
	pdflatex poster.tex
	pdflatex poster.tex

clean:
	rm -rf poster.{aux,bbl,blg,ilg,ind,log,toc,nav,snm,out}
