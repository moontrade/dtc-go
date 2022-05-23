// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const MARKET_DATA_UPDATE_OPEN_INTEREST_FIXED_SIZE: usize = 16;

/// size                  u16            = MarketDataUpdateOpenInterestFixedSize  (16)
/// r#type                u16            = MARKET_DATA_UPDATE_OPEN_INTEREST  (124)
/// symbol_id             u32            = 0
/// open_interest         u32            = 0
/// trading_session_date  DateTime4Byte  = 0
const MARKET_DATA_UPDATE_OPEN_INTEREST_FIXED_DEFAULT: [u8; 16] = [16, 0, 124, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// The MarketDataUpdateOpenInterestFixed message is sent by the Server to
/// the Client to update the OpenInterest field previously sent through the
/// MarketDataSnapshotFixed message.
pub trait MarketDataUpdateOpenInterest {
    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// The open interest for the symbol.
    fn open_interest(&self) -> u32;

    fn trading_session_date(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// The open interest for the symbol.
    fn set_open_interest(&mut self, value: u32) -> &mut Self;

    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self;

    fn copy_to(&self, to: &mut impl MarketDataUpdateOpenInterest) {
        to.set_symbol_id(self.symbol_id());
        to.set_open_interest(self.open_interest());
        to.set_trading_session_date(self.trading_session_date());
    }
}

/// The MarketDataUpdateOpenInterestFixed message is sent by the Server to
/// the Client to update the OpenInterest field previously sent through the
/// MarketDataSnapshotFixed message.
pub struct MarketDataUpdateOpenInterestFixed {
    data: *const MarketDataUpdateOpenInterestFixedData
}

pub struct MarketDataUpdateOpenInterestFixedUnsafe {
    data: *const MarketDataUpdateOpenInterestFixedData
}

#[repr(packed, C)]
pub struct MarketDataUpdateOpenInterestFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    open_interest: u32,
    trading_session_date: DateTime4Byte,
}

impl MarketDataUpdateOpenInterestFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: MARKET_DATA_UPDATE_OPEN_INTEREST.to_le(),
            symbol_id: 0,
            open_interest: 0,
            trading_session_date: 0,
        }
    }
}

unsafe impl Send for MarketDataUpdateOpenInterestFixed {}
unsafe impl Send for MarketDataUpdateOpenInterestFixedUnsafe {}
unsafe impl Send for MarketDataUpdateOpenInterestFixedData {}

impl Drop for MarketDataUpdateOpenInterestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateOpenInterestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateOpenInterestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateOpenInterestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateOpenInterestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateOpenInterestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateOpenInterestFixed {
    type Target = MarketDataUpdateOpenInterestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateOpenInterestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateOpenInterestFixedUnsafe {
    type Target = MarketDataUpdateOpenInterestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateOpenInterestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateOpenInterestFixed {
    type Safe = MarketDataUpdateOpenInterestFixed;
    type Unsafe = MarketDataUpdateOpenInterestFixedUnsafe;
    type Data = MarketDataUpdateOpenInterestFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateOpenInterestFixedData::new())
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
            data: data as *const MarketDataUpdateOpenInterestFixedData
        }
    }

}

impl crate::Message for MarketDataUpdateOpenInterestFixedUnsafe {
    type Safe = MarketDataUpdateOpenInterestFixed;
    type Unsafe = MarketDataUpdateOpenInterestFixedUnsafe;
    type Data = MarketDataUpdateOpenInterestFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateOpenInterestFixedData::new())
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
            data: data as *const MarketDataUpdateOpenInterestFixedData
        }
    }

}

/// The MarketDataUpdateOpenInterestFixed message is sent by the Server to
/// the Client to update the OpenInterest field previously sent through the
/// MarketDataSnapshotFixed message.
impl MarketDataUpdateOpenInterest for MarketDataUpdateOpenInterestFixed {
    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// The open interest for the symbol.
    fn open_interest(&self) -> u32 {
        u32::from_le(self.open_interest)
    }

    fn trading_session_date(&self) -> DateTime4Byte {
        u32::from_le(self.trading_session_date)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }


    /// The open interest for the symbol.
    fn set_open_interest(&mut self, value: u32) -> &mut Self {
        self.open_interest = value.to_le();
        self
    }


    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self {
        self.trading_session_date = value.to_le();
        self
    }

}

/// The MarketDataUpdateOpenInterestFixed message is sent by the Server to
/// the Client to update the OpenInterest field previously sent through the
/// MarketDataSnapshotFixed message.
impl MarketDataUpdateOpenInterest for MarketDataUpdateOpenInterestFixedUnsafe {
    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    /// The open interest for the symbol.
    fn open_interest(&self) -> u32 {
        if self.is_out_of_bounds(12) {
            0
        } else {
            u32::from_le(self.open_interest)
        }
    }

    fn trading_session_date(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(16) {
            0
        } else {
            u32::from_le(self.trading_session_date)
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


    /// The open interest for the symbol.
    fn set_open_interest(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.open_interest = value.to_le();
        }
        self
    }


    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.trading_session_date = value.to_le();
        }
        self
    }

}

