// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const HISTORICAL_PRICE_DATA_RECORD_RESPONSE_FIXED_SIZE: usize = 88;

/// size             u16                          = HistoricalPriceDataRecordResponseFixedSize  (88)
/// r#type           u16                          = HISTORICAL_PRICE_DATA_RECORD_RESPONSE  (803)
/// request_id       i32                          = 0
/// start_date_time  DateTimeWithMicrosecondsInt  = 0
/// open_price       f64                          = 0
/// high_price       f64                          = 0
/// low_price        f64                          = 0
/// last_price       f64                          = 0
/// volume           f64                          = 0
/// bid_volume       f64                          = 0
/// ask_volume       f64                          = 0
/// is_final_record  bool                         = false
const HISTORICAL_PRICE_DATA_RECORD_RESPONSE_FIXED_DEFAULT: [u8; 88] = [88, 0, 35, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// The HistoricalPriceDataTickRecordResponseFixed message is used when the
/// RecordInterval field in a historical data request message is set to a
/// value greater than INTERVAL_TICK. For example, if the RecordInterval is
/// INTERVAL_1_MINUTE, then a message of this type will contain data for a
/// 1 minute timeframe with a start time specified by the StartDateTime field.
/// 1 minute timeframe with a start time specified by the StartDateTime field.
///
/// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponseFixed
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
///
/// This message can be part of a compressed series of messages of this same
/// type, if the Client requested compression be used.
pub trait HistoricalPriceDataRecordResponse {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32;

    /// The starting Date-Time in UTC of the data record in this message.
    ///
    /// It is part of the DTC Protocol specification that this must be the starting
    /// Date-Time of the data record.
    fn start_date_time(&self) -> DateTimeWithMicrosecondsInt;

    /// The Open price of the data record in this message.
    fn open_price(&self) -> f64;

    /// The High price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
    /// price at the time of the trade. In this case the OpenPrice field needs
    /// to be 0 in this case.
    fn high_price(&self) -> f64;

    /// The Low price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the LowPrice field can be the Bid price
    /// at the time of the trade. In this case the OpenPrice field needs to be
    /// 0 in this case.
    fn low_price(&self) -> f64;

    /// The Last price of the data record in this message.
    fn last_price(&self) -> f64;

    /// The Volume of this data record of this message.
    ///
    /// (union)
    fn volume(&self) -> f64;

    fn open_interest(&self) -> u32;

    /// The Open Interest or Number of Trades of this data record in this message.
    /// The Open Interest or Number of Trades of this data record in this message.
    fn num_trades(&self) -> u32;

    /// The volume of trades at the bid price or lower of the data record in this
    /// message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Ask, then BidVolume must be zero.
    fn bid_volume(&self) -> f64;

    /// The volume of trades at the ask price or higher of the data record in
    /// this message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Bid, then AskVolume must be zero.
    fn ask_volume(&self) -> f64;

    /// Set to 1 to indicate final record in response to the historical price
    /// data request.
    ///
    /// The default is 0 meaning there are more records to follow.
    fn is_final_record(&self) -> bool;

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// The starting Date-Time in UTC of the data record in this message.
    ///
    /// It is part of the DTC Protocol specification that this must be the starting
    /// Date-Time of the data record.
    fn set_start_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self;

    /// The Open price of the data record in this message.
    fn set_open_price(&mut self, value: f64) -> &mut Self;

    /// The High price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
    /// price at the time of the trade. In this case the OpenPrice field needs
    /// to be 0 in this case.
    fn set_high_price(&mut self, value: f64) -> &mut Self;

    /// The Low price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the LowPrice field can be the Bid price
    /// at the time of the trade. In this case the OpenPrice field needs to be
    /// 0 in this case.
    fn set_low_price(&mut self, value: f64) -> &mut Self;

    /// The Last price of the data record in this message.
    fn set_last_price(&mut self, value: f64) -> &mut Self;

    /// The Volume of this data record of this message.
    ///
    /// (union)
    fn set_volume(&mut self, value: f64) -> &mut Self;

    fn set_open_interest(&mut self, value: u32) -> &mut Self;
    /// The Open Interest or Number of Trades of this data record in this message.
    /// The Open Interest or Number of Trades of this data record in this message.
    fn set_num_trades(&mut self, value: u32) -> &mut Self;

    /// The volume of trades at the bid price or lower of the data record in this
    /// message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Ask, then BidVolume must be zero.
    fn set_bid_volume(&mut self, value: f64) -> &mut Self;

    /// The volume of trades at the ask price or higher of the data record in
    /// this message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Bid, then AskVolume must be zero.
    fn set_ask_volume(&mut self, value: f64) -> &mut Self;

    /// Set to 1 to indicate final record in response to the historical price
    /// data request.
    ///
    /// The default is 0 meaning there are more records to follow.
    fn set_is_final_record(&mut self, value: bool) -> &mut Self;

    fn copy_to(&self, to: &mut impl HistoricalPriceDataRecordResponse) {
        to.set_request_id(self.request_id());
        to.set_start_date_time(self.start_date_time());
        to.set_open_price(self.open_price());
        to.set_high_price(self.high_price());
        to.set_low_price(self.low_price());
        to.set_last_price(self.last_price());
        to.set_volume(self.volume());
        to.set_(self.());
        to.set_bid_volume(self.bid_volume());
        to.set_ask_volume(self.ask_volume());
        to.set_is_final_record(self.is_final_record());
    }
}

/// The HistoricalPriceDataTickRecordResponseFixed message is used when the
/// RecordInterval field in a historical data request message is set to a
/// value greater than INTERVAL_TICK. For example, if the RecordInterval is
/// INTERVAL_1_MINUTE, then a message of this type will contain data for a
/// 1 minute timeframe with a start time specified by the StartDateTime field.
/// 1 minute timeframe with a start time specified by the StartDateTime field.
///
/// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponseFixed
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
///
/// This message can be part of a compressed series of messages of this same
/// type, if the Client requested compression be used.
pub struct HistoricalPriceDataRecordResponseFixed {
    data: *const HistoricalPriceDataRecordResponseFixedData
}

pub struct HistoricalPriceDataRecordResponseFixedUnsafe {
    data: *const HistoricalPriceDataRecordResponseFixedData
}

#[repr(packed, C)]
pub struct HistoricalPriceDataRecordResponseFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    start_date_time: DateTimeWithMicrosecondsInt,
    open_price: f64,
    high_price: f64,
    low_price: f64,
    last_price: f64,
    volume: f64,
    open_interest: u32,
    bid_volume: f64,
    ask_volume: f64,
    is_final_record: bool,
}

impl HistoricalPriceDataRecordResponseFixedData {
    pub fn new() -> Self {
        Self {
            size: 88u16.to_le(),
            r#type: HISTORICAL_PRICE_DATA_RECORD_RESPONSE.to_le(),
            request_id: 0,
            start_date_time: 0,
            open_price: 0.0f64,
            high_price: 0.0f64,
            low_price: 0.0f64,
            last_price: 0.0f64,
            volume: 0.0f64,
            open_interest: ,
            bid_volume: 0.0f64,
            ask_volume: 0.0f64,
            is_final_record: false,
        }
    }
}

unsafe impl Send for HistoricalPriceDataRecordResponseFixed {}
unsafe impl Send for HistoricalPriceDataRecordResponseFixedUnsafe {}
unsafe impl Send for HistoricalPriceDataRecordResponseFixedData {}

impl Drop for HistoricalPriceDataRecordResponseFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalPriceDataRecordResponseFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalPriceDataRecordResponseFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalPriceDataRecordResponseFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataRecordResponseFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataRecordResponseFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalPriceDataRecordResponseFixed {
    type Target = HistoricalPriceDataRecordResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataRecordResponseFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalPriceDataRecordResponseFixedUnsafe {
    type Target = HistoricalPriceDataRecordResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataRecordResponseFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for HistoricalPriceDataRecordResponseFixed {
    type Safe = HistoricalPriceDataRecordResponseFixed;
    type Unsafe = HistoricalPriceDataRecordResponseFixedUnsafe;
    type Data = HistoricalPriceDataRecordResponseFixedData;
    const BASE_SIZE: usize = 88;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataRecordResponseFixedData::new())
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const HistoricalPriceDataRecordResponseFixedData
        }
    }

}

impl crate::Message for HistoricalPriceDataRecordResponseFixedUnsafe {
    type Safe = HistoricalPriceDataRecordResponseFixed;
    type Unsafe = HistoricalPriceDataRecordResponseFixedUnsafe;
    type Data = HistoricalPriceDataRecordResponseFixedData;
    const BASE_SIZE: usize = 88;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataRecordResponseFixedData::new())
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const HistoricalPriceDataRecordResponseFixedData
        }
    }

}

/// The HistoricalPriceDataTickRecordResponseFixed message is used when the
/// RecordInterval field in a historical data request message is set to a
/// value greater than INTERVAL_TICK. For example, if the RecordInterval is
/// INTERVAL_1_MINUTE, then a message of this type will contain data for a
/// 1 minute timeframe with a start time specified by the StartDateTime field.
/// 1 minute timeframe with a start time specified by the StartDateTime field.
///
/// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponseFixed
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
///
/// This message can be part of a compressed series of messages of this same
/// type, if the Client requested compression be used.
impl HistoricalPriceDataRecordResponse for HistoricalPriceDataRecordResponseFixed {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// The starting Date-Time in UTC of the data record in this message.
    ///
    /// It is part of the DTC Protocol specification that this must be the starting
    /// Date-Time of the data record.
    fn start_date_time(&self) -> DateTimeWithMicrosecondsInt {
        i64::from_le(self.start_date_time)
    }

    /// The Open price of the data record in this message.
    fn open_price(&self) -> f64 {
        crate::f64_le(self.open_price)
    }

    /// The High price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
    /// price at the time of the trade. In this case the OpenPrice field needs
    /// to be 0 in this case.
    fn high_price(&self) -> f64 {
        crate::f64_le(self.high_price)
    }

    /// The Low price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the LowPrice field can be the Bid price
    /// at the time of the trade. In this case the OpenPrice field needs to be
    /// 0 in this case.
    fn low_price(&self) -> f64 {
        crate::f64_le(self.low_price)
    }

    /// The Last price of the data record in this message.
    fn last_price(&self) -> f64 {
        crate::f64_le(self.last_price)
    }

    /// The Volume of this data record of this message.
    ///
    /// (union)
    fn volume(&self) -> f64 {
        crate::f64_le(self.volume)
    }

    fn open_interest(&self) -> u32 {
        u32::from_le(self.open_interest)
    }

    /// The Open Interest or Number of Trades of this data record in this message.
    /// The Open Interest or Number of Trades of this data record in this message.
    fn num_trades(&self) -> u32 {
        u32::from_le(self.open_interest)
    }

    /// The volume of trades at the bid price or lower of the data record in this
    /// message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Ask, then BidVolume must be zero.
    fn bid_volume(&self) -> f64 {
        crate::f64_le(self.bid_volume)
    }

    /// The volume of trades at the ask price or higher of the data record in
    /// this message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Bid, then AskVolume must be zero.
    fn ask_volume(&self) -> f64 {
        crate::f64_le(self.ask_volume)
    }

    /// Set to 1 to indicate final record in response to the historical price
    /// data request.
    ///
    /// The default is 0 meaning there are more records to follow.
    fn is_final_record(&self) -> bool {
        self.is_final_record
    }

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    /// The starting Date-Time in UTC of the data record in this message.
    ///
    /// It is part of the DTC Protocol specification that this must be the starting
    /// Date-Time of the data record.
    fn set_start_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self {
        self.start_date_time = value.to_le();
        self
    }


    /// The Open price of the data record in this message.
    fn set_open_price(&mut self, value: f64) -> &mut Self {
        self.open_price = f64_le(value);
        self
    }


    /// The High price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
    /// price at the time of the trade. In this case the OpenPrice field needs
    /// to be 0 in this case.
    fn set_high_price(&mut self, value: f64) -> &mut Self {
        self.high_price = f64_le(value);
        self
    }


    /// The Low price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the LowPrice field can be the Bid price
    /// at the time of the trade. In this case the OpenPrice field needs to be
    /// 0 in this case.
    fn set_low_price(&mut self, value: f64) -> &mut Self {
        self.low_price = f64_le(value);
        self
    }


    /// The Last price of the data record in this message.
    fn set_last_price(&mut self, value: f64) -> &mut Self {
        self.last_price = f64_le(value);
        self
    }


    /// The Volume of this data record of this message.
    ///
    /// (union)
    fn set_volume(&mut self, value: f64) -> &mut Self {
        self.volume = f64_le(value);
        self
    }


    fn set_open_interest(&self) -> u32 {
        self.open_interest = value.to_le();
        self
    }

    /// The Open Interest or Number of Trades of this data record in this message.
    /// The Open Interest or Number of Trades of this data record in this message.
    fn set_num_trades(&self) -> u32 {
        self.open_interest = value.to_le();
        self
    }


    /// The volume of trades at the bid price or lower of the data record in this
    /// message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Ask, then BidVolume must be zero.
    fn set_bid_volume(&mut self, value: f64) -> &mut Self {
        self.bid_volume = f64_le(value);
        self
    }


    /// The volume of trades at the ask price or higher of the data record in
    /// this message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Bid, then AskVolume must be zero.
    fn set_ask_volume(&mut self, value: f64) -> &mut Self {
        self.ask_volume = f64_le(value);
        self
    }


    /// Set to 1 to indicate final record in response to the historical price
    /// data request.
    ///
    /// The default is 0 meaning there are more records to follow.
    fn set_is_final_record(&mut self, value: bool) -> &mut Self {
        self.is_final_record = value;
        self
    }

}

/// The HistoricalPriceDataTickRecordResponseFixed message is used when the
/// RecordInterval field in a historical data request message is set to a
/// value greater than INTERVAL_TICK. For example, if the RecordInterval is
/// INTERVAL_1_MINUTE, then a message of this type will contain data for a
/// 1 minute timeframe with a start time specified by the StartDateTime field.
/// 1 minute timeframe with a start time specified by the StartDateTime field.
///
/// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponseFixed
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
/// message can still be used instead of HistoricalPriceDataTickRecordResponseFixed.
///
/// This message can be part of a compressed series of messages of this same
/// type, if the Client requested compression be used.
impl HistoricalPriceDataRecordResponse for HistoricalPriceDataRecordResponseFixedUnsafe {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// The starting Date-Time in UTC of the data record in this message.
    ///
    /// It is part of the DTC Protocol specification that this must be the starting
    /// Date-Time of the data record.
    fn start_date_time(&self) -> DateTimeWithMicrosecondsInt {
        if self.is_out_of_bounds(16) {
            0
        } else {
            i64::from_le(self.start_date_time)
        }
    }

    /// The Open price of the data record in this message.
    fn open_price(&self) -> f64 {
        if self.is_out_of_bounds(24) {
            0.0f64
        } else {
            crate::f64_le(self.open_price)
        }
    }

    /// The High price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
    /// price at the time of the trade. In this case the OpenPrice field needs
    /// to be 0 in this case.
    fn high_price(&self) -> f64 {
        if self.is_out_of_bounds(32) {
            0.0f64
        } else {
            crate::f64_le(self.high_price)
        }
    }

    /// The Low price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the LowPrice field can be the Bid price
    /// at the time of the trade. In this case the OpenPrice field needs to be
    /// 0 in this case.
    fn low_price(&self) -> f64 {
        if self.is_out_of_bounds(40) {
            0.0f64
        } else {
            crate::f64_le(self.low_price)
        }
    }

    /// The Last price of the data record in this message.
    fn last_price(&self) -> f64 {
        if self.is_out_of_bounds(48) {
            0.0f64
        } else {
            crate::f64_le(self.last_price)
        }
    }

    /// The Volume of this data record of this message.
    ///
    /// (union)
    fn volume(&self) -> f64 {
        if self.is_out_of_bounds(56) {
            0.0f64
        } else {
            crate::f64_le(self.volume)
        }
    }

    fn open_interest(&self) -> u32 {
        if self.is_out_of_bounds(60) {
            0
        } else {
            u32::from_le(self.open_interest)
        }
    }

    /// The Open Interest or Number of Trades of this data record in this message.
    /// The Open Interest or Number of Trades of this data record in this message.
    fn num_trades(&self) -> u32 {
        if self.is_out_of_bounds(60) {
            0
        } else {
            u32::from_le(self.open_interest)
        }
    }

    /// The volume of trades at the bid price or lower of the data record in this
    /// message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Ask, then BidVolume must be zero.
    fn bid_volume(&self) -> f64 {
        if self.is_out_of_bounds(72) {
            0.0f64
        } else {
            crate::f64_le(self.bid_volume)
        }
    }

    /// The volume of trades at the ask price or higher of the data record in
    /// this message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Bid, then AskVolume must be zero.
    fn ask_volume(&self) -> f64 {
        if self.is_out_of_bounds(80) {
            0.0f64
        } else {
            crate::f64_le(self.ask_volume)
        }
    }

    /// Set to 1 to indicate final record in response to the historical price
    /// data request.
    ///
    /// The default is 0 meaning there are more records to follow.
    fn is_final_record(&self) -> bool {
        if self.is_out_of_bounds(81) {
            false
        } else {
            self.is_final_record
        }
    }

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }


    /// The starting Date-Time in UTC of the data record in this message.
    ///
    /// It is part of the DTC Protocol specification that this must be the starting
    /// Date-Time of the data record.
    fn set_start_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.start_date_time = value.to_le();
        }
        self
    }


    /// The Open price of the data record in this message.
    fn set_open_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.open_price = f64_le(value);
        }
        self
    }


    /// The High price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
    /// price at the time of the trade. In this case the OpenPrice field needs
    /// to be 0 in this case.
    fn set_high_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(32) {
            self.high_price = f64_le(value);
        }
        self
    }


    /// The Low price of the data record in this message.
    ///
    /// In the case where NumTrades is 1, the LowPrice field can be the Bid price
    /// at the time of the trade. In this case the OpenPrice field needs to be
    /// 0 in this case.
    fn set_low_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            self.low_price = f64_le(value);
        }
        self
    }


    /// The Last price of the data record in this message.
    fn set_last_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.last_price = f64_le(value);
        }
        self
    }


    /// The Volume of this data record of this message.
    ///
    /// (union)
    fn set_volume(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(56) {
            self.volume = f64_le(value);
        }
        self
    }


    fn set_open_interest(&self) -> u32 {
        if !self.is_out_of_bounds(60) {
            self.open_interest = value.to_le();
        }
        self
    }

    /// The Open Interest or Number of Trades of this data record in this message.
    /// The Open Interest or Number of Trades of this data record in this message.
    fn set_num_trades(&self) -> u32 {
        if !self.is_out_of_bounds(60) {
            self.open_interest = value.to_le();
        }
        self
    }


    /// The volume of trades at the bid price or lower of the data record in this
    /// message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Ask, then BidVolume must be zero.
    fn set_bid_volume(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(72) {
            self.bid_volume = f64_le(value);
        }
        self
    }


    /// The volume of trades at the ask price or higher of the data record in
    /// this message.
    ///
    /// In the case where this message consists of a single trade, if the trade
    /// was at the Bid, then AskVolume must be zero.
    fn set_ask_volume(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(80) {
            self.ask_volume = f64_le(value);
        }
        self
    }


    /// Set to 1 to indicate final record in response to the historical price
    /// data request.
    ///
    /// The default is 0 meaning there are more records to follow.
    fn set_is_final_record(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(81) {
            self.is_final_record = value;
        }
        self
    }

}

