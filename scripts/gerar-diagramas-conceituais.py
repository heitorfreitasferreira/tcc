#!/usr/bin/env python3
"""Gera diagramas conceituais em SVG e PNG para a monografia."""

from __future__ import annotations

import math
import shutil
import subprocess
import tempfile
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
FIGS = ROOT / "monografia" / "figs"


def svg(width: int, height: int, body: list[str]) -> str:
    return "\n".join(
        [
            f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
            '<defs><marker id="arrow" markerWidth="10" markerHeight="10" refX="8" refY="3" orient="auto" markerUnits="strokeWidth"><path d="M0,0 L0,6 L9,3 z" fill="#333"/></marker></defs>',
            '<rect width="100%" height="100%" fill="white"/>',
            '<style>text{font-family:Arial,Helvetica,sans-serif;fill:#303030}.title{font-size:23px;font-weight:700}.label{font-size:14px}.small{font-size:11px}.box{fill:#f6f8fa;stroke:#6b7280;stroke-width:1.4}.arrow{stroke:#333;stroke-width:1.8;fill:none;marker-end:url(#arrow)}.math{font-size:16px;font-style:italic}</style>',
            *body,
            "</svg>",
        ]
    )


def write(name: str, content: str) -> None:
    FIGS.mkdir(parents=True, exist_ok=True)
    svg_path = FIGS / f"{name}.svg"
    png_path = FIGS / f"{name}.png"
    svg_path.write_text(content, encoding="utf-8")
    converter = shutil.which("rsvg-convert")
    if converter:
        subprocess.run([converter, "-o", str(png_path), str(svg_path)], check=True)


def tex_escape(text: str) -> str:
    replacements = {
        "\\": r"\textbackslash{}",
        "&": r"\&",
        "%": r"\%",
        "$": r"\$",
        "#": r"\#",
        "_": r"\_",
        "{": r"\{",
        "}": r"\}",
        "^": r"\textasciicircum{}",
        "~": r"\textasciitilde{}",
    }
    return "".join(replacements.get(ch, ch) for ch in text)


def write_tikz(name: str, content: str) -> None:
    FIGS.mkdir(parents=True, exist_ok=True)
    tex_path = FIGS / f"{name}.tex"
    pdf_path = FIGS / f"{name}.pdf"
    svg_path = FIGS / f"{name}.svg"
    png_path = FIGS / f"{name}.png"
    tex_path.write_text(content, encoding="utf-8")
    with tempfile.TemporaryDirectory() as tmp:
        tmp_path = Path(tmp)
        work_tex = tmp_path / tex_path.name
        work_tex.write_text(content, encoding="utf-8")
        subprocess.run(
            ["pdflatex", "-interaction=nonstopmode", work_tex.name],
            cwd=tmp_path,
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
            check=True,
        )
        generated_pdf = tmp_path / f"{name}.pdf"
        if generated_pdf.exists():
            shutil.copyfile(generated_pdf, pdf_path)
            subprocess.run(["pdftocairo", "-svg", str(pdf_path), str(svg_path)], check=True)
            subprocess.run(["pdftocairo", "-png", "-singlefile", "-r", "180", str(pdf_path), str(png_path.with_suffix(""))], check=True)


def angular_penalty() -> None:
    width, height = 760, 470
    a = (160, 300)
    b = (360, 230)
    c = (600, 170)
    body = [f'<text class="title" x="{width/2}" y="36" text-anchor="middle">Penalidade angular dependente da sequência</text>']
    body.append(f'<line class="arrow" x1="{a[0]}" y1="{a[1]}" x2="{b[0]}" y2="{b[1]}"/>')
    body.append(f'<line class="arrow" x1="{b[0]}" y1="{b[1]}" x2="{c[0]}" y2="{c[1]}"/>')
    for label, p, fill in [("k", a, "#0a66c2"), ("i", b, "#d62828"), ("j", c, "#0a66c2")]:
        body.append(f'<circle cx="{p[0]}" cy="{p[1]}" r="8" fill="{fill}"/>')
        body.append(f'<text class="label" x="{p[0]+14}" y="{p[1]-10}">{label}</text>')
    body.append('<path d="M327 242 A56 56 0 0 0 410 218" fill="none" stroke="#e07810" stroke-width="3"/>')
    body.append('<text class="label" x="360" y="185" fill="#e07810">θ</text>')
    body.append('<text class="math" x="110" y="390">G[k][i][j] = distância(i,j) + (θ / π)</text>')
    body.append('<text class="small" x="110" y="418">O custo de i → j depende do nó anterior k; por isso o tensor precisa de três índices.</text>')
    write("diagram-angular-penalty", svg(width, height, body))


def tensor_3d() -> None:
    width, height = 780, 520
    body = [f'<text class="title" x="{width/2}" y="36" text-anchor="middle">Tensor de custo 3D: G[anterior][atual][próximo]</text>']
    origin = (220, 360)
    d = (210, -70)
    vx, vy = (210, 0)
    colors = ["#e6f0ff", "#d7ebff", "#c8e4ff", "#b9ddff", "#aad6ff"]
    for layer in range(5):
        ox = origin[0] + layer * 24
        oy = origin[1] - layer * 16
        fill = colors[layer]
        body.append(f'<polygon points="{ox},{oy} {ox+vx},{oy+vy} {ox+vx+d[0]},{oy+vy+d[1]} {ox+d[0]},{oy+d[1]}" fill="{fill}" stroke="#4b5563" opacity="0.88"/>')
        for i in range(1, 4):
            x1 = ox + i * vx / 4
            y1 = oy + i * vy / 4
            body.append(f'<line x1="{x1}" y1="{y1}" x2="{x1+d[0]}" y2="{y1+d[1]}" stroke="#9ca3af"/>')
            x2 = ox + i * d[0] / 4
            y2 = oy + i * d[1] / 4
            body.append(f'<line x1="{x2}" y1="{y2}" x2="{x2+vx}" y2="{y2+vy}" stroke="#9ca3af"/>')
    body.append('<line class="arrow" x1="220" y1="390" x2="450" y2="390"/>')
    body.append('<text class="label" x="330" y="414" text-anchor="middle">próximo</text>')
    body.append('<line class="arrow" x1="190" y1="360" x2="190" y2="160"/>')
    body.append('<text class="label" x="160" y="260" transform="rotate(-90 160 260)" text-anchor="middle">atual</text>')
    body.append('<line class="arrow" x1="455" y1="350" x2="560" y2="285"/>')
    body.append('<text class="label" x="565" y="282">anterior</text>')
    body.append('<text class="math" x="110" y="465">Pré-computação O(n³) → avaliação da rota em O(n)</text>')
    write("diagram-tensor-3d", svg(width, height, body))


def flowchart(name: str, title: str, steps: list[str], color: str) -> None:
    color_hex = color.lstrip("#")
    nodes = []
    edges = []
    for i, step in enumerate(steps):
        pos = "" if i == 0 else f", below=0.8cm of step{i}"
        nodes.append(
            f"\\node[flowbox{pos}] (step{i+1}) "
            f"{{\\strut {tex_escape(step)}}};"
        )
        if i > 0:
            edges.append(f"\\draw[arrow] (step{i}) -- (step{i+1});")

    content = "\n".join(
        [
            r"\documentclass[tikz,border=6pt]{standalone}",
            r"\usepackage[T1]{fontenc}",
            r"\usepackage[utf8]{inputenc}",
            r"\usepackage{xcolor}",
            r"\usetikzlibrary{arrows.meta,positioning,shapes.geometric}",
            f"\\definecolor{{methodcolor}}{{HTML}}{{{color_hex}}}",
            r"\begin{document}",
            r"\begin{tikzpicture}[",
            r"  font=\sffamily,",
            r"  flowbox/.style={rectangle, rounded corners=3pt, draw=gray!70, fill=gray!6, very thick, minimum width=9.8cm, minimum height=0.95cm, align=center, text width=9.2cm, left color=methodcolor!18, right color=gray!4},",
            r"  arrow/.style={-{Latex[length=3mm]}, thick, gray!70}",
            r"]",
            f"\\node[font=\\sffamily\\bfseries\\Large] (title) {{{tex_escape(title)}}};",
            r"\node[flowbox, below=0.75cm of title] (step1) {\strut " + tex_escape(steps[0]) + r"};",
            *nodes[1:],
            *edges,
            r"\end{tikzpicture}",
            r"\end{document}",
        ]
    )
    write_tikz(name, content)


def flowcharts() -> None:
    flowchart(
        "flowchart-ga",
        "Algoritmo Genético (GA)",
        [
            "Inicializa população de permutações",
            "Avalia makespan via tensor G",
            "Seleciona pais por torneio",
            "Aplica crossover OX + mutação swap",
            "Preserva elitismo e repete gerações",
        ],
        "#0a66c2",
    )
    flowchart(
        "flowchart-pso",
        "Particle Swarm Optimization (PSO)",
        [
            "Inicializa posições random-keys",
            "Ordena chaves para obter permutação",
            "Avalia pBest e gBest",
            "Atualiza velocidade e posição",
            "Repete até o limite de iterações",
        ],
        "#20913c",
    )
    flowchart(
        "flowchart-aco",
        "Ant Colony Optimization (ACO)",
        [
            "Inicializa feromônio 3D tau[k][i][j]",
            "Cada formiga constrói uma rota",
            "Escolhe próximo nó por roleta tau^alpha eta^beta",
            "Evapora e deposita feromônio",
            "Retorna melhor rota encontrada",
        ],
        "#e07810",
    )
    flowchart(
        "flowchart-lowerbound",
        "Lower Bound (AP)",
        [
            "Reduz tensor 3D para matriz 2D c'[j][k]",
            "Resolve Assignment Problem (Hungarian O(n³))",
            "Obtém bound como soma das atribuições",
            "Retorna limitante inferior do makespan",
        ],
        "#7b2d8e",
    )
    flowchart(
        "flowchart-bruteforce",
        "Busca Exaustiva",
        [
            "Gera permutações por algoritmo de Heap",
            "Avalia cada rota via tensor G",
            "Mantém o menor makespan",
            "Retorna ótimo global",
        ],
        "#333333",
    )


def main() -> None:
    angular_penalty()
    tensor_3d()
    flowcharts()
    print(f"Diagramas conceituais gerados em {FIGS}")


if __name__ == "__main__":
    main()
