---
description: Busca literatura acadêmica em múltiplas bases simultaneamente
---

# /consultar-academico

Executa busca acadêmica em paralelo nas bases disponíveis via MCPs e consolida resultados.

## Fluxo

1. **Disparar buscas paralelas**:
   - `crossref_search_works(query)` — metadados + citações (~155M works)
   - `semantic-scholar` (search_papers / search_by_topic) — 16 tools, TL;DR, PDF OA
   - `arxiv_search_papers(query)` — arXiv específico
   - `google-scholar` — Google Scholar (se relevante)
2. **Consolidar**: agrupa resultados por DOI/título, deduplica, ordena por relevância
3. **Apresentar**: tabela com título, ano, DOI, fonte, OA status

## Uso

```
/consultar-academico "particle swarm optimization TSP survey"
/consultar-academico "hyperparameter tuning metaheuristics"
```

## Dependências

- MCPs: `crossref`, `semantic-scholar`, `arxiv`, `google-scholar`, `scholar-sidekick`
