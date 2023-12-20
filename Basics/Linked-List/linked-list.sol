pragma solidity ^0.8.0;

contract LinkedList {
    struct Node {
        uint256 data;
        uint256 next;
    }

    mapping(uint256 => Node) public nodes;
    uint256 public head;

    constructor() {
        head = 0;
    }

    function insert(uint256 data) public {
        nodes[head] = Node(data, head);
        head = nodes[head].next;
    }

    function delete(uint256 key) public {
        uint256 current = head;
        uint256 prev = 0;

        while (current != 0 && nodes[current].data != key) {
            prev = current;
            current = nodes[current].next;
        }

        if (current == 0) {
            return;
        }

        nodes[prev].next = nodes[current].next;
    }
}
