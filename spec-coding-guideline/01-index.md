# Coding Guidelines Internal Specification Index

> **/goal** Master and enforce the internal repository specifications and slide deck presentation system for Coding Guidelines.
> **/learn** Read the sequentially ordered specification files in this directory when maintaining internal tools, presentation systems, or visual docs.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Maintain internal repository-specific specifications without polluting consumer bundles.
- [ ] `/learn` Adhere to Vite build pipeline and verify packaging to `slides-app/dist.zip`.
- [ ] `/goal` Verify zero explicit `true` boolean checks across slide components.
- [ ] `/learn` Run quality gate verification via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py`.

## Directory Contents

* [02-architecture.md](02-architecture.md): Slides system component and store architecture.
* [03-slide-authoring.md](03-slide-authoring.md): Guidelines for authoring interactive presentation slides.
* [04-design-tokens.md](04-design-tokens.md): Slide theme, colors, and typography tokens.
* [05-animation-primitives.md](05-animation-primitives.md): Slide deck transitions and frame animations.
* [06-curriculum.md](06-curriculum.md): Slide curriculum content mapping.
* [07-build-and-zip-pipeline.md](07-build-and-zip-pipeline.md): CI packaging to `dist.zip` for release assets.
* [08-gif-generation.md](08-gif-generation.md): Automated GIF generation pipeline for slide previews.
* [09-quality-and-offline.md](09-quality-and-offline.md): Offline viewer support and quality benchmarks.
