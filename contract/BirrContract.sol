pragma solidity ^0.4.24;

contract BirrContract {

  struct Object {
    address owner;      // Address of msg.sender
    uint32 asNumber;    // asNumber
    string hash;       // IPFS Hash

    /* string[] ipv4;  // IPv4 Prefixes
    string[] ipv6;  // IPv6 Prefixes */
  }

  mapping(address => Object) public objects;

  address[] public ownerAddresses;

  constructor(uint32 asNumber, string ipfsHash) public {
    objects[msg.sender].owner = msg.sender;
    objects[msg.sender].asNumber = asNumber;
    objects[msg.sender].hash = ipfsHash;

    ownerAddresses.push(msg.sender);
  }

  function getHowManyOwners() view public returns (uint) {
    return ownerAddresses.length;
  }

  function getOwner(uint32 i) view public returns (address) {
    return ownerAddresses[i];
  }

  function getObjects() view public returns (uint32, string, address) {
    address addr = msg.sender;

    return (objects[addr].asNumber, objects[addr].hash, objects[addr].owner);
  }
}
