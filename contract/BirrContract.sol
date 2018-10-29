pragma solidity ^0.4.24;

contract BirrContract {
  string public storedData;

  constructor() public {
    // initialize things
    storedData = "init";
  }

  function set(string x) public payable {
    storedData = x;
  }
}
