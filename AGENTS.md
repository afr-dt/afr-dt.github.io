# Repository Guidelines

## Project Structure & Module Organization
- `hugo.toml`: Site config (baseURL, params, Hugo Module import for theme).
- `content/`: Markdown pages. Home lives at `content/_index.md`.
- `assets/`: Theme/site assets (SCSS, images) processed by Hugo.
- `layouts/`: Custom templates overriding the theme.
- `static/`: Files served as‑is (e.g., `static/robots.txt`).
- `archetypes/`: Front‑matter templates for new content.
- `data/`: YAML/TOML/JSON data files for templates.
- Theme: provided via Hugo Module (`github.com/LordMathis/hugo-theme-nightfall`), no `themes/` folder needed.
- Output: Hugo builds to `public/` (uploaded by CI; do not commit).

## Build, Test, and Development Commands
- Local dev: `hugo server -D` — start live server with drafts.
- Build: `hugo` — generate static site into `public/`.
- Clean build: `rm -rf public && hugo --gc --minify` — garbage collect and minify.
- Resolve modules: `hugo mod get` — fetch/refresh Hugo Modules.
- Tooling: Use Hugo Extended 0.146.0+. Dart Sass is required by the theme.

## Coding Style & Naming Conventions
- Markdown: one H1 per page; wrap at ~100 cols; use fenced code blocks.
- Front matter: YAML (`---`) with concise titles; keep slugs lowercase.
- Filenames: lowercase, hyphen‑separated (e.g., `my-page.md`).
- Templates: prefer partials; keep logic minimal in templates.
- CSS/SCSS: place in `assets/`; use variables; avoid inline styles.

## Testing Guidelines
- Build checks: `hugo --printI18nWarnings --printPathWarnings` to catch missing keys/links.
- Links: verify internal links resolve locally; avoid absolute URLs unless needed.
- Optional lint: if installed, run `markdownlint **/*.md` for Markdown style.

## Commit & Pull Request Guidelines
- Commits: follow Conventional Commits (`feat:`, `fix:`, `docs:`, `ci:`, `refactor:`).
- PRs: include purpose, screenshots of visual changes, and link related issues.
- CI: `.github/workflows/gh-pages.yml` builds `public/` and deploys to GitHub Pages.

## Security & Configuration Tips
- Base URL: keep `baseURL` in `hugo.toml` accurate for correct asset links.
- Do not commit `public/` or secrets. Theme changes should be made via overrides in `layouts/` or `assets/`.
