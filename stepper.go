package src

type Attributes map[string]string

type ElementData struct {
	tagName string
	attributes Attributes
}

type NodeType interface {}

type TextType struct {
	string
}

type ElementType struct {
	ElementData
}

type Node struct {
	children []Node
	nodeType NodeType
}

func Text(data string) Node {
	return Node{
		nodeType: data,
	}
}

func Element(name string, attributes Attributes, children[]Node) Node {
	return Node{
		children: children,
		nodeType: ElementData{
			tagName:    name,
			attributes: attributes,
		},
	}
}
