pragma solidity ^0.5.0;

contract Store {
  address payable public owner;
  string public version;

  struct MultiHash {
    bytes32 digest;
    uint8 hashFunction;
    uint8 size;
  }

  struct Object {
    uint32 asNumber;
    MultiHash route;
    MultiHash route6;
    MultiHash autNum;
    MultiHash asSet;
  }

  mapping (address => MultiHash) public items;
  mapping (address => Object) private objects;

  event ItemSet (
    address indexed key,
    bytes32 digest,
    uint8 hashFunction,
    uint8 size
  );

  constructor (string memory _version) public {
    version = _version;
    owner = msg.sender;
  }

  function setRoute(bytes32 _digest, uint8 _hashFunction, uint8 _size) public {
    MultiHash memory item = MultiHash(_digest, _hashFunction, _size);
    objects[msg.sender].route = item;
    emit ItemSet(msg.sender, _digest, _hashFunction, _size);
  }

  function setRoute6(bytes32 _digest, uint8 _hashFunction, uint8 _size) public {
    MultiHash memory item = MultiHash(_digest, _hashFunction, _size);
    objects[msg.sender].route6 = item;
    emit ItemSet(msg.sender, _digest, _hashFunction, _size);
  }

  function setAutNum(bytes32 _digest, uint8 _hashFunction, uint8 _size) public {
    MultiHash memory item = MultiHash(_digest, _hashFunction, _size);
    objects[msg.sender].autNum = item;
    emit ItemSet(msg.sender, _digest, _hashFunction, _size);
  }

  function setAsSet(bytes32 _digest, uint8 _hashFunction, uint8 _size) public {
    MultiHash memory item = MultiHash(_digest, _hashFunction, _size);
    objects[msg.sender].asSet = item;
    emit ItemSet(msg.sender, _digest, _hashFunction, _size);
  }

  function getRoute(address addr) public view returns(bytes32 digest, uint8 hashFunction, uint8 size) {
    MultiHash memory item = objects[addr].route;
    return (item.digest, item.hashFunction, item.size);
  }

  function getRoute6(address addr) public view returns(bytes32 digest, uint8 hashFunction, uint8 size) {
    MultiHash memory item = objects[addr].route6;
    return (item.digest, item.hashFunction, item.size);
  }

  function getAutNum(address addr) public view returns(bytes32 digest, uint8 hashFunction, uint8 size) {
    MultiHash memory item = objects[addr].autNum;
    return (item.digest, item.hashFunction, item.size);
  }

  function getAsSet(address addr) public view returns(bytes32 digest, uint8 hashFunction, uint8 size) {
    MultiHash memory item = objects[addr].asSet;
    return (item.digest, item.hashFunction, item.size);
  }

  function kill() public {
    require(msg.sender == owner);
    selfdestruct(owner); // send ether to address inside parenthis
  }
}
