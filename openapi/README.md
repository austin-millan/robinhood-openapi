# Go API client for openapi

Robinhood API Documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://api.robinhood.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**GetAccounts**](docs/AccountsApi.md#getaccounts) | **Get** /accounts/ | getAccounts
*AccountsApi* | [**GetPortfolio**](docs/AccountsApi.md#getportfolio) | **Get** /accounts/{accountId}/portfolio/ | getPortfolio
*AccountsApi* | [**GetPosition**](docs/AccountsApi.md#getposition) | **Get** /accounts/{accountId}/positions/{positionId}/ | getPosition
*AccountsApi* | [**GetPositions**](docs/AccountsApi.md#getpositions) | **Get** /accounts/{accountId}/positions/ | getPositions
*AuthenticationApi* | [**Login**](docs/AuthenticationApi.md#login) | **Post** /api-token-auth/ | login
*AuthenticationApi* | [**Logout**](docs/AuthenticationApi.md#logout) | **Post** /api-token-logout/ | logout
*DividendsApi* | [**GetDividend**](docs/DividendsApi.md#getdividend) | **Get** /dividends/{id}/ | getDividend
*FundamentalsApi* | [**GetFundamentals**](docs/FundamentalsApi.md#getfundamentals) | **Get** /fundamentals/ | getFundamentals
*FundamentalsApi* | [**GetSymbolFundamentals**](docs/FundamentalsApi.md#getsymbolfundamentals) | **Get** /fundamentals/{symbol}/ | getSymbolFundamentals
*InstrumentsApi* | [**GetInstrument**](docs/InstrumentsApi.md#getinstrument) | **Get** /instruments/{instrument_id}/ | getInstrument
*InstrumentsApi* | [**GetInstrumentSplits**](docs/InstrumentsApi.md#getinstrumentsplits) | **Get** /instruments/{instrument_id}/splits/ | getInstrumentSplits
*InstrumentsApi* | [**GetInstruments**](docs/InstrumentsApi.md#getinstruments) | **Get** /instruments/ | getInstruments
*MarketsApi* | [**GetMArketHours**](docs/MarketsApi.md#getmarkethours) | **Get** /markets/{mic}/hours/{date}/ | getMArketHours
*MarketsApi* | [**GetMarkets**](docs/MarketsApi.md#getmarkets) | **Get** /markets | getMarkets
*MoversApi* | [**GetMovers**](docs/MoversApi.md#getmovers) | **Get** /midlands/movers/sp500/ | getMovers
*OrderApi* | [**CancelOrder**](docs/OrderApi.md#cancelorder) | **Post** /orders/{order_id}/cancel/ | cancelOrder
*OrderApi* | [**GetOptionsOrders**](docs/OrderApi.md#getoptionsorders) | **Get** /options/orders | getOptionsOrders
*OrderApi* | [**GetOrder**](docs/OrderApi.md#getorder) | **Get** /orders/{order_id}/ | getOrder
*OrderApi* | [**GetOrders**](docs/OrderApi.md#getorders) | **Get** /orders/ | getOrders
*OrderApi* | [**PlaceOrder**](docs/OrderApi.md#placeorder) | **Post** /orders/ | placeOrder
*QuoteApi* | [**GetQuotes**](docs/QuoteApi.md#getquotes) | **Get** /quotes/ | getQuotes
*QuoteApi* | [**GetSymbolQuote**](docs/QuoteApi.md#getsymbolquote) | **Get** /quotes/{symbol}/ | getSymbolQuote
*UserApi* | [**GetBasicUserInfo**](docs/UserApi.md#getbasicuserinfo) | **Get** /user/basic_info/ | getBasicUserInfo
*UserApi* | [**GetInvestmentProfile**](docs/UserApi.md#getinvestmentprofile) | **Get** /user/investment_profile/ | getInvestmentProfile
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /user/ | getUser
*UserApi* | [**GetUserID**](docs/UserApi.md#getuserid) | **Get** /user/id/ | getUserID
*WatchlistsApi* | [**BulkAddWatchlists**](docs/WatchlistsApi.md#bulkaddwatchlists) | **Post** /watchlists/Default/bulk_add/ | bulkAddWatchlists
*WatchlistsApi* | [**CreateWatchlist**](docs/WatchlistsApi.md#createwatchlist) | **Post** /watchlists/ | createWatchlist
*WatchlistsApi* | [**DeleteInstrumentFromWatchlist**](docs/WatchlistsApi.md#deleteinstrumentfromwatchlist) | **Delete** /watchlists/{name}/{instrumentId} | deleteInstrumentFromWatchlist
*WatchlistsApi* | [**GetWatchlistByName**](docs/WatchlistsApi.md#getwatchlistbyname) | **Get** /watchlists/{name}/ | getWatchlistByName
*WatchlistsApi* | [**GetWatchlists**](docs/WatchlistsApi.md#getwatchlists) | **Get** /watchlists/ | getWatchlists


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountInfo](docs/AccountInfo.md)
 - [AccountType](docs/AccountType.md)
 - [Auth](docs/Auth.md)
 - [BasicInfo](docs/BasicInfo.md)
 - [CashBalances](docs/CashBalances.md)
 - [CryptoAccount](docs/CryptoAccount.md)
 - [CryptoAssetCurrency](docs/CryptoAssetCurrency.md)
 - [CryptoCurrencyPair](docs/CryptoCurrencyPair.md)
 - [CryptoOrder](docs/CryptoOrder.md)
 - [CryptoOrderOptions](docs/CryptoOrderOptions.md)
 - [CryptoOrderOutput](docs/CryptoOrderOutput.md)
 - [CryptoPortfolio](docs/CryptoPortfolio.md)
 - [Direction](docs/Direction.md)
 - [Execution](docs/Execution.md)
 - [ExecutionType](docs/ExecutionType.md)
 - [FundamentalsData](docs/FundamentalsData.md)
 - [GetOptionOrdersResponse](docs/GetOptionOrdersResponse.md)
 - [InstrumentData](docs/InstrumentData.md)
 - [InstrumentSplit](docs/InstrumentSplit.md)
 - [InstrumentState](docs/InstrumentState.md)
 - [InvestmentExperience](docs/InvestmentExperience.md)
 - [InvestmentObjective](docs/InvestmentObjective.md)
 - [InvestmentProfile](docs/InvestmentProfile.md)
 - [Leg](docs/Leg.md)
 - [LiquidityNeeds](docs/LiquidityNeeds.md)
 - [MarginBalances](docs/MarginBalances.md)
 - [MaritalStatus](docs/MaritalStatus.md)
 - [MarketData](docs/MarketData.md)
 - [MarketHours](docs/MarketHours.md)
 - [MinTicks](docs/MinTicks.md)
 - [OpenCloseStrategy](docs/OpenCloseStrategy.md)
 - [OptionChain](docs/OptionChain.md)
 - [OptionInstrument](docs/OptionInstrument.md)
 - [OptionOrder](docs/OptionOrder.md)
 - [OptionOrderInput](docs/OptionOrderInput.md)
 - [OptionOrderLeg](docs/OptionOrderLeg.md)
 - [OptionOrderLegExecutions](docs/OptionOrderLegExecutions.md)
 - [OptionsMarketData](docs/OptionsMarketData.md)
 - [Order](docs/Order.md)
 - [OrderAction](docs/OrderAction.md)
 - [OrderState](docs/OrderState.md)
 - [OrderTotalNotional](docs/OrderTotalNotional.md)
 - [PaginatedAccountInfo](docs/PaginatedAccountInfo.md)
 - [PaginatedFundamentalsData](docs/PaginatedFundamentalsData.md)
 - [PaginatedInstrumentData](docs/PaginatedInstrumentData.md)
 - [PaginatedInstrumentSplit](docs/PaginatedInstrumentSplit.md)
 - [PaginatedMarketData](docs/PaginatedMarketData.md)
 - [PaginatedMovers](docs/PaginatedMovers.md)
 - [PaginatedOptionChain](docs/PaginatedOptionChain.md)
 - [PaginatedOptionInstrument](docs/PaginatedOptionInstrument.md)
 - [PaginatedOrder](docs/PaginatedOrder.md)
 - [PaginatedPosition](docs/PaginatedPosition.md)
 - [PaginatedQuoteData](docs/PaginatedQuoteData.md)
 - [PaginatedWatchListCreateResponse](docs/PaginatedWatchListCreateResponse.md)
 - [PaginatedWatchListsData](docs/PaginatedWatchListsData.md)
 - [Portfolio](docs/Portfolio.md)
 - [Position](docs/Position.md)
 - [PositionEffect](docs/PositionEffect.md)
 - [Quote](docs/Quote.md)
 - [QuoteCurrency](docs/QuoteCurrency.md)
 - [QuoteData](docs/QuoteData.md)
 - [RiskTolerance](docs/RiskTolerance.md)
 - [Side](docs/Side.md)
 - [SourceOfFunds](docs/SourceOfFunds.md)
 - [TaxBracket](docs/TaxBracket.md)
 - [TimeHorizon](docs/TimeHorizon.md)
 - [TimeInForce](docs/TimeInForce.md)
 - [TotalNetWorth](docs/TotalNetWorth.md)
 - [Trigger](docs/Trigger.md)
 - [UnderlyingInstrument](docs/UnderlyingInstrument.md)
 - [UserId](docs/UserId.md)
 - [UserInfo](docs/UserInfo.md)
 - [WatchListCreateResponse](docs/WatchListCreateResponse.md)
 - [WatchListsData](docs/WatchListsData.md)
 - [Watchlist](docs/Watchlist.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author

austin.millan@gmail.com

