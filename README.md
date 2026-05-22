# Hugo Static Site on GitHub Pages

This repository hosts a static site built with [Hugo](https://gohugo.io/) and the [hugo-coder](https://github.com/luizdepra/hugo-coder) theme. It is automatically deployed to [GitHub Pages](https://pages.github.com/) via GitHub Actions.

## ✨ Features
- **Hugo (Extended):** Fast static site generation with SCSS support.
- **Auto Deployment:** Seamless CI/CD using GitHub Actions on every push to `main`.
- **Hugo Modules:** Theme management as a module (no need for `themes/` folder).

## 🛠 Prerequisites
- **Hugo Extended:** ≥ 0.146.0
- **Dart Sass:** Available in your `PATH` (e.g., `npm i -g sass`)
- **Go:** Required for Hugo Modules.

## 📂 Project Structure
```text
.
├── .github/workflows/  # Deployment automation
├── archetypes/         # Content templates
├── content/            # Markdown content
│   ├── about.md        # "About Me" page
│   └── posts/          # Blog articles
└── hugo.toml           # Site configuration
```

## 🚀 Getting Started

### Local Development
1. Install dependencies:
   ```bash
   hugo mod tidy
   ```
2. Run the development server:
   ```bash
   hugo server -D
   ```
3. Open [http://localhost:1313](http://localhost:1313) in your browser.

### Create New Content
```bash
hugo new posts/my-new-article.md
```

## 🚢 Deployment
Deployment is automatic. When you push to the `main` branch, the [GitHub Actions workflow](.github/workflows/gh-pages.yml) builds the site and deploys it to GitHub Pages.

> **Note:** Ensure your repository settings under **Settings → Pages** have "Build and deployment → Source" set to **GitHub Actions**.

## 📝 License
This project is personal content. The theme follows its own [license](https://github.com/luizdepra/hugo-coder/blob/main/LICENSE).

---
*Created by [Alejandro Flores](https://afr-dt.github.io/)*
