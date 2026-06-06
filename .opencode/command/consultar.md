---
description: Busca literatura acadêmica em múltiplas bases simultaneamente (renomeado de /consultar-academico)
---

# /consultar

Executa busca acadêmica em paralelo nas bases disponíveis via MCPs e consolida resultados.

## Fluxo

1. **Disparar buscas paralelas**:
   - `academic-search_search_papers(query)` — fachada unificada (Crossref + Semantic Scholar)
   - `arxiv_search_papers(query)` — arXiv específico (preprints recentes)
   - `google-scholar` — Google Scholar (literatura cinzenta, se relevante)
2. **Consolidar**: agrupa por DOI/título, deduplica, ordena por relevância
3. **Apresentar**: tabela com título, ano, DOI, fonte, OA status

## Uso

```
/consultar "particle swarm optimization TSP survey"
/consultar "hyperparameter tuning metaheuristics"
```

## Dependências
- MCPs: `academic-search`, `arxiv`, `google-scholar`, `scholar-sidekick`
