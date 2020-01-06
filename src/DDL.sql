DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer` (
  `CustomerID` int unsigned NOT NULL AUTO_INCREMENT,
  `CompanyName` varchar(50),
  `FirstName` varchar(50),
  `LastName` varchar(50),
  `BillingAddress` varchar(255),
  `City` varchar(50),
  `StateOrProvince` varchar(20),
  `ZIPCode` varchar(20),
  `Email` varchar(75),
  `CompanyWebsite` varchar(200),
  `PhoneNumber` varchar(30),
  `FaxNumber` varchar(30),
  `ShipAddress` varchar(255),
  `ShipCity` varchar(50),
  `ShipStateOrProvince` varchar(50),
  `ShipZIPCode` varchar(20),
  `ShipPhoneNumber` varchar(30),
  PRIMARY KEY (`CustomerID`)
) AUTO_INCREMENT=1 ;

DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
`EmployeeID` int unsigned NOT NULL AUTO_INCREMENT,
`FirstName` varchar(50),
`LastName` varchar(50),
`Title` varchar(50),
`WorkPhone` varchar(30),
PRIMARY KEY (`EmployeeID`)
) AUTO_INCREMENT=1 ;

DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
`ProductID` int unsigned NOT NULL AUTO_INCREMENT,
`ProductName` varchar(50),
`UnitPrice` float(6),
`InStock` char(1),
PRIMARY KEY (`ProductID`)
)AUTO_INCREMENT=1 ;

DROP TABLE IF EXISTS `shipping_method`;
CREATE TABLE `shipping_method` (
`ShippingMethodID` int unsigned NOT NULL AUTO_INCREMENT,
`ShippingMethod` varchar(20),
PRIMARY KEY (`ShippingMethodID`)
)AUTO_INCREMENT=1 ;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
`OrderID` int unsigned NOT NULL AUTO_INCREMENT,
`CustomerID` int unsigned,
`EmployeeID` int unsigned,
`ShippingMethodID` int unsigned,
`OrderDate` Date,
`ShipDate` Date,
`FreightCharge` int,
`Taxes` int,
`PaymentReceived` char(1),
`Comment` varchar(150),
PRIMARY KEY (`OrderID`),
FOREIGN KEY(CustomerID) REFERENCES customer(CustomerID),
FOREIGN KEY(EmployeeID) REFERENCES employee(EmployeeID),
FOREIGN KEY(ShippingMethodID) REFERENCES shipping_method(ShippingMethodID)
)AUTO_INCREMENT=1 ;

DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail` (
`OrderDetailID` int unsigned NOT NULL AUTO_INCREMENT,
`OrderID` int unsigned,
`ProductID` int unsigned, 
`Quantity` int,
`UnitPrice` int,
`Discount` int,
PRIMARY KEY (`OrderDetailID`),
FOREIGN KEY (OrderID) REFERENCES orders(OrderID),
FOREIGN KEY (ProductID) REFERENCES product(ProductID)
)AUTO_INCREMENT=1 ;