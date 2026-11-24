# Children Sum Property

> Striver SDE Sheet 130 · Day 19 Binary Tree III — a local equation checked globally: each parent equals its children's sum, leaves exempt, nil counts zero. Go, stdlib only, tested.

## Problem

Does every non-leaf node's value equal the sum of its children's values? Missing children count as zero; leaves (no children) satisfy the property by definition.

## Approach

A direct conjunction of local checks: at each node, compare its value against the children's sum, then require both subtrees to hold too. The subtlety is entirely in the exemptions — a *leaf* must pass unconditionally (otherwise only zero-leaved trees would qualify), while a single-child node must still balance against that one child. Short-circuit `&&` stops at the first violation. (The sheet's companion exercise — *converting* a tree to satisfy the property using only value increases — builds on this check: push parent values down, then fix parents on the way back up.)

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
