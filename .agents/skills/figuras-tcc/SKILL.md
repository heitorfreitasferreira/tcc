---
name: figuras-tcc
description: Cria figuras, gráficos e artefatos visuais para a monografia TCC. Cobre convenções do projeto (render server Go, batch scripts, estilo ABNT pt-BR) e princípios Tufte de design científico (data-ink ratio, small multiples, LaTeX-first composition). Consolida project-figures e scientific-figures.
---

# Figuras TCC

Skill consolidada para geração de figuras da monografia.

---

## 1. Convenções do Projeto

Adaptado de `project-figures`.

### Fonte Canônica

O **web server** (`tcc serve`) é a fonte canônica para imagens de análise visual. Expõe `/api/render?map=X&run=Y&iteration=N` retornando `image/png`.

- `bash scripts/gerar-figuras.sh` gera figuras padrão em `monografia/figs/`
- Se o endpoint não tiver o que precisa, **estenda** `src/web/render/` e `src/web/handlers/render.go`
- **Python é aceitável** para gráficos estatísticos (matplotlib/seaborn). Código em `scripts/` ou `scripts/visualizacoes.ipynb`
- Primeiro verifique se o web server já expõe os dados necessários

### Estilo

- **Toda figura deve expor sua escala**: eixos rotulados, unidades, ticks, legenda para múltiplas séries
- **Prefira composição LaTeX**: renderize artefatos de imagem mínimos e organize painéis em `.tex`
- Figuras autocontidas: nomes de métodos, IDs de instância, unidades (`makespan`, `%`, `ms`, `n`)
- Monografia em **português (pt-BR)**, ABNT: `Figura`, `Fonte`
- **Fluxogramas: TikZ `.tex`** em `monografia/figs/` com preview `.svg`/`.png`
- Scripts fonte em `scripts/`, saída vetorial (`.svg`/`.tex`) em `monografia/figs/`, PNG para inspeção rápida
- Cores muted para métodos, sem fundos decorativos, spines mínimos, labels de painel (`A`, `B`, `C`...)
- `scripts/gerar-analises.sh` é o comando canônico de batch para artefatos visuais finais

---

## 2. Princípios Tufte

Adaptado de `scientific-figures`.

### Regra Central

Toda figura deve responder uma pergunta de pesquisa. Se a imagem não apoia um claim, revela um padrão, compara alternativas ou documenta um método, não a gere.

### Data-Ink Ratio

- Maximize data-ink: mantenha marcas que codificam dados; remova bordas, sombras, gradientes, grids pesados, legendas redundantes
- Evite chartjunk: sem barras 3D, texturas, ícones, cores ornamentais, pie charts explodidos
- Prefira labels diretos sobre legendas quando couberem perto dos dados
- Use small multiples para comparar mesma estrutura entre métodos, instâncias, seeds
- Use escalas consistentes entre painéis quando a comparação depende de magnitude
- Revele variação, não só tendência central: inclua CI, desvio padrão, boxplots, violin plots ou pontos individuais
- Cores muted com propósito: codificam método, grupo, estado ou incerteza
- Use codificações redundantes para impressão grayscale: estilo de linha, forma de marcador, labels

### Requisitos de Figura

- Todo gráfico expõe sua escala
- Todo eixo tem label e unidades
- Toda figura multi-série identifica cada método/grupo
- Todo gráfico estatístico mostra tamanho da amostra
- Toda imagem de rota/mapa inclui eixos, ticks de coordenada ou indicador de escala
- Toda figura é legível no tamanho que aparecerá no documento
- Nome de arquivo descreve o conteúdo, não o passo do script

---

## 3. Tipos de Figura Recomendados

| Pergunta | Prefira | Evite |
|---|---|---|
| Comparar métodos entre instâncias | Dot plot, slopegraph, small multiples | Pie chart, barras 3D |
| Mostrar convergência | Linha com banda de incerteza ou small multiples | Uma linha ruidosa por seed sem sumário |
| Mostrar distribuição entre seeds | Boxplot + pontos, violin + pontos, ECDF | Barras só com média |
| Mostrar scaling de runtime | Gráfico log-scale linha/ponto com unidades | Barras sem labels |
| Mostrar solução de rota | Mapa com rota, nós, escala, método, instância | Fundo decorativo de drone |
| Mostrar relações | Scatter com fit só se justificado, resíduos | Curva suavizada sem descrição |
| Mostrar processo/algoritmo | Fluxograma TikZ | Caixas decorativas com setas |

---

## 4. Workflow de Análise Visual

### 1. Defina o Claim

Antes de gerar figura, escreva: `Esta figura apoia o claim de que...`

### 2. Inspecione os Dados

- Verifique arquivos fonte
- Cheque unidades, definições de métrica, missing values, seeds repetidas
- Confirme se valores menores ou maiores são melhores
- Confirme se comparações devem ser absolutas, relativas, normalizadas, ranqueadas

### 3. Escolha Codificações

- Posição em escala comum é a codificação mais forte
- Comprimento é aceitável para magnitudes simples
- Cor codifica categorias, não magnitude ordenada (a menos que use colormap perceptual)
- Área, volume e ângulo são codificações fracas; evite

### 4. Gere Reproduzivelmente

- Constantes reutilizáveis no topo: paths, tamanho, cores, ordem de métodos, labels, unidades
- Uma função por painel em figuras compostas
- SVG primeiro, PNG depois para charts
- Para mapas/rotas: renderize painéis individuais via endpoint Go, componha labels/captions em `.tex`

---

## 5. Composição LaTeX-First

- LaTeX é a camada de composição preferida
- Não coloque títulos de figura, texto explicativo longo, captions ou notas ABNT em imagens raster
- Gere os menores artefatos de imagem úteis: um mapa, um painel de rota, um chart
- Componha painéis em `.tex` com `figure`, `subfigure`/`subcaption` ou `minipage`
- Deixe LaTeX controlar largura final com `\includegraphics[width=0.9\textwidth]{...}`
- Fragments `.tex` gerados em `monografia/figs/`

---

## 6. Baseline Matplotlib

```python
import matplotlib.pyplot as plt

METHOD_COLORS = {
    "bruteforce": "#4c4c4c",
    "ga": "#4c78a8",
    "pso": "#f58518",
    "aco": "#54a24b",
}

plt.rcParams.update({
    "figure.dpi": 120, "savefig.dpi": 300,
    "font.size": 9, "axes.titlesize": 10,
    "axes.labelsize": 9, "xtick.labelsize": 8,
    "ytick.labelsize": 8, "legend.fontsize": 8,
    "axes.spines.top": False, "axes.spines.right": False,
    "axes.grid": True, "grid.color": "#d9d9d9",
    "grid.linewidth": 0.5, "grid.alpha": 0.7,
    "legend.frameon": False,
})
```

---

## 7. Específico TSP / Patrulha de Drones

- Enquadre figuras em torno de: qualidade de rota, makespan, convergência, custo computacional, robustez entre seeds
- Imagens de rota: inclua instance ID, método, seed/run ID, iteração, valor objetivo
- Nomes de método consistentes: `bruteforce`, `ga`, `pso`, `aco`
- Gap de otimalidade como `%`: `(método - ótimo) / ótimo * 100` para métricas de minimização
- Runtime: unidades `ms`, `s` ou `min`; use log scale só quando ordens de magnitude diferem

---

## 8. Honestidade Estatística

- Não use bar charts para distribuições (a menos que sejam contagens/proporções)
- Não trunque eixos para amplificar efeitos pequenos (a menos que explícito e justificado)
- Não suavize curvas de convergência sem declarar método e janela
- Não esconda runs falhos; codifique ou reporte
- Não misture métricas com unidades diferentes no mesmo eixo
- Não compare métodos estocásticos com single seed a menos que explicitamente ilustrativo

---

## 9. Checklist de Qualidade

- [ ] A figura apoia um claim específico
- [ ] Fonte de dados é real e documentada
- [ ] Eixos, ticks, unidades e legendas/labels presentes
- [ ] Escala visível para cada plot ou mapa
- [ ] Cores consistentes, muted e significativas
- [ ] Incerteza ou distribuição mostrada para resultados estocásticos
- [ ] Clutter visual removido sem remover evidência
- [ ] Imagem legível no tamanho do documento
- [ ] Output reproduzível de script ou endpoint
- [ ] SVG/PDF/TEX fonte e PNG preview salvos
- [ ] Texto e composição em `.tex` quando melhorar legibilidade

## Anti-Patterns

- Dashboards decorativos como figuras acadêmicas
- Barras só com média para otimizadores estocásticos
- Figuras sem unidades ou tamanho de amostra
- Legendas que repetem informação já visível em labels
- Paletas que dependem só de contraste vermelho/verde
- Screenshots de plots em vez de artefatos gerados
- Mapas de rota sem escala de coordenadas
- Claims no texto mais fortes que o que a figura mostra
