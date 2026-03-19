# Weave vs Git: Grafana Feature Toggle Demo

This repo demonstrates how [weave](https://github.com/Ataraxy-Labs/weave) resolves merge conflicts that git can't.

It uses a simplified version of Grafana's `registry.go` where every branch that adds a feature toggle appends to the same slice. When two branches add different toggles at the end, git always conflicts. Weave merges them cleanly.

## Try it yourself

### 1. See git fail

```bash
git checkout main
git merge branch-a    # conflict!
git merge --abort

git merge branch-b    # conflict!
git merge --abort
```

### 2. Install weave and try again

```bash
brew install ataraxy-labs/tap/weave
weave setup

git merge branch-a    # clean merge, both toggles included
```

### What happened?

Both branches add a new feature toggle to the end of `standardFeatureFlags`:

- **branch-a** adds `adhocFilterLabelsFromPanels`
- **branch-b** adds `vizPresets`

Git sees two insertions at the same line and conflicts. Weave parses the Go code, identifies each struct entry as an independent addition, and includes both.

This is the exact pattern described in [grafana/grafana#118134](https://github.com/grafana/grafana/issues/118134).
