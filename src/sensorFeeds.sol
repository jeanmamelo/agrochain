// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./chainlinkInterface.sol";

contract SensorConsumerV3 {

    /**
     * Please abstract the data:
     * BTC/USD would be Temperature sensor feed
     * ETH/USD would be Humidity sensor feed
     * LINK/USD would be Pressure sensor feed
     */

    AggregatorV3Interface internal TemperatureSensorFeed;
    AggregatorV3Interface internal HumiditySensorFeed;
    AggregatorV3Interface internal PressureSensorFeed;

    /**
     * Network: Sepolia Testnet
     * According to Chainlink docs: https://docs.chain.link/data-feeds/price-feeds/addresses?network=ethereum&page=1

     * BTC/USD Address: 0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43
     * ETH/USD Address: 0x694AA1769357215DE4FAC081bf1f309aDC325306
     * LINK/USD Address:0x42585eD362B3f1BCa95c640FdFf35Ef899212734
     */

    constructor() {
        TemperatureSensorFeed = AggregatorV3Interface(0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43);
        HumiditySensorFeed = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        PressureSensorFeed = AggregatorV3Interface(0x42585eD362B3f1BCa95c640FdFf35Ef899212734);
    }

    /**
     * Returns the latest sensor data
     */
    function LatestTemperatureSensorData() public view returns (uint80,int) {
        (
            uint80 roundID,
            int price,
            uint startedAt,
            uint timeStamp,
            uint80 answeredInRound
        ) = TemperatureSensorFeed.latestRoundData();
        return (roundID,price);
    }

    function LatestHumiditySensorData() public view returns (uint80,int) {
        (
            uint80 roundID,
            int price,
            uint startedAt,
            uint timeStamp,
            uint80 answeredInRound
        ) = HumiditySensorFeed.latestRoundData();
        return (roundID,price);
}

 function LatestPressureSensorData() public view returns (uint80,int) {
        (
            uint80 roundID,
            int price,
            uint startedAt,
            uint timeStamp,
            uint80 answeredInRound
        ) = PressureSensorFeed.latestRoundData();
        return (roundID,price);
    }
}
