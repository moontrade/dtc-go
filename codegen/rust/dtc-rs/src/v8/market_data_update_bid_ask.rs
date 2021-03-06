// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const MARKET_DATA_UPDATE_BID_ASK_FIXED_SIZE: usize = 40;

/// size          u16            = MarketDataUpdateBidAskFixedSize  (40)
/// type          u16            = MARKET_DATA_UPDATE_BID_ASK  (108)
/// symbol_id     u32            = 0
/// bid_price     f64            = f64::MAX
/// bid_quantity  f32            = 0
/// ask_price     f64            = f64::MAX
/// ask_quantity  f32            = 0
/// date_time     DateTime4Byte  = 0
pub(crate) const MARKET_DATA_UPDATE_BID_ASK_FIXED_DEFAULT: [u8; 40] = [
    40, 0, 108, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 255,
    255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// The Server sends this market data feed message to the Client when the
/// best bid or ask price or size changes.
pub trait MarketDataUpdateBidAsk: Message {
    type Safe: MarketDataUpdateBidAsk;
    type Unsafe: MarketDataUpdateBidAsk;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// The current Bid price. Leave unset if there is no price available.
    fn bid_price(&self) -> f64;

    /// The current number of contracts/shares at the bid price.
    fn bid_quantity(&self) -> f32;

    /// The current ask or offer price. Leave unset if there is no price available.
    /// The current ask or offer price. Leave unset if there is no price available.
    fn ask_price(&self) -> f64;

    /// The current number of contracts/shares at the ask price.
    fn ask_quantity(&self) -> f32;

    /// The Date-Time of the Bid and Ask update.
    fn date_time(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// The current Bid price. Leave unset if there is no price available.
    fn set_bid_price(&mut self, value: f64) -> &mut Self;

    /// The current number of contracts/shares at the bid price.
    fn set_bid_quantity(&mut self, value: f32) -> &mut Self;

    /// The current ask or offer price. Leave unset if there is no price available.
    /// The current ask or offer price. Leave unset if there is no price available.
    fn set_ask_price(&mut self, value: f64) -> &mut Self;

    /// The current number of contracts/shares at the ask price.
    fn set_ask_quantity(&mut self, value: f32) -> &mut Self;

    /// The Date-Time of the Bid and Ask update.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketDataUpdateBidAsk) {
        to.set_symbol_id(self.symbol_id());
        to.set_bid_price(self.bid_price());
        to.set_bid_quantity(self.bid_quantity());
        to.set_ask_price(self.ask_price());
        to.set_ask_quantity(self.ask_quantity());
        to.set_date_time(self.date_time());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 4 {
            return Err(crate::Error::Malformed("need more data"));
        }
        let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
        let base_size = if Self::BASE_SIZE_OFFSET == 0 {
            size
        } else {
            let base_size = unsafe {
                u16::from_le(*(data.as_ptr().offset(Self::BASE_SIZE_OFFSET) as *const u16))
            };
            if base_size > size {
                return Err(crate::Error::Malformed("base_size is greater than size"));
            }
            base_size
        };
        if (base_size as usize) >= Self::BASE_SIZE {
            let msg = unsafe { Self::Safe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Safe(&msg));
            core::mem::forget(msg);
            r
        } else {
            let msg = unsafe { Self::Unsafe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Unsafe(&msg));
            core::mem::forget(msg);
            r
        }
    }
}

/// The Server sends this market data feed message to the Client when the
/// best bid or ask price or size changes.
pub struct MarketDataUpdateBidAskFixed {
    data: *const MarketDataUpdateBidAskFixedData,
}

pub struct MarketDataUpdateBidAskFixedUnsafe {
    data: *const MarketDataUpdateBidAskFixedData,
}

#[repr(packed(8), C)]
pub struct MarketDataUpdateBidAskFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    bid_price: f64,
    bid_quantity: f32,
    ask_price: f64,
    ask_quantity: f32,
    date_time: DateTime4Byte,
}

impl MarketDataUpdateBidAskFixedData {
    pub fn new() -> Self {
        Self {
            size: 40u16.to_le(),
            r#type: MARKET_DATA_UPDATE_BID_ASK.to_le(),
            symbol_id: 0u32,
            bid_price: f64_le(f64::MAX),
            bid_quantity: 0.0f32,
            ask_price: f64_le(f64::MAX),
            ask_quantity: 0.0f32,
            date_time: 0u32,
        }
    }
}

unsafe impl Send for MarketDataUpdateBidAskFixed {}
unsafe impl Send for MarketDataUpdateBidAskFixedUnsafe {}
unsafe impl Send for MarketDataUpdateBidAskFixedData {}

impl Drop for MarketDataUpdateBidAskFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateBidAskFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateBidAskFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateBidAskFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateBidAskFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateBidAskFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateBidAskFixed {
    type Target = MarketDataUpdateBidAskFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateBidAskFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateBidAskFixedUnsafe {
    type Target = MarketDataUpdateBidAskFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateBidAskFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for MarketDataUpdateBidAskFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateBidAskFixed(size: {}, type: {}, symbol_id: {}, bid_price: {}, bid_quantity: {}, ask_price: {}, ask_quantity: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.bid_price(), self.bid_quantity(), self.ask_price(), self.ask_quantity(), self.date_time()).as_str())
    }
}

impl core::fmt::Debug for MarketDataUpdateBidAskFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateBidAskFixed(size: {}, type: {}, symbol_id: {}, bid_price: {}, bid_quantity: {}, ask_price: {}, ask_quantity: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.bid_price(), self.bid_quantity(), self.ask_price(), self.ask_quantity(), self.date_time()).as_str())
    }
}

impl core::fmt::Display for MarketDataUpdateBidAskFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateBidAskFixedUnsafe(size: {}, type: {}, symbol_id: {}, bid_price: {}, bid_quantity: {}, ask_price: {}, ask_quantity: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.bid_price(), self.bid_quantity(), self.ask_price(), self.ask_quantity(), self.date_time()).as_str())
    }
}

impl core::fmt::Debug for MarketDataUpdateBidAskFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateBidAskFixedUnsafe(size: {}, type: {}, symbol_id: {}, bid_price: {}, bid_quantity: {}, ask_price: {}, ask_quantity: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.bid_price(), self.bid_quantity(), self.ask_price(), self.ask_quantity(), self.date_time()).as_str())
    }
}

impl crate::Message for MarketDataUpdateBidAskFixed {
    type Data = MarketDataUpdateBidAskFixedData;

    const TYPE: u16 = MARKET_DATA_UPDATE_BID_ASK;
    const BASE_SIZE: usize = 40;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateBidAskFixedData::new()),
        }
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
            data: data as *const MarketDataUpdateBidAskFixedData,
        }
    }
}
impl crate::Message for MarketDataUpdateBidAskFixedUnsafe {
    type Data = MarketDataUpdateBidAskFixedData;

    const TYPE: u16 = MARKET_DATA_UPDATE_BID_ASK;
    const BASE_SIZE: usize = 40;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateBidAskFixedData::new()),
        }
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
            data: data as *const MarketDataUpdateBidAskFixedData,
        }
    }
}
/// The Server sends this market data feed message to the Client when the
/// best bid or ask price or size changes.
impl MarketDataUpdateBidAsk for MarketDataUpdateBidAskFixed {
    type Safe = MarketDataUpdateBidAskFixed;
    type Unsafe = MarketDataUpdateBidAskFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// The current Bid price. Leave unset if there is no price available.
    fn bid_price(&self) -> f64 {
        f64_le(self.bid_price)
    }

    /// The current number of contracts/shares at the bid price.
    fn bid_quantity(&self) -> f32 {
        f32_le(self.bid_quantity)
    }

    /// The current ask or offer price. Leave unset if there is no price available.
    /// The current ask or offer price. Leave unset if there is no price available.
    fn ask_price(&self) -> f64 {
        f64_le(self.ask_price)
    }

    /// The current number of contracts/shares at the ask price.
    fn ask_quantity(&self) -> f32 {
        f32_le(self.ask_quantity)
    }

    /// The Date-Time of the Bid and Ask update.
    fn date_time(&self) -> DateTime4Byte {
        u32::from_le(self.date_time)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    /// The current Bid price. Leave unset if there is no price available.
    fn set_bid_price(&mut self, value: f64) -> &mut Self {
        self.bid_price = f64_le(value);
        self
    }

    /// The current number of contracts/shares at the bid price.
    fn set_bid_quantity(&mut self, value: f32) -> &mut Self {
        self.bid_quantity = f32_le(value);
        self
    }

    /// The current ask or offer price. Leave unset if there is no price available.
    /// The current ask or offer price. Leave unset if there is no price available.
    fn set_ask_price(&mut self, value: f64) -> &mut Self {
        self.ask_price = f64_le(value);
        self
    }

    /// The current number of contracts/shares at the ask price.
    fn set_ask_quantity(&mut self, value: f32) -> &mut Self {
        self.ask_quantity = f32_le(value);
        self
    }

    /// The Date-Time of the Bid and Ask update.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        self.date_time = value.to_le();
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }
}

/// The Server sends this market data feed message to the Client when the
/// best bid or ask price or size changes.
impl MarketDataUpdateBidAsk for MarketDataUpdateBidAskFixedUnsafe {
    type Safe = MarketDataUpdateBidAskFixed;
    type Unsafe = MarketDataUpdateBidAskFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    /// The current Bid price. Leave unset if there is no price available.
    fn bid_price(&self) -> f64 {
        if self.is_out_of_bounds(16) {
            f64_le(f64::MAX)
        } else {
            f64_le(self.bid_price)
        }
    }

    /// The current number of contracts/shares at the bid price.
    fn bid_quantity(&self) -> f32 {
        if self.is_out_of_bounds(20) {
            0.0f32
        } else {
            f32_le(self.bid_quantity)
        }
    }

    /// The current ask or offer price. Leave unset if there is no price available.
    /// The current ask or offer price. Leave unset if there is no price available.
    fn ask_price(&self) -> f64 {
        if self.is_out_of_bounds(32) {
            f64_le(f64::MAX)
        } else {
            f64_le(self.ask_price)
        }
    }

    /// The current number of contracts/shares at the ask price.
    fn ask_quantity(&self) -> f32 {
        if self.is_out_of_bounds(36) {
            0.0f32
        } else {
            f32_le(self.ask_quantity)
        }
    }

    /// The Date-Time of the Bid and Ask update.
    fn date_time(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(40) {
            0u32
        } else {
            u32::from_le(self.date_time)
        }
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.symbol_id = value.to_le();
        }
        self
    }

    /// The current Bid price. Leave unset if there is no price available.
    fn set_bid_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.bid_price = f64_le(value);
        }
        self
    }

    /// The current number of contracts/shares at the bid price.
    fn set_bid_quantity(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.bid_quantity = f32_le(value);
        }
        self
    }

    /// The current ask or offer price. Leave unset if there is no price available.
    /// The current ask or offer price. Leave unset if there is no price available.
    fn set_ask_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(32) {
            self.ask_price = f64_le(value);
        }
        self
    }

    /// The current number of contracts/shares at the ask price.
    fn set_ask_quantity(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(36) {
            self.ask_quantity = f32_le(value);
        }
        self
    }

    /// The Date-Time of the Bid and Ask update.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            self.date_time = value.to_le();
        }
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }
}

#[cfg(test)]
pub(crate) mod tests {
    use super::*;

    #[test]
    pub(crate) fn layout() {
        unsafe {
            assert_eq!(
                40usize,
                core::mem::size_of::<MarketDataUpdateBidAskFixedData>(),
                "MarketDataUpdateBidAskFixedData sizeof expected {:} but was {:}",
                40usize,
                core::mem::size_of::<MarketDataUpdateBidAskFixedData>()
            );
            assert_eq!(
                40u16,
                MarketDataUpdateBidAskFixed::new().size(),
                "MarketDataUpdateBidAskFixed sizeof expected {:} but was {:}",
                40u16,
                MarketDataUpdateBidAskFixed::new().size(),
            );
            assert_eq!(
                MARKET_DATA_UPDATE_BID_ASK,
                MarketDataUpdateBidAskFixed::new().r#type(),
                "MarketDataUpdateBidAskFixed type expected {:} but was {:}",
                MARKET_DATA_UPDATE_BID_ASK,
                MarketDataUpdateBidAskFixed::new().r#type(),
            );
            assert_eq!(
                108u16,
                MarketDataUpdateBidAskFixed::new().r#type(),
                "MarketDataUpdateBidAskFixed type expected {:} but was {:}",
                108u16,
                MarketDataUpdateBidAskFixed::new().r#type(),
            );
            let d = MarketDataUpdateBidAskFixedData::new();
            let p = (&d as *const _ as *const u8).offset(0) as usize;
            assert_eq!(
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
                "size offset expected {:} but was {:}",
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
            );
            assert_eq!(
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
                "type offset expected {:} but was {:}",
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
            );
            assert_eq!(
                4usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
                "symbol_id offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.bid_price) as usize) - p,
                "bid_price offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.bid_price) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.bid_quantity) as usize) - p,
                "bid_quantity offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.bid_quantity) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.ask_price) as usize) - p,
                "ask_price offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.ask_price) as usize) - p,
            );
            assert_eq!(
                32usize,
                (core::ptr::addr_of!(d.ask_quantity) as usize) - p,
                "ask_quantity offset expected {:} but was {:}",
                32usize,
                (core::ptr::addr_of!(d.ask_quantity) as usize) - p,
            );
            assert_eq!(
                36usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                36usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
        }
    }
}
