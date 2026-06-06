# AGENTS — Visualização e Artefatos

**Servidor web é a fonte canônica para imagens de análise visual.** O comando `tcc serve` expõe `/api/render?map=X&run=Y&iteration=N` retornando `image/png` — renderização Go pura (sem dependência de browser). Todas as figuras de análise da monografia devem ser geradas por este endpoint quando possível.

- Use `bash scripts/gerar-figuras.sh` para gerar lote das figuras padrão em `monografia/figs/`. O script sobe o servidor, chama `/api/render` para cada figura desejada, depois para o servidor.
- **Se o endpoint de render não tem a visualização que você precisa, estenda-o em vez de duplicar lógica.** Adicione funções de render em `src/web/render/` e conecte em `src/web/handlers/render.go`.
- **Python também é aceitável** para gerar imagens, plotagens, tabelas ou artefatos de análise quando for genuinamente mais fácil (ex.: plots estatísticos com matplotlib/seaborn, transformações complexas com pandas). Prefira adicionar código em `scripts/` ou no notebook Jupyter `scripts/visualizacoes.ipynb`. **Não duplique lógica de visualização** — se o servidor web já renderiza o que você precisa, use-o.
- Quando uma nova análise/visualização for necessária, primeiro verifique se o servidor web já expõe os dados necessários (evolution, summary, timing, graph structure) e considere adicionar uma nova função de render ou parâmetro de query antes de recorrer a outra ferramenta.
- **Todo gráfico/plot deve expor sua escala.** Imagens de rota/mapa devem incluir eixos de coordenadas ou indicador de escala; gráficos estatísticos devem ter eixos rotulados, unidades quando aplicável, tick labels e legenda quando múltiplas séries/métodos são mostrados.
- **Prefira texto e composição em LaTeX.** Quando uma figura precisar de títulos, rótulos de painel, texto explicativo, legendas ou organização multi-painel, renderize os menores artefatos de imagem úteis separadamente e organize no `.tex` depois.
- Figuras acadêmicas devem ser específicas e autocontidas: nomes de método, IDs de instância, unidades (`makespan`, `%`, `ms`, `n`) e codificações visuais documentadas por rótulos ou legendas.
- Figuras para a monografia devem ser escritas em português (pt-BR) seguindo redação acadêmica ABNT: use `Figura`, `Fonte`, nomes de métricas em português quando natural (`tempo de otimização`, `taxa de acerto`, `iteração`), mantendo `makespan` apenas quando for a métrica objetivo definida.
- **Fluxogramas são preferencialmente artefatos TikZ `.tex`.** Mantenha fluxogramas avulsos em `monografia/figs/` como `.tex` com previews `.svg`/`.png` quando útil.
- Para gráficos gerados, prefira manter fonte e saída: script em `scripts/`, vetor (`.svg` ou `.tex`) em `monografia/figs/`, PNG para inspeção rápida.
- Figuras estilo publicação científica devem usar padrão visual compartilhado: largura fixa (aproximadamente largura total do texto), tamanhos de fonte consistentes, cores de método suaves, sem fundos decorativos, sem spines superiores/direitos desnecessários, rótulos de painel (`A`, `B`, `C`, ...).
- Código gerador de figuras deve ser organizado em torno de constantes/helpers reutilizáveis e funções de desenho por painel. Salve SVG como artefato editável primário e PNG como preview.
- Use a skill `scientific-figures` para qualquer nova pilha de figuras ou revisão substancial.
- `scripts/gerar-analises.sh` é o comando batch canônico para artefatos visuais finais.
