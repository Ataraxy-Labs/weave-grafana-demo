# Weave vs Git: Grafana Feature Toggle Demo

This repo demonstrates how [weave](https://github.com/Ataraxy-Labs/weave) resolves merge conflicts that git can't.

It uses a simplified version of Grafana's `registry.go` where every branch that adds a feature toggle appends to the same slice. When two branches add different toggles at the end, git always conflicts. Weave merges them cleanly.

## Try it yourself

### 1. See git fail

```bash
git clone https://github.com/Ataraxy-Labs/weave-grafana-demo.git
cd weave-grafana-demo

git merge origin/branch-a         # clean (fast-forward)
git merge origin/branch-b         # CONFLICT! git can't merge both additions
git merge --abort
```

### 2. Install weave and try again

```bash
brew install ataraxy-labs/tap/weave
weave setup

git merge origin/branch-b         # clean! weave merges both toggles
```

### 3. Verify

```bash
grep 'Name:' registry.go          # 12 toggles: 10 base + adhocFilter + vizPresets
```

### 4. Revert to normal git merging

```bash
weave unsetup
```

### What happened?

Both branches add a new feature toggle to the end of `standardFeatureFlags`:

- **branch-a** adds `adhocFilterLabelsFromPanels`
- **branch-b** adds `vizPresets`

Git sees two insertions at the same line and conflicts. Weave parses the Go code, identifies each struct entry as an independent addition, and includes both.

This is the exact pattern described in [grafana/grafana#118134](https://github.com/grafana/grafana/issues/118134).
