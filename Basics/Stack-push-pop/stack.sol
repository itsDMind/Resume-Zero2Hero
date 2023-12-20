pragma solidity ^0.8.0;

contract Stack {
    uint256[] private items;

    function push(uint256 item) public {
        items.push(item);
    }

    function pop() public returns (uint256) {
        require(!isEmpty(), "pop from an empty stack");
        uint256 item = items[items.length - 1];
        items.pop();
        return item;
    }

    function isEmpty() public view returns (bool) {
        return items.length == 0;
    }
}
