package models

// Node describes tree's node
type Node struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Root    int    `json:"root"`
	Parent  int    `json:"parent"`
	Height  int    `json:"height"`
	Path    string `json:"path"`
}
