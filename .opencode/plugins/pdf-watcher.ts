import type { Plugin } from "@opencode-ai/plugin"

export const PdfWatcher: Plugin = async ({ $, client }) => {
  await client.app.log({
    body: {
      service: "pdf-watcher",
      level: "info",
      message: "PDF watcher plugin iniciado — monitorando vault/papers/pdfs/",
    },
  })

  return {
    "file.watcher.updated": async (input) => {
      const paths = Array.isArray(input.paths) ? input.paths : [input.paths ?? input.path]
      for (const p of paths) {
        if (!p || typeof p !== "string") continue
        if (!p.endsWith(".pdf")) continue
        if (!p.includes("vault/papers/pdfs/")) continue

        const filename = p.split("/").pop()!
        const key = filename.replace(/\.pdf$/i, "")

        const existingNote = await $`ls vault/papers/${key}.md 2>/dev/null || true`.text()
        const existingTask = await $`grep -l "acao.*incorporar.*${key}" vault/roadmap/tarefas/P*.md 2>/dev/null || true`.text()

        if (existingNote.trim() || existingTask.trim()) {
          await client.app.log({
            body: {
              service: "pdf-watcher",
              level: "debug",
              message: `Nota ou tarefa já existe para ${key}, ignorando`,
            },
          })
          continue
        }

        const doi = await $`bash scripts/extract-pdf-doi.sh "${p}" 2>/dev/null || true`.text()
        const doiStr = doi.trim()

        await client.app.log({
          body: {
            service: "pdf-watcher",
            level: "info",
            message: `Novo PDF detectado: ${filename}${doiStr ? " (DOI: " + doiStr + ")" : ""} — criando tarefa na fila`,
            extra: { path: p, key, doi: doiStr },
          },
        })

        const saida = doiStr
          ? `Nota vault/papers/${key}.md enriquecida com metadados de ${doiStr}`
          : `Nota vault/papers/${key}.md criada e enriquecida (DOI não detectado automaticamente)`

        const result = await $`bash scripts/roadmap.sh tarefa criar "Incorporar artigo: ${key}" "${saida}" alta literatura`.text()
        await client.app.log({
          body: {
            service: "pdf-watcher",
            level: "info",
            message: `Tarefa criada na fila: ${result.trim()}`,
          },
        })
      }
    },
  }
}
