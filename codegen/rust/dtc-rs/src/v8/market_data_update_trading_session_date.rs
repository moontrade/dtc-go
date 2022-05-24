// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-24 10:12:33.526761 +0800 WITA m=+0.004576126
use super::*;
use crate::message::*;

const MARKET_DATA_UPDATE_TRADING_SESSION_DATE_FIXED_SIZE: usize = 12;

/// size       u16            = MarketDataUpdateTradingSessionDateFixedSize  (12)
/// type       u16            = MARKET_DATA_UPDATE_TRADING_SESSION_DATE  (136)
/// symbol_id  u32            = 0
/// date       DateTime4Byte  = 0
const MARKET_DATA_UPDATE_TRADING_SESSION_DATE_FIXED_DEFAULT: [u8; 12] = [12, 0, 136, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// Sent by the Server to the Client to update the trading session Date.
pub trait MarketDataUpdateTradingSessionDate {
    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// The date of the current trading session. The time component is not normally
    /// considered relevant in this case.
    fn date(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// The date of the current trading session. The time component is not normally
    /// considered relevant in this case.
    fn set_date(&mut self, value: DateTime4Byte) -> &mut Self;

    fn copy_to(&self, to: &mut impl MarketDataUpdateTradingSessionDate) {
        to.set_symbol_id(self.symbol_id());
        to.set_date(self.date());
    }
}

/// Sent by the Server to the Client to update the trading session Date.
pub struct MarketDataUpdateTradingSessionDateFixed {
    data: *const MarketDataUpdateTradingSessionDateFixedData
}

pub struct MarketDataUpdateTradingSessionDateFixedUnsafe {
    data: *const MarketDataUpdateTradingSessionDateFixedData
}

#[repr(packed, C)]
pub struct MarketDataUpdateTradingSessionDateFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    date: DateTime4Byte,
}

impl MarketDataUpdateTradingSessionDateFixedData {
    pub fn new() -> Self {
        Self {
            size: 12u16.to_le(),
            r#type: MARKET_DATA_UPDATE_TRADING_SESSION_DATE.to_le(),
            symbol_id: 0u32,
            date: 0u32,
        }
    }
}

unsafe impl Send for MarketDataUpdateTradingSessionDateFixed {}
unsafe impl Send for MarketDataUpdateTradingSessionDateFixedUnsafe {}
unsafe impl Send for MarketDataUpdateTradingSessionDateFixedData {}

impl Drop for MarketDataUpdateTradingSessionDateFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateTradingSessionDateFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateTradingSessionDateFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateTradingSessionDateFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateTradingSessionDateFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateTradingSessionDateFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateTradingSessionDateFixed {
    type Target = MarketDataUpdateTradingSessionDateFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateTradingSessionDateFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateTradingSessionDateFixedUnsafe {
    type Target = MarketDataUpdateTradingSessionDateFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateTradingSessionDateFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateTradingSessionDateFixed {
    type Safe = MarketDataUpdateTradingSessionDateFixed;
    type Unsafe = MarketDataUpdateTradingSessionDateFixedUnsafe;
    type Data = MarketDataUpdateTradingSessionDateFixedData;
    const BASE_SIZE: usize = 12;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateTradingSessionDateFixedData::new())
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
            data: data as *const MarketDataUpdateTradingSessionDateFixedData
        }
    }

}

impl crate::Message for MarketDataUpdateTradingSessionDateFixedUnsafe {
    type Safe = MarketDataUpdateTradingSessionDateFixed;
    type Unsafe = MarketDataUpdateTradingSessionDateFixedUnsafe;
    type Data = MarketDataUpdateTradingSessionDateFixedData;
    const BASE_SIZE: usize = 12;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateTradingSessionDateFixedData::new())
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
            data: data as *const MarketDataUpdateTradingSessionDateFixedData
        }
    }

}

/// Sent by the Server to the Client to update the trading session Date.
impl MarketDataUpdateTradingSessionDate for MarketDataUpdateTradingSessionDateFixed {
    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// The date of the current trading session. The time component is not normally
    /// considered relevant in this case.
    fn date(&self) -> DateTime4Byte {
        u32::from_le(self.date)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }


    /// The date of the current trading session. The time component is not normally
    /// considered relevant in this case.
    fn set_date(&mut self, value: DateTime4Byte) -> &mut Self {
        self.date = value.to_le();
        self
    }

}

/// Sent by the Server to the Client to update the trading session Date.
impl MarketDataUpdateTradingSessionDate for MarketDataUpdateTradingSessionDateFixedUnsafe {
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

    /// The date of the current trading session. The time component is not normally
    /// considered relevant in this case.
    fn date(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(12) {
            0u32
        } else {
            u32::from_le(self.date)
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


    /// The date of the current trading session. The time component is not normally
    /// considered relevant in this case.
    fn set_date(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.date = value.to_le();
        }
        self
    }

}
