---
title: Estrutura dos Apêndices
tags:
  - status/atualizado
  - tipo/writing
  - topico/monografia
  - topico/apendices
status: pronto
created: 2026-06-05
updated: 2026-06-05
type: writing
---

# Estrutura dos Apêndices

Este documento define o conteúdo, a organização e as regras de compilação dos apêndices da monografia. A estrutura segue as recomendações do [[roadmap-monografia]] (P14) e as figuras/tabelas mínimas catalogadas em [[figuras-tabelas-monografia]].

## Apêndices Definidos

### Apêndice A — Pseudocódigo dos Métodos (`ape_metodos/pseudocodigo.tex`)

**Status:** Pronto (herdado de P14).

Conteúdo:
- Algoritmo Genético (crossover OX, mutação swap, elitismo)
- PSO com random keys (vetor real → permutação)
- ACO com feromônio 3D (transições τ(i,j,k))
- Busca exaustiva
- Lower bound via relaxação AP (Hungarian)

### Apêndice B — Resultados Experimentais Completos (`ape_resultados/resultados.tex`)

**Status:** Esqueleto criado em P14; preenchimento automático via `scripts/gerar-metricas-monografia.py` futuramente.

Conteúdo planejado:
- Tabela completa de ótimos brute-force (`10a` a `15c`)
- Tabela de gaps (BF vs AP bound) por instância e método
- Tabela de makespan médio, melhor e desvio padrão nas 30 instâncias
- Tabela de tempos computacionais por instância e método
- Estatísticas descritivas completas

Regra: gerar as tabelas com `\input{generated/tables/...}` quando os scripts estiverem integrados. Enquanto não houver geração automática, usar fragmentos manuais comentados.

### Apêndice C — Figuras Adicionais (`ape_figuras/figuras.tex`)

**Status:** Esqueleto criado em P14; figuras individuais geradas conforme necessidade.

Conteúdo planejado:
- Curvas de convergência por instância (overlays que não couberam no corpo)
- Painéis de rotas por instância não incluídos no Capítulo 4
- Diagrama CD com anotação completa
- Gráficos de estabilidade (boxplots) detalhados por instância

Regra: manter no corpo apenas as figuras que sustentam claims centrais; as demais vão para este apêndice como material suplementar.

### Apêndice D — Instâncias de Teste (`ape_instancias/instancias.tex`)

**Status:** Esqueleto criado em P14; preenchimento depende de script de extração.

Conteúdo planejado:
- Descrição das 30 instâncias sintéticas (10a–100c)
- Tabela com número de pontos (n), variante (a/b/c) e tipo
- Exemplo de coordenadas de uma instância pequena (10a)
- Exemplo de matriz de penalidade angular (triplas)
- Parâmetros de geração (sementes, distribuição espacial)

Regra: não incluir coordenadas completas de todas as instâncias no texto — apenas exemplo representativo.

## Regras de Compilação

- Os apêndices são incluídos via `\include` em `main_ppgco_ufu.tex`, dentro do ambiente `apendicesenv`.
- A ordem de inclusão segue a numeração alfabética: A (pseudocódigo), B (resultados), C (figuras), D (instâncias).
- Cada apêndice vive em seu próprio diretório `ape_<nome>/` com um único arquivo `.tex`.
- O preenchimento de tabelas e figuras geradas por script deve esperar a integração com `scripts/gerar-metricas-monografia.py`.

## Dependências

| Apêndice | Depende de | Bloqueante? |
|---|---|---|
| A (pseudocódigo) | Nenhuma | Não |
| B (resultados) | Dados em `src/data/results/` | Não (esqueleto não precisa) |
| C (figuras) | Figuras em `monografia/figs/` | Não (esqueleto não precisa) |
| D (instâncias) | Arquivos `src/data/*.points` | Não (esqueleto não precisa) |
