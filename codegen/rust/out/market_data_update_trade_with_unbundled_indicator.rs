// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_FIXED_SIZE: usize = 40;

/// size                       u16                          = MarketDataUpdateTradeWithUnbundledIndicatorFixedSize  (40)
/// r#type                     u16                          = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR  (137)
/// symbol_id                  u32                          = 0
/// at_bid_or_ask              AtBidOrAskEnum8              = BID_ASK_UNSET_8  (0)
/// unbundled_trade_indicator  UnbundledTradeIndicatorEnum  = UNBUNDLED_TRADE_NONE  (0)
/// trade_condition            u8                           = 0
/// reserve1                   u8                           = 0
/// reserve2                   u32                          = 0
/// price                      f64                          = 0
/// volume                     u32                          = 0
/// reserve3                   u32                          = 0
/// date_time                  DateTimeWithMilliseconds     = 0
const MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_FIXED_DEFAULT: [u8; 40] = [40, 0, 137, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait MarketDataUpdateTradeWithUnbundledIndicator {
    fn symbol_id(&self) -> u32;

    fn at_bid_or_ask(&self) -> AtBidOrAskEnum8;

    fn unbundled_trade_indicator(&self) -> UnbundledTradeIndicatorEnum;

    fn trade_condition(&self) -> u8;

    fn reserve1(&self) -> u8;

    fn reserve2(&self) -> u32;

    fn price(&self) -> f64;

    fn volume(&self) -> u32;

    fn reserve3(&self) -> u32;

    fn date_time(&self) -> DateTimeWithMilliseconds;

    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    fn set_at_bid_or_ask(&mut self, value: AtBidOrAskEnum8) -> &mut Self;

    fn set_unbundled_trade_indicator(&mut self, value: UnbundledTradeIndicatorEnum) -> &mut Self;

    fn set_trade_condition(&mut self, value: u8) -> &mut Self;

    fn set_reserve1(&mut self, value: u8) -> &mut Self;

    fn set_reserve2(&mut self, value: u32) -> &mut Self;

    fn set_price(&mut self, value: f64) -> &mut Self;

    fn set_volume(&mut self, value: u32) -> &mut Self;

    fn set_reserve3(&mut self, value: u32) -> &mut Self;

    fn set_date_time(&mut self, value: DateTimeWithMilliseconds) -> &mut Self;

    fn copy_to(&self, to: &mut impl MarketDataUpdateTradeWithUnbundledIndicator) {
        to.set_symbol_id(self.symbol_id());
        to.set_at_bid_or_ask(self.at_bid_or_ask());
        to.set_unbundled_trade_indicator(self.unbundled_trade_indicator());
        to.set_trade_condition(self.trade_condition());
        to.set_reserve1(self.reserve1());
        to.set_reserve2(self.reserve2());
        to.set_price(self.price());
        to.set_volume(self.volume());
        to.set_reserve3(self.reserve3());
        to.set_date_time(self.date_time());
    }
}

pub struct MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    data: *const MarketDataUpdateTradeWithUnbundledIndicatorFixedData
}

pub struct MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    data: *const MarketDataUpdateTradeWithUnbundledIndicatorFixedData
}

#[repr(packed, C)]
pub struct MarketDataUpdateTradeWithUnbundledIndicatorFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    at_bid_or_ask: AtBidOrAskEnum8,
    unbundled_trade_indicator: UnbundledTradeIndicatorEnum,
    trade_condition: u8,
    reserve1: u8,
    reserve2: u32,
    price: f64,
    volume: u32,
    reserve3: u32,
    date_time: DateTimeWithMilliseconds,
}

impl MarketDataUpdateTradeWithUnbundledIndicatorFixedData {
    pub fn new() -> Self {
        Self {
            size: 40u16.to_le(),
            r#type: MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR.to_le(),
            symbol_id: 0,
            at_bid_or_ask: AtBidOrAskEnum8::BidAskUnset8.to_le(),
            unbundled_trade_indicator: UnbundledTradeIndicatorEnum::UnbundledTradeNone.to_le(),
            trade_condition: 0,
            reserve1: 0,
            reserve2: 0,
            price: 0.0f64,
            volume: 0,
            reserve3: 0,
            date_time: 0.0f64,
        }
    }
}

unsafe impl Send for MarketDataUpdateTradeWithUnbundledIndicatorFixed {}
unsafe impl Send for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {}
unsafe impl Send for MarketDataUpdateTradeWithUnbundledIndicatorFixedData {}

impl Drop for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    type Target = MarketDataUpdateTradeWithUnbundledIndicatorFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    type Target = MarketDataUpdateTradeWithUnbundledIndicatorFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    type Safe = MarketDataUpdateTradeWithUnbundledIndicatorFixed;
    type Unsafe = MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe;
    type Data = MarketDataUpdateTradeWithUnbundledIndicatorFixedData;
    const BASE_SIZE: usize = 40;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateTradeWithUnbundledIndicatorFixedData::new())
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
            data: data as *const MarketDataUpdateTradeWithUnbundledIndicatorFixedData
        }
    }

}

impl crate::Message for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    type Safe = MarketDataUpdateTradeWithUnbundledIndicatorFixed;
    type Unsafe = MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe;
    type Data = MarketDataUpdateTradeWithUnbundledIndicatorFixedData;
    const BASE_SIZE: usize = 40;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateTradeWithUnbundledIndicatorFixedData::new())
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
            data: data as *const MarketDataUpdateTradeWithUnbundledIndicatorFixedData
        }
    }

}

impl MarketDataUpdateTradeWithUnbundledIndicator for MarketDataUpdateTradeWithUnbundledIndicatorFixed {
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn at_bid_or_ask(&self) -> AtBidOrAskEnum8 {
        AtBidOrAskEnum8::from_le(self.at_bid_or_ask)
    }

    fn unbundled_trade_indicator(&self) -> UnbundledTradeIndicatorEnum {
        UnbundledTradeIndicatorEnum::from_le(self.unbundled_trade_indicator)
    }

    fn trade_condition(&self) -> u8 {
        self.trade_condition
    }

    fn reserve1(&self) -> u8 {
        self.reserve1
    }

    fn reserve2(&self) -> u32 {
        u32::from_le(self.reserve2)
    }

    fn price(&self) -> f64 {
        crate::f64_le(self.price)
    }

    fn volume(&self) -> u32 {
        u32::from_le(self.volume)
    }

    fn reserve3(&self) -> u32 {
        u32::from_le(self.reserve3)
    }

    fn date_time(&self) -> DateTimeWithMilliseconds {
        crate::f64_le(self.date_time)
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }


    fn set_at_bid_or_ask(&mut self, value: AtBidOrAskEnum8) -> &mut Self {
        self.at_bid_or_ask = unsafe { core::mem::transmute((value as u8).to_le()) };
        self
    }


    fn set_unbundled_trade_indicator(&mut self, value: UnbundledTradeIndicatorEnum) -> &mut Self {
        self.unbundled_trade_indicator = unsafe { core::mem::transmute((value as i8).to_le()) };
        self
    }


    fn set_trade_condition(&mut self, value: u8) -> &mut Self {
        self.trade_condition = value;
        self
    }


    fn set_reserve1(&mut self, value: u8) -> &mut Self {
        self.reserve1 = value;
        self
    }


    fn set_reserve2(&mut self, value: u32) -> &mut Self {
        self.reserve2 = value.to_le();
        self
    }


    fn set_price(&mut self, value: f64) -> &mut Self {
        self.price = f64_le(value);
        self
    }


    fn set_volume(&mut self, value: u32) -> &mut Self {
        self.volume = value.to_le();
        self
    }


    fn set_reserve3(&mut self, value: u32) -> &mut Self {
        self.reserve3 = value.to_le();
        self
    }


    fn set_date_time(&mut self, value: DateTimeWithMilliseconds) -> &mut Self {
        self.date_time = f64_le(value);
        self
    }

}

impl MarketDataUpdateTradeWithUnbundledIndicator for MarketDataUpdateTradeWithUnbundledIndicatorFixedUnsafe {
    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn at_bid_or_ask(&self) -> AtBidOrAskEnum8 {
        if self.is_out_of_bounds(9) {
            AtBidOrAskEnum8::BidAskUnset8.to_le()
        } else {
            AtBidOrAskEnum8::from_le(self.at_bid_or_ask)
        }
    }

    fn unbundled_trade_indicator(&self) -> UnbundledTradeIndicatorEnum {
        if self.is_out_of_bounds(10) {
            UnbundledTradeIndicatorEnum::UnbundledTradeNone.to_le()
        } else {
            UnbundledTradeIndicatorEnum::from_le(self.unbundled_trade_indicator)
        }
    }

    fn trade_condition(&self) -> u8 {
        if self.is_out_of_bounds(11) {
            0
        } else {
            self.trade_condition
        }
    }

    fn reserve1(&self) -> u8 {
        if self.is_out_of_bounds(12) {
            0
        } else {
            self.reserve1
        }
    }

    fn reserve2(&self) -> u32 {
        if self.is_out_of_bounds(16) {
            0
        } else {
            u32::from_le(self.reserve2)
        }
    }

    fn price(&self) -> f64 {
        if self.is_out_of_bounds(24) {
            0.0f64
        } else {
            crate::f64_le(self.price)
        }
    }

    fn volume(&self) -> u32 {
        if self.is_out_of_bounds(28) {
            0
        } else {
            u32::from_le(self.volume)
        }
    }

    fn reserve3(&self) -> u32 {
        if self.is_out_of_bounds(32) {
            0
        } else {
            u32::from_le(self.reserve3)
        }
    }

    fn date_time(&self) -> DateTimeWithMilliseconds {
        if self.is_out_of_bounds(40) {
            0.0f64
        } else {
            crate::f64_le(self.date_time)
        }
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.symbol_id = value.to_le();
        }
        self
    }


    fn set_at_bid_or_ask(&mut self, value: AtBidOrAskEnum8) -> &mut Self {
        if !self.is_out_of_bounds(9) {
            self.at_bid_or_ask = unsafe { core::mem::transmute((value as u8).to_le()) };
        }
        self
    }


    fn set_unbundled_trade_indicator(&mut self, value: UnbundledTradeIndicatorEnum) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.unbundled_trade_indicator = unsafe { core::mem::transmute((value as i8).to_le()) };
        }
        self
    }


    fn set_trade_condition(&mut self, value: u8) -> &mut Self {
        if !self.is_out_of_bounds(11) {
            self.trade_condition = value;
        }
        self
    }


    fn set_reserve1(&mut self, value: u8) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.reserve1 = value;
        }
        self
    }


    fn set_reserve2(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.reserve2 = value.to_le();
        }
        self
    }


    fn set_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.price = f64_le(value);
        }
        self
    }


    fn set_volume(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(28) {
            self.volume = value.to_le();
        }
        self
    }


    fn set_reserve3(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(32) {
            self.reserve3 = value.to_le();
        }
        self
    }


    fn set_date_time(&mut self, value: DateTimeWithMilliseconds) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            self.date_time = f64_le(value);
        }
        self
    }

}

