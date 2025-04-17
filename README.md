# mdtree

Convert markdown lists into ASCII trees.

![screenshot](https://github.com/user-attachments/assets/e233ad76-2ee1-4395-a8be-88fabc551b67)

For example, if you run:

```bash
cat <<EOF | mdtree
- a
- b
  - ba
EOF
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
cat <<EOF | mdtree --root ⁜ --style rounded
- foo
- bar
  - hi
EOF
```

Resulting in:

```
⁜
├── foo
╰── bar
    ╰── hi
```

## Install

- macOS: `brew install caarlos0/tap/mdtree`
- nix: use [this NUR](https://github.com/caarlos0/nur)
- others: download from the [latest release](https://github.com/caarlos0/mdtree/releases/latest)

## License

MIT
