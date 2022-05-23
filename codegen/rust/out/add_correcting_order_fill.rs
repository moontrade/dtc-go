// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const ADD_CORRECTING_ORDER_FILL_VLS_SIZE: usize = 56;

const ADD_CORRECTING_ORDER_FILL_FIXED_SIZE: usize = 216;

/// size             u16          = AddCorrectingOrderFillVLSSize  (56)
/// r#type           u16          = ADD_CORRECTING_ORDER_FILL  (309)
/// base_size        u16          = AddCorrectingOrderFillVLSSize  (56)
/// symbol           string       = ""
/// exchange         string       = ""
/// trade_account    string       = ""
/// client_order_id  string       = ""
/// buy_sell         BuySellEnum  = BUY_SELL_UNSET  (0)
/// fill_price       f64          = 0
/// fill_quantity    f64          = 0
/// free_form_text   string       = ""
const ADD_CORRECTING_ORDER_FILL_VLS_DEFAULT: [u8; 56] = [56, 0, 53, 1, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size             u16          = AddCorrectingOrderFillFixedSize  (216)
/// r#type           u16          = ADD_CORRECTING_ORDER_FILL  (309)
/// symbol           string64     = ""
/// exchange         string16     = ""
/// trade_account    string32     = ""
/// client_order_id  string32     = ""
/// buy_sell         BuySellEnum  = BUY_SELL_UNSET  (0)
/// fill_price       f64          = 0
/// fill_quantity    f64          = 0
/// free_form_text   string48     = ""
const ADD_CORRECTING_ORDER_FILL_FIXED_DEFAULT: [u8; 216] = [216, 0, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait AddCorrectingOrderFill {
    fn symbol(&self) -> &str;

    fn exchange(&self) -> &str;

    fn trade_account(&self) -> &str;

    fn client_order_id(&self) -> &str;

    fn buy_sell(&self) -> BuySellEnum;

    fn fill_price(&self) -> f64;

    fn fill_quantity(&self) -> f64;

    fn free_form_text(&self) -> &str;

    fn set_symbol(&mut self, value: &str) -> &mut Self;

    fn set_exchange(&mut self, value: &str) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn set_client_order_id(&mut self, value: &str) -> &mut Self;

    fn set_buy_sell(&mut self, value: BuySellEnum) -> &mut Self;

    fn set_fill_price(&mut self, value: f64) -> &mut Self;

    fn set_fill_quantity(&mut self, value: f64) -> &mut Self;

    fn set_free_form_text(&mut self, value: &str) -> &mut Self;

    fn copy_to(&self, to: &mut impl AddCorrectingOrderFill) {
        to.set_symbol(self.symbol());
        to.set_exchange(self.exchange());
        to.set_trade_account(self.trade_account());
        to.set_client_order_id(self.client_order_id());
        to.set_buy_sell(self.buy_sell());
        to.set_fill_price(self.fill_price());
        to.set_fill_quantity(self.fill_quantity());
        to.set_free_form_text(self.free_form_text());
    }
}

pub struct AddCorrectingOrderFillVLS {
    data: *const AddCorrectingOrderFillVLSData,
    capacity: usize
}

pub struct AddCorrectingOrderFillVLSUnsafe {
    data: *const AddCorrectingOrderFillVLSData,
    capacity: usize
}

#[repr(packed, C)]
pub struct AddCorrectingOrderFillVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    symbol: VLS,
    exchange: VLS,
    trade_account: VLS,
    client_order_id: VLS,
    buy_sell: BuySellEnum,
    fill_price: f64,
    fill_quantity: f64,
    free_form_text: VLS,
}

pub struct AddCorrectingOrderFillFixed {
    data: *const AddCorrectingOrderFillFixedData
}

pub struct AddCorrectingOrderFillFixedUnsafe {
    data: *const AddCorrectingOrderFillFixedData
}

#[repr(packed, C)]
pub struct AddCorrectingOrderFillFixedData {
    size: u16,
    r#type: u16,
    symbol: [u8; 64],
    exchange: [u8; 16],
    trade_account: [u8; 32],
    client_order_id: [u8; 32],
    buy_sell: BuySellEnum,
    fill_price: f64,
    fill_quantity: f64,
    free_form_text: [u8; 48],
}

impl AddCorrectingOrderFillVLSData {
    pub fn new() -> Self {
        Self {
            size: 56u16.to_le(),
            r#type: ADD_CORRECTING_ORDER_FILL.to_le(),
            base_size: 56u16.to_le(),
            symbol: crate::message::VLS::new(),
            exchange: crate::message::VLS::new(),
            trade_account: crate::message::VLS::new(),
            client_order_id: crate::message::VLS::new(),
            buy_sell: BuySellEnum::BuySellUnset.to_le(),
            fill_price: 0,
            fill_quantity: 0,
            free_form_text: crate::message::VLS::new(),
        }
    }
}

impl AddCorrectingOrderFillFixedData {
    pub fn new() -> Self {
        Self {
            size: 216u16.to_le(),
            r#type: ADD_CORRECTING_ORDER_FILL.to_le(),
            symbol: [0; 64],
            exchange: [0; 16],
            trade_account: [0; 32],
            client_order_id: [0; 32],
            buy_sell: BuySellEnum::BuySellUnset.to_le(),
            fill_price: 0.0f64,
            fill_quantity: 0.0f64,
            free_form_text: [0; 48],
        }
    }
}

unsafe impl Send for AddCorrectingOrderFillFixed {}
unsafe impl Send for AddCorrectingOrderFillFixedUnsafe {}
unsafe impl Send for AddCorrectingOrderFillFixedData {}
unsafe impl Send for AddCorrectingOrderFillVLS {}
unsafe impl Send for AddCorrectingOrderFillVLSUnsafe {}
unsafe impl Send for AddCorrectingOrderFillVLSData {}

impl Drop for AddCorrectingOrderFillFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AddCorrectingOrderFillFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AddCorrectingOrderFillVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AddCorrectingOrderFillVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for AddCorrectingOrderFillFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AddCorrectingOrderFillFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AddCorrectingOrderFillVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AddCorrectingOrderFillVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for AddCorrectingOrderFillFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AddCorrectingOrderFillFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AddCorrectingOrderFillVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AddCorrectingOrderFillVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for AddCorrectingOrderFillFixed {
    type Target = AddCorrectingOrderFillFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AddCorrectingOrderFillFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AddCorrectingOrderFillFixedUnsafe {
    type Target = AddCorrectingOrderFillFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AddCorrectingOrderFillFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AddCorrectingOrderFillVLS {
    type Target = AddCorrectingOrderFillVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AddCorrectingOrderFillVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AddCorrectingOrderFillVLSUnsafe {
    type Target = AddCorrectingOrderFillVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AddCorrectingOrderFillVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for AddCorrectingOrderFillFixed {
    type Safe = AddCorrectingOrderFillFixed;
    type Unsafe = AddCorrectingOrderFillFixedUnsafe;
    type Data = AddCorrectingOrderFillFixedData;
    const BASE_SIZE: usize = 216;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AddCorrectingOrderFillFixedData::new())
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
            data: data as *const AddCorrectingOrderFillFixedData
        }
    }

}

impl crate::Message for AddCorrectingOrderFillFixedUnsafe {
    type Safe = AddCorrectingOrderFillFixed;
    type Unsafe = AddCorrectingOrderFillFixedUnsafe;
    type Data = AddCorrectingOrderFillFixedData;
    const BASE_SIZE: usize = 216;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AddCorrectingOrderFillFixedData::new())
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
            data: data as *const AddCorrectingOrderFillFixedData
        }
    }

}

impl crate::Message for AddCorrectingOrderFillVLS {
    type Safe = AddCorrectingOrderFillVLS;
    type Unsafe = AddCorrectingOrderFillVLSUnsafe;
    type Data = AddCorrectingOrderFillVLSData;
    const BASE_SIZE: usize = 56;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AddCorrectingOrderFillVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY
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
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const AddCorrectingOrderFillVLSData,
            capacity
        }
    }

}

impl crate::VLSMessage for AddCorrectingOrderFillVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const AddCorrectingOrderFillVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }

}impl crate::Message for AddCorrectingOrderFillVLSUnsafe {
    type Safe = AddCorrectingOrderFillVLS;
    type Unsafe = AddCorrectingOrderFillVLSUnsafe;
    type Data = AddCorrectingOrderFillVLSData;
    const BASE_SIZE: usize = 56;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AddCorrectingOrderFillVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY
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
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const AddCorrectingOrderFillVLSData,
            capacity
        }
    }

}

impl crate::VLSMessage for AddCorrectingOrderFillVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const AddCorrectingOrderFillVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }

}impl AddCorrectingOrderFill for AddCorrectingOrderFillVLS {
    fn symbol(&self) -> &str {
        crate::get_vls(self, self.symbol)
    }

    fn exchange(&self) -> &str {
        crate::get_vls(self, self.exchange)
    }

    fn trade_account(&self) -> &str {
        crate::get_vls(self, self.trade_account)
    }

    fn client_order_id(&self) -> &str {
        crate::get_vls(self, self.client_order_id)
    }

    fn buy_sell(&self) -> BuySellEnum {
        BuySellEnum::from_le(self.buy_sell)
    }

    fn fill_price(&self) -> f64 {
        crate::f64_le(self.fill_price)
    }

    fn fill_quantity(&self) -> f64 {
        crate::f64_le(self.fill_quantity)
    }

    fn free_form_text(&self) -> &str {
        crate::get_vls(self, self.free_form_text)
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        self.symbol = crate::set_vls(self, self.symbol, value);
        self
    }


    fn set_exchange(&mut self, value: &str) -> &mut Self {
        self.exchange = crate::set_vls(self, self.exchange, value);
        self
    }


    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = crate::set_vls(self, self.trade_account, value);
        self
    }


    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        self.client_order_id = crate::set_vls(self, self.client_order_id, value);
        self
    }


    fn set_buy_sell(&mut self, value: BuySellEnum) -> &mut Self {
        self.buy_sell = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }


    fn set_fill_price(&mut self, value: f64) -> &mut Self {
        self.fill_price = f64_le(value);
        self
    }


    fn set_fill_quantity(&mut self, value: f64) -> &mut Self {
        self.fill_quantity = f64_le(value);
        self
    }


    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        self.free_form_text = crate::set_vls(self, self.free_form_text, value);
        self
    }

}

impl AddCorrectingOrderFill for AddCorrectingOrderFillVLSUnsafe {
    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            crate::get_vls(self, self.symbol)
        }
    }

    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            crate::get_vls(self, self.exchange)
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(18) {
            ""
        } else {
            crate::get_vls(self, self.trade_account)
        }
    }

    fn client_order_id(&self) -> &str {
        if self.is_out_of_bounds(22) {
            ""
        } else {
            crate::get_vls(self, self.client_order_id)
        }
    }

    fn buy_sell(&self) -> BuySellEnum {
        if self.is_out_of_bounds(28) {
            BuySellEnum::BuySellUnset.to_le()
        } else {
            BuySellEnum::from_le(self.buy_sell)
        }
    }

    fn fill_price(&self) -> f64 {
        if self.is_out_of_bounds(40) {
            0
        } else {
            crate::f64_le(self.fill_price)
        }
    }

    fn fill_quantity(&self) -> f64 {
        if self.is_out_of_bounds(48) {
            0
        } else {
            crate::f64_le(self.fill_quantity)
        }
    }

    fn free_form_text(&self) -> &str {
        if self.is_out_of_bounds(52) {
            ""
        } else {
            crate::get_vls(self, self.free_form_text)
        }
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.symbol = crate::set_vls(self, self.symbol, value);
        }
        self
    }


    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.exchange = crate::set_vls(self, self.exchange, value);
        }
        self
    }


    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.trade_account = crate::set_vls(self, self.trade_account, value);
        }
        self
    }


    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(22) {
            self.client_order_id = crate::set_vls(self, self.client_order_id, value);
        }
        self
    }


    fn set_buy_sell(&mut self, value: BuySellEnum) -> &mut Self {
        if !self.is_out_of_bounds(28) {
            self.buy_sell = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }


    fn set_fill_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            self.fill_price = f64_le(value);
        }
        self
    }


    fn set_fill_quantity(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.fill_quantity = f64_le(value);
        }
        self
    }


    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(52) {
            self.free_form_text = crate::set_vls(self, self.free_form_text, value);
        }
        self
    }

}

impl AddCorrectingOrderFill for AddCorrectingOrderFillFixed {
    fn symbol(&self) -> &str {
        crate::get_fixed(&self.symbol[..])
    }

    fn exchange(&self) -> &str {
        crate::get_fixed(&self.exchange[..])
    }

    fn trade_account(&self) -> &str {
        crate::get_fixed(&self.trade_account[..])
    }

    fn client_order_id(&self) -> &str {
        crate::get_fixed(&self.client_order_id[..])
    }

    fn buy_sell(&self) -> BuySellEnum {
        BuySellEnum::from_le(self.buy_sell)
    }

    fn fill_price(&self) -> f64 {
        crate::f64_le(self.fill_price)
    }

    fn fill_quantity(&self) -> f64 {
        crate::f64_le(self.fill_quantity)
    }

    fn free_form_text(&self) -> &str {
        crate::get_fixed(&self.free_form_text[..])
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.symbol[..], value);
        self
    }


    fn set_exchange(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.exchange[..], value);
        self
    }


    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.trade_account[..], value);
        self
    }


    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.client_order_id[..], value);
        self
    }


    fn set_buy_sell(&mut self, value: BuySellEnum) -> &mut Self {
        self.buy_sell = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }


    fn set_fill_price(&mut self, value: f64) -> &mut Self {
        self.fill_price = f64_le(value);
        self
    }


    fn set_fill_quantity(&mut self, value: f64) -> &mut Self {
        self.fill_quantity = f64_le(value);
        self
    }


    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.free_form_text[..], value);
        self
    }

}

impl AddCorrectingOrderFill for AddCorrectingOrderFillFixedUnsafe {
    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(68) {
            ""
        } else {
            crate::get_fixed(&self.symbol[..])
        }
    }

    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(84) {
            ""
        } else {
            crate::get_fixed(&self.exchange[..])
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(116) {
            ""
        } else {
            crate::get_fixed(&self.trade_account[..])
        }
    }

    fn client_order_id(&self) -> &str {
        if self.is_out_of_bounds(148) {
            ""
        } else {
            crate::get_fixed(&self.client_order_id[..])
        }
    }

    fn buy_sell(&self) -> BuySellEnum {
        if self.is_out_of_bounds(152) {
            BuySellEnum::BuySellUnset.to_le()
        } else {
            BuySellEnum::from_le(self.buy_sell)
        }
    }

    fn fill_price(&self) -> f64 {
        if self.is_out_of_bounds(160) {
            0.0f64
        } else {
            crate::f64_le(self.fill_price)
        }
    }

    fn fill_quantity(&self) -> f64 {
        if self.is_out_of_bounds(168) {
            0.0f64
        } else {
            crate::f64_le(self.fill_quantity)
        }
    }

    fn free_form_text(&self) -> &str {
        if self.is_out_of_bounds(216) {
            ""
        } else {
            crate::get_fixed(&self.free_form_text[..])
        }
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(68) {
            crate::set_fixed(&mut self.symbol[..], value);
        }
        self
    }


    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(84) {
            crate::set_fixed(&mut self.exchange[..], value);
        }
        self
    }


    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(116) {
            crate::set_fixed(&mut self.trade_account[..], value);
        }
        self
    }


    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(148) {
            crate::set_fixed(&mut self.client_order_id[..], value);
        }
        self
    }


    fn set_buy_sell(&mut self, value: BuySellEnum) -> &mut Self {
        if !self.is_out_of_bounds(152) {
            self.buy_sell = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }


    fn set_fill_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(160) {
            self.fill_price = f64_le(value);
        }
        self
    }


    fn set_fill_quantity(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(168) {
            self.fill_quantity = f64_le(value);
        }
        self
    }


    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(216) {
            crate::set_fixed(&mut self.free_form_text[..], value);
        }
        self
    }

}

