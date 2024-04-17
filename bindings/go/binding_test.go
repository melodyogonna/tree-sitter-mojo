package tree_sitter_MOJO_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-MOJO"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_MOJO.Language())
	if language == nil {
		t.Errorf("Error loading Mojo grammar")
	}
}
