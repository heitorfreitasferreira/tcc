---
tags: [projeto, implementacao, brute-force, exaustivo, siglas]
sigla: "BF"
definicao: "Busca Exaustiva (Brute Force)"
incluir: sim
ocorrencias_ac: 0
ocorrencias_texto: 0
arquivos_ac: ""
---

# Busca Exaustiva — Implementação

Implementação em `src/optimization/brute/`. Usada como **baseline ótima** para instâncias pequenas ($n \leq 11$).

## Algoritmo

`src/optimization/brute/main.go` — `Optimize()`:

1. Gera **todas as $(n-1)!$ permutações** dos nós $1 \dots n-1$
2. Avalia cada uma via `g.Makespan(perm)`
3. Mantém a melhor (menor makespan)
4. Retorna a solução ótima global

Geração de permutações via **Heap's algorithm** — `generatePermutations()`:

```go
for perm := range generatePermutations(places) {
    mksp := g.Makespan(perm)
    if bestMksp > mksp {
        // novo melhor
    }
}
```

## Complexidade

- $(n-1)!$ permutações avaliadas
- Cada avaliação: $O(n)$ lookups no tensor
- Total: $O(n! \cdot n)$ — viável apenas para $n \leq 11$

## Uso

- CLM (Comprimento do Caminho Mínimo) para instâncias pequenas
- Validação da qualidade das metaheurísticas (gap para o ótimo)
- Instâncias usadas nos experimentos: `10a`, `10b`, `10c` (n=10)

## Conexões

- [[problem-formulation]] — tensor 3D usado na avaliação
- [[experiment-pipeline]] — script `run_experiments_bruteforce_missing.sh`
- [[ga]], [[pso]], [[aco]] — metaheurísticas comparadas com o ótimo

## Código-fonte

- `src/optimization/brute/main.go` — `Optimize()`, `generatePermutations()`
- Heap's algorithm com goroutine e canal para streaming de permutações
