pragma solidity ^0.4.24;

contract BirrContract {

  bytes32 ipfsHash;

  struct Object {
      uint32 asNumber;
      bytes32 routeObj;
      bytes32 route6Obj;
      bytes32 autNumObj;
      bytes32 asNumObj;
  }

  mapping(address => Object) objects;

  function newObject(uint32 asNumber) public {
      objects[msg.sender] = Object(asNumber, "", "", "", "");
  }

  function getIpfsHash() public view returns (string) {
      return bytes32ToString(ipfsHash);
  }

  function setIpfsHash(string x) public {
      ipfsHash = stringToBytes32(x);
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

  function bytes32ToString(bytes32 x) private pure returns (string) {
      bytes memory bytesString = new bytes(32);
      uint charCount = 0;
      for (uint j = 0; j < 32; j++) {
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
