package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func main() {
	var style string
	var root string
	cmd := &cobra.Command{
		Use:   "mdtree",
		Short: "Convert markdown lists into ASCII trees",
		Long: `Convert markdown lists into ASCII trees.

For example, if you run:

  echo -e "- a\n- b\n  - ba" | mdtree

It will output:
  .
  ├── a
  └── b
      └── ba

Which you can then use to express a file tree, or anything else, really.`,
		Example: `  mdtree <file.md
  mdtree --root ⁜ --style rounded <file.md
  echo -e "- foo\n- bar\n  - hi" | mdtree
  cat file.txt | mdtree`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(renderTree(cmd.InOrStdin(), root, enumFor(style)))
		},
	}

	cmd.Flags().StringVarP(&style, "style", "s", "default", "Which tree style to use, either default, rounded")
	cmd.Flags().StringVarP(&root, "root", "r", ".", "String to use as the root node")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func enumFor(style string) tree.Enumerator {
	switch style {
	case "rounded":
		return tree.RoundedEnumerator
	default:
		return tree.DefaultEnumerator
	}
}

func renderTree(r io.Reader, root string, enum tree.Enumerator) string {
	source, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	node := goldmark.DefaultParser().Parse(text.NewReader(source))

	// Document
	if !node.HasChildren() {
		return ""
	}

	// List
	node = node.FirstChild()
	if node.Kind() != ast.KindList {
		return ""
	}
	if !node.HasChildren() {
		return ""
	}

	return buildTree(source, node.FirstChild(), root, enum).String()
}

func buildTree(source []byte, node ast.Node, root string, enum tree.Enumerator) *tree.Tree {
	tree := tree.Root(root).Enumerator(enum)
	for {
		if node.Kind() == ast.KindListItem {
			if node.LastChild().Kind() == ast.KindList {
				subtree := buildTree(source, node.LastChild().FirstChild(), root, enum)
				subtree.Root(string(node.FirstChild().Text(source)))
				tree.Child(subtree)
			} else {
				tree.Child(string(node.FirstChild().Text(source)))
			}
		}
		node = node.NextSibling()
		if node == nil {
			break
		}
	}

	return tree
}
