// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const MARKET_DATA_UPDATE_TRADE_COMPACT_FIXED_SIZE: usize = 24;

/// size           u16             = MarketDataUpdateTradeCompactFixedSize  (24)
/// r#type         u16             = MARKET_DATA_UPDATE_TRADE_COMPACT  (112)
/// price          f32             = 0
/// volume         f32             = 0
/// date_time      DateTime4Byte   = 0
/// symbol_id      u32             = 0
/// at_bid_or_ask  AtBidOrAskEnum  = BID_ASK_UNSET  (0)
const MARKET_DATA_UPDATE_TRADE_COMPACT_FIXED_DEFAULT: [u8; 24] = [24, 0, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// Sent by the Server to the Client when a trade occurs. This message is
/// a more compact MarketDataUpdateTradeFixed. For the price it uses a 4 byte
/// float.
pub trait MarketDataUpdateTradeCompact {
    /// The price of the trade.
    fn price(&self) -> f32;

    /// The volume of the trade.
    fn volume(&self) -> f32;

    /// The timestamp of the trade in UNIX time format. This does not contain
    /// the milliseconds for compactness.
    fn date_time(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// Indicator whether the trade occurred at the Bid or Ask price.
    fn at_bid_or_ask(&self) -> AtBidOrAskEnum;

    /// The price of the trade.
    fn set_price(&mut self, value: f32) -> &mut Self;

    /// The volume of the trade.
    fn set_volume(&mut self, value: f32) -> &mut Self;

    /// The timestamp of the trade in UNIX time format. This does not contain
    /// the milliseconds for compactness.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// Indicator whether the trade occurred at the Bid or Ask price.
    fn set_at_bid_or_ask(&mut self, value: AtBidOrAskEnum) -> &mut Self;

    fn copy_to(&self, to: &mut impl MarketDataUpdateTradeCompact) {
        to.set_price(self.price());
        to.set_volume(self.volume());
        to.set_date_time(self.date_time());
        to.set_symbol_id(self.symbol_id());
        to.set_at_bid_or_ask(self.at_bid_or_ask());
    }
}

/// Sent by the Server to the Client when a trade occurs. This message is
/// a more compact MarketDataUpdateTradeFixed. For the price it uses a 4 byte
/// float.
pub struct MarketDataUpdateTradeCompactFixed {
    data: *const MarketDataUpdateTradeCompactFixedData
}

pub struct MarketDataUpdateTradeCompactFixedUnsafe {
    data: *const MarketDataUpdateTradeCompactFixedData
}

#[repr(packed, C)]
pub struct MarketDataUpdateTradeCompactFixedData {
    size: u16,
    r#type: u16,
    price: f32,
    volume: f32,
    date_time: DateTime4Byte,
    symbol_id: u32,
    at_bid_or_ask: AtBidOrAskEnum,
}

impl MarketDataUpdateTradeCompactFixedData {
    pub fn new() -> Self {
        Self {
            size: 24u16.to_le(),
            r#type: MARKET_DATA_UPDATE_TRADE_COMPACT.to_le(),
            price: 0.0f32,
            volume: 0.0f32,
            date_time: 0,
            symbol_id: 0,
            at_bid_or_ask: AtBidOrAskEnum::BidAskUnset.to_le(),
        }
    }
}

unsafe impl Send for MarketDataUpdateTradeCompactFixed {}
unsafe impl Send for MarketDataUpdateTradeCompactFixedUnsafe {}
unsafe impl Send for MarketDataUpdateTradeCompactFixedData {}

impl Drop for MarketDataUpdateTradeCompactFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateTradeCompactFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateTradeCompactFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateTradeCompactFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateTradeCompactFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateTradeCompactFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateTradeCompactFixed {
    type Target = MarketDataUpdateTradeCompactFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateTradeCompactFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateTradeCompactFixedUnsafe {
    type Target = MarketDataUpdateTradeCompactFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateTradeCompactFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateTradeCompactFixed {
    type Safe = MarketDataUpdateTradeCompactFixed;
    type Unsafe = MarketDataUpdateTradeCompactFixedUnsafe;
    type Data = MarketDataUpdateTradeCompactFixedData;
    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateTradeCompactFixedData::new())
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
            data: data as *const MarketDataUpdateTradeCompactFixedData
        }
    }

}

impl crate::Message for MarketDataUpdateTradeCompactFixedUnsafe {
    type Safe = MarketDataUpdateTradeCompactFixed;
    type Unsafe = MarketDataUpdateTradeCompactFixedUnsafe;
    type Data = MarketDataUpdateTradeCompactFixedData;
    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateTradeCompactFixedData::new())
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
            data: data as *const MarketDataUpdateTradeCompactFixedData
        }
    }

}

/// Sent by the Server to the Client when a trade occurs. This message is
/// a more compact MarketDataUpdateTradeFixed. For the price it uses a 4 byte
/// float.
impl MarketDataUpdateTradeCompact for MarketDataUpdateTradeCompactFixed {
    /// The price of the trade.
    fn price(&self) -> f32 {
        crate::f32_le(self.price)
    }

    /// The volume of the trade.
    fn volume(&self) -> f32 {
        crate::f32_le(self.volume)
    }

    /// The timestamp of the trade in UNIX time format. This does not contain
    /// the milliseconds for compactness.
    fn date_time(&self) -> DateTime4Byte {
        u32::from_le(self.date_time)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// Indicator whether the trade occurred at the Bid or Ask price.
    fn at_bid_or_ask(&self) -> AtBidOrAskEnum {
        AtBidOrAskEnum::from_le(self.at_bid_or_ask)
    }

    /// The price of the trade.
    fn set_price(&mut self, value: f32) -> &mut Self {
        self.price = f32_le(value);
        self
    }


    /// The volume of the trade.
    fn set_volume(&mut self, value: f32) -> &mut Self {
        self.volume = f32_le(value);
        self
    }


    /// The timestamp of the trade in UNIX time format. This does not contain
    /// the milliseconds for compactness.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        self.date_time = value.to_le();
        self
    }


    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }


    /// Indicator whether the trade occurred at the Bid or Ask price.
    fn set_at_bid_or_ask(&mut self, value: AtBidOrAskEnum) -> &mut Self {
        self.at_bid_or_ask = unsafe { core::mem::transmute((value as u16).to_le()) };
        self
    }

}

/// Sent by the Server to the Client when a trade occurs. This message is
/// a more compact MarketDataUpdateTradeFixed. For the price it uses a 4 byte
/// float.
impl MarketDataUpdateTradeCompact for MarketDataUpdateTradeCompactFixedUnsafe {
    /// The price of the trade.
    fn price(&self) -> f32 {
        if self.is_out_of_bounds(8) {
            0.0f32
        } else {
            crate::f32_le(self.price)
        }
    }

    /// The volume of the trade.
    fn volume(&self) -> f32 {
        if self.is_out_of_bounds(12) {
            0.0f32
        } else {
            crate::f32_le(self.volume)
        }
    }

    /// The timestamp of the trade in UNIX time format. This does not contain
    /// the milliseconds for compactness.
    fn date_time(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(16) {
            0
        } else {
            u32::from_le(self.date_time)
        }
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(20) {
            0
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    /// Indicator whether the trade occurred at the Bid or Ask price.
    fn at_bid_or_ask(&self) -> AtBidOrAskEnum {
        if self.is_out_of_bounds(22) {
            AtBidOrAskEnum::BidAskUnset.to_le()
        } else {
            AtBidOrAskEnum::from_le(self.at_bid_or_ask)
        }
    }

    /// The price of the trade.
    fn set_price(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.price = f32_le(value);
        }
        self
    }


    /// The volume of the trade.
    fn set_volume(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.volume = f32_le(value);
        }
        self
    }


    /// The timestamp of the trade in UNIX time format. This does not contain
    /// the milliseconds for compactness.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.date_time = value.to_le();
        }
        self
    }


    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.symbol_id = value.to_le();
        }
        self
    }


    /// Indicator whether the trade occurred at the Bid or Ask price.
    fn set_at_bid_or_ask(&mut self, value: AtBidOrAskEnum) -> &mut Self {
        if !self.is_out_of_bounds(22) {
            self.at_bid_or_ask = unsafe { core::mem::transmute((value as u16).to_le()) };
        }
        self
    }

}

