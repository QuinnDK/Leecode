package main

import "fmt"

type Node struct {
	Key   string
	Value string
	pre   *Node
	next  *Node
}

func (n *Node) Init(key string, value string) {
	n.Key = key
	n.Value = value
}

var head *Node
var end *Node

var limit int

type LRUCache struct {
	limit   int
	HashMap map[string]*Node
}

func GetLRUCache(limit int) *LRUCache {
	lruCache := LRUCache{limit: limit}
	lruCache.HashMap = make(map[string]*Node, limit)
	return &lruCache
}

func (l *LRUCache) get(key string) string {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	} else {
		return ""
	}
}

func (l *LRUCache) put(key, value string) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(head)
			delete(l.HashMap, oldKey)
		}
		node := Node{Key: key, Value: value}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

func (l *LRUCache) refreshNode(node *Node) {
	if node == end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) removeNode(node *Node) string {
	if node == end {
		end = end.pre
	} else if node == head {
		head = head.next
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

func (l *LRUCache) addNode(node *Node) {
	if end != nil {
		end.next = node
		node.pre = end
		node.next = nil
	}
	end = node
	if head == nil {
		head = node
	}
}

func main() {
	lruCache := GetLRUCache(5)
	lruCache.put("001", "用户１信息")
	lruCache.put("002", "用户１信息")
	lruCache.put("003", "用户１信息")
	lruCache.put("004", "用户１信息")
	lruCache.put("005", "用户１信息")
	lruCache.get("002")
	lruCache.put("004", "用户２信息更新")
	lruCache.put("006", "用户6信息更新")
	fmt.Println(lruCache.get("001"))

	fmt.Println(lruCache.get("006"))
	fmt.Print(lruCache.HashMap)
}
