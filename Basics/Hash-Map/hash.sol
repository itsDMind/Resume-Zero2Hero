pragma solidity ^0.8.0;

contract HashMap {
    mapping(string => string) myMap;

    function insert(string memory key, string memory value) public {
        myMap[key] = value;
    }

    function retrieve(string memory key) public view returns (string memory) {
        return myMap[key];
    }
}
