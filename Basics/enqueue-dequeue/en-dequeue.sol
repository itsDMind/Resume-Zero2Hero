pragma solidity ^0.8.0;

contract Queue {
    uint256[] private items;

    function enqueue(uint256 item) public {
        items.push(item);
    }

    function dequeue() public returns (uint256) {
        require(!isEmpty(), "dequeue from an empty queue");
        uint256 item = items[0];
        delete items[0];
        items.pop();
        return item;
    }

    function isEmpty() public view returns (bool) {
        return items.length == 0;
    }
}
