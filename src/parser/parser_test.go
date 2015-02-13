package main

import "testing"

//
/*
terminal text nodes cannot be pruned
only siblings prune the weak links
Requirements:
I want to depth first search a node tree
At the root node
    return Terminal = true
    and
    if a text node
        return text length
    else
        return 0
At a node
    return Terminal = false
    if a text node

*/

func TestParseNode(t *testing.T) {
}
