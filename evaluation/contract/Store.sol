pragma solidity ^0.5.1;

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

    // for managing route objects
    mapping(bytes32 => MultiHash) routes;
    mapping(bytes32 => bool) existRouteKeys;
    bytes32[] routeKeys;
  }

  // mapping (address => MultiHash) public items;
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

  function find(address addr, bytes32 key) private view returns(uint i) {
    i = 0;
    while (objects[addr].routeKeys[i] != key) {
      i++;
    }
    return i;
  }

  function contains(address addr, bytes32 key) private view returns (bool) {
    return objects[addr].existRouteKeys[key];
  }

  function removeByIndex(address addr, uint i) private {
    while (i < objects[addr].routeKeys.length-1) {
      objects[addr].routeKeys[i] = objects[addr].routeKeys[i+1];
      i++;
    }

    objects[addr].routeKeys.length--;
  }

  function removeByKey(address addr, bytes32 key) private {
    uint i = find(addr, key);
    removeByIndex(addr, i);
    delete objects[addr].routes[key];
    delete objects[addr].existRouteKeys[key];
  }

  function removeRoute(bytes32 key) public {
    removeByKey(msg.sender, key);
  }

  // under delelopping
  function setRoute(bytes32 key, bytes32 _digest, uint8 _hashFunction, uint8 _size) public {
    MultiHash memory item = MultiHash(_digest, _hashFunction, _size);
    objects[msg.sender].routes[key] = item;
    objects[msg.sender].routeKeys.push(key);
    objects[msg.sender].existRouteKeys[key] = true;
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

  function getRoute(address addr, bytes32 key) public view returns(bytes32 digest, uint8 hashFunction, uint8 size) {
    MultiHash memory item = objects[addr].routes[key];
    return (item.digest, item.hashFunction, item.size);
  }

  function getRouteKeys() public view returns(bytes32[] memory) {
    return objects[msg.sender].routeKeys;
  }

  function existsRouteKey(bytes32 key) public view returns(bool) {
    return objects[msg.sender].existRouteKeys[key];
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
