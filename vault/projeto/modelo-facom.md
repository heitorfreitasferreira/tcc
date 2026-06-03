---
tags:
  - facom
  - monografia
  - modelo
  - abnt
---

# Modelo de Monografia FACOM/UFU

Fonte oficial: `modelo_de_monografia_da_facom.pdf` (abnTeX2, classe `ppgco.cls`).

## Estrutura Obrigatória

| Elemento | Ordem | Descrição |
|---|---|---|
| Capa | 1 | Nome do autor, título, faculdade, cidade, ano |
| Folha de Rosto | 2 | Autor, título, natureza (TCC apresentado à FACOM...), área, orientador |
| Dedicatória | 3 | Opcional |
| Agradecimentos | 4 | Opcional |
| Epígrafe | 5 | Opcional |
| Resumo | 6 | Objetivo, método, resultados, conclusões. Palavras-chave ao final |
| Abstract | 7 | Mesmo padrão, em inglês. Keywords ao final |
| Lista de Ilustrações | 8 | Figuras, gráficos |
| Lista de Tabelas | 9 | Tabelas |
| Lista de Siglas | 10 | Siglas usadas (ex.: ABNT, TCP) |
| Sumário | 11 | Seções numeradas |

## Capítulos

### 1. Introdução
- **Motivação**: estado da arte, problemas não resolvidos, _gancho_ para objetivos
- **Problema**: desafios e justificativa da escolha
- **Hipótese**: suposição para solução; deve ser comprovada nos experimentos
- **Objetivos**: usar verbos como "contribuir", "analisar", "investigar", "comparar" (evitar "desenvolver um sistema")
- **Contribuições**: científicas (publicações ficam em seção própria)
- **Organização da Monografia**: descrição breve de cada capítulo

### 2. Fundamentação Teórica
- Pode ser dividido em Conceitos Básicos + Trabalhos Correlatos (capítulos separados)

### 3. Experimentos e Análise dos Resultados
- **Método de Avaliação**: métricas, parâmetros, bases de dados, _baselines_
- **Experimentos**: gráficos/tabelas com clareza
- **Avaliação dos Resultados**: acertos, limitações, justificativas, evidências da hipótese

### 4. Conclusão
- Relacionar objetivos da introdução com o que foi alcançado
- **Principais Contribuições**: hipótese validada pelos experimentos
- **Trabalhos Futuros**: melhorias do método, novos projetos
- **Contribuições Bibliográficas**: publicações resultantes

## Elementos Pós-Textuais
- **Referências** (ABNT)
- **Apêndices** (material do autor)
- **Anexos** (material de terceiros)

## Formatação Técnica
- Classe: **abnTeX2** (ABNT NBR 14724:2011)
- Compilação: `pdflatex + bibtex + makeindex + pdflatex (2x)`
- Citações >3 linhas: recuo 4cm, fonte menor, sem aspas
- Tabelas: conforme normas da universidade
- Figuras: preferir vetoriais (PDF); ferramentas recomendadas InkScape / Gimp
- Expressões matemáticas: ambiente `equation`
- Alíneas: letras minúsculas com parêntese; subalíneas com travessão
- Siglas: usar `\ac{ref}`, `\acl{ref}`, `\acs{ref}`

## Observações Importantes
- Banca questiona: *"Algum conhecimento novo foi produzido?"*
- Objetivos triviais ("desenvolver um sistema") são **evitados**
- Hipóteses devem ser **testáveis** via experimentos
- Monografia divide-se em TCC1 (introdução + fundamentação + metodologia) e TCC2 (experimentos + conclusão)

## Compilação LaTeX
```bash
pdflatex ARQUIVO_PRINCIPAL.tex
bibtex ARQUIVO_PRINCIPAL.aux
makeindex ARQUIVO_PRINCIPAL.idx
makeindex ARQUIVO_PRINCIPAL.nlo -s nomencl.ist -o ARQUIVO_PRINCIPAL.nls
pdflatex ARQUIVO_PRINCIPAL.tex
pdflatex ARQUIVO_PRINCIPAL.tex
```

## Links Úteis
- [abnTeX2 no GitHub](https://github.com/abntex/abntex2)
- [Ferramentas abnTeX2](https://github.com/abntex/abntex2/wiki/Ferramentas)
- [Tables Generator (LaTeX)](http://www.tablesgenerator.com)
- InkScape: <http://inkscape.org/>
- Gimp: <http://www.gimp.org/>
