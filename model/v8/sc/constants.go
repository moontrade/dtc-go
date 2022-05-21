// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-14 08:04:12.929995 +0800 WITA m=+0.027582959

package sc

const (
	CURRENT_VERSION                                         int32  = 8
	USERNAME_PASSWORD_LENGTH                                int32  = 32
	SYMBOL_EXCHANGE_DELIMITER_LENGTH                        int32  = 4
	SYMBOL_LENGTH                                           int32  = 64
	EXCHANGE_LENGTH                                         int32  = 16
	UNDERLYING_SYMBOL_LENGTH                                int32  = 32
	SYMBOL_DESCRIPTION_LENGTH                               int32  = 64
	EXCHANGE_DESCRIPTION_LENGTH                             int32  = 48
	ORDER_ID_LENGTH                                         int32  = 32
	TRADE_ACCOUNT_LENGTH                                    int32  = 32
	TEXT_DESCRIPTION_LENGTH                                 int32  = 96
	TEXT_MESSAGE_LENGTH                                     int32  = 256
	ORDER_FREE_FORM_TEXT_LENGTH                             int32  = 48
	CLIENT_SERVER_NAME_LENGTH                               int32  = 48
	GENERAL_IDENTIFIER_LENGTH                               int32  = 64
	CURRENCY_CODE_LENGTH                                    int32  = 8
	ORDER_FILL_EXECUTION_LENGTH                             int32  = 64
	PRICE_STRING_LENGTH                                     int32  = 16
	LOGON_REQUEST                                           uint16 = 1
	LOGON_RESPONSE                                          uint16 = 2
	HEARTBEAT                                               uint16 = 3
	LOGOFF                                                  uint16 = 5
	ENCODING_REQUEST                                        uint16 = 6
	ENCODING_RESPONSE                                       uint16 = 7
	MARKET_DATA_REQUEST                                     uint16 = 101
	MARKET_DATA_REJECT                                      uint16 = 103
	MARKET_DATA_SNAPSHOT                                    uint16 = 104
	MARKET_DATA_SNAPSHOT_INT                                uint16 = 125
	MARKET_DATA_UPDATE_TRADE                                uint16 = 107
	MARKET_DATA_UPDATE_TRADE_COMPACT                        uint16 = 112
	MARKET_DATA_UPDATE_TRADE_INT                            uint16 = 126
	MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT                  uint16 = 134
	MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR       uint16 = 137
	MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2     uint16 = 146
	MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP                   uint16 = 142
	MARKET_DATA_UPDATE_BID_ASK                              uint16 = 108
	MARKET_DATA_UPDATE_BID_ASK_COMPACT                      uint16 = 117
	MARKET_DATA_UPDATE_BID_ASK_INT                          uint16 = 127
	MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP                 uint16 = 143
	MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS      uint16 = 144
	MARKET_DATA_UPDATE_SESSION_OPEN                         uint16 = 120
	MARKET_DATA_UPDATE_SESSION_OPEN_INT                     uint16 = 128
	MARKET_DATA_UPDATE_SESSION_HIGH                         uint16 = 114
	MARKET_DATA_UPDATE_SESSION_HIGH_INT                     uint16 = 129
	MARKET_DATA_UPDATE_SESSION_LOW                          uint16 = 115
	MARKET_DATA_UPDATE_SESSION_LOW_INT                      uint16 = 130
	MARKET_DATA_UPDATE_SESSION_VOLUME                       uint16 = 113
	MARKET_DATA_UPDATE_OPEN_INTEREST                        uint16 = 124
	MARKET_DATA_UPDATE_SESSION_SETTLEMENT                   uint16 = 119
	MARKET_DATA_UPDATE_SESSION_SETTLEMENT_INT               uint16 = 131
	MARKET_DATA_UPDATE_SESSION_NUM_TRADES                   uint16 = 135
	MARKET_DATA_UPDATE_TRADING_SESSION_DATE                 uint16 = 136
	MARKET_DEPTH_REQUEST                                    uint16 = 102
	MARKET_DEPTH_REJECT                                     uint16 = 121
	MARKET_DEPTH_SNAPSHOT_LEVEL                             uint16 = 122
	MARKET_DEPTH_SNAPSHOT_LEVEL_INT                         uint16 = 132
	MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT                       uint16 = 145
	MARKET_DEPTH_UPDATE_LEVEL                               uint16 = 106
	MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS       uint16 = 140
	MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP                  uint16 = 141
	MARKET_DEPTH_UPDATE_LEVEL_INT                           uint16 = 133
	MARKET_DATA_FEED_STATUS                                 uint16 = 100
	MARKET_DATA_FEED_SYMBOL_STATUS                          uint16 = 116
	TRADING_SYMBOL_STATUS                                   uint16 = 138
	MARKET_ORDERS_REQUEST                                   uint16 = 150
	MARKET_ORDERS_REJECT                                    uint16 = 151
	MARKET_ORDERS_ADD                                       uint16 = 152
	MARKET_ORDERS_MODIFY                                    uint16 = 153
	MARKET_ORDERS_REMOVE                                    uint16 = 154
	MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY                 uint16 = 155
	SUBMIT_NEW_SINGLE_ORDER                                 uint16 = 208
	SUBMIT_NEW_SINGLE_ORDER_INT                             uint16 = 206
	SUBMIT_NEW_OCO_ORDER                                    uint16 = 201
	SUBMIT_NEW_OCO_ORDER_INT                                uint16 = 207
	SUBMIT_FLATTEN_POSITION_ORDER                           uint16 = 209
	FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT                     uint16 = 210
	CANCEL_ORDER                                            uint16 = 203
	CANCEL_REPLACE_ORDER                                    uint16 = 204
	CANCEL_REPLACE_ORDER_INT                                uint16 = 205
	OPEN_ORDERS_REQUEST                                     uint16 = 300
	OPEN_ORDERS_REJECT                                      uint16 = 302
	ORDER_UPDATE                                            uint16 = 301
	HISTORICAL_ORDER_FILLS_REQUEST                          uint16 = 303
	HISTORICAL_ORDER_FILLS_REJECT                           uint16 = 308
	HISTORICAL_ORDER_FILL_RESPONSE                          uint16 = 304
	CURRENT_POSITIONS_REQUEST                               uint16 = 305
	CURRENT_POSITIONS_REJECT                                uint16 = 307
	POSITION_UPDATE                                         uint16 = 306
	ADD_CORRECTING_ORDER_FILL                               uint16 = 309
	CORRECTING_ORDER_FILL_RESPONSE                          uint16 = 310
	TRADE_ACCOUNTS_REQUEST                                  uint16 = 400
	TRADE_ACCOUNT_RESPONSE                                  uint16 = 401
	EXCHANGE_LIST_REQUEST                                   uint16 = 500
	EXCHANGE_LIST_RESPONSE                                  uint16 = 501
	SYMBOLS_FOR_EXCHANGE_REQUEST                            uint16 = 502
	UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST                 uint16 = 503
	SYMBOLS_FOR_UNDERLYING_REQUEST                          uint16 = 504
	SECURITY_DEFINITION_FOR_SYMBOL_REQUEST                  uint16 = 506
	SECURITY_DEFINITION_RESPONSE                            uint16 = 507
	SYMBOL_SEARCH_REQUEST                                   uint16 = 508
	SECURITY_DEFINITION_REJECT                              uint16 = 509
	ACCOUNT_BALANCE_REQUEST                                 uint16 = 601
	ACCOUNT_BALANCE_REJECT                                  uint16 = 602
	ACCOUNT_BALANCE_UPDATE                                  uint16 = 600
	ACCOUNT_BALANCE_ADJUSTMENT                              uint16 = 607
	ACCOUNT_BALANCE_ADJUSTMENT_REJECT                       uint16 = 608
	ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE                     uint16 = 609
	HISTORICAL_ACCOUNT_BALANCES_REQUEST                     uint16 = 603
	HISTORICAL_ACCOUNT_BALANCES_REJECT                      uint16 = 604
	HISTORICAL_ACCOUNT_BALANCE_RESPONSE                     uint16 = 606
	USER_MESSAGE                                            uint16 = 700
	GENERAL_LOG_MESSAGE                                     uint16 = 701
	ALERT_MESSAGE                                           uint16 = 702
	JOURNAL_ENTRY_ADD                                       uint16 = 703
	JOURNAL_ENTRIES_REQUEST                                 uint16 = 704
	JOURNAL_ENTRIES_REJECT                                  uint16 = 705
	JOURNAL_ENTRY_RESPONSE                                  uint16 = 706
	HISTORICAL_PRICE_DATA_REQUEST                           uint16 = 800
	HISTORICAL_PRICE_DATA_RESPONSE_HEADER                   uint16 = 801
	HISTORICAL_PRICE_DATA_REJECT                            uint16 = 802
	HISTORICAL_PRICE_DATA_RECORD_RESPONSE                   uint16 = 803
	HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE              uint16 = 804
	HISTORICAL_PRICE_DATA_RECORD_RESPONSE_INT               uint16 = 805
	HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT          uint16 = 806
	HISTORICAL_PRICE_DATA_RESPONSE_TRAILER                  uint16 = 807
	HISTORICAL_MARKET_DEPTH_DATA_REQUEST                    uint16 = 900
	HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER            uint16 = 901
	HISTORICAL_MARKET_DEPTH_DATA_REJECT                     uint16 = 902
	HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE            uint16 = 903
	HISTORICAL_TRADES_REQUEST                               uint16 = 10100
	HISTORICAL_TRADES_REJECT                                uint16 = 10101
	HISTORICAL_TRADES_RESPONSE                              uint16 = 10102
	REPLAY_CHART_DATA                                       uint16 = 10104
	REPLAY_CHART_DATA_PERFORM_ACTION                        uint16 = 10105
	REPLAY_CHART_DATA_STATUS                                uint16 = 10106
	REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS                  uint16 = 10107
	NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE                 uint16 = 10108
	SC_CONFIGURATION_SYNCHRONIZATION                        uint16 = 10109
	SC_TRADE_ORDER                                          uint16 = 10110
	INDIVIDUAL_TRADE_POSITION                               uint16 = 10112
	TRADE_POSITION_CONSOLIDATED                             uint16 = 10113
	TRADE_ACTIVITY_DATA                                     uint16 = 10114
	TRADE_ACCOUNT_DATA_REQUEST                              uint16 = 10115
	TRADE_ACCOUNT_DATA_RESPONSE                             uint16 = 10116
	TRADE_ACCOUNT_DATA_UPDATE                               uint16 = 10117
	TRADE_ACCOUNT_DATA_DELETE                               uint16 = 10118
	TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_RESPONSE           uint16 = 10119
	TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE             uint16 = 10120
	TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_RESPONSE               uint16 = 10121
	TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE                 uint16 = 10122
	TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_RESPONSE         uint16 = 10124
	TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_ADD              uint16 = 10125
	TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE           uint16 = 10126
	TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_RESPONSE      uint16 = 10127
	TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD           uint16 = 10128
	TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE        uint16 = 10129
	TRADE_ACCOUNT_DATA_RESPONSE_TRAILER                     uint16 = 10130
	TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE            uint16 = 10131
	PROCESSED_FILL_IDENTIFIER                               uint16 = 10132
	INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST           uint16 = 10133
	INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE               uint16 = 10134
	USER_INFORMATION                                        uint16 = 10136
	INTERPROCESS_SYNCHRONIZATION_TRADE_ACTIVITY_REQUEST     uint16 = 10137
	DOWNLOAD_HISTORICAL_ORDER_FILL_AND_ACCOUNT_BALANCE_DATA uint16 = 10138
	CLIENT_DEVICE_UPDATE                                    uint16 = 10139
	WRITE_INTRADAY_DATA_FILE_SESSION_VALUE                  uint16 = 10140
	MARGIN_DATA_REQUEST                                     uint16 = 10141
	MARGIN_DATA_RESPONSE                                    uint16 = 10142
)