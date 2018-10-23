pragma solidity ^0.4.24;

contract SimpleStorage {
  string public storedData;

  constructor() public {
    // initialize things
    storedData = "init";
  }

  // payableをつけると中身のデータを書き換えられるようになった
  function set(string x) public payable {
    storedData = x;
  }
}
