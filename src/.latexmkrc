#!/usr/bin/env perl
$pdflatex = 'pdflatex -synctex=1 -interaction=nonstopmode';
$pdf_mode = 1;
$do_cd = 1;
$out_dir = build;
$clean_ext = 'snm nav synctex.gz';
@default_files = (
  'folien.tex',
  '00_organisatorisches.tex',
  '01_themenueberblick.tex',
  'rekursion.tex',
);
