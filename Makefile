default:
	wget -c https://raw.githubusercontent.com/charmbracelet/freeze/main/configurations/full.json
	mdtree --root ⋯ --style rounded <tree.txt >tree_ascii.txt
	freeze -c ./full.json tree_ascii.txt
