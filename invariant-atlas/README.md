# Invariant Atlas

A concise study website for exercises 082–120.

From this directory, run:

```sh
node build.mjs
node serve.mjs
```

Open http://localhost:4173. You can also open the generated `dist/index.html` directly. No package installation is required.

- Search by pattern, keyword, or exercise number.
- Select a card for the invariant, preservation step, stopping rule, example, cost, and source.
- Use Recall mode to hide answers. Review marks stay in the current browser.
- Prefix sum and three-way partition have interactive traces.

Edit `src/data.js` for study notes, `src/style.css` for design, and `src/app.js` for interactions. Run `node build.mjs` after changes. Git tracks `src/`; generated `dist/` files are ignored. Fonts have system fallbacks.

The build also copies Go source snapshots from the parent repository. To refresh only those snapshots, run `node sync-sources.mjs`. Neither command regenerates the editorial notes. Update the visible snapshot date in `src/index.html` when refreshing. The Go implementations are not modified.

The prose uses short, direct technical English inspired by STE. It is not a formally certified ASD-STE100 document.
