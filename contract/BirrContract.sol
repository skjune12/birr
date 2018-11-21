pragma solidity ^0.5.0;

contract BirrContract {

  struct Object {
    uint32 asNumber;
    bytes32 routeObj;
    bytes32 route6Obj;
    bytes32 autNumObj;
    bytes32 asSetObj;
  }

  mapping(address => Object) objects;
  address[] public addressList;

  function newObject(uint32 asNumber) public {
    objects[msg.sender] = Object(asNumber, "", "", "", "");
    addressList.push(msg.sender);
  }

  function getObject(address addr) public view
  returns (uint32, string memory, string memory, string memory, string memory)
  {
    Object memory obj = objects[addr];
    string memory routeObj = bytes32ToString(obj.routeObj);
    string memory route6Obj = bytes32ToString(obj.route6Obj);
    string memory autNumObj = bytes32ToString(obj.autNumObj);
    string memory asSetObj = bytes32ToString(obj.asSetObj);

    return (obj.asNumber, routeObj, route6Obj, autNumObj, asSetObj);
  }

  function setRouteObj(string memory ipfsHash) public {
    objects[msg.sender].routeObj = stringToBytes32(ipfsHash);
  }

  function setRoute6Obj(string memory ipfsHash) public {
    objects[msg.sender].route6Obj = stringToBytes32(ipfsHash);
  }

  function setAutNumObj(string memory ipfsHash) public {
    objects[msg.sender].autNumObj = stringToBytes32(ipfsHash);
  }

  function setAsSetObj(string memory ipfsHash) public {
    objects[msg.sender].asSetObj = stringToBytes32(ipfsHash);
  }

  function numOfAddrs() public view returns (uint) {
    return addressList.length;
  }

  function stringToBytes32(string memory source) private pure returns (bytes32 result) {
    bytes memory tempEmptyStringTest = bytes(source);
    if (tempEmptyStringTest.length == 0) {
      return 0x0;
    }
    assembly {
      result := mload(add(source, 32))
    }
  }

  function bytes32ToString(bytes32 x) private pure returns (string memory) {
    bytes memory bytesString = new bytes(32);
    uint charCount = 0;
    uint j = 0;
    for (j = 0; j < 32; j++) {
      byte char = byte(bytes32(uint(x) * 2 ** (8 * j)));
      if (char != 0) {
        bytesString[charCount] = char;
        charCount++;
      }
    }
    bytes memory bytesStringTrimmed = new bytes(charCount);
    for (j = 0; j < charCount; j++) {
      bytesStringTrimmed[j] = bytesString[j];
    }
    return string(bytesStringTrimmed);
  }
}
