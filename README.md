# mdtree

Convert markdown lists into ASCII trees.

For example, if you run:

```bash
echo -e "- a\n- b\n - ba" | mdtree
```

It will output:

```
.
├── a
└── b
    └── ba
```

Which you can then use to express a file tree, or anything else, really.

You can also customize the tree style with `--style`, and change the root
element with `--root`, for example:

```bash
$ echo -e "- foo\n- bar\n  - hi" | mdtree --root ⁜ --style rounded
```

Resulting in:

```
⁜
├── foo
╰── bar
    ╰── hi
```
