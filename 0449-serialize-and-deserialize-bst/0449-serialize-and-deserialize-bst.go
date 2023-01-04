package main

import (
	"bytes"
	"fmt"
)

type Codec struct {
	buf *bytes.Buffer
}

func Constructor() Codec {
	return Codec{
		buf: &bytes.Buffer{},
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.buf.Reset()
	var dfs func(node *TreeNode)
	dfs= func(node *TreeNode) {
		if node== nil {
			return
		}
		this.buf.WriteRune(rune(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return this.buf.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var build func(i,j int) *TreeNode
	var k int
	myData:=bytes.Runes([]byte(data))
	build= func(i, j int) *TreeNode{
		if i>= j {
			return nil
		}
		tmp:=myData[i]
		root:=&TreeNode{Val: int(tmp)}
		for k = i+1; k < j&&myData[k]<myData[i]; k++ {}
		root.Left= build(i+1,k)
		root.Right=build(k,j)
		return root
	}
	return build(0,len(myData))
}