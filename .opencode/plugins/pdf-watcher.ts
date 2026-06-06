import type { Plugin } from "@opencode-ai/plugin"

export const PdfWatcher: Plugin = async ({ $, directory, client }) => {
  const LOG = ".opencode/log/pdf-events.log"

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

        const existing = await $`ls vault/papers/${key}.md 2>/dev/null || true`.text()
        if (existing.trim()) {
          await client.app.log({
            body: {
              service: "pdf-watcher",
              level: "debug",
              message: `Nota vault já existe para ${key}, ignorando`,
            },
          })
          continue
        }

        await client.app.log({
          body: {
            service: "pdf-watcher",
            level: "info",
            message: `Novo PDF detectado: ${filename} — execute /incorporar "${p}"`,
            extra: { path: p, key },
          },
        })

        const doi = await $`bash scripts/extract-pdf-doi.sh "${p}" 2>/dev/null || true`.text()
        const doiInfo = doi.trim() ? ` DOI: ${doi.trim()}` : ""

        await $`mkdir -p .opencode/log`
        await $`echo "$(date -Iseconds) | ${p}${doiInfo}" >> ${LOG}`
      }
    },
  }
}
