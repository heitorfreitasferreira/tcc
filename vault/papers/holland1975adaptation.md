---
title: "Adaptation in Natural and Artificial Systems"
authors: [Holland, John H.]
year: 1975
doi: ""
bibtex_key: holland1975adaptation
bibtex-key: holland1975adaptation
pdf: "papers/pdfs/holland1975adaptation.pdf"
tags: [ga foundational]
status: lido-parcial
rating: 4
---

## PDF

![[holland1975adaptation.pdf]]

## Resumo

Holland formula uma teoria geral de adaptação em sistemas naturais e artificiais. O livro define adaptação como modificação progressiva de estruturas para melhorar desempenho em um ambiente, com ênfase em problemas nos quais há incerteza, não linearidade e conflito entre explorar alternativas e explorar soluções conhecidas. A contribuição central é tratar populações como portadoras de muitos esquemas simultâneos, permitindo paralelismo intrínseco na avaliação de blocos parciais de solução.

> [!warning] Leitura parcial
> O PDF local tem texto extraível, mas com OCR irregular e espaçamento degradado. Citações devem ser conferidas visualmente antes de uso final.

## Contribuições Principais

- Generaliza o conceito de adaptação para biologia, aprendizagem, economia, controle, inteligência artificial e otimização.
- Introduz o conceito de schema como generalização de conjuntos coadaptados de genes.
- Formula operadores genéticos generalizados: crossing-over, inversão e mutação.
- Defende o paralelismo intrínseco: uma estrutura avaliada fornece informação sobre muitos esquemas.
- Relaciona algoritmos genéticos, alocação ótima de tentativas e o problema do bandido multiarmado.

## Relevância para o TCC

Holland fornece a base conceitual do GA usado no TCC. Para o problema de rotas de drones, cada rota candidata pode ser vista como uma estrutura avaliada por desempenho, enquanto subestruturas de rota funcionam como blocos parciais que podem ser preservados, recombinados ou descartados. Essa leitura justifica o uso de população, seleção, crossover e mutação no [[tsp]], especialmente quando o espaço de busca é grande demais para enumeração exaustiva.

## Métodos e Abordagens

- Formalização de estruturas, ambientes, operadores e planos adaptativos.
- Uso de esquemas para decompor informação parcial de soluções.
- Crossover como recombinação de segmentos entre estruturas.
- Mutação como fonte de variação.
- População como base compacta de informações sobre muitos esquemas.

## Conexões

- [[goldberg1989genetic]]
- [[genetic-algorithms]]
- [[tsp]]
- [[oliver1987crossover]]
- [[potvin1996ga]]
- [[larranaga1999ga]]
- [[kennedy1995particle]]

## Notas e Insights

- O prefácio de 1992 registra a mudança terminológica de “genetic plan” para “genetic algorithm”.
- Holland enfatiza melhoria adaptativa mais do que otimização global estrita, útil para explicar metaheurísticas em instâncias grandes ou dinâmicas.
- A teoria de esquemas ajuda a escrever a fundamentação do GA sem reduzi-lo a uma metáfora biológica superficial.

## Citações-chave

> “Basically, adaptive processes are optimization processes...”

> “The possibility of ‘intrinsic parallelism’ — the testing of many schemata by testing a single structure — is a direct offshoot of this approach.”
